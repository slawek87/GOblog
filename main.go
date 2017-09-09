package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slawek87/GOblog/blog"
	"github.com/slawek87/GOblog/auth"
)

func main() {
	r := gin.Default()
	goauth := auth.GOauth()

	v1 := r.Group("api/v1/blog/")
	{
		v1.POST("/posts", goauth.AuthenticationMiddleware, blog.PostBlogPostAPI)
		v1.GET("/posts", blog.GetBlogPostsAPI)
		v1.GET("/posts/:id", blog.GetBlogPostAPI)
		v1.PUT("/posts/:id", goauth.AuthenticationMiddleware, blog.UpdateBlogPostAPI)
		v1.DELETE("/posts/:id", goauth.AuthenticationMiddleware, blog.DeleteBlogPostAPI)
	}

	r.Run(":8080")
}

