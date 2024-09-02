package domain

type UserType string

const (
	COMUM  UserType = "COMUM"
	SELLER UserType = "SELLER"
)

func (u UserType) String() string {
	switch u {
	case COMUM:
		return "COMUM"
	case SELLER:
		return "SELLER"
	}
	return ""
}
