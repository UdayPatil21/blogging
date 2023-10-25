package routes

import (
	"github.com/UdayPatil21/blogging/blogs"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	//Create all the routes in your application
	r.GET("blogs", blogs.GetAllBlogs)
	r.GET("blog:id", blogs.GetBlog)
	r.POST("createBlog", blogs.CreateBlog)
	r.POST("editBlog", blogs.EditBlog)
	r.DELETE("blog:id", blogs.DeleteBlog)

}
