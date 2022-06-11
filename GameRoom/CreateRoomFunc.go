package GameRoom

import (
	"WebChess/Utils"
	"github.com/gin-gonic/gin"
)

func CreateRoomFunc(c *gin.Context) {

	info, err := Utils.GetUserFromToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(403, Utils.GetErrorRes(gin.H{"Response": "TokenError"}))
		return
	}

	password := c.PostForm("password")

	err = CreateRoom(info+"Room", info, password)
	PlayingGame[info+"Room"].Prepare[0] = 1

	if err != nil {
		c.JSON(403, Utils.GetErrorRes(gin.H{"Info": "RoomExist"}))
		return
	}

	c.JSON(200, Utils.GetNormalRes(gin.H{"RoomName": info + "Room"}))
	return

}
