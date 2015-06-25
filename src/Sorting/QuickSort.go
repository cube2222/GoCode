package main
import "fmt"

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
			if (tab[i] > tab[j]) {
				helper := tab[i]
				tab[i] = tab[j]
				tab[j] = helper
			}
		}
	}
	return tab
}

func main() {
	var tab []int
	tab = []int{0, 2, 1, 7, 4, 3, 8}
	fmt.Println(tab)
	tab = sort(tab, 0, len(tab)-1)
	fmt.Println(tab)
}
