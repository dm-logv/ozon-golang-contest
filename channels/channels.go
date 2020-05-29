package main

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func(f func(int) int, in1 <-chan int,
		in2 <-chan int, out chan<- int, n int) {

		calcs := make([]*int, n*2)
		completed := make(chan bool, n*2)

		go read_channel(calcs, 0, in1, n, f, completed)
		go read_channel(calcs, 1, in2, n, f, completed)

		go printer(calcs, completed, out, n)
	}(f, in1, in2, out, n)
}

func read_channel(calcs []*int, shift int, ch <-chan int,
	n int, f func(int) int, completed chan bool) {

	for i := 0; i < n; i++ {
		x := <-ch

		go func(x int, i int) {
			fx := f(x)
			calcs[i+shift] = &fx
			completed <- true
		}(x, i)
	}
}

func printer(calcs []*int, completed chan bool, out chan<- int, n int) {
	printed := 0
	for <-completed {
		x1, x2 := calcs[printed], calcs[printed+1]

		if x1 != nil && x2 != nil {
			out <- *x1 + *x2

			printed += 2
		}

		if printed == n*2 {
			return
		}

	}
}
