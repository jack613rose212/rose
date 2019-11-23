package study_models

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//只有models层是没有绑定接受者的
//都是以接受者的身份进来函数
//在models层实现的

//实现对切片的增删改查
func Go切片的增删改查(c *gin.Context) {
	var slice = []int{1, 99, 88, 77}
	//方式一:初始化增加或者指定下标进行增加/替换
	//slice[4] = 111  Todo  错误的写法  只有append才可以进行自动扩容
	slice[0] = 999 //999 99  88  77
	fmt.Println("初始化进行指定下标的添加:", slice)
	//方式二:append
	slice = append(slice, 888)              //追加
	fmt.Println("append添加888以后的元素:", slice) //999 99  88  77  888
	var ii int = 1                          //手动设置指定下标
	//删除指定下标1
	slice2 := append(slice[:ii], slice[ii+1:]...) //注意  多个元素要使用....
	fmt.Println(slice2)                           //output==> 999 88 77 888  删除了99
	//Todo 注意  slice2删除了底层的slice切片99
	//查找切片里的元素
	for i, v := range slice {
		if v == 88 {
			slice[i] = 888888 //output:999  888888   77  888
		}
	}
	fmt.Println("查找切片里面的元素:", slice)
}

//生产者
func Producter(in chan<- int) {
	for i := 1; i <= 5; i++ {
		in <- i
		fmt.Println("生产:", i)
	}
	close(in) //生产完毕以后关闭管道
}

//消费者
func Consumer(out <-chan int) {

	//两种方式判断管道是否关闭
	//1 if  vales,ok:=<-ch;ok==true{ }   关闭管道时候为false  一边写 一边读  读到数据是true
	//2  range  range操作是直到chanel关闭
	for num := range out {
		fmt.Println("消费:", num)
	}
}
