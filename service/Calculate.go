package service

import (
	"Calculator/model"
	"strconv"
	"unicode"
)

// 中缀表达式转后缀表达式
func Convert(exp string) string {
	stack := model.Stack{}
	postfix := ""
	expLen := len(exp)

	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])

		switch char {
		case " ":
			continue

			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1

		default:
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for !stack.IsEmpty() {
				top := stack.Top()
				if top == "(" || IsLower(top, char) {
					break
				}
				postfix += top
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}

	// 栈不空则全部输出
	for !stack.IsEmpty() {
		postfix += stack.Pop()
	}

	return postfix
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func IsLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}

// 后缀表达式进行运算
func Calculate(postfix string) int {
	// 定义一个栈
	a := model.Stack{}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		if unicode.IsDigit(rune(postfix[i])) {
			a.Push(nextChar)
		} else if len(a.Data) < 2 {
			continue
		} else {
			// 操作符：取出两个数字计算值，再将结果压栈
			//a, _ = stack.Pop()
			//num1,_:= strconv.Atoi(a)
			num1, _ := strconv.Atoi(a.Pop())
			num2, _ := strconv.Atoi(a.Pop())

			switch nextChar {
			case "+":
				a.Push(strconv.Itoa(num1 + num2))
			case "-":
				a.Push(strconv.Itoa(num2 - num1))
			case "*":
				a.Push(strconv.Itoa(num1 * num2))
			case "/":
				a.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	result, _ := strconv.Atoi(a.Top())
	return result
}
