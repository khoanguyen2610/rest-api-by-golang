package models

type User struct {
	BaseModel
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

/*
|--------------------------------------------------------------------------
| Declare table name of Model | optional
| Default table name with "s" eg. users
|--------------------------------------------------------------------------
*/
func (User) TableName() string {
	return "users"
}
