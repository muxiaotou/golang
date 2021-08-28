package main

import (
	"fmt"
)

func jiaoji(s, s1 []byte) []byte {
	m := make(map[byte]byte)
	var s2 []byte
	for _, v := range s {
		m[v] = v
	}

	for _, v := range s1 {
		_, ok := m[v]
		if ok {
			s2 = append(s2, v)
		}
	}
	return s2
}

func chaji(s, s1 []byte) []byte {
	m := make(map[byte]byte)
	var s2 []byte
	for _, v := range s {
		m[v] = v
	}
	for _, v := range s1 {
		_, ok := m[v]
		if !ok {
			s2 = append(s2, v)
		}
	}
	return s2
}

func bingji(s, s1 []byte) []byte {
	m := make(map[byte]byte)
	for _, v := range s {
		m[v] = v
	}
	for _, v := range s1 {
		_, ok := m[v]
		if !ok {
			m[v] = v
		}
	}
	var tmp []byte
	for _, v := range m {
		tmp = append(tmp, v)
	}
	return tmp
}

func main() {
	s := "abcdefg"
	s1 := "abdf"
	fmt.Println(string(jiaoji([]byte(s), []byte(s1)))) //交集,长的字符串在前
	fmt.Println(string(chaji([]byte(s1), []byte(s))))  //差集，短的字符串在前
	fmt.Println(string(bingji([]byte(s1), []byte(s)))) //差集，短的字符串在前

}
