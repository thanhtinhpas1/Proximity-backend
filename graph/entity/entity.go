package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlockList struct {
	ID      uuid.UUID `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	From    uuid.UUID `json:"from" gorm:"type:uuid;uniqueIndex:idx_from"`
	To      uuid.UUID `json:"to" gorm:"type:uuid;uniqueIndex:idx_to"`
	Deleted gorm.DeletedAt
}

type Chat struct {
	ID           uuid.UUID `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	UserID       uuid.UUID `json:"userID,omitempty" gorm:"type:uuid;"`
	Author       *User     `json:"author,omitempty" gorm:"foreignKey:UserID;"`
	Participants []User    `json:"participants" gorm:"many2many:chat_participants;"`
	Messages     []Message `json:"messages"`
}

type Participant struct {
	ChatID uuid.UUID `json:"chat_id,omitempty"`
	UserID uuid.UUID `json:"user_id,omitempty"`
}

type Comment struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Body      string    `json:"body,omitempty"`
	PostID    uuid.UUID `json:"postID,omitempty" gorm:"type:uuid;"`
	UserID    uuid.UUID `json:"userId,omitempty" gorm:"type:uuid;"`
	Author    *User     `json:"author,omitempty" gorm:"foreignKey:UserID;"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Deleted   gorm.DeletedAt
}

type Message struct {
	ID        uuid.UUID   `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Type      MessageType `json:"type,string,omitempty" gorm:"type:string"`
	Asset     *string     `json:"asset,omitempty"`
	Body      *string     `json:"body,omitempty"`
	CreatedAt time.Time   `json:"createdAt,omitempty"`
	Seen      *bool       `json:"seen,omitempty"`
	UserID    uuid.UUID   `json:"userID,omitempty" gorm:"type:uuid;"`
	Author    *User       `json:"author,omitempty" gorm:"foreignKey:UserID;"`
	ChatId    uuid.UUID   `json:"chatId,omitempty" gorm:"type:uuid; "`
}

type Notification struct {
	ID           uuid.UUID        `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	ResourceID   string           `json:"resourceId,omitempty"`
	UserID       uuid.UUID        `json:"userID,omitempty" gorm:"type:uuid;"`
	ActionUserID uint             `json:"actionUserID,string,omitempty"`
	User         *User            `json:"user,omitempty" gorm:"foreignKey:UserID;"`
	ActionUser   *User            `json:"actionUser,omitempty" gorm:"foreignKey:ActionUserID;"`
	CreatedAt    time.Time        `json:"createdAt,omitempty"`
	Type         NotificationType `json:"type,omitempty"`
	Deleted      gorm.DeletedAt
}

type Post struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Caption   *string   `json:"caption,omitempty"`
	URI       string    `json:"uri,omitempty"`
	Reports   int       `json:"reports,omitempty"`
	UserID    uuid.UUID `json:"userID,omitempty" gorm:"type:uuid;"`
	Author    *User     `json:"author,omitempty" gorm:"foreignKey:UserID;"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:PostID"`
	Likes     []Like    `json:"likes" gorm:"foreignKey:PostID"`
	Deleted   gorm.DeletedAt
}

type Like struct {
	PostID uuid.UUID `json:"postId,omitempty" gorm:"primarykey;type:uuid;"`
	UserID uuid.UUID `json:"userId,omitempty" gorm:"primarykey;type:uuid;"`
	Action string    `json:"action,omitempty"`
}

type Story struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	URI       string    `json:"uri,omitempty"`
	UserID    uuid.UUID `json:"userID,omitempty" gorm:"type:uuid;"`
	Author    *User     `json:"author,omitempty" gorm:"foreignKey:UserID;"`
	Views     string    `json:"views,omitempty"`
	Type      StoryType `json:"type,omitempty"`
}

type User struct {
	ID            uuid.UUID      `json:"id,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Token         string         `json:"token,omitempty" gorm:"uniqueIndex:idx_token"`
	FcmToken      *string        `json:"fcmToken,omitempty" gorm:"unique"`
	Name          string         `json:"name,omitempty" gorm:"uniqueIndex:idx_name"`
	Handle        string         `json:"handle,omitempty"`
	Avatar        *string        `json:"avatar,omitempty"`
	Email         string         `json:"email,omitempty" gorm:"unique"`
	LastSeen      float64        `json:"lastSeen,omitempty"`
	About         *string        `json:"about,omitempty"`
	Posts         []Post         `json:"posts" gorm:"foreignKey:UserID"`
	Stories       []Story        `json:"stories" gorm:"foreignKey:UserID"`
	Following     []User         `json:"following" gorm:"many2many:user_followings;"`
	Followers     []User         `json:"followers" gorm:"many2many:user_followers;"`
	Chats         []Chat         `json:"chats" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UserID"`
}

type UserFollowing struct {
	UserID      uuid.UUID `json:"user_id,omitempty"`
	FollowingID uuid.UUID `json:"following_id,omitempty"`
}

type UserFollower struct {
	UserID     uuid.UUID `json:"user_id,omitempty"`
	FollowerID uuid.UUID `json:"follower_id,omitempty"`
}

type ConnectionsType string

const (
	ConnectionsTypeFollowing ConnectionsType = "FOLLOWING"
	ConnectionsTypeFollowers ConnectionsType = "FOLLOWERS"
)

var AllConnectionsType = []ConnectionsType{
	ConnectionsTypeFollowing,
	ConnectionsTypeFollowers,
}

type MessageType string

const (
	MessageTypeText  MessageType = "TEXT"
	MessageTypeImage MessageType = "IMAGE"
	MessageTypeVideo MessageType = "VIDEO"
)

var AllMessageType = []MessageType{
	MessageTypeText,
	MessageTypeImage,
	MessageTypeVideo,
}

type StoryType string

const (
	StoryTypeImage StoryType = "IMAGE"
	StoryTypeVideo StoryType = "VIDEO"
)

var AllStoryType = []StoryType{
	StoryTypeImage,
	StoryTypeVideo,
}

type NotificationType string

const (
	NotificationTypeFollow  NotificationType = "FOLLOW"
	NotificationTypeComment NotificationType = "COMMENT"
	NotificationTypeLike    NotificationType = "LIKE"
)

var AllNotificationType = []NotificationType{
	NotificationTypeFollow,
	NotificationTypeComment,
	NotificationTypeLike,
}

type LikeAction string

const (
	LikeActionLike   LikeAction = "LIKE"
	LikeActionUnlike LikeAction = "UNLIKE"
)

var AllLikeAction = []LikeAction{
	LikeActionLike,
	LikeActionUnlike,
}
