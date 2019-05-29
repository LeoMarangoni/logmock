package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var level = os.Getenv("LOGMOCK_LEVEL")
var l = getLevel()

func getLevel() int {
	x := 2
	if level == "TRACE" {
		x = 0
	} else if level == "DEBUG" {
		x = 1
	} else if level == "INFO" {
		x = 2
	} else if level == "WARN" {
		x = 3
	} else if level == "ERROR" {
		x = 4
	}
	return x
}

var levels = [5]io.Writer{
	os.Stdout,
	os.Stdout,
	os.Stdout,
	os.Stdout,
	os.Stderr}

var (
	Trace *log.Logger
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func init() {
	for i := 0; i < l; i++ {
		levels[i] = ioutil.Discard
	}
	Trace = log.New(levels[0],
		"[TRACE] ",
		log.Ldate|log.Ltime)

	Debug = log.New(levels[1],
		"[DEBUG] ",
		log.Ldate|log.Ltime)

	Info = log.New(levels[2],
		"[INFO] ",
		log.Ldate|log.Ltime)

	Warn = log.New(levels[3],
		"[WARN] ",
		log.Ldate|log.Ltime)

	Error = log.New(levels[4],
		"[ERROR] ",
		log.Ldate|log.Ltime)
}
