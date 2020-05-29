package main

type Calc struct {
	x1 *int
	x2 *int
}

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func(f func(int) int, in1 <-chan int,
		in2 <-chan int, out chan<- int, n int) {

		calcs := make([]*Calc, n)

		for i := 0; i < n; i++ {
			calcs[i] = &Calc{nil, nil}
		}

		completed := make(chan *Calc, n*2)
		go read_channel(&calcs, 1, in1, n, f, completed)
		go read_channel(&calcs, 2, in2, n, f, completed)

		go printer(completed, out, n)
	}(f, in1, in2, out, n)
}

func read_channel(calcs *[]*Calc, field int, ch <-chan int,
	n int, f func(int) int, completed chan *Calc) {
	for i := 0; i < n; i++ {
		x := <-ch

		go func(x int, i int) {
			fx := f(x)
			c := (*calcs)[i]

			if field == 1 {
				c.x1 = &fx
			} else if field == 2 {
				c.x2 = &fx
			}

			completed <- c
		}(x, i)
	}
}

func printer(completed chan *Calc, out chan<- int, n int) {
	printed := 0
	for {
		c := <-completed

		if c.x1 != nil && c.x2 != nil {
			out <- *c.x1 + *c.x2
			printed++
		}

		if printed == n {
			break
		}
	}
}
