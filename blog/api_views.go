package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostBlogPostAPI(c *gin.Context) {
	post := Blog{}
	c.Bind(&post)

	_, err := CreateBlogPost(&post)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, &post)
}

func GetBlogPostsAPI(c *gin.Context) {
	posts, err := ListBlogPosts()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &posts)
}

func GetBlogPostAPI(c *gin.Context) {
	id := c.Param("id")
	post, err := GetBlogPost(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &post)
}

func UpdateBlogPostAPI(c *gin.Context) {
	id := c.Param("id")

	post := Blog{}
	c.Bind(&post)

	record, err := UpdateBlogPost(&post, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, &record)
}

func DeleteBlogPostAPI(c *gin.Context) {
	id := c.Param("id")
	post, err := DeleteBlogPost(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, &post)
}

