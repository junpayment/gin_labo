package index

import (
	"github.com/gin-gonic/gin"
	"github.com/junpayment/go-slim"
	"log"
	"path/filepath"
	"os"
)

func Index(c *gin.Context) {
	parentDir := os.Getenv("PARENT_DIR")
	if "" == parentDir {
		parentDir, _ = os.Getwd()
	}

	viewPath, _ := filepath.Abs(parentDir + "/apps/views/index/index.slim")
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
