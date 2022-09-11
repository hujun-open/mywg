// mywg_test
package mywg

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestMyWGFinishChan(t *testing.T) {
	wg := NewMyWG()
	n := new(uint32)
	atomic.StoreUint32(n, 0)
	f := func(wg *MyWG, sleepdur int, n *uint32) {
		defer t.Logf("return after %d secs", sleepdur)
		time.Sleep(time.Duration(sleepdur) * time.Second)
		atomic.AddUint32(n, 1)
		wg.Done()
	}
	var target uint32 = 10
	wg.Add(uint(target))
	for i := 0; i < int(target); i++ {
		go f(wg, rand.Intn(10), n)
	}
	t.Log("start waiting")
	<-wg.FinishChan
	if atomic.LoadUint32(n) != target {
		t.Fatal("wait finishes before all is done")
	}
	t.Log("Done")
}

func TestMyWGCancel(t *testing.T) {
	wg := NewMyWG()
	wg.Add(2)
	f := func(wg *MyWG, sleepdur int) {
		defer t.Logf("return after %d secs", sleepdur)
		time.Sleep(time.Duration(sleepdur) * time.Second)
		wg.Done()
		wg.Cancel()
	}
	waitf := func(wg *MyWG) {
		t.Log("start waiting")
		<-wg.FinishChan
		t.Logf("stop waiting")
	}
	go f(wg, 3)
	waitf(wg)
	t.Log("Done")
}
