package user

const (
	// CreateUserSQL ...
	CreateUserSQL = `
		INSERT INTO users (
			id,
			username,
			full_name,
			email,
			password,
			created_at
		) VALUES (:id, :username, :full_name, :email, crypt(:password, gen_salt('bf')), NOW())
	`

	// GetUserSQL ...
	GetUserSQL = `
		SELECT
			id,
			username,
			full_name,
			email,
			created_at
		FROM
			users
		WHERE id = $1
	`
)