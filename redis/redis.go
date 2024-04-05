package redis

import (
	"github.com/andReyM228/lib/log"
	"github.com/go-redis/redis"
)

type cache struct {
	client *redis.Client
	log    log.Logger
}

type Config struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func NewCache(cfg Config, log log.Logger) Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return cache{
		client: client,
		log:    log,
	}
}

func (c cache) Set(key string, value interface{}) error {
	err := c.client.Set(key, value, 0).Err()
	if err != nil {
		c.log.Error(err.Error())
		return err
	}

	return nil
}

func (c cache) GetBytes(key string) ([]byte, error) {
	val, err := c.client.Get(key).Bytes()
	if err != nil {
		c.log.Error(err.Error())
		return nil, err
	}

	return val, nil
}

func (c cache) GetString(key string) (string, error) {
	val, err := c.client.Get(key).Result()
	if err != nil {
		c.log.Error(err.Error())
		return "", err
	}

	return val, nil
}
