//+build ignore
package main

import "fmt"

//解题思路 ,暴力破解，逐个对比

func longsubstring(s []string) (res string) {
	res = s[0]                    //假定第一个item是最长子串
	for i := 1; i < len(s); i++ { //遍历slice当中从index为1开始的的所有item
		for j := 0; j < len(res) && j < len(s[i]); j++ {
			if res[j] == s[i][j] {
				continue
			} else {
				res = res[:j]
				break
			}
		}
	}
	return
}

func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longsubstring(s), len(longsubstring(s)))

	s = []string{"dog", "racecar", "car"}
	fmt.Println(longsubstring(s), len(longsubstring(s)))
}
