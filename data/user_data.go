package datas

type User struct {
	Name     string
	Email    string
	Password string
}
type SessionData struct{
	SessionId string
	Email string
}

var UserData = make(map[string]User)
var Sessions =make(map[string]SessionData)

func init() {
	// user Datas
	UserData["vajidhussain77@gmail.com"] = User{"Vajid", "vajidhussain77@gmail.com", "123"}
	UserData["arun@gmail.com"] = User{"arun", "arun@gmail.com", "456"}
	UserData["shibil.@gmail.com"]=User{"shibil","shibil.@gmail.com","789"}

	// SessionData Id and email
}
