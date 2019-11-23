package main

import (
	"fmt"
	"sync"
	"time"
)
//wait Group配合  sync.Mutex互斥锁 或者读写锁
/*
wait Group
它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
sync.WaitGroup只有3个方法，Add()，Done()，Wait()。
其中Done()是Add(-1)的别名。简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
*/
func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	fmt.Println("Locked")
	mutex.Lock()  //加了这两个以后就有顺序了  ???  TODO  为甚么加了这两个以后就有了执行顺序
	//加了锁以后  执行顺序[相对]会好很多
	for i := 1; i <= 3; i++ {
		wait.Add(1)//sync.WaitGroup  只添加一个go程
		go func(i int) {
			fmt.Println("Not lock:", i)
			mutex.Lock()
			fmt.Println("Lock:", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock:", i)
			mutex.Unlock()
			defer wait.Done()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()
	wait.Wait() //阻塞主程  等待

}
