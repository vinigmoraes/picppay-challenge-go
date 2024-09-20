package domain

type UserType string

const (
	COMMON UserType = "COMMON"
	SELLER UserType = "SELLER"
)

func (u UserType) String() string {
	switch u {
	case COMMON:
		return "COMMON"
	case SELLER:
		return "SELLER"
	}
	return ""
}
