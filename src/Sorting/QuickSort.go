package main
import (
	"fmt"
	"math/rand"
)

func sort(tab []int, left, right int) []int {
	i, j, m := left, right, (left+right)/2
	if (j - i + 1) == 2 {
		if (tab[i] <= tab[j]) {
			return tab
		} else {
			helper := tab[i]
			tab[i] = tab[j]
			tab[j] = helper
		}
	} else if (j - i + 1) < 2 {
		return tab
	} else {
		for i < j {
			for tab[i] < tab[m] {
				i++
			};
			for tab[j] > tab[m] {
				j--
			};
			if tab[i] > tab[j] {
				helper := tab[i]
				tab[i] = tab[j]
				tab[j] = helper
			}
			if tab[i] == tab[j] {
				if tab[i] <= tab[m] {
					i++
				} else {
					j--
				}
			}
		}
		tab = sort(tab, left, m)
		tab = sort(tab, m, right)
	}
	return tab
}

func main() {
	var tab []int
	tab = []int{}
	for i := 0; i < 1000; i++ {
		tab = append(tab, rand.Intn(100))
	}
	fmt.Println(tab)
	fmt.Println("***************************************************************")
	tab = sort(tab, 0, len(tab)-1)
	fmt.Println(tab)
}
