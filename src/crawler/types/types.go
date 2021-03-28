package types

// Request 请求，定义了请求链接以及对应的解析方法
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult 解析结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser 返回空的解析结
func NilParser(body []byte) ParseResult {
	return ParseResult{}
}
