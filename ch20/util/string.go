package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Test1() {
	fmt.Println("先导入fmt包，才能使用")
	r := gin.Default()
	r.Run()
}
