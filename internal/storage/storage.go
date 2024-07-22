package storage

import "errors"

// Здесь хранится общая для всех реализаций storage информация
// Общий интерфейс
// Здесь не будет описан интерфейс, потому что он будет описан по месту использования

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url already exists")
)
