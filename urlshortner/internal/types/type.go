package types

import "math/rand"

type UrlStructure struct {
	Link string `json:"link"`
}

var LinkAndUniqueValue = make(map[string]string)

var tempV = "123"
var isItUniqueString = map[string]bool{tempV: true}

func UniqueStringGenaration() string {
	data := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	for {
		var udata [5]rune

		for i := range 5 {
			udata[i] = data[rand.Intn(len(data))]
		}

		// Store the newly generated string
		tempV = string(udata[:])

		// If already used, generate another
		if isItUniqueString[tempV] {
			continue
		}

		// Mark the value as used
		isItUniqueString[tempV] = true
		return tempV

	}
}
