package client

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/prashantcfc/gocowin/log"
	"io/ioutil"
	"net/http"
)

func Client(url, districtId, date, stateId string) (map[string]interface{}, error) {

    log := log.GetLoggerWithName("Client")
	client := &http.Client {
	}
	method := "GET"
	if stateId != "" {
		url = url + "/" + stateId
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Error(err, "%V", err)
		return nil, err
	}
	id := uuid.New()
	req.Header.Add("User-Agent", id.String())

	q := req.URL.Query()
	if districtId != "" {
		q.Add("district_id", districtId)
	}
	if date != "" {
		q.Add("date", date)
	}

	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Error(err,"%V",err)
		return nil, err
	}
	defer res.Body.Close()


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err,"%V",err)
		return nil, err
	}

	var raw map[string]interface{}
	err = json.Unmarshal(body, &raw)
	if err != nil {
		log.Error(err, "Error while unmarshalling..")
	}

	return raw, nil

}

