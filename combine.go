package product

type Combine struct {
	use []bool
	arr []int
	wid int
	het int
	end bool
}

func NewCombine(l, t int) *Combine {
	c := &Combine{}
	if 0 < l && l <= t {
		c.wid = l
		c.het = t
		c.arr = make([]int, l)
		c.use = make([]bool, t)
		c.end = false
		for i := 0; i < l; i++ {
			c.arr[i] = i
			c.use[i] = true
		}
	} else {
		c.end = true
	}
	return c
}

func (c *Combine) Init(l, t int) {
	if 0 < l && l <= t {
		c.wid = l
		c.het = t
		c.arr = make([]int, l)
		c.use = make([]bool, t)
		c.end = false
		for i := 0; i < l; i++ {
			c.arr[i] = i
			c.use[i] = true
		}
	} else {
		c.wid = 0
		c.het = 0
		c.arr = nil
		c.use = nil
		c.end = true
	}
}

func (c *Combine) Pick(t int) int {
	if c.end {
		return -1
	}
	for t < 0 {
		t += c.wid
	}
	for t >= c.wid {
		t -= c.wid
	}
	return c.arr[t]
}

func (c *Combine) Next() {
	if c.end {
		return
	}
	t := c.wid - 1
	for t >= 0 {
		if c.arr[t] >= 0 {
			c.use[c.arr[t]] = false
		} else {
			c.arr[t] = c.arr[t-1]
		}
		for {
			c.arr[t]++
			if c.arr[t] >= c.het || !c.use[c.arr[t]] {
				break
			}
		}
		if c.arr[t] < c.het {
			c.use[c.arr[t]] = true
			t++
			if t == c.wid {
				break
			}
			c.arr[t] = -1
		} else {
			t--
		}
	}
	if t < 0 {
		c.end = true
	}
}

func (c *Combine) Over() bool {
	return c.end
}

func (c *Combine) Length() int {
	return c.wid
}
