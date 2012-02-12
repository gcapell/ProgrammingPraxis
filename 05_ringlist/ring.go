package main

import (
	"log"
	"container/ring"
)


func main() {
	log.SetFlags(0) // quieter logging

	size := 41
	advance := 3
	
	r := ring.New(size)
	for j := 0; j< size; j++ {
		r.Value = j
		r = r.Next()
	}
	
	log.Print("initial len:", r.Len())
	for r.Len()>1 {
		r = r.Move(advance-1)
		deleted := r.Unlink(1)
		log.Print("del: ", val(deleted))
	}
	log.Print("survivor:", val(r))
}

func val(r *ring.Ring) int{
	a, ok := r.Value.(int)
	if ok {
		return a
	}
	return -1	
}