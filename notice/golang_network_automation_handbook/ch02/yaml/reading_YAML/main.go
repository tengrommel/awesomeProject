package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Host  string `yaml:"host"`
	Ports []int  `yaml:"ports"`
}

func main() {
	data, _ := ioutil.ReadFile("file.yaml")
	conf := Config{}
	_ = yaml.Unmarshal(data, &conf)
	fmt.Printf("%s:%d\n", conf.Host, conf.Ports)
}
