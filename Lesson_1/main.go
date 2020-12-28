package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const (
	OK = 0
	ERROR = 1
)


func main(){
	var ntpServer string = "0.beevik-ntp.pool.ntp.org"
	var maxCount int = 10
	if result, err := now(ntpServer, maxCount); err != nil{
		os.Exit(ERROR)
	} else {
		fmt.Printf("Now is: %s\n", result )
		os.Exit(OK)
	}
}

func now(host string, max_count int) (string, error) {
	count := 0
	for count < max_count{
		cTime, err := ntp.Time(host)
		if err != nil{
			count++
			fmt.Printf("Wait ! Error %e\n", err.Error())
			time.Sleep(1 * time.Second)
			if count >= max_count {
				return "", err
			}
		} else{
			return cTime.Format("2006-January-02 15:04:05"), err
		}
	}
	return "", nil
}