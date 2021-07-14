package main

import "Calculator/route"

func main() {
	// 实现思路：先将中缀表达式转化为后缀表达式，然后再对后缀表达式进行计算即可。

	//验证方法：在postman中输入url：http://localhost:8080/calculate key设置为name  在value中输入想要求的表达式字符串即可。

	route.Getroute()
}
