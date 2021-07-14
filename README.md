# 第二题计算器技术文档

### 1.整体框架

根据用户的请求的参数字符串在浏览器输出表达式的结果，对核心运算函数进行单元测试。

### 2.目录结构

![image-20210714170728456](/Users/fuyuanfu/Library/Application Support/typora-user-images/image-20210714170728456.png)

main.go  --代码入口

post.go --转发路由信息

Calculate.go --计算器的中缀转后缀和后缀计算算法

Getroute.go --接受路由和网页参数

Calculate_test.go --计算器中缀转后缀和后缀计算算法测试方法

Testpressure.py 测试脚本

### 3.代码逻辑分层

程序入口：app/main.go

应用层：/service/Calculate.go

路由层：route/Getroute.go

控制层：ctrl/post.go

模型层：model/Stack.go



### 4.接口设计

供客户端调用的接口

请求方法：http post

请求地址：localhost:8080/calculate

请求参数

string类型

Name:"5+8-9"

一，用户请求

 

登录

 

1，http请求方式：post (使用https协议)

2，  [https://localhost:8080/字符串表达式](https://localhost:8080/表达式)

| 参数  | 是否必须 | 说明                   | 类型   |
| ----- | -------- | ---------------------- | ------ |
| input | 是       | 用户输入的字符串表达式 | string |

 

返回的数据

| 参数   | 是否必须 | 说明                               | 类型 |
| ------ | -------- | ---------------------------------- | ---- |
| output | 是       | 根据用户输入字符串输出相对应的结果 | int  |

### 

供客户端调用的接口

请求方法：http post

请求地址：localhost:8080/calculate

请求参数

string类型

Name:"5+8-9" 





| 方法名                                     | 说明                                          | 返回结果                                     | 形式参数                 |
| ------------------------------------------ | --------------------------------------------- | -------------------------------------------- | ------------------------ |
| Func convert (exe string) string           | 中缀表达式转为后缀表达式                      | 一个string类型的后缀表达式变量               | string类型               |
| Func IsLower(top string,newTop string)bool | 比较运算符栈顶top和新运算符newTop的优先级高低 | 一个bool类型变量                             | top string,newTop string |
| Func calculate(postfix string) int         | 后缀表达式进行计算                            | 一个int类型变量                              | postfix string           |
| func(stack *Stack)push (element string)    | 将字符入栈                                    | 无                                           | element string           |
| func(stack*Stack)Pop()string               | 将栈中字符出栈                                | 一个string 字符                              | 无                       |
| Func(stack *Stack) Top() string            | 取栈顶字符                                    | 一个string字符                               | 无                       |
| func(stack *Stack) IsEmpty()bool           | 判断栈是否为空                                | bool                                         | 无                       |
| Func Test_Calculate(t *testing.T)          | 单元测试函数                                  | 控制台输出Pass或者test failed                | 无                       |
| func Getroute()                            | 转发路由                                      | 将路由收到后转发给控制层，并将数据在网页渲染 | 无                       |
| Func Getstring()                           | 接受路由和参数，返回数据                      | 将数据返回给路由层                           | 无                       |

 

 



### 5.第三方库

```
"github.com/gin-gonic/gin"  --用于构建gin框架
```

```
"strconv"   --用于调用strconv.Iota和strconv.Atoi方法将字符转化为数字和将数字转化为字符
```

```
"testing"  --用于单元测试
```

### 6.编译执行

执行app.main.go文件即可

### ![image-20210714183446462](/Users/fuyuanfu/Library/Application Support/typora-user-images/image-20210714183446462.png)

