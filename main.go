package main

import (
	"math/rand"
	"os"
	"time"

	"go.uber.org/zap"
)

func randomLog(logger *zap.SugaredLogger, level string, msg string, err err) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := 3
	switch level {
	case "DEBUG":
		n = 4
	case "INFO":
		n = 3
	case "WARN":
		n = 2
	case "ERROR":
		n = 1
	}

	switch x := r1.Intn(n); x {
	case 3:
		logger.Debugw(msg)
	case 2:
		logger.Infow(msg)
	case 1:
		logger.Warnw(msg)
	case 0:
		logger.Errorw(msg, "error", err.name, "description", err.reason)
	}
}

func randomMsg() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	x := r1.Intn(len(msgs))
	return msgs[x]
}

func randomErr() err {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	x := r1.Intn(len(errs))
	return errs[x]
}

func main() {
	interval := time.Duration(500) * time.Millisecond
	logger := zap.NewExample().Sugar()

	logger.Infof("Starting application")
	time.Sleep(interval * 2)

	env := os.Getenv("LOGMOCK_ENV")
	if "LOGMOCK_ENV" == "" {
		env = "dev"
	}
	level := os.Getenv("LOGMOCK_LOG_LEVEL")
	if level == "" {
		switch env {
		case "dev":
			level = "DEBUG"
		case "prd":
			level = "INFO"
		}

	}

	logger.Infof("Environment defined to %s", env)
	logger.Infof("Log interval defined to %s miliseconds", interval)
	logger.Infof("Log level defined to %s", level)

	for {
		randomLog(logger, level, randomMsg(), randomErr())
		time.Sleep(interval)
	}

}
