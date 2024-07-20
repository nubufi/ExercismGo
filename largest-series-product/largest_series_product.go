package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	series,err := getSeries(digits,span)
	var highest int64

	if err != nil {
		return highest,err
	}
	for _,serie := range series {
		product := calcProduct(serie)
		if product> highest{
			highest = product
		}
	}

	return highest,nil
}

func getSeries(digits string,span int) ([][]int64,error) {
	if span<0 {
		return nil, errors.New("span can't me negative")
	}
	var series  [][]int64
	slice,err := stringToIntSlice(digits)
	if err != nil {
		return nil,err
	}
	if len(slice)<span {
		return nil, errors.New("span must be smaller than string length")
	}
	if len(slice) == span {
		return [][]int64{slice},nil
	}
	i:=0
	j:=span
	for j!=len(slice)+1{
		series = append(series,slice[i:j])
		i++
		j++
	}
	return series,nil
}

func stringToIntSlice(s string) ([]int64, error) {
	var result []int64
	for _, char := range s {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, errors.New("digits input must only contain digits")
		}
		result = append(result, int64(num))
	}
	return result, nil
}

func calcProduct(serie []int64) int64 {
	var product int64 = 1
	for _,n := range serie {
		product *= n
	}
	return product
}
