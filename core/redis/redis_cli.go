package redis

import "github.com/gomodule/redigo/redis"

type RedisCli struct {
	conn redis.Conn
}

var redisCliInstance *RedisCli = nil

func Connect() (conn *RedisCli) {
	if redisCliInstance == nil {
		redisCliInstance = &RedisCli{}

		var err error
		redisCliInstance.conn, err = redis.Dial("tcp", ":6379")
		if err != nil {
			panic(err)
		}
		if _, err := redisCliInstance.conn.Do("AUTH", "Britannica"); err != nil {
			redisCliInstance.conn.Close()
			panic(err)
		}
	}
	return redisCliInstance
}

func (r *RedisCli) SetValue(key, value string, expiration ...interface{}) error {
	_, err := r.conn.Do("SET", key, value)
	if err == nil && expiration != nil {
		r.conn.Do("EXPIRE", key, expiration[0])
	}
	return err
}

func (r *RedisCli) GetValue(key string) (interface{}, error) {
	return r.conn.Do("GET", key)
}
