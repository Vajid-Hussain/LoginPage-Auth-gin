package handlers

import (
	"fmt"
	datas "ginserver/data"


	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "github.com/spf13/pflag"
)

func DeleteCookie(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", false, true)
	cookie,err:=c.Cookie("session")
	if err!=nil{
		fmt.Println("wrog in DeleteCookie function",err)
	}
	delete(datas.Sessions,cookie)
	fmt.Println("cookie deleted")
}

func CreateCookie(c *gin.Context) {
	fmt.Printf("cookie created")
	id := uuid.NewString()
	email := c.PostForm("email")
    fmt.Println("cookie id :",id,"email is :",email)

	// Set cookie and session datas
	session:=datas.SessionData{SessionId: id,Email: email}
	datas.Sessions[id] = session
	c.SetCookie("session", id, 3600, "/", "", false, true)
	fmt.Println(datas.Sessions)
}
