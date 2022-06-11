package User

import "github.com/gin-gonic/gin"

type User struct {
	Username string `gorm:"primary_key"`
	Password string
}

func LoadUserFuncs(r *gin.Engine) {

	r.POST("/user", UserRegisterFunc)
	r.GET("/user", UserLoginFunc)

}
