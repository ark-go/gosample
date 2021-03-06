package internal

// взято тут https://stackru.com/questions/45347682/golang-shifrovanie-stroki-s-pomoschyu-aes-i-base64
// смотреть еще  https://gist.github.com/kkirsche/e28da6754c39d5e7ea10
// https://gist.github.com/manishtpatel/8222606
//
import (
	"crypto/aes"
	"crypto/cipher"
	_ "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	_ "io"
	"log"
	mrand "math/rand"
	"time"
)

func CriptoAes256() {

	//--
	mrand.Seed(time.Now().UnixNano())
	keyPrefix := []byte("ПреФиксКоде")
	keyMain := make([]byte, 32-len(keyPrefix))

	mrand.Read(keyMain) // заполняем весь []byte,  никогда не вернет ошибку (documentation)
	key := append(keyPrefix, keyMain...)
	//--
	//key := []byte("a very very very very secret key") // 32 bytes
	plaintext := []byte("Эту строку я пытаюсь зашифровать и посмотреть что будет или не будет")
	fmt.Printf("Шифруем строку: %s\n", plaintext)
	ciphertext, err := encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0x\n", ciphertext)
	result, err := decrypt(key, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Расшифруем ее: %s\n", result)
}

// See alternate IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) // 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize] // вектор инициализации для XOR
	println("aes.BlockSize", aes.BlockSize, len(key))
	// if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 	return nil, err
	// }
	mrand.Seed(time.Now().UnixNano())
	mrand.Read(iv) // я так понимаю тут вписываем вначало ciphertext
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b)) // а тут подаем конец ciphertext
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("Слишком маленький текст")
	}
	iv := text[:aes.BlockSize]  // вытаскиваем вектор инициализации
	text = text[aes.BlockSize:] // за ним шифрованный текст
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
