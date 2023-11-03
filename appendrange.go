package piscine

func AppendRange(min int, max int) []int {
	if min >= max {
		return nil
	}
	var mmax []int
	for i := min; i < max; i++ {
		mmax = append(mmax, i)
	}
	return mmax
}
