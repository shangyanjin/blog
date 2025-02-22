package model

// Setting represents a blog setting
type Setting struct {
	ID          int       `gorm:"primarykey" json:"id"`                                               // Primary key
	CreatedAt   LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // Creation time
	UpdatedAt   LocalTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // Last update time
	DeletedAt   LocalTime `gorm:"type:datetime;default:null" json:"deleted_at"`                       // Deletion time for soft delete
	Group       string    `gorm:"size:50;not null;default:''" json:"group"`                           // Setting group (system, site, theme)
	Name        string    `gorm:"size:100;not null;default:''" json:"name"`                           // Setting name
	Value       string    `gorm:"type:text;not null;default:''" json:"value"`                         // Setting value
	Type        string    `gorm:"size:20;not null;default:'string'" json:"type"`                      // Value type (string, int, bool, json)
	Title       string    `gorm:"size:100;not null;default:''" json:"title"`                          // Display title
	Description string    `gorm:"size:500;not null;default:''" json:"description"`                    // Setting description
	Order       int       `gorm:"not null;default:0" json:"order"`                                    // Display order
	Status      string    `gorm:"size:20;not null;default:1" json:"status"`                           // Setting status (0: disabled, 1: enabled)
}

// TableName specifies the table name for Setting model
func (Setting) TableName() string {
	return GetTableName("settings")
}
