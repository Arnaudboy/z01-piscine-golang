package main

func IsValid(sol *[]int) bool {
	// Check uniqueness -> No queen on the same line
	for _, v := range sol {
		for j := 1; j < len(sol)-1; j++ {
			if v == sol[j] {
				return false
			}
		}
	}
	//Check diagonal
	if len(sol) > 1 {
		for j := 1; j < len(sol)-2; j++ {
			if sol[j] == sol[j+1]+j || sol[j] == sol[j-1]-j {
				return false
			}
		}
	}
	return true
}

func Solve(sol *[]int) {
	if len(sol) == 8 && IsValid(sol) {
		return true
	}
	for i := 1; i < 9; i++ {
		for _, x := range sol {
			x = i
			if IsValid(sol) {
				if Solve(sol) {
					return true
				}
			}
			x = 0
		}
	}
}
