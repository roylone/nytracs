package msg

import (
	_ "encoding/json"
	"encoding/xml"
	_ "io/ioutil"
	_ "log"
)

/*
<?xml version="1.0" encoding="UTF-8"?>
<resources>
	<string name="VideoLoading">Loading video…</string>
	<string name="ApplicationName">what</string>
</resources>
*/
type RLGG struct {
	XMLName  xml.Name `xml:"resources"`
	InnerXml []Inner  `xml:"string"`
}

type Inner struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:"name,attr"`
	InnerText  string   `xml:",innerxml"`
}
