package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}
	first := nums[1] - nums[0]
	out := Stats{Count: len(nums) - 1, Sum: first, Min: first, Max: first}
	for i := 2; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		out.Sum += d
		if d < out.Min {
			out.Min = d
		}
		if d > out.Max {
			out.Max = d
		}
	}

	return out
}
