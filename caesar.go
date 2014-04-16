// Package caesar provides a simple way to encrypt and decrypt messages using the Caesar cipher.
// It works with UTF-8 strings but can only encrypt/decrypt the latin ASCII characters ([A..Z], [a..z]) in a string.
package caesar

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Encrypts a plaintext string by shifting each character with the provided key
func EncryptPlaintext(plaintext string, key int) string {
  return rotateText(plaintext, key)
}

// Decrypts ciphertext by reverse shifting each character in the ciphertext with the key provided
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
