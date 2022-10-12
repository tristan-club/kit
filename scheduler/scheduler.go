package scheduler

import (
	"bytes"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/tristan-club/kit/config"
	"github.com/tristan-club/kit/log"
	"github.com/tristan-club/kit/rds"
	"time"
)

type ScheduleHandler func(id string, data interface{})

type Scheduler struct {
	name string
}

func NewScheduler(name string) *Scheduler {
	if !rds.IsConnected() {
		log.Panic().Msgf("redis is not connected, please connect redis first")
	}

	return &Scheduler{
		name: name,
	}
}

func (s *Scheduler) Schedule(delay time.Duration, data interface{}) error {
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

	_ = c.Send("zadd", config.RedisScheduleJobKey(s.name), time.Now().Add(delay).Unix(), jobId)

	_, err = c.Do("exec")

	if err != nil {
		log.Error().Str("type", "schedule").Err(err).Send()
		return err
	}

	return nil
}
func (s *Scheduler) delJob(JobId string) {
	log.Info().Str("type", "schedule").Msgf("delete Job %s", JobId)

	c := rds.GetClient()
	defer c.Close()

	_ = c.Send("MULTI")

	_ = c.Send("del", config.RedisScheduleDataKey(JobId))

	_ = c.Send("zrem", config.RedisScheduleJobKey(s.name), JobId)

	_, err := c.Do("exec")

	if err != nil {
		log.Error().Str("type", "schedule").Err(err).Send()
	}
}

func (s *Scheduler) Run(handler ScheduleHandler) {
	runSchedule := func(jobId string, handler ScheduleHandler) {

		log.Info().Str("type", "schedule").Msgf("get Job %s", jobId)

		b, err := redis.Bytes(rds.Do("get", config.RedisScheduleDataKey(jobId)))
		if err != nil {
			log.Error().
				Str("type", "schedule").
				Err(err).Send()
			return
		}

		err = rds.Lock(config.RedisScheduleLockKey(jobId), time.Hour)
		// lock failed, Job has been run by another process
		if err != nil {
			log.Info().
				Str("type", "schedule").
				Msgf("lock Job %s failed, job is being processed by another process", jobId)
			return
		}

		decoder := json.NewDecoder(bytes.NewBuffer(b))
		decoder.UseNumber()

		var data interface{}
		err = decoder.Decode(&data)

		if err != nil {
			log.Error().
				Str("type", "data").
				Err(err).Send()
		} else {
			handler(jobId, data)
		}

		s.delJob(jobId)
		_ = rds.UnLock(config.RedisScheduleLockKey(jobId))
	}

	for {
		select {
		case _ = <-time.After(time.Second):
		}

		reply, err := redis.Strings(rds.Do("ZRANGEBYSCORE", config.RedisScheduleJobKey(s.name), 0, time.Now().Unix()))

		if err != nil {
			log.Error().Str("type", "schedule").Err(err)
			continue
		}

		for _, jobId := range reply {
			go runSchedule(jobId, handler)
		}
	}
}
