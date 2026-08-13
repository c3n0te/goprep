package main

import (
	"errors"
	"strings"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = uint64(len(alphabet))
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Encode(id uint64) string {
	if id == 0 {
		return string(alphabet[0])
	}

	var builder strings.Builder
	for id > 0 {
		remainder := id % base
		builder.WriteByte(alphabet[remainder])
		id = id / base
	}

	return reverse(builder.String())
}

func Decode(slug string) (uint64, error) {
	var id uint64
	for i := 0; i < len(slug); i++ {
		pos := strings.IndexByte(alphabet, slug[i])
		if pos == -1 {
			return 0, errors.New("invalid character in base62 slug")
		}
		id = id*base + uint64(pos)
	}
	return id, nil
}
