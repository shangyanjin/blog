package model

// Tag represents a blog tag
type Tag struct {
	ID          int    `gorm:"primarykey" json:"id"`                            // Primary key
	Name        string `gorm:"size:50;not null;default:''" json:"name"`         // Tag name
	Slug        string `gorm:"size:100;not null;default:''" json:"slug"`        // Tag slug for URL
	Description string `gorm:"size:500;not null;default:''" json:"description"` // Tag description

	Order int `gorm:"not null;default:0" json:"order"` // Display order
	Count int `gorm:"not null;default:0" json:"count"` // Number of posts with this tag

	CreatedAt LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // Creation time
	UpdatedAt LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // Last update time
	DeletedAt LocalTime `gorm:"type:datetime;default:null" json:"deleted_at"`                       // Deletion time for soft delete
}

// TableName specifies the table name for Tag model
func (Tag) TableName() string {
	return "tag"
}
