// psáno v jazyce golang
package main

import (
	"fmt"
)

func back(konici []int, ground []int, spaces []int) (police []int) { // vrací poníky na polici
	if len(ground) == 0 { // došli poníci na zemi => skonči, je vráceno
		return konici
	}
	for i, kun := range ground { // projdi poníky na zemi
		if spaces[0] == 0 { // mezera na začátku
			if konici[1] == kun {
				continue
			}
		} else if spaces[0] == (len(konici) - 1) { // mezera na konci
			if konici[len(konici)-2] == kun {
				continue
			}
		} else if konici[spaces[0]-1] == kun || konici[spaces[0]+1] == kun { // mezera uprostřed
			continue
		} // kolem všech mezer jsou jiné barvy, zkus vrátit poníka
		next := konici[:]
		next[spaces[0]] = kun
		ground[i], ground[0] = ground[0], ground[i]
		police = back(next, ground[1:], spaces[1:])
		if police == nil { // pokud další poník už nejde přidat, zkus jiného
			continue
		} else { // vyšlo to vrať co máš
			break
		}
	}
	return
}

func main() {
	var N int
	var konici []int
	var spaces []int
	var down int
	var ground []int
	fmt.Scan(&N) // načte data
	konici = make([]int, N)
	spaces = make([]int, 0)
	for i := 0; i < N; i++ {
		fmt.Scan(&konici[i])
		if konici[i] == -1 {
			spaces = append(spaces, i)
			down++
		}
	}
	ground = make([]int, down)
	for i := 0; i < down; i++ {
		fmt.Scan(&ground[i])
	}
	out := fmt.Sprint(back(konici, ground, spaces))
	out = out[1 : len(out)-1] // odstraní [] na začátku a na konci
	if len(out) == 0 {
		fmt.Println("Musis to preskladat!")
	} else {
		fmt.Println(out)
	}
}
