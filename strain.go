package strain

type Ints []int

type Strings []string

type Lists []Ints

type strain interface {
	Keep(func() bool) Ints
	Discard(func() bool) Ints
}

func (c Ints) Keep(f func(int) bool) Ints {
	var r Ints
	for _, i := range c {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

func (c Ints) Discard(f func(int) bool) Ints {
	var r Ints
	for _, i := range c {
		if !(f(i)) {
			r = append(r, i)
		}
	}
	return r
}
func (c Strings) Keep(f func(string) bool) Strings {
	var r Strings
	for _, i := range c {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

func (c Strings) Discard(f func(string) bool) Strings {
	var r Strings
	for _, i := range c {
		if !(f(i)) {
			r = append(r, i)
		}
	}
	return r
}

func (c Lists) Keep(f func([]int) bool) Lists {
	var r Lists
	for _, i := range c {
		if f(i) {
			r = append(r, i)
		}
	}
	return r
}

func (c Lists) Discard(f func([]int) bool) Lists {
	var r Lists
	for _, i := range c {
		if !(f(i)) {
			r = append(r, i)
		}
	}
	return r
}
