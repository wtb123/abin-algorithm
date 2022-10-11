package main

import (
	"sort"
	"strconv"
)

func isDigit(b byte) bool { return '0' <= b && b <= '9' }
func isLower(b byte) bool { return 'a' <= b && b <= 'z' }

// 时间复杂度：
// 最坏情况下栈有O(n)层，每次出栈时需要更新O(n)个原子数量，因此遍历化学式时间复杂度为O(n^2)

// 空间复杂度：
// 空间复杂度取决于栈中所有哈希表中的元素个数之和，而这不会超过化学式 formula 的长度，因此空间复杂度为 O(n)

func countOfAtoms(formula string) string {
	i, n := 0, len(formula)
	stack := []map[string]int{{}}

	// 匿名函数，解析原子名称
	parseAtom := func() string {
		start := i
		i++
		for i < n && isLower(formula[i]) {
			i++
		}

		return formula[start:i]
	}

	// 匿名函数，解析原子个数
	parseNum := func() (num int) {
		if i == n || !isDigit(formula[i]) {
			return 1
		}

		for ; i < n && isDigit(formula[i]); i++ {
			num = num*10 + int(formula[i]-'0')
		}
		return
	}

	// 解析 formula 计算原子个数
	for i < n {
		switch formula[i] {
		case '(':
			i++
			stack = append(stack, map[string]int{})
		case ')':
			i++
			num := parseNum()              // 获取右括号右侧的数字
			atomMap := stack[len(stack)-1] // 保存栈 map
			stack = stack[:len(stack)-1]   // 栈顶元素退栈
			for atom, v := range atomMap { // 将栈顶 map 结算结果存入上一层 map
				stack[len(stack)-1][atom] += v * num
			}
		default:
			atom := parseAtom()
			num := parseNum()
			stack[len(stack)-1][atom] += num
		}
	}

	atomMap := stack[0] // 最后所有的结果都会保存到栈底
	atomArr := make([]string, 0, len(atomMap))
	for k, _ := range atomMap {
		atomArr = append(atomArr, k)
	}
	sort.Strings(atomArr)

	ans := ""
	for _, atom := range atomArr {
		ans += atom
		if num := atomMap[atom]; num > 1 {
			ans += strconv.Itoa(num)
		}
	}

	return ans
}
