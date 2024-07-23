package stringset

import "fmt"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set []string

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	set := Set{}
	for _, v := range l {
		if !set.Has(v) {
			set = append(set, v)
		}
	}
	return set
}

func (s Set) String() string {
	if len(s) == 0 {
		return "{}"
	}
	if len(s) == 1 {
		return fmt.Sprintf(`{"%s"}`, s[0])
	}
	string := "{"
	for i, v := range s {
		if i == len(s)-1 {
			string += fmt.Sprintf(`"%s"`, v)
		} else {
			string += fmt.Sprintf(`"%s", `, v)
		}
	}
	return string + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		*s = append(*s, elem)
	}
}

func Subset(s1, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}
	if len(s1) > len(s2) {
		return false
	}

	for _, v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}

	for _, v := range s1 {
		if s2.Has(v) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	intersection := Set{}
	if len(s1)>len(s2){
		s1, s2 = s2, s1
	}

	for _, v := range s1 {
		if s2.Has(v) {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

func Difference(s1, s2 Set) Set {
	difference := Set{}
	for _, v := range s1 {
		if !s2.Has(v) {
			difference = append(difference, v)
		}
	}

	return difference
}

func Union(s1, s2 Set) Set {
	union := Set{}
	for _, v := range s1 {
		if !union.Has(v) {
			union = append(union, v)
		}
	}
	for _, v := range s2 {
		if !union.Has(v) {
			union = append(union, v)
		}
	}

	return union
}
