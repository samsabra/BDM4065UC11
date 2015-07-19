BDM4065UC11
==============

This is a RS-232C library for [BDM4065UC/11](http://www.philips.co.jp/c/pc-monitor/brilliance-led-backlit-lcd-display-bdm4065uc_11/prd/).

[![GoDoc](https://godoc.org/github.com/pocke/BDM4065UC11?status.svg)](https://godoc.org/github.com/pocke/BDM4065UC11)
[![Build Status](https://travis-ci.org/pocke/BDM4065UC11.svg?branch=master)](https://travis-ci.org/pocke/BDM4065UC11)
[![Coverage Status](https://coveralls.io/repos/pocke/BDM4065UC11/badge.svg?branch=master&service=github)](https://coveralls.io/github/pocke/BDM4065UC11?branch=master)


Installation
---------------

```sh
go get github.com/pocke/BDM4065UC11
```


Usage
------

```go
c, err := bdm.New("/dev/ttyUSB0", 9600)
if err != nil {
  panic(err)
}
defer c.Close()

// Get model number
data, err := c.Send([]byte{0xA2, 0x01})
if err != nil {
  panic(err)
}

fmt.Println(data.String())  // => BDM4065UC
```

Refs
--------

### Official manual

- http://img.billiger.de/dynimg/Rh5uazBxyvkX90aFpmTDmHt4H9rycHjeBa9LTdXG2DAdykIt2twci304WNNeiecNpF9N2LdMexGE-4B9pS958M/Bedienungsanleitung.pdf
- (japanese) http://download.p4c.philips.com/files/b/bdm4065uc_11/bdm4065uc_11_dfu_jpn.pdf
- SICP protocol http://epi-na.com/wp-content/uploads/2014/10/SICP_Command_Document_master.pdf
