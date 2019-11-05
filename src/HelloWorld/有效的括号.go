package main

import (
	"regexp"
	"strconv"
)

// 有效的括号
func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack []byte = make([]byte, len(s))
	j := 0
	for i := 0; i < len(s); i++ {

		if s[i] == '(' || s[i] == '{' || s[i] == '[' {

			stack[j] = s[i]
			j++
			continue
		}

		if j > 0 && stack[j-1] == '(' && s[i] == ')' {
			j--
			continue
		}
		if j > 0 && stack[j-1] == '{' && s[i] == '}' {
			j--
			continue
		}
		if j > 0 && stack[j-1] == '[' && s[i] == ']' {
			j--
			continue
		}

		stack[j] = s[i]
		j++

	}
	if j == 0 {
		return true
	} else {
		return false
	}
}

// 最长有效括号
func longestValidParentheses(s string) int {
	var stack []int = make([]int, len(s))
	var j = 0
	var sum = 0
	var max = 0
	var temp = 0
	var i = 0
	var flag = false
	var count = 0

	for i < len(s) {

		if s[i] == '(' {
			stack[j] = 0
			j++
			i++
			continue
		}

		for j > 0 && stack[j-1] != 0 {
			if j == 1 {
				break
			}
			// 取栈顶元素
			temp = stack[j-1]
			// 栈顶指针减去1
			if j > 1 {
				j--
			}
			// 这里栈中元素为0，并且j>0，表示是碰到'('这个在栈里面，一旦碰到这个表示，不能再继续累加连续的有效
			// 对括号
			if stack[j-1] == 0 {

				break
			}
			// 否则将栈中连续的数相加，就是有效括号的对数
			stack[j-1] = stack[j-1] + temp
			stack[j] = 0

		}
		// 这里就是当栈中只有一个数字时，此时又来了一个')',这时应该记录当前栈中有效括号的对数，与当前最大值比较，如果
		// 大于当前最大值，则有效的括号对数置为当前值，然后栈清空，有效括号对数重新计数
		if j == 1 && s[i] == ')' && stack[j-1] != 0 {
			sum = stack[j-1]
			if sum >= max {
				max = sum

			}
			j--
			stack[j] = 0
			sum = 0
		}
		temp = 0
		count = 0
		for j > 0 && stack[j-1] == 0 && i < len(s) && s[i] == ')' {

			if stack[j] != 0 && count == 0 {
				// 这个当栈中的元素为 0，4 ，并且 j指向栈顶元素4时，本来栈指针总是指向栈顶元素的下一位，但是通过
				// 上面连续有效括号对数循环累加，会出现这样的情况，当栈是这样的情况时，恰好s[i]==')'
				// 这个时候相当于这样的一个例子（（）（））
				temp = stack[j]
			}
			count++

			stack[j-1] = temp + 2
			temp = stack[j-1]

			if count > 1 {
				// 处理栈指针的下移，只有当进行当前循环第二次时，才需要下移，因为如果进行一次循环，j无需下移
				j--
				// 当j是指向栈顶的下一位时，将栈顶置为0
				stack[j+1] = 0
			} else {
				// 如果j指向栈顶当前元素，将栈顶置为0
				stack[j] = 0
			}

			i++
			if !flag {
				flag = true
			}

		}

		if flag {
			flag = false
			continue
		}
		i++

	}
	for w := len(stack) - 1; w >= 0; w-- {
		for w >= 0 && stack[w] == 0 {
			w--
		}
		for w >= 0 && stack[w] != 0 {
			sum = sum + stack[w]
			w--

		}
		if sum >= max {
			max = sum

		}
		sum = 0
	}

	return max
}

// 逆波兰表达式求值
func evalRPN(tokens []string) int {
	nums := make([]int, 3)
	var j = 0
	for i := 0; i < len(tokens); i++ {
		match, _ := regexp.MatchString("^-?\\d+$", tokens[i])
		if match {
			if j == len(nums) {
				// 如果栈满，动态扩容
				nums = append(nums, 1)
			}
			nums[j], _ = strconv.Atoi(tokens[i])
			j++
			continue
		}
		switch tokens[i] {
		case "+":
			if j > 1 {
				nums[j-2] = nums[j-2] + nums[j-1]
				j--
			}

		case "-":
			{
				if j == 1 {
					nums[j-1] = -nums[j-1]
				} else {
					nums[j-2] = nums[j-2] - nums[j-1]
					j--
				}

			}
		case "*":
			nums[j-2] = nums[j-2] * nums[j-1]
			j--

		case "/":
			nums[j-2] = nums[j-2] / nums[j-1]
			j--

		default:

		}

	}
	return nums[0]
}

// 求k个数的最大值
func maxNum(nums []int, i int, j int) int {
	const INT_MAX = int(^uint(0) >> 1)
	var max = ^INT_MAX

	for k := i; k <= j; k++ {
		if nums[k] >= max {
			max = nums[k]
		}
	}
	return max
}

//滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		result := make([]int, 0)
		return result
	}
	result := make([]int, 1)
	var j = 0
	var w = 0
	var i int
	var max = maxNum(nums, 0, k-1)
	result[0] = max
	for i = k; i < len(nums); i++ {
		w = i - k
		result = append(result, 1)
		if nums[w] != max && nums[i] < max {
			j++
			result[j] = max

			continue
		}

		if nums[i] >= max {
			j++
			result[j] = nums[i]
			max = nums[i]

			continue

		}
		if nums[w] == max && nums[i] < max {
			max = maxNum(nums, w+1, i)
			j++
			result[j] = max

		}
	}
	return result
}

// 爬楼梯
func climbStairs(n int) int {

	result := make([]int, 2)

	result[0] = 1
	result[1] = 2

	if n > 2 {
		for i := 2; i < n; i++ {
			result = append(result, 1)
			result[i] = result[i-1] + result[i-2]
		}

	}
	return result[n-1]
}

// func main() {
// 	print(climbStairs(45))
// 	// var nums []int = []int{}
// 	// // int k =3
// 	// result := maxSlidingWindow(nums, 0)
// 	// for i := 0; i < len(result); i++ {
// 	// 	print(result[i])
// 	// }

// }
