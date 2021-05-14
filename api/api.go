package api

import (
	"github.com/prashantcfc/gocowin/client"
	"github.com/prashantcfc/gocowin/constants"
	"github.com/prashantcfc/gocowin/log"
)

func GetSessionsbyDistrictIdAndDate(districtId, date string) (map[string]interface{}, error) {

	log := log.GetLoggerWithName("getSessionsbyDistrictIdAndDate()")

	res, err := client.Client(constants.GetSessionsByDistrictIdAndDate, districtId, date, "")
	if err != nil {
		log.Error(err, "Failed to get sessions by districtId and date")
		return nil, err
	}
	return res, nil
}

func GetSessionsbyPinAndDate(pin, date string) (map[string]interface{}, error) {

	log := log.GetLoggerWithName("GetSessionsbyPinAndDate()")

	res, err := client.Client(constants.GetSessionsByDistrictIdAndDate, pin, date, "")
	if err != nil {
		log.Error(err, "Failed to get sessions by pin and date")
		return nil, err
	}
	return res, nil
}

func GetStateCodes() (map[string]interface{}, error) {
	log := log.GetLoggerWithName("GetStateCodes()")

	res, err := client.Client(constants.StatesList, "","", "")
	if err != nil {
		log.Error(err, "Failed to get list of state codes")
		return nil, err
	}
	return res, nil
}

func GetDistrictCodesInState(stateId string) (map[string]interface{}, error) {
	log := log.GetLoggerWithName("GetDistrictCodesInState()")

	res, err := client.Client(constants.DistrictList, "","", stateId)
	if err != nil {
		log.Error(err, "Failed to get list of district codes by state")
		return nil, err
	}
	return res, nil
}




