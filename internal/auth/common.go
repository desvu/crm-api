package auth

// User struct to store user data in request context
type User struct {
	ID    int
	Roles map[string]bool
}

func (u *User) IsEmpty() bool {
	return u.ID == 0
}

// Fingerprint struct to store data in request context
type Fingerprint struct {
	UA   string
	IP   string
	HWID string
}
