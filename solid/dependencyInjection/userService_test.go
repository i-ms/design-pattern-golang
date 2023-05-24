package dependencyInjection

import "testing"

func TestUserRepository(t *testing.T) {
	userRepository := &UserRepositoryImpl{}
	userService := NewUserService(userRepository)
	_ = userService.CreateUser(User{ID: 1})
	_ = userService.DeleteUser(1)
}
