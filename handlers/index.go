package handlers

import (
	"fmt"
	datas "ginserver/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	ClearCache(c)

	// checking cookie exist also fetchig user details
	var name string
	cookie, err := c.Cookie("session")
	if err != nil {
		fmt.Println("error in index rooter :", err)
	}
	
	session, exist := datas.Sessions[cookie]
	if exist{
		data,exist:=datas.UserData[session.Email]
		if exist{
			name=data.Name
			c.HTML(http.StatusOK, "index.html", name)
		}
	}else{
		// c.HTML(http.StatusOK, "index.html", "Guest User")
		c.Redirect(http.StatusFound, "/login")
	}
	
}
