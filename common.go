package chitchat

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type RedisConfig struct {
	Host string
	Port int
	DB   int
	Key  string
}

func (config *RedisConfig) ConnString() string {
	return fmt.Sprintf("%v:%v", config.Host, config.Port)
}

type Config struct {
	RedisConfig
	Voice string
}

func Connect(config *RedisConfig) redis.Conn {
	c, err := redis.Dial("tcp", config.ConnString())

	if err != nil {
		log.Fatal(err)
	}

	// switch to the correct database
	if _, err = c.Do("SELECT", config.DB); err != nil {
		log.Fatal(err)
	}

	return c
}

// Parse the config
func ParseConfig() *Config {
	// define the command line params
	host := flag.String("host", "127.0.0.1", "Redis host")
	port := flag.Int("port", 6379, "Redis port")
	db := flag.Int("db", 0, "Redis database")
	key := flag.String("key", "speech", "Key to broadcast to")

	voice := flag.String("voice", "", "Voice to use for speech")

	// parse it all
	flag.Parse()

	// grab redis config
	redisConfig := &RedisConfig{
		Host: *host,
		Port: *port,
		DB:   *db,
		Key:  *key,
	}

	// grab other config
	config := &Config{
		RedisConfig: *redisConfig,
		Voice:       *voice,
	}

	return config
}
