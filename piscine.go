package piscine

//UltimateDivMod ...
func UltimateDivMod(a *int, b *int) {
	pa := a
	pb := b

	*pb = *a % *pb

	*pa = *a / *pb
}
