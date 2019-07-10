package main

import (
	"fmt"
	"saprfc/helper"
	"time"
)

func abapSystem() saprfc.ConnectionParameter {
	return saprfc.ConnectionParameter{
		Dest:      "",
		Client:    "800",
		User:      "HEX",
		Passwd:    "123456",
		Lang:      "ZH",
		Ashost:    "127.0.0.1",
		Sysnr:     "00",
		Saprouter: "",
	}
}

func main() {
	c, err := saprfc.ConnectionFromParams(abapSystem())
	if err != nil {
		fmt.Printf("Connection Error: %#v \n", err)
	}

	in_date, err := time.Parse("2006-01-02", "2019-07-01")
	params := map[string]interface{}{
		"IN_DATE": in_date,
	}

	r, err := c.Call("ZMM_MATNR", params)
	if err != nil {
		fmt.Printf("Call Error: %#v \n", err)
	}

	fmt.Printf("Response: %#v \n", r)

	c.Close()
}
