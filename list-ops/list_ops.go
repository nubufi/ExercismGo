package listops


// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	// If the list is empty, return the initial value
	if len(s) == 0 {
		return initial
	}

	// Iterate through the list and apply the function to the accumulator and the current element
	for _, v := range s {
		initial = fn(initial, v)
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	var foldrHelper func([]int, int) int
	foldrHelper = func(lst []int, acc int) int {
		if len(lst) == 0 {
			return acc
		}
		return fn(lst[0], foldrHelper(lst[1:], acc))
	}
	
	return foldrHelper(s, initial)
}

func (s IntList) Filter(fn func(int) bool) IntList {
	newList := []int{}

	for _,n := range s {
		if fn(n) {
			newList = append(newList,n)
		}
	}
	
	return newList
}

func (s IntList) Length() int {
	var length int

	for i := range s {
		length = i+1
	}

	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	newList := make([]int,len(s))

	for i,n := range s {
		newList[i] = fn(n)
	}

	return newList
}


func (s IntList) Reverse() IntList {
	newList := make([]int,len(s))

	for i,n := range s {
		newList[len(s)-i - 1] = n
	}

	return newList
}

func (s IntList) Append(lst IntList) IntList {
	newList := make([]int,len(s)+len(lst))

	copy(newList,s)

	for i,n := range lst{
		newList[i+len(s)] = n
	}

	return newList

}

func (s IntList) Concat(lists []IntList) IntList {
	newList := s

	for _,n := range lists {
		newList = newList.Append(n)
	}

	return newList
}
