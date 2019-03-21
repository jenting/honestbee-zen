package cache

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	redigotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/garyburd/redigo"

	"github.com/honestbee/Zen/config"
)

type redisPool struct {
	pool *redis.Pool
}

// NewRedis returns a Redis instance.
func NewRedis(conf *config.Config, dbIndex int) (Cache, error) {
	connectTimeout := time.Duration(conf.Cache.ConnectTimeoutSec) * time.Second
	readTimeout := time.Duration(conf.Cache.ReadTimeoutSec) * time.Second
	writeTimeout := time.Duration(conf.Cache.WriteTimeoutSec) * time.Second
	idleTimeout := time.Duration(conf.Cache.IdleTimeoutSec) * time.Second

	return &redisPool{
		pool: &redis.Pool{
			MaxIdle:     conf.Cache.MaxIdle,
			MaxActive:   conf.Cache.MaxActive,
			IdleTimeout: idleTimeout,
			Wait:        conf.Cache.Wait,
			Dial: func() (redis.Conn, error) {
				c, err := redigotrace.Dial(
					"tcp",
					conf.Cache.Host+":"+conf.Cache.Port,
					redis.DialConnectTimeout(connectTimeout),
					redis.DialReadTimeout(readTimeout),
					redis.DialWriteTimeout(writeTimeout),
					redis.DialPassword(conf.Cache.Password),
					redis.DialDatabase(dbIndex),
					redigotrace.WithServiceName("helpcenter-zendesk-redis"),
				)
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"cache: [Dial] dial failed",
					)
				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return errors.Wrapf(
					err,
					"cache: [TestOnBorrow] PING failed",
				)
			},
		},
	}, nil
}

// do is wrapping pool usage of redigo do.
func (r *redisPool) do(cmd string, args ...interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	reply, err := conn.Do(cmd, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "cache: [do] failed on cmd:%s, args:%v", cmd, args)
	}

	return reply, nil
}

func (r *redisPool) Close() error {
	return errors.Wrapf(r.pool.Close(), "cache: [Close] redis pool close failed")
}

// IntDo is a wrapper returns int type result.
func (r *redisPool) IntDo(cmd string, args ...interface{}) (int, error) {
	return redis.Int(r.do(cmd, args...))
}

// StringDo is a wrapper returns string type result.
func (r *redisPool) StringDo(cmd string, args ...interface{}) (string, error) {
	return redis.String(r.do(cmd, args...))
}

// StringsDo is a wrapper returns string[] type result.
func (r *redisPool) StringsDo(cmd string, args ...interface{}) ([]string, error) {
	return redis.Strings(r.do(cmd, args...))
}

// BoolDo is a wrapper returns boolean type result.
func (r *redisPool) BoolDo(cmd string, args ...interface{}) (bool, error) {
	return redis.Bool(r.do(cmd, args...))
}

// Float64Do is a wrapper returns float64 type result.
func (r *redisPool) Float64Do(cmd string, args ...interface{}) (float64, error) {
	return redis.Float64(r.do(cmd, args...))
}
