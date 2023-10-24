package mdl

/**
**/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Connect struct {
	Data     string
	DataByte []byte
	Error    string
	Time     time.Time
}

func (C *Connect) Conection(method string, url string, jsonData []byte) {
	var err error
	var resp *http.Request
	if method == "POST" {
		resp, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	} else {
		resp, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(resp)
	if error != nil {
		C.Data = error.Error()
	} else {

		body, _ := ioutil.ReadAll(response.Body)
		C.DataByte = body
		C.Data = string(body)
	}

}
