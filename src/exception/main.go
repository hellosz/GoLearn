package main

import (
	"exception/filelist"
	"log"
	"net/http"
	"os"
)

type pathError interface {
	Error() string
	Message() string
}

type myFuncHandler func(http.ResponseWriter, *http.Request) error

// 错误处理程序
func ErrorWrapper(handler myFuncHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// panic处理
		defer func() {
			if err := recover(); err != nil {
				// 记录日志
				log.Printf("Server Internal Error, Message:%v", err)

				// 范会报错结果
				code := http.StatusInternalServerError
				http.Error(writer, http.StatusText(code), code)
			}
		}()

		err := handler(writer, request)

		// 自定义异常处理
		if err, ok := err.(pathError); ok {
			http.Error(writer, err.Message(), http.StatusBadRequest)
			return
		}

		// 处理非法请求类型的错误
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusUnauthorized
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}

}

func main() {
	log.Print("start listen to 9090 port....")

	// 处理请求
	http.HandleFunc("/", ErrorWrapper(filelist.BusinessHandler))

	// 处理服务器监听错误
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
