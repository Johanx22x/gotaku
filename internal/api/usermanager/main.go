package usermanager

// API is the interface that wraps the basic methods of the user info API
type API interface {
    Authenticate()  (string, error)
}

// GetManager returns the manager for the given manager name
func GetManager(name string) (API, error) {
    switch name {
    default:
        return nil, nil
    }
}
