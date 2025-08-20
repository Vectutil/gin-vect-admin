package response

import (
	"errors"
	"fmt"
	"gin-vect-admin/pkg/robot"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

type buildConfig struct {
	data       interface{}
	sourceData interface{}
	code       int32
	msg        string
	httpCode   int
	sendErrMsg bool
}

type BuildOption func(*buildConfig)

func WithData(data interface{}) BuildOption {
	return func(c *buildConfig) {
		c.data = data
	}
}

func WithErrCode(code int32) BuildOption {
	return func(c *buildConfig) {
		c.code = code
	}
}

func WithSourceData(data interface{}) BuildOption {
	return func(c *buildConfig) {
		c.sourceData = data
	}
}

func WithHTTPCode(code int) BuildOption {
	return func(c *buildConfig) {
		c.httpCode = code
	}
}

func WithSendErrMsg() BuildOption {
	return func(c *buildConfig) {
		c.sendErrMsg = true
	}
}

type Response struct {
	Code int         `json:"code"`           // 业务码
	Msg  string      `json:"msg"`            // 提示信息
	Data interface{} `json:"data,omitempty"` // 数据内容（成功时）
}

type qwError struct {
	TimeStamp int64       `json:"timeStamp"`
	Code      int         `json:"code"`
	Api       string      `json:"api"`
	Msg       string      `json:"msg"`
	Stack     string      `json:"stack"`
	Request   interface{} `json:"request"`
}

func NewError(code int, msg string) error {
	return &appError{code: code, msg: msg}
}

// appError 在项目中定义统一错误类型
type appError struct {
	code int
	msg  string
}

func (e *appError) Code() int {
	return e.code
}
func (e *appError) Error() string {
	return e.msg
}

// Success 成功返回
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

// HandleDefault ，返回延迟处理函数
func HandleDefault(ctx *gin.Context, opts ...BuildOption) func(*error, any) {
	// 定义延迟处理函数
	handler := func(err *error, r any) {
		conf := &buildConfig{}

		if opts != nil {
			for _, opt := range opts {
				opt(conf)
			}
		}
		if r != nil {
			*err = errors.New(fmt.Sprintf("%v", r))
		}
		if *err != nil {
			errVal := fmt.Sprintf("%+v", *err)
			code := http.StatusInternalServerError
			var e *appError
			if errors.As(*err, &e) {
				code = e.Code()
				if e.Error() != "" {
					errVal = e.Error()
				}
			}
			if conf.sendErrMsg {
				sendErrMsg(ctx, code, err)
			}
			ctx.JSON(http.StatusOK, Response{
				Code: code,
				Msg:  errVal,
			})
			return
		}
		if conf.sourceData != nil {
			Success(ctx, conf.sourceData)
			return
		}
		Success(ctx, conf.data)
	}

	return handler
}

//func HandleListDefault(ctx *gin.Context, res common.IBaseListResp) func(*error, any) {
//	// 定义延迟处理函数
//	handler := func(err *error, r any) {
//		if r != nil {
//			*err = errors.New(fmt.Sprintf("%v", r))
//		}
//		if *err != nil {
//			resValue := fmt.Sprintf("%v", res)
//			code := http.StatusInternalServerError
//			var e *appError
//			if errors.As(*err, &e) {
//				code = e.Code()
//				if e.Error() != "" {
//					resValue = e.Error()
//				}
//			}
//			Error(ctx, err, code, resValue)
//			return
//		}
//		res.Adjust()
//		Success(ctx, res)
//	}
//
//	return handler
//}

func Stack(err error) string {
	stack := string(debug.Stack())
	// 先替换 \n\t 组合
	all := ">" + strings.ReplaceAll(stack, "\n\t", "\n>")
	return all
}

func sendErrMsg(c *gin.Context, code int, err *error) {
	stack := Stack(*err)
	markdown := fmt.Sprintf(
		`
## 🚨 实时新增接口异常，请相关同事注意 \n
> **时间**：%d  
> **接口**：%s  
> **状态码**：%d  
> **错误信息**：%v

### 📚 堆栈：
%s`, time.Now().UnixNano(), c.Request.URL.Path, code, fmt.Sprintf("%+v", *err), stack)

	robot.CallQWAssistant(c, markdown, robot.QWRobotMsgTypeText)
}
