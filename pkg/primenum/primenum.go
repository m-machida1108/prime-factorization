package primenum

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

	res := []int64{}
	i := int64(2)
	for {
		if target == i {
			res = append(res, target)
			break
		}
		if !Is(i) || target%i != 0 {
			i++
			continue
		}
		target = target / i
		res = append(res, i)
	}

	return res
}
