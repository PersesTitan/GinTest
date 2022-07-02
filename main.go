package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		//"upper": strings.ToUpper,
		//"safe": func(s string) template.HTML {
		//	return template.HTML(s)
		//},
	})

	//router.LoadHTMLFiles("gin_HTML/templates/about.html")

	//router.Static("/assets", "./assets")
	//router.LoadHTMLGlob("templates/*.html")
	router.LoadHTMLFiles("templates/footer.html",
		"templates/header.html",
		"templates/index.html",
		"templates/about.html")

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "메인페이지 입니다.",
			"message": "제목",
			//"url":     "/about",
		})
	})

	//router.GET("/about", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "about.html", gin.H{
	//		"about":   "서브 페이지",
	//		"content": "안녕",
	//	})
	//})

	//router := setupRouter()
	_ = router.Run(":8080")
}
