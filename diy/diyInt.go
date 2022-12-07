package diy

const (
	MaxInt = 1<<31 - 1
	MinInt = -1 << 31
)

func Min(i, j int) int {
	if i <= j {
		return i
	} else {
		return j
	}
}

func Max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
