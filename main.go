package main

import (
	"encoding/hex"

	"github.com/elizarpif/camelia/camelia"
)

func fromHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func main() {
	_, err := camelia.NewCipher(fromHex("0123456789abcdeffedcba9876543210"))
	if err != nil {
		panic(err)
	}

	plaintext := fromHex("0123456789abcdeffedcba9876543210")
	ciphertext := fromHex("67673138549669730857065648eabe43")


}

/*
Алгоритм необходимо реализовать в оконном или Web-приложении.
(Де-)Шифрование должно производиться асинхронно (не допускается отвисание UI-потока) и многопоточно(для увеличения скорости,если это возможно).
Необходимо реализовать режимы шифрования: электронной кодовой книги(ECB),
сцепления блоков шифротекста(CBC),
обратной связи по шифротексту(CFB),
обратной связи по выходу(OFB).
Предусмотреть возможности ввода ключа шифрования и вектора инициализации для режимов шифрования из файла.
Приложение должно уметь шифровать любые файлы, также необходимо реализовать отображение прогресса операции шифрованияи(опционально)
отмену операции шифрования по запросу пользователя.
*/
