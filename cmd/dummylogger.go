package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	conf := parseArgs()

	if conf.endAfterSeconds >= 0 {
		time.AfterFunc(time.Duration(conf.endAfterSeconds)*time.Millisecond, func() {
			os.Exit(0)
		})
	}

	if conf.failAfterSeconds >= 0 {
		time.AfterFunc(time.Duration(conf.failAfterSeconds)*time.Millisecond, func() {
			os.Exit(1)
		})
	}

	var logFile *os.File
	if len(conf.logToFilePath) > 0 {
		var err error
		if logFile, err = os.Create(conf.logToFilePath); err != nil {
			panic(err)
		}
		defer logFile.Close()
	}

	ticker := time.NewTicker(time.Duration(conf.logInterval) * time.Millisecond)
	for range ticker.C {
		now := time.Now().UnixNano()
		_, _ = fmt.Fprintf(os.Stdout, "[%d] stdout\n", now)
		_, _ = fmt.Fprintf(os.Stderr, "[%d] stderr\n", now)
		_, _ = fmt.Fprintf(logFile, "[%d] file\n", now)
	}
}

func parseArgs() *conf {
	var conf = &conf{}

	flag.IntVar(&conf.failAfterSeconds, "fail", -1, "number of seconds after program should fail with !=0")
	flag.IntVar(&conf.endAfterSeconds, "end", -1, "number of seconds after program should end with 0")
	flag.StringVar(&conf.logToFilePath, "file", "", "path of the file to log into")
	flag.IntVar(&conf.logInterval, "loginterval", 100, "log each <loginterval> ms")

	flag.Parse()

	fmt.Printf("%+v\n", *conf)

	return conf
}

type conf struct {
	failAfterSeconds int
	endAfterSeconds  int
	logInterval      int
	logToFilePath    string
}
