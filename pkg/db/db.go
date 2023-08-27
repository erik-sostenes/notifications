package db

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/notifications-api/pkg/common"
	"github.com/redis/go-redis/v9"
)

// Type represents an uint for the type of DataBase
type Type uint

const (
	// SQL represents MySQL database
	SQL Type = iota
	// NoSQL represents MongoDB database
	NoSQL
)

// Configuration represents the settings of the type of storage
type Configuration struct {
	// Type defines the type of storage to be used.
	Type
	Driver   string
	Host     string
	Port     string
	User     string
	Database string
	Password string
}

// NewRedisDBConfiguration returns an instance of Configuration with all the settings
// to make the connection to the database
func NewRedisDBConfiguration() Configuration {
	return Configuration{
		Type: NoSQL,
		Host: common.GetEnv("REDIS_HOST"),
		Port: common.GetEnv("REDIS_PORT"),
	}
}

// NewRedisClient method that will connect a redis client and returns an instance of redis.Client
func NewRedisClient(config Configuration) (*redis.Client, error) {
	switch config.Type {
	case NoSQL:
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprint(config.Host, ":", config.Port),
			Password: config.Password,
			DB:       0,
		})

		return client, client.Ping(context.Background()).Err()
	default:
		panic(fmt.Sprintf("%T type is not supported", config.Type))
	}
}

// NewRedisDataBase method that returns an instance of redis.Client
// if an error occurs a panic will be launched
func NewRedisDataBase(config Configuration) (db *redis.Client) {
	db, err := NewRedisClient(config)
	if err != nil {
		panic(err)
	}
	return
}
