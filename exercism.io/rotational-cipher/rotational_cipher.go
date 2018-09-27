package rotationalcipher

import "bytes"

func RotationalCipher(str string, shift int) string {
	b := bytes.Buffer{}
	for i := 0; i < len(str); i++ {
		switch {
		case str[i] >= 'A' && str[i] <= 'Z':
			b.WriteByte(byte('A' + (int(str[i]-'A')+shift)%26))
		case str[i] >= 'a' && str[i] <= 'z':
			b.WriteByte(byte('a' + (int(str[i]-'a')+shift)%26))
		default:
			b.WriteByte(str[i])
		}
	}
	return b.String()
}
