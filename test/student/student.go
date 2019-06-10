package student

import "fmt"

//Raid1c ...
func Raid1c(x, y int) {
	hch := make([]string, x)
	hcb := make([]string, x)
	vc := make([]string, x)

	//horizontal haut
	for i := 0; i < x; i++ {
		if i == 0 {
			hch[i] = "A"
		} else if i == (x - 1) {
			hch[i] = "A"
		} else {
			hch[i] = "B"
		}
	}

	//horizontal bas
	for i := 0; i < x; i++ {
		if i == 0 {
			hcb[i] = "C"
		} else if i == (x - 1) {
			hcb[i] = "C"
		} else {
			hcb[i] = "B"
		}
	}

	//verticale
	for i := 0; i < len(vc); i++ {
		if i == 0 {
			vc[i] = "B"
		} else if i == (len(vc) - 1) {
			vc[i] = "B"
		} else {
			vc[i] = " "
		}
	}

	for i := 0; i < x; i++ {
		if i == (x - 1) {
			fmt.Println(hch[i])
		} else {
			fmt.Print(hch[i])
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
				fmt.Println(hcb[i])
			} else {
				fmt.Print(hcb[i])
			}
		}
	}

}
