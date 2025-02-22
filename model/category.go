package model

// Category represents a blog category
type Category struct {
	ID          int    `gorm:"primarykey" json:"id"`                            // Primary key
	Name        string `gorm:"size:50;not null;default:''" json:"name"`         // Category name
	Slug        string `gorm:"size:100;not null;default:''" json:"slug"`        // Category slug for URL
	Description string `gorm:"size:500;not null;default:''" json:"description"` // Category description
	ParentID    int    `gorm:"not null;default:0" json:"parent_id"`             // Parent category ID, 0 for root

	Order int `gorm:"not null;default:0" json:"order"` // Display order
	Count int `gorm:"not null;default:0" json:"count"` // Number of posts in this category

	CreatedAt LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // Creation time
	UpdatedAt LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // Last update time
	DeletedAt LocalTime `gorm:"type:datetime;default:null" json:"deleted_at"`                       // Deletion time for soft delete
}

// TableName specifies the table name for Category model
func (Category) TableName() string {
	return "category"
}
