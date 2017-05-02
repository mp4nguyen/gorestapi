package calendarCtrl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"bitbucket.org/restapi/models/calendarMdl"
)

func GetCalendar(w http.ResponseWriter, r *http.Request) {

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
	cals, err := calendarMdl.GetCalendar(calParams.Id, calParams.From, calParams.To)

	log.Printf("sql duration = %s", time.Since(start))
	start = time.Now()

	prevCal := calendarMdl.Calendar{}
	currentDateIndex := -1
	for _, cal := range cals.Calendars {
		//fmt.Println("			- date = ", cal)
		if prevCal.CalendarDate != cal.CalendarDate {
			rDate := returnDate{}
			rDate.Date = cal.CalendarDate
			rDate.FormatedDate = cal.CalendarDate
			rDate.Doctors = append(rDate.Doctors, &calDoctor{cal.DoctorId, cal.DoctorName, []calendarMdl.Calendar{cal}, []*slotSegment{}})
			buildCals.Dates = append(buildCals.Dates, &rDate)
			currentDateIndex++
		} else {
			buildDoctorCal(cal, buildCals.Dates[currentDateIndex])
		}

		prevCal = cal
	}

	//make slotSegments for each doctor
	for _, date := range buildCals.Dates {
		//fmt.Println("			- date = ", key)
		for _, doctor := range date.Doctors {

			if len(doctor.SlotSegments) == 0 {
				doctor.SlotSegments = append(doctor.SlotSegments, &slotSegment{})
			}
			currentSegment := 0
			//fmt.Println("					+ doctorId = ", doctorId)
			prevSlot := calendarMdl.Calendar{}
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

	//build returnSlots
	for _, date := range buildCals.Dates {
		for _, doctor := range date.Doctors {
			for _, segment := range doctor.SlotSegments {
				if segment.FirstSlotBlock != nil {
					firstSLot := segment.FirstSlotBlock.FirstSlot
					slot := returnSlot{}
					slot.CalId = firstSLot.CalId
					slot.CalendarDate = firstSLot.CalendarDate
					slot.CalendarTime = firstSLot.CalendarTime
					slot.DoctorId = firstSLot.DoctorId
					slot.DoctorName = firstSLot.DoctorName
					slot.SiteId = firstSLot.SiteId
					slot.CalendarDate = firstSLot.CalendarDate
					for _, followingSlot := range segment.FirstSlotBlock.Slots {
						slot.FollowingSlots = append(slot.FollowingSlots, followingSlot)
					}
					date.Slots = append(date.Slots, &slot)
				}

				if segment.LastSlotBlock != nil {
					firstSLot := segment.LastSlotBlock.FirstSlot
					slot := returnSlot{}
					slot.CalId = firstSLot.CalId
					slot.CalendarDate = firstSLot.CalendarDate
					slot.CalendarTime = firstSLot.CalendarTime
					slot.DoctorId = firstSLot.DoctorId
					slot.DoctorName = firstSLot.DoctorName
					slot.SiteId = firstSLot.SiteId
					slot.CalendarDate = firstSLot.CalendarDate
					for _, followingSlot := range segment.LastSlotBlock.Slots {
						slot.FollowingSlots = append(slot.FollowingSlots, followingSlot)
					}
					date.Slots = append(date.Slots, &slot)
				}

			}
		}
		//sort slots in date object
		sort.Sort(ByCalendarTime(date.Slots))
	}

	//log.Println("buildCals = ", buildCals)

	log.Printf("build cal duration = %s", time.Since(start))

	outputCals, err := json.Marshal(buildCals)
	if err != nil {
		fmt.Println("failed to convert to JSON ", err)
	}
	//fmt.Println(" JSON  = ", string(outputCals))
	fmt.Fprintln(w, string(outputCals))

}

func buildDoctorCal(cal calendarMdl.Calendar, existingCal *returnDate) {
	prevDoctor := existingCal.Doctors[len(existingCal.Doctors)-1]
	if prevDoctor.DoctorID != cal.DoctorId {
		existingCal.Doctors = append(existingCal.Doctors, &calDoctor{cal.DoctorId, cal.DoctorName, []calendarMdl.Calendar{cal}, []*slotSegment{}})
	} else {
		prevDoctor.Slots = append(prevDoctor.Slots, cal)
		//existingCal.doctors[cal.DoctorId] = doctorMap
	}
}

func assignSlotIntoSegment(doctor *calDoctor, currentSegment int, slot calendarMdl.Calendar) {
	segment := doctor.SlotSegments[currentSegment]
	segment.Slots = append(segment.Slots, &slot)
	doctor.SlotSegments[currentSegment] = segment
}

func printBuildCal(buildCals returnCal) {
	//for debug

	for key, date := range buildCals.Dates {
		fmt.Println("			- date = ", key, " : ", date.Date)
		for doctorId, doctor := range date.Doctors {
			fmt.Println("					+ doctorId = ", doctorId, " : ", doctor.DoctorName)
			//fmt.Println("							# Slots ")
			// for slotId, slot := range doctor.Slots {
			// 	fmt.Println("									* slotId = ", slotId, " slot =", slot)
			// }
			fmt.Println("							# Segments ")
			for segmentIndex, segment := range doctor.SlotSegments {
				//fmt.Println("									* segmentIndex = ", segmentIndex, " segment =", segment)
				fmt.Println("							   # Segments #", segmentIndex)
				fmt.Println("							   # firstSlotBlock #", segment.FirstSlotBlock.FirstSlot)
				for slotIndex, slot := range segment.FirstSlotBlock.Slots {
					fmt.Println("									* slotIndex = ", segmentIndex, ":", slotIndex, " slot =", slot)
				}
				fmt.Println("							   # LastSlotBlock #", segment.LastSlotBlock.FirstSlot)
				for slotIndex, slot := range segment.LastSlotBlock.Slots {
					fmt.Println("									* slotIndex = ", segmentIndex, ":", slotIndex, " slot =", slot)
				}
				fmt.Println("							   # allSlots #")
				for slotIndex, slot := range segment.Slots {
					fmt.Println("									* slotIndex = ", segmentIndex, ":", slotIndex, " slot =", slot)
				}
			}
		}
	}
}
