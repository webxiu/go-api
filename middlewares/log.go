package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func (middleware Middleware) Middleware1(c *gin.Context) {
	fmt.Println("log1 - 全局中间件1")
	c.Set("username", "我是中间件数据")
	// 开启协程不会阻塞代码的执行, 相当于异步
	// 定义一个 goroutine 统计日志, 不能使用上下文的c *gin.Context, 必须使用其只读副本 c.Copy()
	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("过了5秒才会输出 Done! in path " + cCp.Request.URL.Path)
	}()
	c.Next()
	fmt.Println("log1 - 全局中间件2")
}
func (middleware Middleware) Middleware2(c *gin.Context) {
	fmt.Println("log2 - 全局中间件1")
	c.Next()
	fmt.Println("log2 - 全局中间件2")
}
