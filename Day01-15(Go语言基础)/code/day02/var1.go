package main

import "fmt"

func main() {
	/*
	变量：variable
	概念：一小块内存，用于存储数据，在程序运行过程中数值可以改变

	使用：
		step1：变量的声明，也叫定义
			第一种：var 变量名 数据类型
					变量名 = 赋值

					var 变量名 数据类型 = 赋值

			第二种：类型推断，省略数据类型
					var 变量名 = 赋值

			第三种：简短声明，省略var
					变量名 := 赋值

		step2：变量的访问，赋值和取值
				直接根据变量名访问


	go的特性：
		静态语言：强类型语言
			go，java，c++，c#。。

		动态语言：弱类型语言
			JavaScript，php，python，ruby。。
	 */
	 //第一种：定义变量，然后进行赋值
	 var num1 int
	 num1 = 30
	 fmt.Printf("num1的数值是：%d\n",num1)//num1的数值是：30
	 //写在一行
	 var num2 int = 15
	 fmt.Printf("num2的数值是：%d\n",num2)//num2的数值是：15

	 var num3 float64 = 13.2
	 fmt.Printf("num3的数值是：%f\n",num3)//num3的数值是：13.200000

	 //第二种：类型推断,不直接声明变量的类型，由编译器帮我们推断
	 var name = "王二狗"
	 fmt.Printf("类型是：%T，数值是:%s\n",name,name)//类型是：string，数值是:王二狗

	 //第三种，简短定义，也叫简短声明,只能在函数内使用
	 sum := 100
	 fmt.Println(sum)//100

	 //多个变量同时定义
	 var a, b, c int
	 a = 1
	 b = 2
	 c = 3
	 fmt.Println(a,b,c)//1 2 3

	 var m, n int = 100,200
	 fmt.Println(m,n)//100 200

	 var n1,f1,s1 = 100,3.14,"Go"
	 fmt.Println(n1,f1,s1)//100 3.14 Go

	 var (
	 	studentName = "李小花"
	 	age = 18
	 	sex = "女"
	 	school = "No.1 Middle School"
	 )
	 fmt.Printf("学生姓名：%s，年龄：%d，性别：%s, 学校：%s\n",studentName,age,sex,school)//学生姓名：李小花，年龄：18，性别：女, 学校：No.1 Middle School
}
