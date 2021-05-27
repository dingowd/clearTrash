package main

import (
	"fmt"
	"strconv"
	"strings"
)

func adding(s1 string, s2 string) int64 {
	check := map[string]string{"0":"0", "1":"1", "2":"2", "3":"3", "4":"4", "5":"5", "6":"6", "7":"7", "8":"8", "9":"9",}
	a1 := strings.Split(s1, "")
	b1 := strings.Split(s2, "")
	ar := ""
	br := ""
	for i:= range a1 {
		if value, InMap := check[a1[i]]; InMap {
			ar += value
		}
	}
	for i:= range b1 {
		if value, InMap := check[b1[i]]; InMap {
			br += value
		}
	}
	res1,_ := strconv.Atoi(ar)
	res2,_ := strconv.Atoi(br)
	return int64(res1+res2)
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}