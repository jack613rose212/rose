package servers

import (
	"github.com/gin-gonic/gin"
	"rose212/src/rose212/handle/study_handle"
)

//定义一个接口
type IGoABC interface {
	//接口实现的方法
	StarRoute(ctx *gin.Engine) //路由
	GoABC(ctx *gin.Context)
}

//实现这个接口  这个是一个对象
type User struct {
	//这个对象进来的  找到的其他的对象
	//直接继承
	study_handle.Go
}

//启动-->功能
func (u *User) StarRoute(c *gin.Engine) {
	HTTP := c.Group("/golang/studay")
	HTTP.GET("/1", u.GoABC)
	HTTP.GET("/2", u.ProducterConsumer)

}

//go基础模块
func (u *User) GoABC(c *gin.Context) {
	u.Go.GoABC(c)
}
//生产者 消费者
func(u  *User) ProducterConsumer(c  *gin.Context){
	u.Go.ProducterConsumer(c)
}
//都要在这里展示出来一个接口


//还是从对象继承的对象上分叉出去  实现各自的代码业务逻辑


//都需要一个接口进行调用





/*

//服务继承--go的算法模块
	handle.Arithmetic()
	//设计模式


这个是直接可以自己定义进行运行的接口
需要一个能够接收前端清楚  进行处理的


*/
