package scheduler

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/tristan-club/kit/config"
	"github.com/tristan-club/kit/log"
	"github.com/tristan-club/kit/rds"
	"time"
)

type ScheduleHandler func(id string, data []byte)

func Schedule(delay time.Duration, data interface{}) error {
	c := rds.GetClient()
	defer c.Close()

	jobId := uuid.NewV4().String()

	dataJson, err := json.Marshal(data)

	if err != nil {
		log.Error().Str("type", "schedule").Err(err).Send()
		return err
	}

	log.Debug().Str("type", "schedule").Msgf("schedule job %s", jobId)

	_ = c.Send("MULTI")

	_ = c.Send("set", config.RedisScheduleDataKey(jobId), dataJson)

	_ = c.Send("zadd", config.RedisScheduleJobKey(), time.Now().Add(delay).Unix(), jobId)

	_, err = c.Do("exec")

	if err != nil {
		log.Error().Str("type", "schedule").Err(err).Send()
		return err
	}

	return nil
}
func delJob(JobId string) {
	log.Info().Str("type", "schedule").Msgf("delete Job %s", JobId)

	c := rds.GetClient()
	defer c.Close()

	_ = c.Send("MULTI")

	_ = c.Send("del", config.RedisScheduleDataKey(JobId))

	_ = c.Send("zrem", config.RedisScheduleJobKey(), JobId)

	_, err := c.Do("exec")

	if err != nil {
		log.Error().Str("type", "schedule").Err(err).Send()
	}
}

func LoopResume(handler ScheduleHandler) {
	runSchedule := func(jobId string, handler ScheduleHandler) {

		log.Info().Str("type", "schedule").Msgf("get Job %s", jobId)

		bytes, err := redis.Bytes(rds.Do("get", config.RedisScheduleDataKey(jobId)))
		if err != nil {
			log.Error().
				Str("type", "schedule").
				Err(err).Send()

			delJob(jobId)
			return
		}

		err = rds.Lock(config.RedisScheduleLockKey(jobId), time.Hour)
		// lock failed, Job is been running by another process
		if err != nil {
			log.Info().
				Str("type", "schedule").
				Msgf("lock Job %s failed, job is being processed by another process", jobId)

			delJob(jobId)
			return
		}

		handler(jobId, bytes)

		delJob(jobId)
		_ = rds.UnLock(config.RedisScheduleLockKey(jobId))
	}

	for {
		select {
		case _ = <-time.After(time.Second):
		}

		reply, err := redis.Strings(rds.Do("ZRANGEBYSCORE", config.RedisScheduleJobKey(), 0, time.Now().Unix()))

		if err != nil {
			log.Error().Str("type", "schedule").Err(err)
			continue
		}

		for _, jobId := range reply {
			runSchedule(jobId, handler)
		}
	}
}
