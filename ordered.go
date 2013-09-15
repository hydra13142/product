package product

type Ordered struct {
	arr []int
	wid int
	het int
	end bool
}

func NewOrdered(l, t int) *Ordered {
	o := &Ordered{}
	if 0 < l && 0 < t {
		o.wid = l
		o.het = t
		o.arr = make([]int, l)
		o.end = false
	} else {
		o.end = true
	}
	return o
}

func (o *Ordered) Init(l, t int) {
	if 0 < l && 0 < t {
		o.wid = l
		o.het = t
		o.arr = make([]int, l)
		o.end = false
	} else {
		o.wid = 0
		o.het = 0
		o.arr = nil
		o.end = true
	}
}

func (o *Ordered) Pick(t int) int {
	if o.end {
		return -1
	}
	for t < 0 {
		t += o.wid
	}
	for t >= o.wid {
		t -= o.wid
	}
	return o.arr[t]
}

func (o *Ordered) Next() {
	if o.end {
		return
	}
	t := o.wid - 1
	for t >= 0 {
		o.arr[t]++
		if o.arr[t] < o.het {
			for u, t := o.arr[t], t+1; t < o.wid; t++ {
				o.arr[t] = u
			}
			break
		}
		t--
	}
	if o.arr[0] >= o.het {
		o.end = true
	}
}

func (o *Ordered) Over() bool {
	return o.end
}

func (o *Ordered) Length() int {
	return o.wid
}
