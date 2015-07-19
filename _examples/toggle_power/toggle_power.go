package main

import "github.com/pocke/BDM4065UC11"

func main() {
	c, err := bdm.New("/dev/ttyUSB0", 9600)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	if on, err := c.IsPowerOn(); err != nil {
		panic(err)
	} else if on {
		err := c.PowerOff()
		if err != nil {
			panic(err)
		}
	} else {
		err := c.PowerOn()
		if err != nil {
			panic(err)
		}
	}
}
