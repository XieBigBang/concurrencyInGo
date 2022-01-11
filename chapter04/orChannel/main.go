package main

import (
	"concurrencyInGo/chapter04/orChannel/or"
	"fmt"
	"time"
)

func sig(duration time.Duration) <-chan interface{} {
	tmp := make(chan interface{})
	go func() {
		defer close(tmp)
		time.Sleep(duration)
	}()
	return tmp
}

func main() {
	now := time.Now()

	i1 := sig(time.Second * 191)
	i2 := sig(time.Second * 141)
	i3 := sig(time.Second * 16)
	i4 := sig(time.Second * 3)

	<-or.OrChan(i1, i2, i3, i4)

	fmt.Println(time.Since(now))
	for {
		select {
		case <-i1:
			fmt.Println("i1 closed")
			time.Sleep(time.Second)
		case <-i2:
			fmt.Println("i2 closed")
			time.Sleep(time.Second)
		case <-i3:
			fmt.Println("i3 closed")
			time.Sleep(time.Second)
		case <-i4:
			fmt.Println("i4 closed")
			time.Sleep(time.Second)
		}
	}

}
