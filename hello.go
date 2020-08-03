package main

import (
	"fmt" //导入包
)

func main() {
	// fmt.Println("hello,world")

	// const (
	// 	a = iota
	// 	b = iota
	// 	c = iota
	// )
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// //声明数组
	// var brr [3] int 
	// //初始化数组
	// brr[0]=0
	// brr[1]=1
	// brr[2]=2
	// fmt.Println(brr)

	
	//声明 & 初始化数组 
	// arr := [3]int {1,2,3}   // := 左侧变量不能是已经被声明过得，否则报编译器错误
	// fmt.Println(arr)
	
	// //声明 & 初始化数组（自动推断类型及长度）
	// arr := [...] int {3,4,5}
	// fmt.Println(arr)


	//多重赋值,不需要引入中间变量
	i := 2
	j := 3
	i,j = j,i
	fmt.Println(i," ",j)

}
