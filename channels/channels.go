package main

type Calc struct {
	fx1 *int
	fx2 *int
}

func NewCalc(in1 <-chan int, in2 <-chan int, f func(int) int) *Calc {
	c := Calc{nil, nil}

	go func(c *Calc, in1 <-chan int) {
		fx1 := f(<-in1)
		c.fx1 = &fx1
	}(&c, in1)
	go func(c *Calc, in2 <-chan int) {
		fx2 := f(<-in2)
		c.fx2 = &fx2
	}(&c, in2)

	return &c
}

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func(f func(int) int, in1 <-chan int,
		in2 <-chan int, out chan<- int, n int) {

		calcs := make([]*Calc, n)

		for i := 0; i < n; i++ {
			calcs[i] = NewCalc(in1, in2, f)
		}
	}(f, in1, in2, out, n)
}
