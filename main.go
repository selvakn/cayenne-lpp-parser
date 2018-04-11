package main

import "github.com/selvakn/go-cayenne-lib/cayennelpp"
import "./internal"
import (
	"fmt"
	"net/http"
	"bytes"
	"log"
)
import (
	"encoding/base64"
	"io/ioutil"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
	payload, err := base64.StdEncoding.DecodeString(string(body))

	//payload := "AQD/AgFkAwIVSgQD6rYFZQH0BmYyB2f/ZAhooAlx/lgADwaCCnMp7wuGAWMCMf5mDIgH/YcAvvUACGoNdBVKDnUA8A92FXwQg04g"

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
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
