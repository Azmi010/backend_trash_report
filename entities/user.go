package entities

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Point     string
	CreatedAt string
	UpdatedAt string
	Token     string
	Reports   []Report
}
