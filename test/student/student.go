package student

import "fmt"

//Raid1a ...
func Raid1a(x, y int) {
	hc := make([]string, x)
	vc := make([]string, x)

	//horizontal
	for i := 0; i < x; i++ {
		if i == 0 {
			hc[i] = "o"
		} else if i == (x - 1) {
			hc[i] = "0"
		} else {
			hc[i] = "-"
		}
	}

	//verticale
	for i := 0; i < len(vc); i++ {
		if i == 0 {
			vc[i] = "|"
		} else if i == (len(vc) - 1) {
			vc[i] = "|"
		} else {
			vc[i] = " "
		}
	}

	for i := 0; i < x; i++ {
		if i == (x - 1) {
			fmt.Println(hc[i])
		} else {
			fmt.Print(hc[i])
		}
	}

	for a := 0; a < (y - 2); a++ {
		for i := 0; i < len(vc); i++ {
			fmt.Print(vc[i])
		}
		fmt.Print("\n")
	}

	if y > 1 {
		for i := 0; i < x; i++ {
			if i == (x - 1) {
				fmt.Println(hc[i])
			} else {
				fmt.Print(hc[i])
			}
		}
	}

}
