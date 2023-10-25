package blogs

import (
	"log"

	"github.com/UdayPatil21/blogging/db"
	model "github.com/UdayPatil21/blogging/models"
	"github.com/google/uuid"
)

type BlogInterface interface {
	CreateBlogService(model.BlogPost) error
	EditBlogService()
	DeleteBlogService()
	GetAllBlogsService()
	GetBlogService()
}
type Blogs struct {
}

func (b *Blogs) CreateBlogService(blogPost model.BlogPost) error {
	//Create uuid
	blogPost.ID = uuid.New().String()
	err := db.DBInstance.Table(db.BlogDB).Create(&blogPost).Error
	if err != nil {
		log.Println("Error inserting the data", err)
		return err
	}
	return nil
}

func (b *Blogs) EditBlogService(blogPost model.BlogPost) error {
	// query:=select * from blogs where Id=%d
	result := model.BlogPost{}
	// First find from the database
	err := db.DBInstance.Table(db.BlogDB).Where("Id=?", blogPost.ID).Find(&result).Error
	if err != nil {
		return err
	}
	//Bind and Then update new data to database
	// if result.ID != "" && result.Title != "" && result.Content != "" {

	// }
	result.Title = blogPost.Title
	result.Content = blogPost.Content
	err = db.DBInstance.Table(db.BlogDB).Where("Id=?", blogPost.ID).Update(&result).Error
	if err != nil {
		log.Println("Error updating the data", err)
		return err
	}
	return nil
}
func (b *Blogs) DeleteBlogService(id string) (model.BlogPost, error) {
	blogPost := model.BlogPost{}
	err := db.DBInstance.Table(db.BlogDB).Where("Id=?", id).Find(&blogPost).Error
	if err != nil {
		log.Println("Record not found to delete", err)
		return blogPost, nil
	}
	err = db.DBInstance.Table(db.BlogDB).Where("Id=?", id).Delete(&blogPost).Error
	if err != nil {
		log.Println("Error deleteting the record", err)
		return blogPost, err
	}
	return blogPost, nil
}
func (b *Blogs) GetAllBlogsService() ([]model.BlogPost, error) {
	blogPost := []model.BlogPost{}
	query := "select * from Blogs"
	err := db.DBInstance.Table(db.BlogDB).Raw(query).Scan(&blogPost).Error
	if err != nil {
		log.Println("Error fetching blogs data", err)
		return blogPost, err
	}
	return blogPost, nil
}
func (b *Blogs) GetBlogService(id string) (model.BlogPost, error) {
	blogPost := model.BlogPost{}
	// query := "select * from Blogs"
	err := db.DBInstance.Table(db.BlogDB).Where("Id=?", id).Find(&blogPost).Error
	if err != nil {
		if err.Error() == "record not found" {
			return blogPost, nil
		}
		log.Println("Error fetching blog data", err)
		return blogPost, err
	}
	return blogPost, nil
}
