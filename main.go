package main

import (
	"flag"
	"github.com/prashantcfc/gocowin/api"
	"github.com/prashantcfc/gocowin/log"
)

func main() {

	log := log.GetLoggerWithName("gocowin")

	states := flag.Bool("states", false, "display list of state codes")
	stateId := flag.String("stateId", "", "display list of district IDs by given stateId")
	pin := flag.String("pin", "", "display list of available slots by pin and date")
	districtId := flag.String("districtId", "", "display list of available slots by distrctId and date")
	date := flag.String("date", "", "date for which slots need to be found")

	flag.Parse()

	// Check negative cases

	if *states != false && *stateId != "" {
		log.Info("Please specify either states or stateId. Not both!!!")
		return
	}

	if *pin != "" && *districtId != "" {
		log.Info("Please specify either pin or districtId. Not both!!")
		return
	}

	if *states {
		res, err := api.GetStateCodes()
		if err != nil {
			log.Error(err, "Error while getting states code list")
			return
		}
		log.Info("Successfully fetched available state codes.", "State Codes:", res)
		return
	}

	if *stateId != "" {
		res, err := api.GetDistrictCodesInState(*stateId)
		if err != nil {
			log.Error(err, "Error while getting district code list by state")
			return
		}
		log.Info("Successfully fetched available district codes of a state.", "District Codes:", res)
		return
	}

	if *pin != "" && *date != "" {
		res, err := api.GetSessionsbyPinAndDate(*pin, *date)
		if err != nil {
			log.Error(err, "Error while getting sessions by pin and date")
			return
		}
		obj := res["sessions"].([]interface{})
		if len(obj) > 0 {
			log.Info("!!!!!!!!!THERE ARE SLOTS AVAILABLE!!!!!!!! HURRY UP!!!!!!!!!!!!!")
			log.Info("Successfully fetched available sessions by Pin and Date.", "Sessions:", res)
			return
		}
		log.Info("No available slots found :-( ")
		return
	}

	if *districtId != "" && *date != "" {
		res, err := api.GetSessionsbyDistrictIdAndDate(*districtId, *date)
		if err != nil {
			log.Error(err, "Error while getting sessions by districtId and date")
			return
		}
		obj := res["sessions"].([]interface{})
		if len(obj) > 0 {
			log.Info("!!!!!!!!!THERE ARE SLOTS AVAILABLE!!!!!!!! HURRY UP!!!!!!!!!!!!!")
			log.Info("Successfully fetched available sessions by District and Date.", "Sessions:", res)
			return
		}
		log.Info("No available slot found :-( ")
		return
	}

	log.Info("Please give a valid input!!! Use --help option to see available options")

}
