// String Encode and Decode
// Design an algorithm to encode a list of strings to a single string.
// The encoded string is then decoded back to the original list of strings.

// Please implement encode and decode

// Input: ["neet","code","love","you"]

// Output:["neet","code","love","you"]

// Input: ["we","say",":","yes"]

// Output: ["we","say",":","yes"]

// 0 <= strs.length < 100
// 0 <= strs[i].length < 200
// strs[i] contains only UTF-8 characters.

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func encode(strs []string) string {
	var str strings.Builder

	for _, s := range strs {
		str.WriteString(fmt.Sprintf("%02d", len(s)))
		str.WriteString(s)
	}

	return str.String()
}
	
func decode(str string) []string {
	strs := make([]string, 0)

	i := 0
	for i < len(str) {
		lenSubStr, _ := strconv.Atoi(str[i: i+2])
		i += 2
		strs = append(strs, str[i: i+lenSubStr])
		i += lenSubStr
	}

	return strs
}

func main() {
	testCases := []struct{
		strs []string
		result []string
	}{
		{[]string{"12","code","love","12youadfasdfasd2133&"}, []string{"12","code","love","12youadfasdfasd2133&"}},
		{ []string{"neet","code","love","you"}, []string{"neet","code","love","you"}},
		{ []string{"12", "02", "ias%&^$"}, []string{"12", "02", "ias%&^$"}},
		{ []string{"we","say",":","yes"}, []string{"we","say",":","yes"}},
	}

	for _, testCase := range testCases {
		result := decode(encode(testCase.strs))
		if !reflect.DeepEqual(result, testCase.result) {
			fmt.Printf("Test failed, Expected: %v, Got: %v\n", testCase.result, result)
		}else {
			fmt.Printf("Test passed, testCase: %v\n", testCase.strs)
		}
	}
}