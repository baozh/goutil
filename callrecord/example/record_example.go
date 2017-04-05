package main

import (
	cr "github.com/baozh/goutil/callrecord"
	"time"
)

func main() {

	cc := cr.GetCounter("test")

	go func () {
		t := time.NewTicker(time.Second)
		defer t.Stop()

		for {
			select {
			case <- t.C:
				cc.IncreBy(2);
				cc.IncreBy(2);
			}
		}
	}()

	time.Sleep(time.Hour * 5)
}













