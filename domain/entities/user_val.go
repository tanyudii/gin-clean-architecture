package entities

type UserType int

const (
	UserTypeUnknown UserType = iota
	UserTypeAdmin
	UserTypeUser
)

var userTypeSliceString = []string{
	"Unknown", "ADMIN", "USER",
}

func (u UserType) String() string {
	return userTypeSliceString[u]
}

func (t UserType) Valid() bool {
	return int(t) > 0 && int(t) <= len(userTypeSliceString)-1
}

var mapUserTypeByKey map[string]UserType

func (u *UserType) Scan(src interface{}) error {
	val, ok := src.([]uint8)
	if !ok {
		*u = 0
	}
	*u = GetUserTypeByKey(string(val))
	return nil
}

func GetUserTypeByKey(key string) UserType {
	if mapUserTypeByKey == nil {
		mapUserTypeByKey = make(map[string]UserType)
		for i := range userTypeSliceString {
			mapUserTypeByKey[userTypeSliceString[i]] = UserType(i)
		}
	}
	return mapUserTypeByKey[key]
}
