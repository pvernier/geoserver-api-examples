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
	const gs_url = "http://localhost:8080/geoserver/rest"
	const workspace = "test"

	// The name of the image mosaic that we will create
	const layer = "sst-go"

	client := &http.Client{}

	// 1) Create the layer
	fileobj, err := os.Open("init.zip")
	if err != nil {
		panic(err)
	}
	defer fileobj.Close()

	var url = gs_url + "/workspaces/" + workspace + "/coveragestores/" + layer + "/file.imagemosaic"

	request, err := http.NewRequest("PUT", url, fileobj)
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(user, pwd)

	request.Header.Set("Content-Type", "application/zip")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

	// 2) Enable the time dimension
	const xml = `
	<coverage>
	<enabled>true</enabled>
	<metadata><entry key="time">
	<dimensionInfo>
	<enabled>true</enabled>
	<presentation>LIST</presentation>
	<units>ISO8601</units><defaultValue/>
	</dimensionInfo>
	</entry></metadata>
	</coverage>`
	dataXML := bytes.NewBufferString(xml)

	request, err = http.NewRequest("PUT", gs_url+"/workspaces/"+workspace+"/coveragestores/"+layer+"/coverages/"+layer, dataXML)

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

	// 3) Add new granules
	fileobj, err = os.Open("granules.zip")
	if err != nil {
		panic(err)
	}
	defer fileobj.Close()

	url = gs_url + "/workspaces/" + workspace + "/coveragestores/" + layer + "/file.imagemosaic?recalculate=nativebbox,latlonbbox"

	request, err = http.NewRequest("POST", url, fileobj)
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(user, pwd)

	request.Header.Set("Content-Type", "application/zip")

	response, err = client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.Status)

}
