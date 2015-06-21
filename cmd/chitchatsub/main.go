// sub.go
package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os/exec"

	. "github.com/scottjbarr/chitchat"
)

func say(message string, voice string) {
	args := make([]string, 1)

	if len(voice) > 0 {
		args = append(args, "-v", voice)
	}

	args = append(args, message)

	exec.Command("/usr/bin/say", args...).Run()
}

func subscribe(config *Config) {
	// connect to the redis server
	c := Connect(&config.RedisConfig)

	defer c.Close()

	psc := redis.PubSubConn{c}
	psc.Subscribe(config.Key)

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			log.Printf("#%s : message = %s\n", v.Channel, v.Data)
			say(string(v.Data), config.Voice)

		case redis.Subscription:
			log.Printf("#%s : %s %d\n", v.Channel, v.Kind, v.Count)

		case error:
			log.Fatal(v)
		}
	}
}

func main() {
	config := ParseConfig()
	subscribe(config)
}
