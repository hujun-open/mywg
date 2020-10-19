// ExampleMyWG
package mywg

import (
	"fmt"
	"time"
)

func ExampleMyWG() {
	wg := NewMyWG()
	wg.Add(2)
	f := func(wg *MyWG, dur int) {
		defer wg.Done()
		time.Sleep(time.Duration(dur) * time.Second)
		fmt.Printf("done after %d seconds\n", dur)
	}
	go f(wg, 3)
	go f(wg, 5)
	<-wg.FinishChan
	fmt.Println("all done")
}
