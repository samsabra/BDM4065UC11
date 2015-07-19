package main

import (
	"fmt"

	"github.com/pocke/BDM4065UC11"
)

func main() {
	c, err := bdm.New("/dev/ttyUSB0", 9600)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	v, err := c.PlatformVersion()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
