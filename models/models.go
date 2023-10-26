package model

// BlogPost represents a blog post model.
type BlogPost struct {
	ID          string `gorm:"column:Id";json:"id";"primaryKey"`
	Title       string `gorm:"column:Title" json:"title"`
	Content     string `gorm:"column:Content" json:"content"`
	CreatedDate string `gorm:"column:CreatedDate" json:"createdDate"`
}
