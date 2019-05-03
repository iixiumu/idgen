package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

var id uint64

type Generator interface {
	GenID() uint64
}

func Factory(tp int) Generator {
	if tp == 1 {
		return MemAtomic{}
	} else if tp == 2 {
		return MemMutex{}
	} else if tp == 3 {
		return RedisIncr{}
	} else if tp == 4 {
		return MysqlOptimism{}
	} else if tp == 5 {
		return MysqlPessimism{}
	} else if tp == 6 {
		return MysqlPessimismStep{}
	}
	return MemAtomic{}
}

type Input struct {
	Type int `form:"type" binding:"required"`
}

type Output struct {
	ID uint64 `json:"id"`
}

func genid(c *gin.Context) {
	var input Input
	err := c.ShouldBind(&input)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	generator := Factory(input.Type)
	c.JSON(200, map[string]interface{}{
		"id": generator.GenID(),
	})

}

func main() {

	go func() {
		log.Println(http.ListenAndServe(":8053", nil))
	}()

	r := gin.Default()
	r.GET("/getid", genid)
	r.GET("/id", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"id": id,
		})
	})

	r.Run(":8052")
}
