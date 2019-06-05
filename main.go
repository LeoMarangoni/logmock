package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/leomarangoni/logmock/logger"
)

var levels = []string{
	"TRACE",
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
}

var msgs = []string{
	"Eat your foot.",
	"I hate the green flashing light.",
	"Hello. I have the urge to kill.",
	"Oh no! You’re going to speak again, aren’t you?",
	"DO NOT DISTURB, evil genius at work.",
	"I’m with stupid------àJ",
	"Rubber ducks are planning world domination!",
	"But my tree only hit the car in self-defence!",
	"I know kung fu and 50 other dangerous words.",
	"Did my sarcasm hurt your feels? Get over it.",
	"Love your enemies, it makes them angry.",
	"Fat kids are harder to kidnap.",
	"Shut up voices! Or I will poke you with y pen again!",
	"Save water, drink beer.",
	"Save a tree, eat a beaver.",
	"Get high, climb a tree.",
	"Save a horse, ride a cowboy.",
	"Don’t mess with me! I have a stick!",
	"Go away, evil Mr Scissors!",
	"Think of gingerbread men: are they delicious holiday treats or just another way for children to show off their cannibalism?",
}

var interval = GetInterval()

func GetInterval() time.Duration {
	interval, err := strconv.ParseInt(os.Getenv("LOGMOCK_INTERVAL"), 10, 0)
	if err != nil {
		interval = 500
		logger.Warn.Println("No valid interval detected, defining 500 miliseconds")
	}
	return time.Duration(interval) * time.Millisecond
}
func RandonLog() {
	level := levels[rand.Intn(len(levels))]
	msg := msgs[rand.Intn(len(msgs))]
	logger.Print(msg, level)

}

func ServeLogs() {
	for {
		RandonLog()
		time.Sleep(interval)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	logger.Info.Println("INDEX")
	json.NewEncoder(w).Encode("Hello World")
}
func main() {
	go ServeLogs()
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
