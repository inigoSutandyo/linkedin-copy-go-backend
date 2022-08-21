package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"

	// "github.com/fmpwizard/go-quilljs-delta/delta"
	"github.com/microcosm-cc/bluemonday"
)

func AddPost(c *gin.Context) {
	id := getUserID(c)
	var post model.Post
	c.BindJSON(&post)
	fmt.Print("Post = ")
	p := bluemonday.UGCPolicy()

	html := p.Sanitize(post.Content)
	post.Content = html

	user := model.GetUserById(id)

	dbErr := model.CreatePost(&user, &post)

	if dbErr != nil {
		fmt.Print("ERROR = ")
		fmt.Println(dbErr.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": dbErr.Error(),
			"status":  false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"post":    post,
	})
}
