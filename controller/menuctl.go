package controller

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func menuHandler(c *gin.Context) {
	//解析模板文件
	t := template.Must(template.ParseFiles(HTMLPath + "Menu.html"))
	//声明一个字符串切片

	stars := []string{"马蓉", "李小璐", "白百何"}
	//执行模板
	t.Execute(c.Writer, stars)
}
