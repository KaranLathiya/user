package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	error_handling "user/error"
)

func ExternalURLCall(method string, url string, bodyDataRequest interface{}, bodyDataResponse interface{}) error {
	bodyDataByte, _ := json.MarshalIndent(bodyDataRequest, "", " ")
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyDataByte))
	if err != nil {
		return error_handling.InternalServerError
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode == 400 || res.StatusCode == 404 {
		return error_handling.InternalServerError
	}
	defer res.Body.Close()
	err = BodyRead(res.Body, bodyDataResponse)
	return err
}
