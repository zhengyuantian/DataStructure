package main

import "fmt"

//49 之前想拿每个单词的累计字节数当key,忘记不同字符相加可能也有相同结果
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		fmt.Println(cnt)
		mp[cnt] = append(mp[cnt], str)
		fmt.Println(mp[cnt])
	}
	ans := [][]string{}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
