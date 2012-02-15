package main

// Baker, Cooper, Fletcher, Miller and Smith live on different floors of
// an apartment house that contains only five floors. Baker does not live
// on the top floor. Cooper does not live on the bottom floor. Fletcher
// does not live on either the top or the bottom floor. Miller lives
// on a higher floor than does Cooper. Smith does not live on a floor
// adjacent to Fletcher’s. Fletcher does not live on a floor adjacent to
// Cooper’s. Where does everyone live?


import (
	"log"
)

func main() {
	log.SetFlags(0) // quieter logging

	checkperms([]int{0,1,2,3,4}, 0)
}

func checkperms(base []int, pos int) {
	if pos == len(base)-1 {
		if valid(base[0],base[1],base[2],base[3],base[4]) {
			log.Print(base)
		}
		return
	}
	for j:=pos; j<len(base); j++ {
		base[j], base[pos] = base[pos], base[j]
		checkperms(base, pos+1)
		base[j], base[pos] = base[pos], base[j]
	}
}

func adjacent(a,b int) bool {
	return a == b +1 || a == b - 1
}

func valid(b,c,f,m,s int) bool {
	return b !=4 &&
		c != 0 &&
		f != 4 && f != 0 &&
		m > c &&
		!adjacent(s,f) &&
		!adjacent(f,c)
}