package httplib

type str string

type handlerFunc func(rw, rq string)

func (h handlerFunc) ServeHTTP(responseWriter, request string) {
	h(responseWriter, request)
}

type iHandler interface {
	ServeHTTP(responseWriter, request string)
}

func (s str) HandleFunc(pattern string, handler func(rw, rq string)) {
	s.handle(pattern, handlerFunc(handler))
}

func (s str) handle(pattern string, ihandler iHandler) {

}
