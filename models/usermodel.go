package models

type User struct {
	Id       int
	Username string
	Password string
	Sex      int64
	Email    string
}

type UserDB struct {
}

var (
	users = map[int]User{1: User{1, "jiandong", "123456", 1, "jiandong@tbk.com"}, 2: User{2, "xiaofang", "123456", 0, "xiaofang@tbk.com"}}
)

//根据用户名查询用户
func (this *UserDB) GetUser(username string) User {
	for _, user := range users {
		if username == user.Username {
			return user
		}
	}
	return User{}
}

//添加用户
func (this *UserDB) AddUser(user User) int {
	u1 := this.GetUser(user.Username)
	if u1.Id == 0 {
		user.Id = len(users) + 1
		users[user.Id] = user
		return 1
	}
	return 0
}
