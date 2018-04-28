package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

type Etls struct {
	Etls []struct {
		ID string `json:"id"`
		Filename string `json:"filename"`
		Incomingfilepath string `json:"incomingfilepath"`
		Intermediatefilepath string `json:"intermediatefilepath"`
		Destinationpath string `json:"destinationpath"`
	} `json:"etls"`
}

func main() {
	jsonFile, err := os.Open("/home/navins/go/src/github.com/navinds25/grpcGoExpts/config/config.json")
	if err != nil {
		fmt.Println("Error while reading", err)
		panic(err)
	}
	defer jsonFile.Close()
	rawdata, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var confst Etls
	json.Unmarshal(rawdata, &confst)

	fmt.Println(confst)
}

