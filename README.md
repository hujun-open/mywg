# mywg
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hujun-open/mywg)](https://pkg.go.dev/github.com/hujun-open/mywg)

mywg is an alternative GO waitgrroup that similar to sync.WaitGroup, the key difference is mywg allow rountine to wait on a channel, so that the wait could be in a select statement; 

## example

```
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
```