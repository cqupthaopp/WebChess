package GameRoom

import "github.com/gin-gonic/gin"

func LoadRoomFuncs(r *gin.Engine) {

	r.POST("/chess/create", CreateRoomFunc)

	r.GET("/chess/join/:name", JoinRoomFunc)

	r.POST("/chess/moving", MovingFunc)

	r.POST("/chess/prepare", PrepareFunc)

}
