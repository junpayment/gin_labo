package index

import (
	"github.com/gin-gonic/gin"
	"github.com/junpayment/go-slim"
	"log"
	"path/filepath"
)

func Index(c *gin.Context) {
	viewPath, _ := filepath.Abs("apps/views/index/index.slim")
	tmpl, err := slim.ParseFile(viewPath)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(c.Writer, slim.Values{
		"names": []string{"foo", "bar", "baz"},
	})
	if nil != err {
		log.Fatal(err)
	}
}
