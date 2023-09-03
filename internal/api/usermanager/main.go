package usermanager

// API is the interface that wraps the basic methods of the user info API
type API interface {
    Authenticate()  (string, error)
    GetUser()       (string, error)
}

// Anilist is the API for Anilist
type Anilist struct {
    token   string
}

// Authenticate authenticates the user using the Anilist API
func (a *Anilist) Authenticate() (string, error) {
    return "", nil
}

// GetUser returns the user info using the Anilist API
func (a *Anilist) GetUser() (string, error) {
    return "", nil
}

// GetManager returns the manager for the given manager name
func GetManager(name string) (API, error) {
    switch name {
    case "anilist":
        return &Anilist{}, nil
    default:
        return nil, nil
    }
}
