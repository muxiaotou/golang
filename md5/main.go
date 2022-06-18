package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func FileMD5(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file fail")
		return ""
	}

	defer file.Close()

	r := bufio.NewReader(file)
	hash, buf := md5.New(), make([]byte, 1<<20)
	for {
		n, err := r.Read(buf)
		//pos, _ := file.Seek(0, os.SEEK_CUR)
		//fmt.Println("current offset is ", pos)
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			return ""
		}
		io.Copy(hash, bytes.NewReader(buf[:n])) //io.Copy效率较高
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	var path string
	if len(os.Args) != 2 {
		fmt.Println("please give path to file/disk")
		os.Exit(1)
	}
	path = os.Args[1]
	fmt.Println(path, " md5 is ", FileMD5(path))
}
