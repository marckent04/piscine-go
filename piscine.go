package piscine

import "fmt"

//PrintComb ...
func PrintComb() {
	n1, n2, n3 := 0, 0, 0

	for a := 0; a <= 9; a++ {
		for b := 0; b < 10; b++ {
			if b == 0 {
				n2 = 0
			}
			for c := 0; c < 10; c++ {
				if c == 0 {
					n3 = 0
				}
				if n3 > n2 && n2 > n1 {
					if n1 == 7 && n2 == 8 && n3 == 9 {
						fmt.Print(n1)
						fmt.Print(n2)
						fmt.Print(n3)
					} else {
						fmt.Print(n1)
						fmt.Print(n2)
						fmt.Print(n3)
						fmt.Print(",")
					}
				}
				n3++

			}
			n2++
		}

		n1++
	}
}
