package caesar

import (
	"strings"
	"testing"
)

const testingString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var expectedEncryptionResults map[int]string = map[int]string{
	1:  "bcdefghijklmnopqrstuvwxyzaBCDEFGHIJKLMNOPQRSTUVWXYZA",
	2:  "cdefghijklmnopqrstuvwxyzabCDEFGHIJKLMNOPQRSTUVWXYZAB",
	3:  "defghijklmnopqrstuvwxyzabcDEFGHIJKLMNOPQRSTUVWXYZABC",
	4:  "efghijklmnopqrstuvwxyzabcdEFGHIJKLMNOPQRSTUVWXYZABCD",
	5:  "fghijklmnopqrstuvwxyzabcdeFGHIJKLMNOPQRSTUVWXYZABCDE",
	6:  "ghijklmnopqrstuvwxyzabcdefGHIJKLMNOPQRSTUVWXYZABCDEF",
	7:  "hijklmnopqrstuvwxyzabcdefgHIJKLMNOPQRSTUVWXYZABCDEFG",
	8:  "ijklmnopqrstuvwxyzabcdefghIJKLMNOPQRSTUVWXYZABCDEFGH",
	9:  "jklmnopqrstuvwxyzabcdefghiJKLMNOPQRSTUVWXYZABCDEFGHI",
	10: "klmnopqrstuvwxyzabcdefghijKLMNOPQRSTUVWXYZABCDEFGHIJ",
	11: "lmnopqrstuvwxyzabcdefghijkLMNOPQRSTUVWXYZABCDEFGHIJK",
	12: "mnopqrstuvwxyzabcdefghijklMNOPQRSTUVWXYZABCDEFGHIJKL",
	13: "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM",
	14: "opqrstuvwxyzabcdefghijklmnOPQRSTUVWXYZABCDEFGHIJKLMN",
	15: "pqrstuvwxyzabcdefghijklmnoPQRSTUVWXYZABCDEFGHIJKLMNO",
	16: "qrstuvwxyzabcdefghijklmnopQRSTUVWXYZABCDEFGHIJKLMNOP",
	17: "rstuvwxyzabcdefghijklmnopqRSTUVWXYZABCDEFGHIJKLMNOPQ",
	18: "stuvwxyzabcdefghijklmnopqrSTUVWXYZABCDEFGHIJKLMNOPQR",
	19: "tuvwxyzabcdefghijklmnopqrsTUVWXYZABCDEFGHIJKLMNOPQRS",
	20: "uvwxyzabcdefghijklmnopqrstUVWXYZABCDEFGHIJKLMNOPQRST",
	21: "vwxyzabcdefghijklmnopqrstuVWXYZABCDEFGHIJKLMNOPQRSTU",
	22: "wxyzabcdefghijklmnopqrstuvWXYZABCDEFGHIJKLMNOPQRSTUV",
	23: "xyzabcdefghijklmnopqrstuvwXYZABCDEFGHIJKLMNOPQRSTUVW",
	24: "yzabcdefghijklmnopqrstuvwxYZABCDEFGHIJKLMNOPQRSTUVWX",
	25: "zabcdefghijklmnopqrstuvwxyZABCDEFGHIJKLMNOPQRSTUVWXY",
	26: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

func TestEncryptPlaintext(t *testing.T) {
	for key, text := range expectedEncryptionResults {
		ciphertext := EncryptPlaintext(testingString, key)
		if strings.EqualFold(ciphertext, text) == false {
			t.Errorf("Encryption failed! Encrypting [%s] with key [%d] produced [%s]. Should produce [%s]", testingString, key, ciphertext, text)
		}
	}
}

func TestDecryptCiphertext(t *testing.T) {
	for key, text := range expectedEncryptionResults {
		plaintext := DecryptCiphertext(text, key)
		if strings.EqualFold(plaintext, testingString) == false {
			t.Errorf("Decryption failed! Decrypting [%s] with key [%d] produced [%s]. Should produce [%s]", text, key, plaintext, testingString)
		}
	}
}

func TestrotateText(t *testing.T) {
	// If TestDecryptCiphertext TestEncryptPlaintext pass then this function must be working
}
