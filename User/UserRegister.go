package User

import (
	"WebChess/MySQLUtil"
	"WebChess/Utils"
	"github.com/gin-gonic/gin"
)

func UserRegisterFunc(c *gin.Context) {

	user := MySQLUtil.User{c.PostForm("username"), c.PostForm("password")}

	if err := MySQLUtil.FindUser(user.Username); err == nil {
		c.JSON(405, Utils.GetErrorRes(gin.H{
			"response": "User " + user.Username + " Exist",
		}))
		return
	}

	err := MySQLUtil.AddUser(user)

	if err != nil {
		c.JSON(405, Utils.GetErrorRes(gin.H{
			"response": "Register Failed",
		}))
	} else {
		token, token_err := Utils.GetToken(user.Username)
		retoken, retoken_err := Utils.GetRefreshToken(user.Username)

		if token_err != nil || retoken_err != nil {
			c.JSON(405, Utils.GetErrorRes(gin.H{
				"response": "GetTokenError",
			}))
		}

		c.JSON(200, Utils.GetNormalRes(gin.H{
			"token":         token,
			"refresh_token": retoken,
			"username":      user.Username,
		}))
	}

}
