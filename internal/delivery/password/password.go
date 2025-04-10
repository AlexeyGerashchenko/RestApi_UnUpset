// Package password предоставляет функции для безопасной работы с паролями
package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword хеширует пароль с использованием алгоритма bcrypt
// Принимает строку пароля и возвращает его хеш и возможную ошибку
// Используется уровень стоимости (cost) 12, что обеспечивает хороший баланс
// между безопасностью и производительностью
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckPassword проверяет, соответствует ли пароль хешу
// Принимает строку пароля и его хеш, возвращает true, если пароль соответствует хешу
// Возвращает false, если пароль не соответствует хешу или произошла ошибка
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
