package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type pathError string

func (p pathError) Message() string {
	return string(p)
}

func (p pathError) Error() string {
	return p.Message()
}

func BusinessHandler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, "/list/") != 0 {
		return pathError("不合法的请求路径" + request.URL.Path)
	}

	// 自定义异常
	// if true {
	// 	panic("自定义panic")
	// }

	path := request.URL.Path[len("/list/"):]

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(content)
	return nil
}
