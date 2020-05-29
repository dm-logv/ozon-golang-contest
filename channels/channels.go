package main

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func(f func(int) int, in1 <-chan int,
		in2 <-chan int, out chan<- int, n int) {

		for i := 0; i < n; i++ {
			out <- f(<-in1) + f(<-in2)
		}
	}(f, in1, in2, out, n)
}
