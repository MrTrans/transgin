package transgin

import "net/http"

// Engine 返回引擎
type Engine struct {
	TIRouter
}

// 实现请求

// New 新建引擎
func New() *Engine {
	engine := &Engine{}
	return engine
}

// 尝试封装http请求实现

// GET 获取资源
func (eng *Engine) GET(relativePath string, handlers ...handleFunc) {

}

// POST 传送资源
func (eng *Engine) POST(relativePath string, handlers ...handleFunc) {

}

// Run 运行地址
func (eng *Engine) Run(addr string) {
	http.ListenAndServe(addr, eng)
}

func (eng *Engine) ServeHTTP(respw http.ResponseWriter, req *http.Request) {
	respw.Write([]byte("hello transgin!"))
}
