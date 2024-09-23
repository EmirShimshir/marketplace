package domain

type Token string

func (t Token) String() string {
	return string(t)
}

type AuthDetails struct {
	AccessToken  Token
	RefreshToken Token
}

type AuthPayload struct {
	UserID ID
	Role   UserRole
}
