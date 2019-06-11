package piscine

//DivMod ...
func DivMod(a int, b int, div *int, mod *int) {
	pdiv := div
	pmod := mod

	*pdiv = a / b
	*pmod = a % b
}
