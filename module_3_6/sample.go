package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Entry struct {
	GlobalId       int64  `json: "global_id"`
	SystemObjectID string `json: "system_object_id"`
	SignatureDate  string `json: "signature_date"`
	Razdel         string `json: "Razdel"`
	Kod            string `json: "Kod"`
	Name           string `json: "Name"`
	Idx            string `json: "Idx"`
}

type OKVED struct {
	Entries []Entry
}

func main() {
	var entry Entry
	filepath := "sample.json"
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := io.ReadAll(f)
	// fmt.Printf("%s", bytes)
	err = json.Unmarshal(bytes, &entry)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(entry)

}
