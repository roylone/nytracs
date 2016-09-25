package msg

import (
	"encoding/json"
	"encoding/xml"
	"log"
)

func ParseJson(content string, v interface{}) interface{} {
	if err := json.Unmarshal([]byte(content), &v); err != nil {
		log.Println("parse error!")
	}
	return v
}

func ParseXml(content string, v interface{}) interface{} {
	if err := xml.Unmarshal([]byte(content), &v); err != nil {
		log.Println("parse error!")
	}
	return v
}
