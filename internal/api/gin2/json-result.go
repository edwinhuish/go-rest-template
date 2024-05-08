package gin2

import "net/http"

type JsonResult struct {
	ctx  *Context    `json:"-"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (res *JsonResult) SetCode(code int) *JsonResult {
	res.Code = code
	return res
}

func (res *JsonResult) SetMessage(message string) *JsonResult {
	res.Msg = message
	return res
}

func (res *JsonResult) SetData(data interface{}) *JsonResult {
	res.Data = data
	return res
}

func (res *JsonResult) Success(dataOptional ...interface{}) {

	if len(dataOptional) >= 1 {
		res.Data = dataOptional[0]
	}

	res.ctx.JSON(http.StatusOK, res)
}

func (res *JsonResult) Fail(errOrMsg interface{}) {

	res.Code = 1

	switch t := errOrMsg.(type) {
	case error:
		res.Msg = t.Error()
	case string:
		res.Msg = t
	default:
		panic("errOrMsg 仅支持 error 或 string")
	}

	res.ctx.JSON(http.StatusOK, res)
}

func (res *JsonResult) FailWithData(errOrMsg interface{}, data interface{}) {
	res.Data = data

	res.Fail(errOrMsg)
}
