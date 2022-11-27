package currentTimeNtp

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func GetCurrentTimeNtp() {
	host := "0.beevik-ntp.pool.ntp.org"
	time, err := ntp.Time(host)

	if err != nil {
		log.Fatal("cant get time \n", err)
	}

	fmt.Printf("Current time: %v\n", time)
}
