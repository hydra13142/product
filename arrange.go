package product

type Arrange struct {
	use []bool
	arr []int
	wid int
	het int
	end bool
}

func NewArrange(l, t int) *Arrange {
	a := &Arrange{}
	if 0 < l && l <= t {
		a.wid = l
		a.het = t
		a.arr = make([]int, l)
		a.use = make([]bool, t)
		a.end = false
		for i := 0; i < l; i++ {
			a.arr[i] = i
			a.use[i] = true
		}
	} else {
		a.end = true
	}
	return a
}

func (a *Arrange) Init(l, t int) {
	if 0 < l && l <= t {
		a.wid = l
		a.het = t
		a.arr = make([]int, l)
		a.use = make([]bool, t)
		a.end = false
		for i := 0; i < l; i++ {
			a.arr[i] = i
			a.use[i] = true
		}
	} else {
		a.wid = 0
		a.het = 0
		a.arr = nil
		a.use = nil
		a.end = true
	}
}

func (a *Arrange) Pick(t int) int {
	if a.end {
		return -1
	}
	for t < 0 {
		t += a.wid
	}
	for t >= a.wid {
		t -= a.wid
	}
	return a.arr[t]
}

func (a *Arrange) Next() {
	if a.end {
		return
	}
	t := a.wid - 1
	for t >= 0 {
		if a.arr[t] >= 0 {
			a.use[a.arr[t]] = false
		}
		for {
			a.arr[t]++
			if a.arr[t] >= a.het || !a.use[a.arr[t]] {
				break
			}
		}
		if a.arr[t] < a.het {
			a.use[a.arr[t]] = true
			t++
			if t == a.wid {
				break
			}
			a.arr[t] = -1
		} else {
			t--
		}
	}
	if t < 0 {
		a.end = true
	}
}

func (a *Arrange) Over() bool {
	return a.end
}

func (a *Arrange) Length() int {
	return a.wid
}
