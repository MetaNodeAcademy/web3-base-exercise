package task

import (
	"fmt"
	"strings"
)

// 任务一: 回文数、只出现一次的数字
func TaskOne() {
	fmt.Println("--------任务一")
	/*
		回文数
		给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
		回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
		例如，121 是回文，而 123 不是。

		示例 1：
		输入：x = 121
		输出：true

		示例 2：
		输入：x = -121
		输出：false
		解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

		示例 3：
		输入：x = 10
		输出：false
		解释：从右向左读, 为 01 。因此它不是一个回文数。
	*/
	// a := isPalindrome(121)
	// b := isPalindrome(-121)
	// c := isPalindrome(10)
	// println(a, b, c)

	/*
		只出现一次的数字
		给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
		你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

		示例 1 ：
		输入：nums = [2,2,1]
		输出：1

		示例 2 ：
		输入：nums = [4,1,2,1,2]
		输出：4

		示例 3 ：
		输入：nums = [1]
		输出：1
	*/
	// nums := []int{2, 2, 1}
	// d := singleNumber(nums)
	// fmt.Println(d)
	// nums = []int{4, 1, 2, 1, 2}
	// e := singleNumber(nums)
	// fmt.Println(e)
	// nums = []int{1}
	// f := singleNumber(nums)
	// fmt.Println(f)

	/*
		有效的括号
		给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
		有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		每个右括号都有一个对应的相同类型的左括号。

		示例 1：
		输入：s = "()"
		输出：true

		示例 2：
		输入：s = "()[]{}"
		输出：true

		示例 3：
		输入：s = "(]"
		输出：false

		示例 4：
		输入：s = "([])"
		输出：true

		示例 5：
		输入：s = "([)]"
		输出：false
	*/
	// a := isValid("()")
	// b := isValid("()[]{}")
	// c := isValid("(]")
	// d := isValid("([])")
	// e := isValid("([)]")
	// fmt.Println(a, b, c, d, e)

	/*
		编写一个函数来查找字符串数组中的最长公共前缀。
		如果不存在公共前缀，返回空字符串 ""。

		示例 1：
		输入：strs = ["flower","flow","flight"]
		输出："fl"

		示例 2：
		输入：strs = ["dog","racecar","car"]
		输出：""
		解释：输入不存在公共前缀。
	*/
	// a := longestCommonPrefix([]string{"flower", "flow", "flight"})
	// b := longestCommonPrefix([]string{"dog", "racecar", "car"})
	// fmt.Println(a, b)

	/*
		给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
		将大整数加 1，并返回结果的数字数组。

		示例 1：
		输入：digits = [1,2,3]
		输出：[1,2,4]
		解释：输入数组表示数字 123。
		加 1 后得到 123 + 1 = 124。
		因此，结果应该是 [1,2,4]。

		示例 2：
		输入：digits = [4,3,2,1]
		输出：[4,3,2,2]
		解释：输入数组表示数字 4321。
		加 1 后得到 4321 + 1 = 4322。
		因此，结果应该是 [4,3,2,2]。

		示例 3：
		输入：digits = [9]
		输出：[1,0]
		解释：输入数组表示数字 9。
		加 1 得到了 9 + 1 = 10。
		因此，结果应该是 [1,0]。
	*/
	// a := plusOne([]int{1, 2, 3})
	// b := plusOne([]int{4, 3, 2, 1})
	// c := plusOne([]int{9})
	// fmt.Println(a, b, c)

	/*
		给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
		你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
		你可以按任意顺序返回答案。
		示例 1：
		输入：nums = [2,7,11,15], target = 9
		输出：[0,1]
		解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

		示例 2：
		输入：nums = [3,2,4], target = 6
		输出：[1,2]

		示例 3：
		输入：nums = [3,3], target = 6
		输出：[0,1]
	*/
	a := twoSum([]int{2, 7, 11, 15}, 9)
	b := twoSum([]int{3, 2, 4}, 6)
	c := twoSum([]int{3, 3}, 6)
	fmt.Println(a, b, c)
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}
	return x == reverted || x == reverted/10
}

// 只出现一次的数字 给你一个 非空 整数数组 nums ，
// 除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
func singleNumber(nums []int) int {
	// 数组只有一个元素
	if len(nums) == 1 {
		return nums[0]
	}
	// 对数组里的元素添加进map
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。有效字符串需满足：左括号必须用相同类型的右括号闭合。左括号必须以正确的顺序闭合。每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。将大整数加 1，并返回结果的数字数组。
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 创建一个长度为n+1的切片，并初始化为0
	res := make([]int, n+1)
	res[0] = 1
	return res
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	if len(nums) == 2 {
		return []int{0, 1}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
