package util

import (
	"log"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		log.Printf("ERR: %v", err)
	}
}

func ToInt(value string) int{
	data, err := strconv.Atoi(value)
	CheckErr(err)
	return data
}