package users_transport_http

type UserHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func NewUserHTTPHandler(usersService UsersService) *UserHTTPHandler {
	return &UserHTTPHandler{
		usersService: usersService,
	}
}
