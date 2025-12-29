package main

type Codec struct {
	idx int
}

func Constructor() Codec {
	return Codec{idx: 0}
}
