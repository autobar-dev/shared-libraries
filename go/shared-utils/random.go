package sharedutils

import "math/rand"

var (
	LowercaseCharacterSet                = []rune("abcdefghijklmnopqrstuvwxyz")
	UppercaseCharacterSet                = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NumberCharacterSet                   = []rune("0123456789")
	LowercaseUppercaseCharacterSet       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LowercaseUppercaseNumberCharacterSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	UppercaseNumberCharacterSet          = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func GenerateRandomString(length int, character_set []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = character_set[rand.Intn(len(character_set))]
	}

	return string(b)
}
