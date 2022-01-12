package controllers

import (
	"blog_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 点击删除后重定向到首页
func DeleteaArticleGet(c *gin.Context) {
	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("删除id:", id)

	_, err := models.DeleteArticle(id)
	if err != nil {
		log.Panicln(err)
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}