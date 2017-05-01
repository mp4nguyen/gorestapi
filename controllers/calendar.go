package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"bitbucket.org/restapi/models"
)

var calendarsModel = new(models.CalendarModel)

type CalendarController struct{}

type GetCalendarParams struct {
	Id        int    `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	MaxPeriod int    `json:"maxPeriod"`
}

type tests struct {
	Tests []test `json:"tests"`
}
type test struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type returnCal struct {
	Dates []*returnDate `json:"dates"`
}

type returnDate struct {
	Date         string       `json:"date"`
	FormatedDate string       `json:"formatedDate"`
	Doctors      []*calDoctor `json:"doctors"`
}

type buildCal struct {
	dates map[string]*calDate `json:"dates"`
}

type calDate struct {
	doctors map[int]*calDoctor `json:"doctors"`
}

type calDoctor struct {
	DoctorID     int               `json:"doctorID"`
	DoctorName   string            `json:"doctorName"`
	Slots        []models.Calendar `json:"-"`
	SlotSegments []*slotSegment    `json:"slotSegments"`
}

type slotSegment struct {
	Slots          []*models.Calendar `json:"-"`
	FirstSlotBlock *slotBlock         `json:"firstSlot"`
	LastSlotBlock  *slotBlock         `json:"LastSlot"`
}

type slotBlock struct {
	FirstSlot models.Calendar   `json:"slot"`
	Slots     []models.Calendar `json:"followingSlots"`
}

func buildDoctorCal(cal models.Calendar, existingCal *returnDate) {
	prevDoctor := existingCal.Doctors[len(existingCal.Doctors)-1]
	if prevDoctor.DoctorID != cal.DoctorId {
		existingCal.Doctors = append(existingCal.Doctors, &calDoctor{cal.DoctorId, cal.DoctorName, []models.Calendar{cal}, []*slotSegment{}})
	} else {
		prevDoctor.Slots = append(prevDoctor.Slots, cal)
		//existingCal.doctors[cal.DoctorId] = doctorMap
	}
}

// func buildDoctorCal(cal models.Calendar, existingCal *calDate) {
//
// 	doctorMap, ok := existingCal.doctors[cal.DoctorId]
// 	if ok {
// 		doctorMap.Slots = append(doctorMap.Slots, cal)
// 		existingCal.doctors[cal.DoctorId] = doctorMap
// 	} else {
// 		existingCal.doctors[cal.DoctorId] = &calDoctor{cal.DoctorId, cal.DoctorName, []models.Calendar{cal}, []*slotSegment{}}
// 	}
//
// }

func (u CalendarController) GetCalendar(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	calParams := GetCalendarParams{}
	dec := json.NewDecoder(r.Body)
	//fmt.Println(dec)
	//fmt.Println(r.FormValue("id"))
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

	//buildCals := buildCal{make(map[string]*calDate)}
	buildCals := returnCal{}
	//buildCals.dates = (make(map[string]calDate))
	cals, err := calendarsModel.GetCalendar(calParams.Id, calParams.From, calParams.To)

	log.Printf("sql duration = %s", time.Since(start))
	start = time.Now()

	//group cals by date,doctor
	// for _, cal := range cals.Calendars {
	// 	existingCalendar, ok := buildCals.dates[cal.CalendarDate]
	// 	if ok {
	// 		buildDoctorCal(cal, existingCalendar)
	// 	} else {
	// 		doctorMap := make(map[int]*calDoctor)
	// 		doctorMap[cal.DoctorId] = &calDoctor{cal.DoctorId, cal.DoctorName, []models.Calendar{cal}, []*slotSegment{}}
	// 		//fmt.Println("buildCals = ", buildCals, " doctorMap = ", doctorMap)
	// 		buildCals.dates[cal.CalendarDate] = &calDate{doctorMap}
	// 	}
	// }
	prevCal := models.Calendar{}
	currentDateIndex := -1
	for _, cal := range cals.Calendars {
		//fmt.Println("			- date = ", cal)
		if prevCal.CalendarDate != cal.CalendarDate {
			rDate := returnDate{}
			rDate.Date = cal.CalendarDate
			rDate.FormatedDate = cal.CalendarDate
			//lpCalDoctor := &calDoctor{cal.DoctorId, cal.DoctorName, []models.Calendar{cal}, []*slotSegment{}}
			//rDate.doctors = []*calDoctor{}
			rDate.Doctors = append(rDate.Doctors, &calDoctor{cal.DoctorId, cal.DoctorName, []models.Calendar{cal}, []*slotSegment{}})
			buildCals.Dates = append(buildCals.Dates, &rDate)
			//fmt.Println("			-----> date = ", rDate)
			currentDateIndex++
		} else {
			buildDoctorCal(cal, buildCals.Dates[currentDateIndex])
		}

		prevCal = cal
	}

	//fmt.Println("buildCals = ", buildCals)

	// for key, date := range buildCals.dates {
	// 	fmt.Println("			- date = ", key)
	// 	for doctorId, doctor := range date.doctors {
	// 		fmt.Println("				+ doctor = ", doctorId)
	// 		for _, slot := range doctor.Slots {
	// 			fmt.Println("					* slot = ", slot)
	//
	// 		}
	// 	}
	// }

	//make slotSegments for each doctor
	for _, date := range buildCals.Dates {
		//fmt.Println("			- date = ", key)
		for _, doctor := range date.Doctors {

			if len(doctor.SlotSegments) == 0 {
				doctor.SlotSegments = append(doctor.SlotSegments, &slotSegment{})
			}
			currentSegment := 0
			//fmt.Println("					+ doctorId = ", doctorId)
			prevSlot := models.Calendar{}
			for slotId, slot := range doctor.Slots {
				//fmt.Println("							* slotId = ", slotId, " slot =", slot)
				if slotId == 0 {
					prevSlot = slot
					assignSlotIntoSegment(doctor, currentSegment, slot)
				} else {
					//check if the prevSlot link to slot ?
					//If not, create segment for slot
					if prevSlot.CalendarToTime == slot.CalendarFromTime {
						assignSlotIntoSegment(doctor, currentSegment, slot)
					} else {
						doctor.SlotSegments = append(doctor.SlotSegments, &slotSegment{})
						currentSegment++

						assignSlotIntoSegment(doctor, currentSegment, slot)
					}
				}
				prevSlot = slot
			}
		}
	}

	//build first and last slot

	for _, date := range buildCals.Dates {
		for _, doctor := range date.Doctors {
			for _, segment := range doctor.SlotSegments {
				//segment.FirstSlotBlock = &slotBlock{}
				//segment.LastSlotBlock = &slotBlock{}

				// duration := segment.FirstSlotBlock.FirstSlot.CalendarFromTime.Sub(segment.FirstSlotBlock.FirstSlot.CalendarToTime)
				// fmt.Println("Duration = ", duration)
				// fmt.Println("Duration = ", duration.Seconds())
				// fmt.Println("Duration = ", duration.Minutes())

				lengthOfSlots := len(segment.Slots)
				//Make sure there are slots inside
				if lengthOfSlots > 0 {
					duration := segment.Slots[lengthOfSlots-1].CalendarToTime.Sub(segment.Slots[0].CalendarFromTime)
					durationOfSlot := segment.Slots[0].CalendarToTime.Sub(segment.Slots[0].CalendarFromTime)
					noOfSlotInABlock := float64(calParams.MaxPeriod) / durationOfSlot.Minutes()
					//check if there is enough time for blockslot
					//mean that if MaxPeriod = 60 minutes => need at least 4 slots
					if duration.Minutes() >= float64(calParams.MaxPeriod) {
						isLastSlotBlock := duration.Minutes() > float64(calParams.MaxPeriod)
						for slotIndex, slot := range segment.Slots {
							if slotIndex == 0 {
								segment.FirstSlotBlock = &slotBlock{}
								segment.FirstSlotBlock.FirstSlot = *slot
								segment.FirstSlotBlock.Slots = append(segment.FirstSlotBlock.Slots, *slot)
							} else {
								lpSlot := *slot
								duration := lpSlot.CalendarToTime.Sub(segment.FirstSlotBlock.FirstSlot.CalendarFromTime)
								if duration.Minutes() <= float64(calParams.MaxPeriod) {
									//fmt.Println("Duration = ", duration.Minutes(), " (minutes); slot = ", *slot)
									segment.FirstSlotBlock.Slots = append(segment.FirstSlotBlock.Slots, *slot)
								}
							}
							//build LastSlotBlock
							if slotIndex >= (lengthOfSlots-int(noOfSlotInABlock)) && isLastSlotBlock {
								if slotIndex == (lengthOfSlots - int(noOfSlotInABlock)) {
									segment.LastSlotBlock = &slotBlock{}
									segment.LastSlotBlock.FirstSlot = *slot
									segment.LastSlotBlock.Slots = append(segment.LastSlotBlock.Slots, *slot)
								} else {
									segment.LastSlotBlock.Slots = append(segment.LastSlotBlock.Slots, *slot)
								}
							}
						}
					}
				}
			}
		}
	}

	//for debug
	//
	// for key, date := range buildCals.dates {
	// 	fmt.Println("			- date = ", key, " : ", date.date)
	// 	for doctorId, doctor := range date.doctors {
	// 		fmt.Println("					+ doctorId = ", doctorId, " : ", doctor.doctorName)
	// 		//fmt.Println("							# Slots ")
	// 		// for slotId, slot := range doctor.Slots {
	// 		// 	fmt.Println("									* slotId = ", slotId, " slot =", slot)
	// 		// }
	// 		fmt.Println("							# Segments ")
	// 		for segmentIndex, segment := range doctor.SlotSegments {
	// 			//fmt.Println("									* segmentIndex = ", segmentIndex, " segment =", segment)
	// 			fmt.Println("							   # Segments #", segmentIndex)
	// 			fmt.Println("							   # firstSlotBlock #", segment.FirstSlotBlock.FirstSlot)
	// 			for slotIndex, slot := range segment.FirstSlotBlock.Slots {
	// 				fmt.Println("									* slotIndex = ", segmentIndex, ":", slotIndex, " slot =", slot)
	// 			}
	// 			fmt.Println("							   # LastSlotBlock #", segment.LastSlotBlock.FirstSlot)
	// 			for slotIndex, slot := range segment.LastSlotBlock.Slots {
	// 				fmt.Println("									* slotIndex = ", segmentIndex, ":", slotIndex, " slot =", slot)
	// 			}
	// 			fmt.Println("							   # allSlots #")
	// 			for slotIndex, slot := range segment.Slots {
	// 				fmt.Println("									* slotIndex = ", segmentIndex, ":", slotIndex, " slot =", slot)
	// 			}
	// 		}
	// 	}
	// }

	//log.Println("buildCals = ", buildCals)

	log.Printf("build cal duration = %s", time.Since(start))

	outputCals, err := json.Marshal(buildCals)
	if err != nil {
		fmt.Println("failed to convert to JSON ", err)
	}
	//fmt.Println(" JSON  = ", string(outputCals))
	fmt.Fprintln(w, string(outputCals))

}

func assignSlotIntoSegment(doctor *calDoctor, currentSegment int, slot models.Calendar) {
	segment := doctor.SlotSegments[currentSegment]
	segment.Slots = append(segment.Slots, &slot)
	doctor.SlotSegments[currentSegment] = segment
}
