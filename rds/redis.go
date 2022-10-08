package rds

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/tristan-club/kit/log"
	"time"
)

var redisClient *redis.Pool

var debug = false

func SetDebug(d bool) {
	debug = d
}

func Connect(host string) {
	maxIdle := 20
	maxActive := 100
	// 建立连接池
	redisClient = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 3600 * time.Second,
		Wait:        false,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", host,
				redis.DialDatabase(4),
				redis.DialConnectTimeout(10*time.Second),
				redis.DialReadTimeout(1*time.Second),
				redis.DialWriteTimeout(1*time.Second))
			if err != nil {
				log.Error().Err(err).Send()
				return nil, err
			}
			return con, nil
		},
	}
}

func GetClient() redis.Conn {
	if debug {
		log.Debug().Str("type", "redis").Msgf("status %v", redisClient.Stats())
	}
	return redisClient.Get()
}

func Do(cmd string, args ...interface{}) (interface{}, error) {
	if debug {
		log.Debug().Str("type", "redis").Msgf("status %v", redisClient.Stats())
	}

	rc := redisClient.Get()

	defer rc.Close()

	if debug {
		log.Debug().Str("type", "redis").Msgf("redis[%s], [%v]", cmd, args)
	}

	reply, err := rc.Do(cmd, args...)

	if err != nil {
		log.Error().Err(err).Send()
	}

	return reply, err
}

func DoStr(cmd string, args ...interface{}) string {
	reply, err := Do(cmd, args...)

	if str, err := redis.String(reply, err); err == nil {
		return str
	}

	return ""
}

// try do a lock
func Lock(name string, maxTime time.Duration) error {
	do, err := Do("set", name, 1, "NX", "EX", maxTime.Milliseconds())

	if err != nil || do == nil {
		return fmt.Errorf("lock key %s failed", name)
	}

	return nil
}

// try do a lock
func UnLock(name string) error {
	_, err := Do("del", name)
	if err != nil {
		return err
	}
	return nil
}

func LogStatus() {
	log.Debug().
		Str("type", "status").
		Str("kind", "redis").
		Msgf("redis status %v", redisClient.Stats())
}
