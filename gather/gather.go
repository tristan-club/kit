package gather

func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

func Intersect(slice1, slice2 []string) []string {
	if len(slice1) > len(slice2) {
		return Intersect(slice2, slice1)
	}
	mp := make(map[string]string)
	mp2 := make(map[string]string)
	ans := []string{}
	for _, v := range slice1 {
		mp[v] = v
	}

	for _, v := range slice2 {
		mp2[v] = v
	}

	for _, v := range mp {
		if v, ok := mp2[v]; ok {
			ans = append(ans, v)
		}
	}

	ans = RemoveRepString(ans)
	return ans
}

func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func IntersectInt(slice1, slice2 []int64) []int64 {
	if len(slice1) > len(slice2) {
		return IntersectInt(slice2, slice1)
	}
	mp := make(map[int64]int64)
	mp2 := make(map[int64]int64)
	ans := []int64{}
	for _, v := range slice1 {
		mp[v] = v
	}

	for _, v := range slice2 {
		mp2[v] = v
	}

	for _, v := range mp {
		if v, ok := mp2[v]; ok {
			ans = append(ans, v)
		}
	}

	ans = RemoveRepInt(ans)
	return ans
}

func DifferenceInt(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	nn := make([]int64, 0)
	inter := IntersectInt(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func SortAhead(slice1, slice2 []string) []string {
	m := Intersect(slice1, slice2)
	n := Difference(slice2, m)
	var s []string
	s = m
	s = append(s, n...)
	return s
}

func InArrayString(param string, array []string) bool {
	for _, v := range array {
		if param == v {
			return true
		}
	}
	return false
}

func InArrayInt(param int, array []int) bool {
	for _, v := range array {
		if param == v {
			return true
		}
	}
	return false
}

func RemoveRepString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func RemoveRepInt(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
