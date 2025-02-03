// @Author Gopher
// @Date 2025/2/3 20:48:00
// @Desc
package main

import (
	"github.com/gin-gonic/gin"
	srv "project-common"
)

func main() {
	r := gin.Default()
	srv.Run(r, "project-user", ":80")
}
