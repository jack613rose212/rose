package study_handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rose212/src/rose212/models/study_models"
)

//定义结构体  只是为了绑定和加载
type Go struct {
}

//golang基础模块
func (*Go) GoABC(c *gin.Context) {

	var (
	//json data是返回的数据结构

	)
	//切片的怎删改查
	study_models.Go切片的增删改查(c)

	c.JSON(
		http.StatusOK, gin.H{
			"Status": "1",
			"Code":   "2",
			"Data":   []struct{}{},
			"Msg":    "",
		},
	)
	return
}
