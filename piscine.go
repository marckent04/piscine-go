package piscine

import (
	"fmt"
	"strconv"
)

//PrintComb2 ...
func PrintComb2() {
	var n1s, n2s string
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			n1s = strconv.Itoa(i) + strconv.Itoa(j)
			n1, err1 := strconv.Atoi(n1s)

			for k := 0; k < 10; k++ {
				for f := 0; f < 10; f++ {
					n2s = strconv.Itoa(k) + strconv.Itoa(f)
					n2, err2 := strconv.Atoi(n2s)
					if n1 == 99 && n2 == 98 {
						fmt.Println(n1s + " " + n2s)
					} else if n1 != n2 {
						fmt.Print(n1s + " " + n2s + ", ")
					}

					if f == 10 {
						fmt.Print(err2)
					}
				}
			}
			if j == 10 {
				fmt.Print(err1)
			}
		}
	}
}
