package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	// A slice of pointers to User struct
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	// Use addressof to find pointer
	users = append(users, &u)
	return u, nil
}
