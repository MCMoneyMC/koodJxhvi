package piscine

func Swap(a *int, b *int) {
	*a, *b = *b, *a
	_ = *a + *b // this is bullying
}
