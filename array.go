package main

// 34
func searchRange(nums []int, target int) []int {
	// 寻找开头的第一个target位置
	first := FindFirst(nums, target)
	// 如果长度为0或者数组中没有目标值，返回默认值
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	// 寻找开头的第一个target+1的第一个位置,并-1
	last := FindFirst(nums, target+1) - 1
	return []int{first, last}
}

func FindFirst(nums []int, target int) int {
	// 初始化索引
	left, right := 0, len(nums)-1
	for left <= right {
		// 更新中心值
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

//209 double pointer
func minSubArrayLen(target int, nums []int) int {
	i, j, minl, sum := -1, -1, len(nums)+1, 0
	for j >= i {
		if sum < target {
			j++
			if j > len(nums)-1 {
				break
			}
			sum += nums[j]
		} else {
			i++
			if j-i+1 < minl {
				minl = j - i + 1
			}
			sum -= nums[i]
		}
	}
	if minl == len(nums)+1 {
		return 0
	}
	return minl
}

//904 hashmap
func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	i, ans := 0, 0
	for j, v := range fruits {
		cnt[v]++
		for len(cnt) > 2 {
			x := fruits[i]
			cnt[x]--
			i++
			if cnt[x] == 0 {
				delete(cnt, x)
			}
		}
		ans = max(ans, j-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type D struct{ x, y int }

func SpiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	if n == 0 || m == 0 {
		return []int{}
	}
	ans := make([]int, n*m)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	dirs := []D{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, d := 0, 0, 0
	for i := 0; i < n*m; i++ {
		ans[i] = matrix[row][col]
		dir := dirs[d]
		visited[row][col] = true
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= m || visited[r][c] {
			d = (d + 1) % 4
			dir = dirs[d]
		}
		row += dir.x
		col += dir.y
	}
	return ans
}
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	dirs := []D{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, d := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		ans[row][col] = i
		dir := dirs[d]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || ans[r][c] > 0 {
			d = (d + 1) % 4
			dir = dirs[d]
		}
		row += dir.x
		col += dir.y
	}
	return ans
}
