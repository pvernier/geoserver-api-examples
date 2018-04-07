package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Geoserver parameters
	const user = "admin"
	const pwd = "geoserver"
	const gsURL = "http://localhost:8080/geoserver/rest"
	const workspace = "test"
	const style = "dem"
	const coverage = "dem-go"

	client := &http.Client{}

	// 1) Create the coverage
	// Will create a file dem-go.geotiff in data/test/dem-go/
	data, err := os.Open("dem.tiff")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	const url = gsURL + "/workspaces/" + workspace + "/coveragestores/" + coverage + "/file.geotiff"

	request, err := http.NewRequest("PUT", url, data)
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(user, pwd)

	request.Header.Set("Content-Type", "image/tiff")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

	// 2) Change the style
	const xml = "<layer><defaultStyle><name>" + style + "</name></defaultStyle></layer>"
	payload := bytes.NewBufferString(xml)

	request, err = http.NewRequest("PUT", gsURL+"/layers/"+coverage, payload)
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(user, pwd)

	request.Header.Set("Content-Type", "application/xml; charset=UTF-8")

	response, err = client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

}
