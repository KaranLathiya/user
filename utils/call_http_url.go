package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	error_handling "user/error"
)

func CallHttpURL(method string, url string, bodyDataRequest interface{}, bodyDataResponse map[string]interface{}) (map[string]interface{}, error) {
	bodyDataByte, err := json.MarshalIndent(bodyDataRequest, "", " ")
	if err != nil {
		return nil,error_handling.MarshalError
	}
	var req *http.Request
	if method == http.MethodGet {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(bodyDataByte))
	}
	if err != nil {
		return nil, error_handling.InternalServerError
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode == http.StatusBadRequest || res.StatusCode == http.StatusNotFound ||  res.StatusCode == http.StatusInternalServerError {
		return nil, error_handling.InternalServerError
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resBody, &bodyDataResponse)
	if err != nil{
		return nil, error_handling.UnmarshalError
	}
	return bodyDataResponse, nil
}
