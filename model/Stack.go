package model

type Stack struct {
	Data []string
}

// 入栈
func (stack *Stack) Push(element string) {
	stack.Data = append(stack.Data, element)

}

// 出栈
func (stack *Stack) Pop() string {
	v := stack.Data[(len(stack.Data) - 1)]
	stack.Data = stack.Data[:len(stack.Data)-1]
	return v
}

// 获取栈顶元素
func (stack *Stack) Top() string {
	return stack.Data[len(stack.Data)-1]
}

// 判断栈是否为空
func (stack *Stack) IsEmpty() bool {
	if len(stack.Data) == 0 {
		return true
	}
	return false
}

//type Stack struct {
//	values    []string
//	valueType reflect.Type
//}
//
//
//
//// 判断值是否符合栈类型
//func (stack *Stack) isAcceptableValue(value string) bool {
//	if  reflect.TypeOf(value) != stack.valueType {
//		return false
//	}
//	return true
//}
//
//// 入栈
//func (stack *Stack) Push(v string) bool {
//	if !stack.isAcceptableValue(v) {
//		return false
//	}
//	stack.values = append(stack.values, v)
//	return true
//}
//
//// 出栈
//func (stack *Stack) Pop() (string, error) {
//	if  len(stack.values) == 0 {
//		return nil, errors.New("stack empty")
//	}
//	v := stack.values[len(stack.values)-1]
//	stack.values = stack.values[:len(stack.values)-1]
//	return v, nil
//}
//
//// 获取栈顶元素
//func (stack *Stack) Top() (string, error) {
//	if stack == nil || len(stack.values) == 0 {
//		return nil, errors.New("stack empty")
//	}
//	return stack.values[len(stack.values)-1], nil
//}
//
//// 获取栈内元素个数
//func (stack *Stack) Len() int {
//	return len(stack.values)
//}
//
//// 判断栈是否为空
//func (stack *Stack) Empty() bool {
//	if stack == nil || len(stack.values) == 0 {
//		return true
//	}
//	return false
//}
//
//// 获取栈内元素类型
//func (stack *Stack) ValueType() reflect.Type {
//	return stack.valueType
//}
