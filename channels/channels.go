package main

import "log"

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func(f func(int) int, in1 <-chan int,
		in2 <-chan int, out chan<- int, n int) {

		log.Println("Started, n", n)

		calcs := make([]*int, n*2)
		completed := make(chan bool, n*2)

		go read_channel(calcs, 0, in1, n, f, completed)
		go read_channel(calcs, 1, in2, n, f, completed)

		go printer(calcs, completed, out, n)
	}(f, in1, in2, out, n)
}

func read_channel(calcs []*int, shift int, ch <-chan int,
	n int, f func(int) int, completed chan bool) {

	log.Println("Reader started", shift)

	for i := 0; i <= n; i++ {
		x := <-ch

		go func(x int, i int) {
			fx := f(x)

			log.Println("i, x, fx", i, x, fx)

			calcs[i+shift] = &fx
			completed <- true
		}(x, i)
	}

	log.Println(shift, "read done")
}

func printer(calcs []*int, completed chan bool, out chan<- int, n int) {
	log.Println("Printer started")

	printed := 0
	for <-completed {
		for i := printed; i <= n*2; i += 2 {
			x1, x2 := calcs[i], calcs[i+1]

			if x1 != nil && x2 != nil {
				log.Println("x1, x2, printed", *x1, *x2, printed)

				out <- (*x1 + *x2)
				printed += 2
			} else {
				break
			}

			if printed == n*2 {
				return
			}
		}
	}

	log.Println("Print done")
}
