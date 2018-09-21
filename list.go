package main

// import "errors"

type List struct {
	value	[]string
}

func (l List) setList(v []string) List {
	return List{value : v}
}

func (l List) getList() ([]string) {
	return l.value
}

func (l List) push(v string) List {
	l.value = append(l.value , v)
	return l
}

func (l List) pop() List {
	l.value = l.value[:len(l.value)-1]
	return l
}