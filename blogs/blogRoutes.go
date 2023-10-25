package blogs

import (
	"log"
	"net/http"

	model "github.com/UdayPatil21/blogging/models"
	"github.com/gin-gonic/gin"
)

var s Blogs

func CreateBlog(c *gin.Context) {
	blogPost := model.BlogPost{}
	// s := Blogs{}
	err := c.ShouldBindJSON(&blogPost)
	if err != nil {
		log.Println("Error binding the blog data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	if blogPost.Title == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"Status":  false,
			"Message": "Title should not be empty",
		})
		return
	} else if blogPost.Content == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"Status":  false,
			"Message": "Content should not be empty",
		})
	}
	err = s.CreateBlogService(blogPost)
	if err != nil {
		log.Println("Error in create blog", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  true,
		"Message": "Blog Created Successfully",
	})

}
func EditBlog(c *gin.Context) {
	blogPost := model.BlogPost{}
	// s := Blogs{}
	err := c.ShouldBindJSON(&blogPost)
	if err != nil {
		log.Println("Error binding the blog data", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	if blogPost.ID == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"Status":  false,
			"Message": "Please provide blog post id which you want to update",
		})
		return
	}
	if blogPost.Title == "" || blogPost.Content == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"Status":  false,
			"Message": "Data should not be empty",
		})
		return
	}
	err = s.EditBlogService(blogPost)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"Status":  false,
				"Message": "No Blog To Update! Check ID",
			})
			return
		}
		log.Println("Error in edit blog service", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  true,
		"Message": "Blog Updated Successfully",
	})
}
func DeleteBlog(c *gin.Context) {
	id := c.Param("id")

	data, err := s.DeleteBlogService(id)
	if err != nil {
		log.Println("Error in delete blog service", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	if data.ID == "" && data.Title == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"Status":  false,
			"Message": "No Blog Present with this ID! Please Check",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  true,
		"Message": "Record deleted sucessfully",
	})
}
func GetAllBlogs(c *gin.Context) {
	data, err := s.GetAllBlogsService()
	if err != nil {
		log.Println("Error in get all blog service", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  true,
		"Message": "Success",
		"Data":    data,
	})
}
func GetBlog(c *gin.Context) {
	id := c.Param("id")
	data, err := s.GetBlogService(id)
	if err != nil {
		log.Println("Error in get blog service", err)
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	if data.ID == "" && data.Content == "" {
		c.JSON(http.StatusOK, gin.H{
			"Status":  true,
			"Message": "No Record Found",
			"Data":    data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  true,
		"Message": "Success",
		"Data":    data,
	})
}
