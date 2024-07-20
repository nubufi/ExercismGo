package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey = shiftKey % 26
	shifted := make([]byte, len(plain))
	for i := 0; i < len(plain); i++ {
		if plain[i] >= 'a' && plain[i] <= 'z' {
			shifted[i] = 'a' + (plain[i]-'a'+byte(shiftKey))%26
		} else if plain[i] >= 'A' && plain[i] <= 'Z' {
			shifted[i] = 'A' + (plain[i]-'A'+byte(shiftKey))%26
		} else {
			shifted[i] = plain[i]
		}
	}
	return string(shifted)
}
