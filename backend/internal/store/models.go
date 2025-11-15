package store
import "time"

type User struct{
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;size:64;not null"`
	Password string `gorm:"size:255;not null"` //hashed pwd
	CreatedAT time.Time
	UpdatedAt time.Time
}

type ConversationType string
const (
	ConversationTypeSingle ConversationType = "single"
	ConversationTypeGroup ConversationType = "group"
	ConversationTypeAI ConversationType = "ai"
)

type Conversation struct {
	ID        uint             `gorm:"primaryKey"`
	Type      ConversationType `gorm:"type:varchar(16);not null"`
	Name      string           `gorm:"size:255"` // group name or AI chat title
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ConversationMember struct {
	ID uint `gorm:"primaryKey"`
	ConversationID uint `gorm:"index;not null"`
	UserID uint `gorm:"index;not null"`
	CreatedAt time.Time
}

type Message struct {
	ID uint `gorm:"primaryKey"`
	ConversationID uint `gorm:"index;not null"`
	SenderID uint `gorm:"index;not null"`
	SenderType string `gorm:"type:varchar(16);not null"` // a user or ai?
	ContentType string `gorm:"type:varchar(16);not null"` // text or file
	Content string `gorm:"type:text;not null"` // text or file url
	CreatedAt time.Time
}