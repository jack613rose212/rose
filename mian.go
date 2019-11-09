package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rose212/src/rose212/servers"
	//rose212是我go_mod init的项目名称 放在在入口处就可以了
	//go mod init rose212  创建的全局项目库
	//全局以项目为开始 多个项目是可以共存的
)

type Obj struct {
	User servers.User
}

func init() {
	fmt.Println("Hello world! start running!!")
}

func main() {
	APIStart()
}

//点燃--->启动功能--->模块加载--->实现功能
func APIStart() {
	//gin.SetMode("Debug") //设置启动模式
	r := gin.Default() //gin的默认路由  进行监听
	AddRouter(r)       //加载所有的路由  // 本身就是返回所有的句柄  不需要再进行取地址
	err:=r.Run(":8080")
	if err!=nil{
		fmt.Println("err:",err)
	}
}

func AddRouter(engine *gin.Engine) {
	var (
		IGoABC     servers.IGoABC      // 接口
		UserStruct = new(servers.User) // 结构体
	)

	//结构体赋值给一个接口
	IGoABC = UserStruct
	IGoABC.StarRoute(engine)

}
