// Package mywg is an alternative GO waitgroup that similar to sync.WaitGroup,
// the key difference is mywg allow rountine to wait on a channel,
// so that the wait could be in a select statement;
package mywg

import (
	"sync"
)

type emptyType struct{}

// MyWG is the alternative waitgroup
type MyWG struct {
	val  uint
	lock sync.RWMutex
	// MyWG will send an empty struct into FinishChan once all done
	FinishChan chan emptyType
}

// NewMyWG creates a new MyWG instance
func NewMyWG() *MyWG {
	r := new(MyWG)
	r.FinishChan = make(chan emptyType)
	return r
}

// Add has semantics as sync.WaitGroup.Add(), only difference is this uses uint
func (wg *MyWG) Add(delta uint) {
	wg.lock.Lock()
	defer wg.lock.Unlock()
	wg.val += delta
}

// Wait has semantics as sync.WaitGroup.Add().
// it returns if wg internal value == 0
func (wg *MyWG) Wait() {
	<-wg.FinishChan
}

// Done has semantics as sync.WaitGroup.Add()
func (wg *MyWG) Done() {
	wg.lock.Lock()
	defer wg.lock.Unlock()
	wg.val -= 1
	if wg.val == 0 {
		close(wg.FinishChan)
	}
}

// Cancel stop the wg, cause Wait() to return, regardless if all is done or not;
func (wg *MyWG) Cancel() {
	close(wg.FinishChan)
}
