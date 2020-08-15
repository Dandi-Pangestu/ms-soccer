package domains

type Auth struct {
	UserID string
}

func (a *Auth) GetUserID() *string {
	if a.UserID == "" {
		return nil
	}

	return &a.UserID
}
