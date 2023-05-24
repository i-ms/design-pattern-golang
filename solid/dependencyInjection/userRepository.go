package dependencyInjection

type UserRepository interface {
	Save(user User) error
	Delete(id int) error
}

type UserRepositoryImpl struct {
	//	Database connection or any other dependency
}

func (r *UserRepositoryImpl) Save(user User) error {
	// Database query logic
	return nil
}

func (r *UserRepositoryImpl) Delete(id int) error {
	// Database query logic
	return nil
}
