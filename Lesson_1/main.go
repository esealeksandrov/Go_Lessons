package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)


func main(){
	var ntp_server string = "0.beevik-ntp.pool.ntp.org"
	fmt.Printf("Now is: %s\n", now(ntp_server))
}

func now(host string) string {
	for {
		c_time, err := ntp.Time(host)
		if err != nil{
			fmt.Printf("Wait ! Error %e\n", err.Error())
			time.Sleep(1 * time.Second)
		} else{
			return c_time.Format("2006-January-02 15:04:05")
		}
	}
}