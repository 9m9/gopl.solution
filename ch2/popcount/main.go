package popcount

// exercise 2.3

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

// PopCount (x uint64) count number of 1s in an uint64 number's binary
// representation
func PopCount(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i*8))])
	}
	return sum
}
