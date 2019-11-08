package servers

import "rose212/src/rose212/handle"

type User struct {
}

//启动-->功能
func (u *User) Star() {

	//服务集成--go基础模块
	handle.GoABC()

	//服务继承--go的算法模块
	handle.Arithmetic()

}
