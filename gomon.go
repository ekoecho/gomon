package main

import (
	"fmt"
	//"os"
	//"bytes"
	"io/ioutil"
	"strings"
)

type Loadavg struct {
	one_min, five_min, ten_min string
}

func getLoadAvg() Loadavg {
	load, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		panic(err)
	}

	l := Loadavg{}
	l.one_min = strings.Split(string(load), " ")[0]
	l.five_min = strings.Split(string(load), " ")[1]
	l.ten_min = strings.Split(string(load), " ")[2]

	return l

}

func main() {
	load := getLoadAvg()

	fmt.Println(load.one_min, load.five_min, load.ten_min)

}
