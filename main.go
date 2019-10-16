package main

import (
	"github.com/subosito/gotenv"
	"net/http"
	"netControllers/rest/getSchedule"
	"netControllers/rest/setSchedule"
	"netControllers/rest/successTakeSchedule"
	s "os"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	dbConnect:= s.Getenv("DB_FULL")
	getSchedule.GetSchedule(dbConnect)
	setSchedule.SetSchedule(dbConnect)
	TakeSchedule.SuccessGetSchedule(dbConnect)
	_ = http.ListenAndServe("localhost:80", nil)
}



