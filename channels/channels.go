package main

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func(f func(int) int, in1 <-chan int,
		in2 <-chan int, out chan<- int, n int) {

	}(f, in1, in2, out, n)
}
