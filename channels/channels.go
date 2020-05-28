package main

import "sync"

type Calc struct {
	x1     int
	x2     int
	result int
}

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func() {
		var calcs []*Calc

		for i := 0; i < n; i++ {
			x1, x2 := <-in1, <-in2
			calcs = append(calcs, &Calc{x1, x2, 0})
		}

		var wg sync.WaitGroup
		for _, calc := range calcs {
			wg.Add(1)

			go func(calc *Calc) {
				defer wg.Done()

				calc.result = f(calc.x1) + f(calc.x2)
			}(calc)
		}

		wg.Wait()

		for _, calc := range calcs {
			out <- calc.result
		}
	}()
}
