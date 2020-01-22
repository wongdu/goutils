package main

import (
	"fmt"
	"sync"
	"time"
)

//<<Golang中闭包的理解>>

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	in := Increase()
	fmt.Println(in())
	fmt.Println(in())
}

func main0() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
		//引用了外部主循环中的变量，所以结果可能是5个5，也可能是4个5，一个4，这并不确定
		//下面这个结果是在win10电脑上测试的执行结果
		/*
			5
			4
			5
			5
			5
		*/
		// 调用下面的睡眠语句后，就能依次输出期望的结果，如果不想睡眠影响性能，可以传递参数到闭包中
		time.Sleep(time.Second)
		/*
			0
			1
			2
			3
			4
		*/
	}
	wg.Wait()
}

//Golang并发中的闭包，传递参数到闭包中，可以解决上面遇到的问题，

/*闭包是匿名函数与匿名函数所引用环境的组合。*/
/*匿名函数有动态创建的特性，该特性使得匿名函数不用通过参数传递的方式，
就可以直接引用外部的变量。匿名函数仅仅是存储了一个函数的返回值，
它同时存储了一个闭包的状态。*/

/*闭包作为函数返回值：匿名函数作为返回值，不如理解为闭包作为函数的返回值。
即闭包被返回赋予一个同类型的变量时，同时赋值的是整个闭包的状态，
该状态会一直存在外部被赋值的变量xxx中，直到xxx被销毁，整个闭包也被销毁。*/
func main1() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		// time.Sleep(time.Second)
	}
	wg.Wait()
}
