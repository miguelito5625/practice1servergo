package ApiHelpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	// Meta     interface{}
	Menssage string
	Data     interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}, message string) {
	fmt.Println("status ", status)
	var res ResponseData

	res.Status = status
	//res.Meta = utils.ResponseMessage(status)
	res.Data = payload
	res.Menssage = message

	w.JSON(status, res)
}
