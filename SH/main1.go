package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//声明全局变量  可以导出使用  注意要大写
var Cond sync.Cond //条件变量

//等待   惊群   加锁   解锁
//加锁  解锁===>互斥锁  读写锁  都实现了locker接口

//同步的死锁问题
func main() {
	//同步  容易死锁   需要加锁
	//异步操作 判定条件变量

	//互斥锁  并发的时候 谁抢到 谁先执行  针对的是两个单独的go程  防止出现资源的争抢
	//读写锁  大量读的时候效率低下 是两把锁  一个Lock是写锁  UnLock解锁   一个RLock是读锁  解锁是RUnLock

	//死锁原因  资源争抢
	//数据通信引起的
	//var mutex  *sync.RWMutex

	//mutex.Lock()
	//defer   mutex.Unlock()

	//C := make(chan int)
	//C <- 100
	//<-C
	//time.Sleep(time.Second * 2)
	//fmt.Println(<-C)

	rand.Seed(time.Now().UnixNano()) //种下随机种子

	//Locker接口  互斥锁

	//无效的内存地址或者空指针的引用  把锁的地址赋值过去
	Cond.L = new(sync.Mutex) //new的是一个值类型  返回对这个类型的引用

	c := make(chan int, 5) //开启一个管道  channel

	for i := 1; i <= 5; i++ { //为什么会多呢
		go Product(c) //生产者
		go Consumer(c)

	}
	//close(c)
	//for i:=1;i<=5 ;i++  {
	//	  //开启五个协程的写法
	//}
	//time.Sleep(time.Second * 10000)
	for {
		;
	}

	//select {}

	//go sing() //go程 go dance() //主程序 for { //死循环 ; } }
	//go dance()
	//for {
	//	;
	//}
}

//生产者
func Product(in chan<- int) {
	//管道里的数据要循环取  循环进行读
	for {
		//先加锁
		Cond.L.Lock()

		//判断是否是写满的
		if len(in) == 5 {
			//Cond.Signal() //唤醒一个  唤醒谁并不清楚
			Cond.Broadcast()//唤醒一群 唤醒谁 并不知道
			Cond.Wait() //等待   (阻塞-->解锁)
		}

		//写入
		num := rand.Intn(100) //[0,n)的伪随机  不包括100  rand.Ind  随机整数
		fmt.Println("生产了:", num)
		//Cond.Broadcast()  //惊群

		in <- num     //写满了阻塞了
		//Cond.Signal() //惊群
		Cond.Broadcast()
		//再解锁
		Cond.L.Unlock()
	}
}

//消费者
func Consumer(out chan int) {
	for {
		Cond.L.Lock()

		//消费干净了
		if len(out) == 0 {
			//Cond.Signal()
			Cond.Broadcast()
			Cond.Wait()
		}

		//消费
		//range 管道
		num := <-out
		fmt.Println("消费了:", num)
		//Cond.Signal() //惊群
		Cond.Broadcast()
		Cond.L.Unlock()

	}

}

func sing() {
	for i := 1; i <= 500; i++ {
		fmt.Println(i, "唱歌"); //资源是否发生争抢
		//time.Sleep(time.Millisecond*10)
	}
}

func dance() {
	for i := 1; i <= 500; i++ { //是否有数据通信
		fmt.Println(i, "跳舞")
		//time.Sleep(time.Millisecond*10)
	}

}
