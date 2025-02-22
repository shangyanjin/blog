package model

// User represents a blog user
type User struct {
	ID           int       `gorm:"primarykey" json:"id"`                                               // Primary key
	CreatedAt    LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // Creation time
	UpdatedAt    LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // Last update time
	DeletedAt    LocalTime `gorm:"type:datetime;default:null" json:"deleted_at"`                       // Deletion time for soft delete
	Username     string    `gorm:"size:50;not null;default:''" json:"username"`                        // Username for login
	Password     string    `gorm:"size:255;not null;default:''" json:"password"`                       // Password hash
	Nickname     string    `gorm:"size:50;not null;default:''" json:"nickname"`                        // Display name
	Email        string    `gorm:"size:100;not null;default:''" json:"email"`                          // Email address
	Avatar       string    `gorm:"size:255;not null;default:''" json:"avatar"`                         // Avatar URL
	Role         string    `gorm:"size:20;not null;default:'user'" json:"role"`                        // User role (admin, editor, user)
	Status       string    `gorm:"size:20;not null;default:0" json:"status"`                           // Account status (0: inactive, 1: active)
	Order        int       `gorm:"not null;default:0" json:"order"`                                    // Display order
	LastLoginAt  LocalTime `gorm:"type:datetime;default:null" json:"last_login_at"`                    // Last login time
	LastLoginIP  string    `gorm:"size:50;not null;default:''" json:"last_login_ip"`                   // Last login IP
	PostCount    int       `gorm:"not null;default:0" json:"post_count"`                               // Number of posts
	CommentCount int       `gorm:"not null;default:0" json:"comment_count"`                            // Number of comments
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return GetTableName("users")
}
