package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
“<campus name="campus1">
    <!-- building-comment -->
    <building name="bldg1">
        <!-- device-comment -->
        <device type="router" name="rtr1"></device>
    </building>
</campus>”
*/

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
	data, _ := ioutil.ReadFile("file.xml")
	camp := &Campus{}
	_ = xml.Unmarshal([]byte(data), &camp)
	f1 := fmt.Println
	f1("Campus Name:", camp.Name)
	f1("Building Name:", camp.Building.Name)
	f1("Building Comment:", camp.Comment)
	f1("Device Comment:", camp.Building.Comment)
	f1("Device Name:", camp.Building.Device.Name)
	f1("Device Type:", camp.Building.Device.Type)
}
