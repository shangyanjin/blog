package model

// Comment represents a blog comment
type Comment struct {
	ID         int       `gorm:"primarykey" json:"id"`                                               // Primary key
	CreatedAt  LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // Creation time
	UpdatedAt  LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // Last update time
	DeletedAt  LocalTime `gorm:"type:datetime;default:null" json:"deleted_at"`                       // Deletion time for soft delete
	PostID     int       `gorm:"not null;default:0" json:"post_id"`                                  // Related post ID
	ParentID   int       `gorm:"not null;default:0" json:"parent_id"`                                // Parent comment ID, 0 for root
	UserID     int       `gorm:"not null;default:0" json:"user_id"`                                  // Comment author user ID
	Content    string    `gorm:"type:text;not null;default:''" json:"content"`                       // Comment content
	Status     string    `gorm:"size:20;not null;default:0" json:"status"`                           // Comment status (0: pending, 1: approved)
	Order      int       `gorm:"not null;default:0" json:"order"`                                    // Display order
	IP         string    `gorm:"size:50;not null;default:''" json:"ip"`                              // Commenter IP address
	UserAgent  string    `gorm:"size:500;not null;default:''" json:"user_agent"`                     // Commenter user agent
	LikeCount  int       `gorm:"not null;default:0" json:"like_count"`                               // Number of likes
	ReplyCount int       `gorm:"not null;default:0" json:"reply_count"`                              // Number of replies
}

// TableName specifies the table name for Comment model
func (Comment) TableName() string {
	return GetTableName("comments")
}
