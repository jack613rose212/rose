package main

import (
	"fmt"
	"rose212/src/rose212/servers" //rose212是我go_mod init的项目名称 放在在入口处就可以了
	//全局以项目为开始 多个项目是可以共存的
)

type Obj struct {
	User servers.User
}

func init() {
	fmt.Println("Hello world, 世界你好! 我们即将启程!")
}

func main() {
	Star()
}

//点燃--->启动功能--->模块加载--->实现功能
func Star() {
	var (
		object Obj
	)
	//一起来搭积木玩吧
	object.User.Star()
}
