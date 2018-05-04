package main

import "./internal"
import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
)
import "github.com/selvakn/go-cayenne-lib/cayennelpp"


func parseLPP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
	payload, err := base64.StdEncoding.DecodeString(string(body))

	//payload := "AQD/AgFkAwIVSgQD6rYFZQH0BmYyB2f/ZAhooAlx/lgADwaCCnMp7wuGAWMCMf5mDIgH/YcAvvUACGoNdBVKDnUA8A92FXwQg3hpIg=="

	decoder := cayennelpp.NewDecoder(bytes.NewBuffer(payload))
	target := internal.NewTarget()

	err = decoder.DecodeUplink(target)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(target)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", parseLPP)
	log.Fatal(http.ListenAndServe(addr, nil))
}
