package demo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/permadao/go-demo/demo/schema"
)

func (d *Demo) runAPI(port string) {
	e := gin.Default()

	// middleware
	e.Use(cors.Default())

	// api
	e.GET("/hello", d.hello)
	e.GET("/hello/:to", d.helloTo)
	e.POST("/submit", d.submit)

	// run server
	d.apiServer = &http.Server{
		Addr:    port,
		Handler: e,
	}

	if err := d.apiServer.ListenAndServe(); err != nil {
		log.Warn("http ListenAndServe", "err", err)
	}
}

func (d *Demo) closeAPI() {
	if err := d.apiServer.Shutdown(context.Background()); err != nil {
		log.Error("shutdown api failed", "err", err)
		return
	}
	log.Info("api closed")
}

func errorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, schema.ErrRes{
		Err: err.Error(),
	})
}

// api implementation
func (d *Demo) hello(c *gin.Context) {
	c.JSON(http.StatusOK, schema.HelloRes{"hello world!"})
}

func (d *Demo) helloTo(c *gin.Context) {
	to := c.Param("to")
	c.JSON(http.StatusOK, schema.HelloRes{fmt.Sprintf("hello %v", to)})
}

func (d *Demo) submit(c *gin.Context) {
	sb := schema.SubmitReq{}
	if err := c.ShouldBindJSON(&sb); err != nil {
		log.Warn("error submitting", "err", err, "body", c.Request.Body)
		errorResponse(c, schema.ERR_INVALID_REQUEST)
		return
	}

	// update sql
	// d.sql.UpdateTestTable()
	// update redis
	// d.redis.SetTest(sb.Type, sb.SubmitMsg)

	c.JSON(http.StatusOK, sb)
}
