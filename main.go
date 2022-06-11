package main

import (
	"WebChess/Chess"
	"WebChess/GameRoom"
	"WebChess/User"
	"github.com/gin-gonic/gin"
)

func main() {

	{
		Chess.InitBlock()
	} //init

	r := gin.Default()

	{
		User.LoadUserFuncs(r)
		GameRoom.LoadRoomFuncs(r)
	} //LoadFunc

	r.Run(":8080")

}
