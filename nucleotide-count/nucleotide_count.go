package dna

import "errors"

type Histogram map[byte]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for i:=0; i<len(d);i++ {
		val,ok := h[d[i]]
		if !ok {
			return h,errors.New("Invalid DNA")
		}
		val++
		h[d[i]] = val
	}
	

	return h, nil
}
