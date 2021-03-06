//
// Copyright © 2017 Kungliga Tekniska högskolan
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the “Software”), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

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
