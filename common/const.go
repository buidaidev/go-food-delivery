package common

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	CurrentUser      = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
