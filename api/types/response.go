package types

import (
	"github.com/SimonMorphy/grog/api/infra/const/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
}
type response struct {
	Err  int    `json:"error"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func (r *Response) Resp(c *gin.Context, err error, data interface{}) {
	if err != nil {
		r.Error(c, err)
	} else {
		r.Success(c, data)
	}
}

func (r *Response) Success(c *gin.Context, data interface{}) {
	no, msg := errors.Output(nil)
	c.JSON(http.StatusOK, response{
		Err:  no,
		Msg:  msg,
		Data: data,
	})
}

func (r *Response) Error(c *gin.Context, err error) {
	no, msg := errors.Output(err)
	c.JSON(http.StatusOK, response{
		Err:  no,
		Msg:  msg,
		Data: nil,
	})
}
