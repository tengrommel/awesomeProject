package main

import (
	"encoding/xml"
	"io/ioutil"
)

type Campus struct {
	XMLName  xml.Name `xml:"campus"`
	Name     string   `xml:"name,attr"`
	Comment  string   `xml:",comment"`
	Building Building
}

type Building struct {
	XMLName xml.Name `xml:"building"`
	Name    string   `xml:"name,attr"`
	Comment string   `xml:",comment"`
	Device  Device
}

type Device struct {
	XMLName xml.Name `xml:"device"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

func main() {
	camp := &Campus{Name: "campus1",
		Comment:  "building-comment",
		Building: Building{Name: "bldg1", Comment: "device-comment", Device: Device{Name: "rtr1", Type: "router"}}}
	indt, _ := xml.MarshalIndent(camp, "", " ")
	_ = ioutil.WriteFile("file.xml", indt, 0644)
}
