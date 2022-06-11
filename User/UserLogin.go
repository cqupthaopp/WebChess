package User

import (
	"WebChess/MySQLUtil"
	"WebChess/Utils"
	"github.com/gin-gonic/gin"
)

func UserLoginFunc(c *gin.Context) {

	LoginUser := MySQLUtil.User{c.PostForm("username"), c.PostForm("password")}

	if err := MySQLUtil.FindUser(LoginUser.Username); err != nil {
		c.JSON(405, Utils.GetErrorRes(gin.H{
			"response": "User " + LoginUser.Username + "is not Exist",
		}))
		return
	}

	UserInfo := MySQLUtil.GetUser(LoginUser.Username)

	if UserInfo.Password != LoginUser.Password {
		c.JSON(403, Utils.GetErrorRes(gin.H{
			"response": "Password is not match",
		}))
	} else {

		token, token_err := Utils.GetToken(LoginUser.Username)
		retoken, retoken_err := Utils.GetRefreshToken(LoginUser.Username)

		if token_err != nil || retoken_err != nil {
			c.JSON(405, Utils.GetErrorRes(gin.H{
				"response": "GetTokenError",
			}))
			return
		}

		c.JSON(200, Utils.GetNormalRes(gin.H{
			"token":         token,
			"refresh_token": retoken,
			"username":      LoginUser.Username,
		}))

	}

}
