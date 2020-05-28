package main

type Calc struct {
	x1  *int
	x2  *int
	out *chan<- int
}

func (c *Calc) set(x1 *int, x2 *int) {
	if x1 != nil {
		c.x1 = x1
	}
	if x2 != nil {
		c.x2 = x2
	}
	if c.x1 != nil && c.x2 != nil {
		*c.out <- *c.x1 + *c.x2
	}
}

func Merge2Channels(f func(int) int, in1 <-chan int,
	in2 <-chan int, out chan<- int, n int) {

	go func() {
		for i := 0; i < n; i++ {
			c := Calc{nil, nil, &out}
			x1 := <-in1
			x2 := <-in2
			go func(c *Calc, x1 int) {
				fx1 := f(x1)
				c.set(&fx1, nil)
			}(&c, x1)
			go func(c *Calc, x2 int) {
				fx2 := f(x2)
				c.set(nil, &fx2)
			}(&c, x2)
		}
	}()
}
