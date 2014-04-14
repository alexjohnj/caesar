package caesar

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncryptPlaintext(plaintext string, key int) string {
	return rotateText(plaintext, key)
}

func DecryptCiphertext(ciphertext string, key int) string {
	return rotateText(ciphertext, -key)
}

func rotateText(inputText string, rot int) string {
	rot %= 26
	rotatedText := []byte(inputText)

	for index, byteValue := range rotatedText {
		if byteValue >= 'a' && byteValue <= 'z' {
			rotatedText[index] = lowerCaseAlphabet[(int((26+(byteValue-'a')))+rot)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			rotatedText[index] = upperCaseAlphabet[(int((26+(byteValue-'A')))+rot)%26]
		}
	}
	return string(rotatedText)
}
