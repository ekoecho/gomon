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

type MemInfo struct {
	totalMem string
	freeMem  string
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

func getMemInfo() MemInfo {

	mem, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		panic(err)
	}

	m := MemInfo{}
	m.totalMem = strings.Split(string(mem), "\n")[0]
	m.freeMem = strings.Split(string(mem), "\n")[1]

	return m
}

func main() {
	load := getLoadAvg()
	mem := getMemInfo()

	fmt.Println(load.one_min, load.five_min, load.ten_min)
	fmt.Println(mem.totalMem)
	fmt.Println(mem.freeMem)

}
