package dependencyInjection

type User struct {
	ID int
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user User) error {
	//Business logic for creating a user
	return s.userRepository.Save(user)
}

func (s *UserService) DeleteUser(id int) error {
	//Business logic for deleting a user
	return s.userRepository.Delete(id)
}
