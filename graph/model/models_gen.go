package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type BlockList struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	From string `json:"from"`
	To   string `json:"to"`
}

type Chat struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UserID       uint      `json:"userID"`
	Author       *User     `json:"author" gorm:"foreignKey:UserID;"`
	Participants []User    `json:"participants" gorm:"many2many:chat_pariticipants;"`
	Messages     []Message `json:"messages"`
}

type Comment struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Body      string    `json:"body"`
	PostID    uint      `json:"PostID"`
	UserID    uint      `json:"userId"`
	Author    *User     `json:"author" gorm:"foreignKey:UserID;"`
	CreatedAt time.Time `json:"createdAt"`
}

type Message struct {
	ID        uint        `json:"id" gorm:"primarykey"`
	Type      MessageType `json:"type"`
	Asset     *string     `json:"asset"`
	Body      *string     `json:"body"`
	CreatedAt time.Time   `json:"createdAt"`
	Seen      *bool       `json:"seen"`
	UserID    uint        `json:"userID"`
	Author    *User       `json:"author" gorm:"foreignKey:UserID;"`
	ChatId    uint        `json:"chatId"`
}

type Notification struct {
	ID           uint             `json:"id" gorm:"primarykey"`
	ResourceID   string           `json:"resourceId"`
	UserID       uint             `json:"userID"`
	ActionUserID uint             `json:"actionUserID"`
	User         *User            `json:"user" gorm:"foreignKey:UserID;"`
	ActionUser   *User            `json:"actionUser" gorm:"foreignKey:ActionUserID;"`
	CreatedAt    time.Time        `json:"createdAt"`
	Type         NotificationType `json:"type"`
}

type Post struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Caption   *string   `json:"caption"`
	URI       string    `json:"uri"`
	Reports   int       `json:"reports"`
	UserID    uint      `json:"userID"`
	Author    *User     `json:"author" gorm:"foreignKey:UserID;"`
	CreatedAt time.Time `json:"createdAt"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:PostID"`
	Likes     string    `json:"likes"`
}

type Story struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	URI       string    `json:"uri"`
	UserID    uint      `json:"userID"`
	Author    *User     `json:"author" gorm:"foreignKey:UserID;"`
	Views     string    `json:"views"`
	Type      StoryType `json:"type"`
}

type User struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	Token         string         `json:"token"`
	FcmToken      *string        `json:"fcmToken"`
	Name          string         `json:"name"`
	Handle        string         `json:"handle"`
	Avatar        *string        `json:"avatar"`
	Email         string         `json:"email"`
	LastSeen      float64        `json:"lastSeen"`
	About         *string        `json:"about"`
	Posts         []Post         `json:"posts" gorm:"foreignKey:UserID"`
	Stories       []Story        `json:"stories" gorm:"foreignKey:UserID"`
	Following     []User         `json:"following" gorm:"many2many:user_followings;"`
	Followers     []User         `json:"followers" gorm:"many2many:user_followers;"`
	Chats         []Chat         `json:"chats" gorm:"foreignKey:UserID"`
	Notifications []Notification `json:"notifications" gorm:"foreignKey:UserID"`
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

func (e ConnectionsType) IsValid() bool {
	switch e {
	case ConnectionsTypeFollowing, ConnectionsTypeFollowers:
		return true
	}
	return false
}

func (e ConnectionsType) String() string {
	return string(e)
}

func (e *ConnectionsType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConnectionsType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConnectionsType", str)
	}
	return nil
}

func (e ConnectionsType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

func (e LikeAction) IsValid() bool {
	switch e {
	case LikeActionLike, LikeActionUnlike:
		return true
	}
	return false
}

func (e LikeAction) String() string {
	return string(e)
}

func (e *LikeAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LikeAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LikeAction", str)
	}
	return nil
}

func (e LikeAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

func (e MessageType) IsValid() bool {
	switch e {
	case MessageTypeText, MessageTypeImage, MessageTypeVideo:
		return true
	}
	return false
}

func (e MessageType) String() string {
	return string(e)
}

func (e *MessageType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MessageType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MessageType", str)
	}
	return nil
}

func (e MessageType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

func (e NotificationType) IsValid() bool {
	switch e {
	case NotificationTypeFollow, NotificationTypeComment, NotificationTypeLike:
		return true
	}
	return false
}

func (e NotificationType) String() string {
	return string(e)
}

func (e *NotificationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationType", str)
	}
	return nil
}

func (e NotificationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

func (e StoryType) IsValid() bool {
	switch e {
	case StoryTypeImage, StoryTypeVideo:
		return true
	}
	return false
}

func (e StoryType) String() string {
	return string(e)
}

func (e *StoryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StoryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StoryType", str)
	}
	return nil
}

func (e StoryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UpdateFollowingAction string

const (
	UpdateFollowingActionFollow   UpdateFollowingAction = "FOLLOW"
	UpdateFollowingActionUnfollow UpdateFollowingAction = "UNFOLLOW"
)

var AllUpdateFollowingAction = []UpdateFollowingAction{
	UpdateFollowingActionFollow,
	UpdateFollowingActionUnfollow,
}

func (e UpdateFollowingAction) IsValid() bool {
	switch e {
	case UpdateFollowingActionFollow, UpdateFollowingActionUnfollow:
		return true
	}
	return false
}

func (e UpdateFollowingAction) String() string {
	return string(e)
}

func (e *UpdateFollowingAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UpdateFollowingAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UpdateFollowingAction", str)
	}
	return nil
}

func (e UpdateFollowingAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
