package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/*
The yaml.v2 package supports a Marshal() function that is used to serialized values into YAML format.
While structs are typically used to support the values, maps and pointers are typically used to support the values,
maps and pointers are accepted as well.
*/

type Config struct {
	Host  string `yaml:"host"`
	Ports []int  `yaml:"ports"`
}

func main() {
	conf := Config{
		Host:  "localhost",
		Ports: []int{80, 443},
	}
	data, _ := yaml.Marshal(&conf)
	_ = ioutil.WriteFile("file.yaml", data, 0644)
}
