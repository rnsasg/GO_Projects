package utils

import (
	"math/rand"
)

var entityTypes = []string{"device", "router", "switch", "l2service", "l3service", "mobile", "computer", "tablet", "desktop"}

func RandomEntityType() string {
	return entityTypes[rand.Intn(len(entityTypes))]
}

func RandomString(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, length)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}
