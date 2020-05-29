package main

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func(f func(int) int, in1 <-chan int,
		in2 <-chan int, out chan<- int, n int) {

		x1s := make([]*int, n)
		x2s := make([]*int, n)
		completed := make(chan bool, n*2)

		read := func(xs []*int, ch <-chan int) {
			for i := 0; i < n; i++ {
				x := <-ch

				go func(i int, x int) {
					fx := f(x)
					xs[i] = &fx
					completed <- true
				}(i, x)
			}
		}

		go read(x1s, in1)
		go read(x2s, in2)

		go func() {
			pushed := 0
			for <-completed {
				for i := pushed; i < n; i++ {
					if x1s[i] != nil && x2s[i] != nil {
						out <- *x1s[i] + *x2s[i]
						pushed++
						if pushed == n {
							return
						}
					} else {
						break
					}

				}
			}
		}()
	}(f, in1, in2, out, n)
}
