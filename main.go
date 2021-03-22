// Working with json
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	var err error
	// This 'type' was created by using http://json2struct.mervine.net
	// https://github.com/jmervine/gojson-http
	type Forecast struct {
		ApprovedTime string `json:"approvedTime"`
		Geometry     struct {
			Coordinates [][]float64 `json:"coordinates"`
			Type        string      `json:"type"`
		} `json:"geometry"`
		ReferenceTime string `json:"referenceTime"`
		TimeSeries    []struct {
			Parameters []struct {
				Level     int64     `json:"level"`
				LevelType string    `json:"levelType"`
				Name      string    `json:"name"`
				Unit      string    `json:"unit"`
				Values    []float64 `json:"values"`
			} `json:"parameters"`
			ValidTime string `json:"validTime"`
		} `json:"timeSeries"`
	}

	urlJson := "https://opendata-download-metfcst.smhi.se/api/category/pmp3g/version/2/geotype/point/lon/16.158/lat/58.5812/data.json"
	// Instead of wasting SMHIs bandwidth
	urlJson = "https://raw.githubusercontent.com/pmopmo/workingWithJson/master/data.json"

	var resp *http.Response
	resp, err = http.Get(urlJson)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var bytes []byte
	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	myJsonString := string(bytes)
	fmt.Println(myJsonString)

	var forecast Forecast
	json.Unmarshal([]byte(myJsonString), &forecast)

	fmt.Printf("ApprovedTime: %s, lat: %f, long: %f, Type %s",
		forecast.ApprovedTime,
		forecast.Geometry.Coordinates[0][0],
		forecast.Geometry.Coordinates[0][1],
		forecast.Geometry.Type)

} // EO main
