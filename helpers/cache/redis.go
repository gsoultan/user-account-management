package cache

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

type Redis interface {
	Ping() error
	Set(key string, val string) error
	Get(key string) (string, error)
	Pool() *redis.Pool
	LPush(key string, val string) (int, error)
	RPush(key string, val string) (int, error)
	LRange(key string, start int, end int) ([]byte, error)
	Exists(key string) (int, error)
}

type aRedis struct {
	pool *redis.Pool
}

func (r *aRedis) Exists(key string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("EXISTS", key))
}

func (r *aRedis) LRange(key string, start int, end int) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()
	rs, err := redis.Bytes(conn.Do("LRANGE", start, end))
	return rs, err
}

func (r *aRedis) RPush(key string, val string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	rs, err := redis.Int(conn.Do("RPUSH", key, val))
	return rs, err
}

func (r *aRedis) LPush(key string, val string) (int, error) {
	conn := r.pool.Get()
	defer conn.Close()
	rs, err := redis.Int(conn.Do("LPUSH", key, val))
	return rs, err
}

func (r *aRedis) Ping() error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("PING"))
	return err
}

func (r *aRedis) Set(key string, val string) error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, val)
	return err
}

func (r *aRedis) Get(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func (r *aRedis) Pool() *redis.Pool {
	return r.pool
}

func NewRedisInstance(network string, address string, port int) Redis {
	pool := &redis.Pool{}
	pool.MaxIdle = 80
	pool.MaxActive = 12000
	pool.Dial = func() (redis.Conn, error) {
		conn, err := redis.Dial(network, address+":"+strconv.Itoa(port))
		return conn, err
	}

	return &aRedis{pool: pool}
}
