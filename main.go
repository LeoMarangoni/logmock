package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func getEnv(envvar string, fallback string) string {
	x := os.Getenv(envvar)
	if x == "" {
		return fallback
	} else {
		return x
	}
}
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

func serveLogs(ms int64, env string, level string) {
	interval := time.Duration(ms) * time.Millisecond
	logger := zap.NewExample().Sugar()

	logger.Infof("Starting application")
	time.Sleep(interval)
	logger.Infof("Environment defined to %s", env)
	time.Sleep(interval)
	logger.Infof("Log interval defined to %s", interval)
	time.Sleep(interval)
	logger.Infof("Log level defined to %s", level)

	for {
		randomLog(logger, level, randomMsg(), randomErr())
		time.Sleep(interval)
	}

}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := zap.NewExample().Sugar()

		logger.Infow("Request OK", "Method", r.Method, "URI", r.RequestURI, "RemoteAddr", r.RemoteAddr, "Host", r.Host, "Agent", r.UserAgent())
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := getEnv("LOGMOCK_PORT", "8321")
	env := getEnv("LOGMOCK_ENV", "dev")

	interval, err := strconv.ParseInt(os.Getenv("LOGMOCK_INTERVAL"), 10, 0)
	if err != nil {
		interval = 500
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

	go serveLogs(interval, env, level)

	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld)
	router.HandleFunc("/health", HealthCheckHandler)
	router.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
