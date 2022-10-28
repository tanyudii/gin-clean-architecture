package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vodeacloud/hr-api/pkg/errutil"
	"github.com/vodeacloud/hr-api/pkg/pagination"
	"net/http"
)

type Response struct {
	Data       interface{}            `json:"data,omitempty"`
	Pagination *pagination.Pagination `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func GinResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, &Response{
		Data: data,
	})
}

func GinResponseWithoutData(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func GinResponseWithPagination(c *gin.Context, code int, data interface{}, pagination *pagination.Pagination) {
	c.JSON(code, &Response{
		Data:       data,
		Pagination: pagination,
	})
}

func GinErrorResponse(c *gin.Context, err error) {
	code := http.StatusInternalServerError
	errResp := &ErrorResponse{}

	customErr, ok := (err).(errutil.CustomError)
	if ok {
		code = customErr.GetHTTPCode()
		errResp.Message = customErr.Error()
		if errutil.IsBadRequestError(err) {
			badRequest := err.(*errutil.BadRequestError)
			fields := badRequest.GetFields()
			if len(fields) != 0 {
				errResp.Errors = fields
			}
		}
	} else {
		errResp.Message = http.StatusText(code)
	}

	c.JSON(code, errResp)
}
