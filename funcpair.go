package gosimplex

type FuncPair struct{
	Value *Vector
	Result float64
}

func QuicksortFuncPair(a []FuncPair)[]FuncPair{
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) >> 1

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Result < a[right].Result {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuicksortFuncPair(a[:left])
	QuicksortFuncPair(a[left+1:])

	return a
}

func Reverse(a []FuncPair)[]FuncPair{
	reverse := make([]FuncPair,len(a))
	
	for i := len(a)-1; i >= 0;i--{
		reverse[len(a)-1-i] = a[i]
	}

	return reverse
}

func CheckVectorContainment(v *Vector,a []FuncPair)bool{
	for _,j := range a{
		if j.Value != nil && v.Cmp((j.Value)){
			return true
		}
	}
	return false
}