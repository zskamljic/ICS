package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/zskamljic/ics"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("cal.ics")
	checkErr(err)

	calStr := string(data)
	calStr = strings.Replace(calStr, "\r", "", -1)

	cal, err := ics.NewCalendar(calStr)
	checkErr(err)

	json.NewEncoder(os.Stdout).Encode(cal)
}
