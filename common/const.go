package common

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	CurrentUser      = "user"
)

const (
	TopicUserLikeRestaurant    = "TopicUserLikeRestaurant"
	TopicUserDisLikeRestaurant = "TopicUserDisLikeRestaurant"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
