package product

type Isolate struct {
	arr []int
	wid int
	het int
	end bool
}

func NewIsolate(l, t int) *Isolate {
	i := &Isolate{}
	if 0 < l && 0 < t {
		i.wid = l
		i.het = t
		i.arr = make([]int, l)
		i.end = false
	} else {
		i.end = true
	}
	return i
}

func (i *Isolate) Init(l, t int) {
	if 0 < l && 0 < t {
		i.wid = l
		i.het = t
		i.arr = make([]int, l)
		i.end = false
	} else {
		i.wid = 0
		i.het = 0
		i.arr = nil
		i.end = true
	}
}

func (i *Isolate) Pick(t int) int {
	if i.end {
		return -1
	}
	for t < 0 {
		t += i.wid
	}
	for t >= i.wid {
		t -= i.wid
	}
	return i.arr[t]
}

func (i *Isolate) Next() {
	if i.end {
		return
	}
	t := i.wid - 1
	i.arr[t]++
	for t > 0 && i.arr[t] == i.het {
		i.arr[t] = 0
		t--
		i.arr[t]++
	}
	if i.arr[0] >= i.het {
		i.end = true
	}
}

func (i *Isolate) Over() bool {
	return i.end
}

func (i *Isolate) Length() int {
	return i.wid
}
