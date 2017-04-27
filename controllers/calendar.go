package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"log"
)

type CalendarController struct{}

type GetCalendarParams struct{
	Id int `json:"id"`
	From string `json:"from"`
	To string `json:"to"`
	MaxPeriod int `json:"maxPeriod"`
}

func (u CalendarController) GetCalendar(w http.ResponseWriter, r *http.Request) {

	calParams := GetCalendarParams{}
	dec := json.NewDecoder(r.Body)
	for {

		if err := dec.Decode(&calParams); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	output, err := json.Marshal(calParams)
	log.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}


}
