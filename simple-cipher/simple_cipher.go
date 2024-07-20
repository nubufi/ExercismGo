package cipher

import (
	"regexp"
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return shift{(distance+26)%26}
}

func (c shift) Encode(input string) string {
	// Convert the input to a slice of runes
	input = strings.ToLower(input)
	encoded := []rune{}
	// Loop over the runes
	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			// Shift the rune by the distance
			e := 'a' + (r - 'a' + rune(c.distance)) % 26
			encoded = append(encoded, e)
		}
	}
	
	// Return the shifted runes as a string
	return string(encoded)
}

func (c shift) Decode(input string) string {
	runes := []rune(input)

	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = 'a' + (r - 'a' + rune(26 - c.distance)) % 26
		}
		if r >= 'A' && r <= 'Z' {
			runes[i] = 'A' + (r - 'A' + rune(26 - c.distance)) % 26
		}
	}

	return string(runes)
}

func NewVigenere(key string) Cipher {
	re := regexp.MustCompile("^a*$")
	if key == "" || strings.IndexFunc(key, func(r rune) bool { return r < 'a' || r > 'z' }) != -1  || re.MatchString(key){
		return nil
	}
	return vigenere{key}
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	re := regexp.MustCompile("[^a-zA-Z]+")
	input = re.ReplaceAllString(input, "")
	key := []rune(v.key)
	runes := []rune(input)
	encoded := []rune{}

	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			encoded = append(encoded, 'a' + (r - 'a' + key[i % len(key)] - 'a') % 26)
		}
	}

	return string(encoded)
}

func (v vigenere) Decode(input string) string {
	
	key := []rune(v.key)
	runes := []rune(input)

	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = 'a' + (r - 'a' + 26 - (key[i % len(key)] - 'a')) % 26
		}
		if r >= 'A' && r <= 'Z' {
			runes[i] = 'A' + (r - 'A' + 26 - (key[i % len(key)] - 'a')) % 26
		}
	}

	return string(runes)
}
