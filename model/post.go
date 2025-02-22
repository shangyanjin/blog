package model

// Post represents a blog post
type Post struct {
	ID      int    `gorm:"primarykey" json:"id"`                         // Primary key
	Title   string `gorm:"size:200;not null;default:''" json:"title"`    // Post title
	Content string `gorm:"type:text;not null;default:''" json:"content"` // Post content in HTML format
	Summary string `gorm:"size:500;not null;default:''" json:"summary"`  // Post summary or excerpt
	Cover   string `gorm:"size:255;not null;default:''" json:"cover"`    // Cover image URL
	Author  string `gorm:"size:100;not null;default:''" json:"author"`   // Post author name
	Tags    string `gorm:"size:200;not null;default:''" json:"tags"`     // Post tags, comma separated

	Status       string `gorm:"size:20;not null;default:0" json:"status"`       // Post status (0: draft, 1: published)
	IsTop        string `gorm:"size:20;not null;default:0" json:"is_top"`       // Whether post is pinned to top (0: no, 1: yes)
	IsRecommend  string `gorm:"size:20;not null;default:0" json:"is_recommend"` // Whether post is recommended (0: no, 1: yes)
	Order        int    `gorm:"not null;default:0" json:"order"`                // Display order
	ViewCount    int    `gorm:"not null;default:0" json:"view_count"`           // Number of views
	LikeCount    int    `gorm:"not null;default:0" json:"like_count"`           // Number of likes
	CommentCount int    `gorm:"not null;default:0" json:"comment_count"`        // Number of comments

	CreatedAt LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // Creation time
	UpdatedAt LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // Last update time
	DeletedAt LocalTime `gorm:"type:datetime;default:null" json:"deleted_at"`                       // Deletion time for soft delete
}

// TableName specifies the table name for Post model
func (Post) TableName() string {
	return "post"
}
