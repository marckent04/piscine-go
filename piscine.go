package piscine

import (
	"fmt"
	"strconv"
)

//PrintComb2 ...
func PrintComb2() {
	var n1s, n2s string
	var n1, n2 int
	var err1, err2 error
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			n1s = strconv.Itoa(i) + strconv.Itoa(j)
			n1, err1 = strconv.Atoi(n1s)

			//creation du 2eme nombre
			for k := 0; k < 10; k++ {
				for f := 0; f < 10; f++ {
					n2s = strconv.Itoa(k) + strconv.Itoa(f)
					n2, err2 = strconv.Atoi(n2s)

					if n1 != n2 {
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

			if n1 == 98 && n2 == 99 {
				fmt.Println(n1s + " " + n2s)
				break
			}
		}
	}
}
