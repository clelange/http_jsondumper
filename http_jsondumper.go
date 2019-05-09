// package httpjsondumper
package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type jsonStruct struct {
	Samplename       string      `json:"samplename"`
	Processed        float32     `json:"processed"`
	Mass4mu8TeVlow   [37]float32 `json:"mass4mu_8TeV_low"`
	Mass2mu2e8TeVlow [37]float32 `json:"mass2mu2e_8TeV_low"`
	Mass4e8TeVlow    [37]float32 `json:"mass4e_8TeV_low"`
}

func hash(item jsonStruct) string {
	arrBytes := []byte{}
	jsonBytes, _ := json.Marshal(item)
	arrBytes = append(arrBytes, jsonBytes...)
	var md5Hash = md5.Sum(arrBytes)
	return hex.EncodeToString(md5Hash[:])
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d jsonStruct

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Received invalid format.")
		return
	}
	// fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Samplename))
	var fileHash = hash(d)
	var fileName = d.Samplename + "-" + fileHash + ".json"
	// fmt.Fprintf(w, "Writing to %s", fileName)
	fmt.Println("Writing to " + fileName)
	// file, _ := json.MarshalIndent(d, "", " ")
	file, _ := json.Marshal(d)

	_ = ioutil.WriteFile(fileName, file, 0644)
}

// code below not needed for Google Cloud Function
func main() {
	http.HandleFunc("/", HelloHTTP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
