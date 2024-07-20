package sublist

import (
	"reflect"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if reflect.DeepEqual(l1, l2) {
		return RelationEqual
	}

	if len(l1) == 0 {
		return RelationSublist
	}

	if len(l2) == 0 {
		return RelationSuperlist
	}

	if len(l1) > len(l2) {
	  if containSubsequence(l2, l1) {
			return RelationSuperlist
		} else {
			return RelationUnequal
		}
	} else {
		if containSubsequence(l1, l2) {
			return RelationSublist
		} else {
			return RelationUnequal
		}
	}
}

func containSubsequence(small, big []int) bool {
	for i := 0; i < len(big)-len(small)+1; i++ {
		if reflect.DeepEqual(small, big[i:i+len(small)]) {
			return true
		}
	}

	return false
}
