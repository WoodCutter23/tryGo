package main

import (
	"github.com/subosito/gotenv"
	"net/http"
	"netControllers/rest/getSchedule"
	"netControllers/rest/getStatistic"
	"netControllers/rest/setSchedule"
	"netControllers/rest/successTakeSchedule"
	"netControllers/rest/setProgramm"
	"netControllers/rest/getProgramm"
	"netControllers/rest/successTakeProgramm"
	"netControllers/rest/setStatistic"
	"netControllers/rest/getActualInfo"
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
	setProgramm.SetProgramm(dbConnect)
	getProgramm.GetProgramm(dbConnect)
	takeProgramm.SuccessGetProgramm(dbConnect)
	setStatistic.SetStatistic(dbConnect)
	getStatistic.GetStat(dbConnect)
	getActualInfo.GetInfo(dbConnect)
	_ = http.ListenAndServe("localhost:80", nil)
}



