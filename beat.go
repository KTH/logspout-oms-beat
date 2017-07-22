package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

func getenvOrDefault(key string, defaultVal string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return defaultVal
}

func main() {
	level, err := strconv.Atoi(os.Getenv("LOGSPOUT_BEAT_LEVEL"))
	if err != nil {
		panic(err)
	}

	sleepTime, err := time.ParseDuration(os.Getenv("LOGSPOUT_BEAT_TIME"))
	if err != nil {
		panic(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	for {
		msg := bunyanMessage{
			V:        0,
			Level:    level,
			Name:     "beat",
			Hostname: hostname,
			Pid:      os.Getpid(),
			Time:     time.Now().Format("2006-01-02T15:04:05.999Z"),
			Msg:      "beat",
		}

		out, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(out))

		time.Sleep(sleepTime)
	}
}

type bunyanMessage struct {
	V        int    `json:"v"`
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Pid      int    `json:"pid"`
	Time     string `json:"time"`
	Msg      string `json:"msg"`
}
