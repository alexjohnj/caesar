# caesar.go

Package caesar provides a simple way of encrypting and decrypting messages using the Caesar cipher. It works with UTF-8 strings although it only encrypts/decrypts characters in the English alphabet.

## Usage

caesar provides two functions to work with:

``` go
func EncryptPlaintext(plaintext string, key int) string
```

``` go
func DecryptCiphertext(ciphertext string, key int) string
```

It's fairly obvious how to use them but just in case, here's an example:


``` go
ciphertext := caesar.EncryptPlaintext("Romans, go home.", 13)
fmt.Println(ciphertext)
```

Output:

```
Ebznaf, tb ubzr.
```

## Why Does This Exist?

Obviously this package has few practical applications. I wrote it to learn about making, documenting and publishing packages in golang. I think it was a fairly good learning exercise.
