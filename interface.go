package product

type Iterator interface {
	Pick(int) int
	Next()
	Over() bool
	Length() int
}
