package option

type User struct {
	Id          int
	Name        string
	DateOfBirth string
	Email       string
	Phone       string
}

type UserOption func(*User)

func WithId(id int) UserOption {
	return func(u *User) {
		u.Id = id
	}
}

func WithName(name string) UserOption {
	return func(u *User) {
		u.Name = name
	}
}

func WithDateOfBirth(dateOfBirth string) UserOption {
	return func(u *User) {
		u.DateOfBirth = dateOfBirth
	}
}

func WithEmail(email string) UserOption {
	return func(u *User) {
		u.Email = email
	}
}

func WithPhone(phone string) UserOption {
	return func(u *User) {
		u.Phone = phone
	}
}

func NewUser(opts ...UserOption) *User {
	// default values for user can be defined here
	u := &User{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}
