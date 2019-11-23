package main

import (
	"fmt"
	"sync"
	"time"
)
/*
多线程并发  使用锁解决
TODO 1: 我们在全局上(主线程上)定义一个互斥锁  等待整体执行完毕以后(再释放)
TODO 2: go程中  进行加锁   防止出现资源的争抢
*/
func main() {


	//声明
	var mutex sync.Mutex
	fmt.Println("可以多个加锁  一个解锁的???. (G0)")
	//加锁mutex
	mutex.Lock()

	fmt.Println("可以直接先加锁  不解锁.(G0)")
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("go程内加锁(G%d)\n", i)
			mutex.Lock()
			fmt.Printf("go程内加锁结束. 锁被释放 (G%d)\n", i)
			//mutex.Unlock()  //TODO 只加锁不解锁  只能走一个(后续的关门进不来了!!!)
		}(i)
	}
	//休息一会,等待打印结果
	time.Sleep(time.Second)
	fmt.Println("休息完毕. (G0)")
	//解锁mutex
	mutex.Unlock()

	fmt.Println("主线程解锁. (G0)")
	//休息一会,等待打印结果
	time.Sleep(time.Second)
}
