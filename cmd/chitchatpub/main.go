// pub.go
package main

import (
	"flag"
	"strings"

	. "github.com/scottjbarr/chitchat"
)

func publish(config *Config, phrase string) {
	c := Connect(&config.RedisConfig)

	defer c.Close()

	c.Send("PUBLISH", config.Key, phrase)
	c.Flush()
}

func main() {
	config := ParseConfig()
	phrase := strings.Join(flag.Args(), " ")

	publish(config, phrase)
}
