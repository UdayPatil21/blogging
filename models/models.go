package model

// BlogPost represents a blog post model.
type BlogPost struct {
	ID      string `gorm:"column:Id";json:"id";"primaryKey"`
	Title   string `gorm:"column:Title" json:"title" required`
	Content string `gorm:"column:Content" json:"content" required`
}
