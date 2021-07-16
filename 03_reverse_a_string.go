package main

import "strings"

// ReserveStringR uses runes
// Supports strings that use more than one byte
// like chineese characters
func ReserveStringR(str string) string {
	var res string
	for _, r := range str {
		res = string(r) + res
	}
	return res
}

// ReserveStringB loops over the string from right to left
// and builds a new string in the process
func ReserveStringB(str string) string {
	var res string 
	for i := len(str)-1; i >=0; i-- {
		res = res + string(str[i])
	}
	return res
}

// ReserveStringB loops over the string from left to right
// and builds a new string in the process
func ReserveStringF(str string) string {
	var res string 
	for i := 0; i <= len(str)-1; i++ {
		res = string(str[i]) + res
	}
	return res
}

// ReserveStringSB uses the strings.Builder methods to
// create the new string
func ReserveStringSB(str string) string {
	var sb strings.Builder
	for i := len(str)-1; i >=0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}