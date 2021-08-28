package util

import (
	"encoding/json"
	"log"
)

func LogJson(i interface{}) {
	bs, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	log.Println(string(bs))
}
