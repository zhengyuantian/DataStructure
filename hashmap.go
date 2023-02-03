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
		mp[cnt] = append(mp[cnt], str)
		fmt.Println(mp[cnt])
	}
	ans := [][]string{}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

//76 slideWindow map

func minWindow(s string, t string) string {
	l, r, v := 0, 0, 0
	n, w := map[byte]int{}, map[byte]int{}
	st, L := 0, len(s)+1
	for i, _ := range t {
		n[t[i]]++
	}
	for r < len(s) {
		c := s[r]
		r++
		if n[c] > 0 {
			w[c]++
			if w[c] == n[c] {
				v++
			}
		}
		for len(n) == v {
			if r-l < L {
				st = l
				L = r - l
			}
			d := s[l]
			l++
			if n[d] > 0 {
				if w[d] == n[d] {
					v--
				}
				w[d]--
			}
		}
	}
	if L == len(s)+1 {
		return ""
	}
	return s[st : st+L]
}
