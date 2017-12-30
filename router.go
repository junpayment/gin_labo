package main

import (
	"github.com/gin-gonic/gin"
	"github.com/junpayment/gin_labo/apps/controllers/index"
	"github.com/suapapa/go_sass"
	"log"
	"fmt"
)

const PORT = "8000"
var router *gin.Engine

// initialize routes
func initialRoutes() {
	router = gin.Default()

	// static
	setStatic()

	// route
	initIndex()

	err := router.Run(fmt.Sprintf(":%s", PORT))
	if nil != err {
		log.Fatal(err)
	}
}

func initIndex() {
	r := router.Group("/web")
	r.GET("/", index.Index)
}

func setStatic() {
	// scss一括コンパイル
	var sc sass.Compiler
	err := sc.CompileFolder("apps/scss", "static/css")
	if nil != err {
		log.Fatal(err)
	}

	// contents以下をpublicとして公開
	router.Static("/contents", "./static")
}
