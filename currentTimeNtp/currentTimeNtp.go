package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func GetCurrentTimeNtp(host string) {

	time, err := ntp.Time(host)

	if err != nil {
		log.Fatal("cant get time \n", err)
	}
	fmt.Printf("current time: %v\n", time)
}
