package primenum

import "fmt"

func Is(target int64) bool {
	if target == 0 || target == 1 {
		return false
	}
	if target == 2 {
		return true
	}

	for i := int64(2); target > i; i++ {
		if target%i == 0 {
			return false
		}
	}

	return true
}

func Factorization(target int64) []int64 {
	if target == 0 || target == 1 {
		return []int64{}
	}
	if Is(target) {
		return []int64{target}
	}

	//res := []int64{}
	// 100個以上の要素に分解される数は対応しない
	res := make([]int64, 100)
	idx := 0
	i := int64(2)
	for {
		fmt.Printf("target:%d, i:%d\n", target, i)
		if target == i || Is(target) {
			//res = append(res, target)
			res[idx] = target
			break
		}
		if !Is(i) || target%i != 0 {
			i++
			continue
		}
		target = target / i
		//res = append(res, i)
		res[idx] = i
		idx++
	}

	if len(res) == 100 {
		return res
	}
	return res[:idx+1]
}
