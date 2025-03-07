package auth

type UserStore interface {
	Validate(username, password string) bool
}

type InMemoryUserStore struct {
	users map[string]string
}


func (store *InMemoryUserStore) AddUser(username, password string) {
    store.users[username] = password
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: map[string]string{
			"admin": "adminpass",
			"user":  "userpass",
		},
	}
}

func (s *InMemoryUserStore) Validate(username, password string) bool {
	storedPass, ok := s.users[username]
	return ok && storedPass == password
}