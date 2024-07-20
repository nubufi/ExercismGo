package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func checkN(phonenumber string,i int) bool {
	d1 := phonenumber[i]
	d2 := phonenumber[3+i]

	if d1 < '2' || d2 <'2' {
		return false
	}

	return true
}

func Number(phoneNumber string) (string, error) {
	re := regexp.MustCompile(`[^0-9]`)
    phoneNumber = re.ReplaceAllString(phoneNumber, "")

	if len(phoneNumber) == 10 {
		if !checkN(phoneNumber,0) {
			return "",errors.New("wrong N")
		}
		return phoneNumber,nil
	} else if len(phoneNumber) == 11 && phoneNumber[0] == '1' {
		if !checkN(phoneNumber,1) {
			return "",errors.New("wrong N")
		}
		return string(phoneNumber[1:]),nil
	} else {
		return "",errors.New("wrong format")
	}
}

func AreaCode(phoneNumber string) (string, error) {
	p,err := Number(phoneNumber)
	if err != nil {
		return "",err
	}

	return string(p[:3]),nil
}

func Format(phoneNumber string) (string, error) {
	p,err := Number(phoneNumber)
	code,_ := AreaCode(phoneNumber)

	if err != nil {
		return "",err
	}

	formatted := fmt.Sprintf("(%s) %s-%s",code,string(p[3:6]),string(p[6:]))

	return formatted,nil
}
