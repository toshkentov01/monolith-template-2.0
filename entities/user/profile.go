package user

// User ...
type UserProfileModel struct {
	ID        string
	Username  string
	FullName  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

// CreateUserModel ...
type CreateUserModel struct {
	ID       string
	Username string
	FullName string
	Email    string
	Password string
}

// GetUserModel ...
type GetUserModel struct {
	ID string
}
