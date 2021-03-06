// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type BlockList struct {
	ID   string `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
}

type Chat struct {
	ID           string     `json:"id"`
	CreatedAt    string     `json:"createdAt"`
	UpdatedAt    string     `json:"updatedAt"`
	Participants []*User    `json:"participants"`
	Messages     []*Message `json:"messages"`
}

type Comment struct {
	ID        string `json:"id"`
	Body      string `json:"body"`
	Author    *User  `json:"author"`
	CreatedAt string `json:"createdAt"`
}

type Like struct {
	PostID string     `json:"postId"`
	UserID string     `json:"userId"`
	Action LikeAction `json:"action"`
}

type Message struct {
	ID        string      `json:"id"`
	Type      MessageType `json:"type"`
	Asset     *string     `json:"asset"`
	Body      *string     `json:"body"`
	CreatedAt string      `json:"createdAt"`
	Seen      *bool       `json:"seen"`
	Author    *User       `json:"author"`
}

type Notification struct {
	ID         string           `json:"id"`
	ResourceID string           `json:"resourceId"`
	User       *User            `json:"user"`
	ActionUser *User            `json:"actionUser"`
	CreatedAt  string           `json:"createdAt"`
	Type       NotificationType `json:"type"`
}

type Post struct {
	ID        string     `json:"id"`
	Caption   *string    `json:"caption"`
	URI       string     `json:"uri"`
	Reports   int        `json:"reports"`
	Author    *User      `json:"author"`
	CreatedAt string     `json:"createdAt"`
	Comments  []*Comment `json:"comments"`
	Likes     []*Like    `json:"likes"`
}

type Story struct {
	ID        string    `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	URI       string    `json:"uri"`
	Author    *User     `json:"author"`
	Views     []string  `json:"views"`
	Type      StoryType `json:"type"`
}

type User struct {
	ID            string          `json:"id"`
	Token         string          `json:"token"`
	FcmToken      *string         `json:"fcmToken"`
	Name          string          `json:"name"`
	Handle        string          `json:"handle"`
	Avatar        *string         `json:"avatar"`
	Email         string          `json:"email"`
	LastSeen      float64         `json:"lastSeen"`
	About         *string         `json:"about"`
	Posts         []*Post         `json:"posts"`
	Stories       []*Story        `json:"stories"`
	Following     []*User         `json:"following"`
	Followers     []*User         `json:"followers"`
	Chats         []*Chat         `json:"chats"`
	Notifications []*Notification `json:"notifications"`
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
