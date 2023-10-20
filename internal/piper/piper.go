package piper

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type Request struct {
	Body io.ReadCloser
}

type ResponseWriter interface {
	Write([]byte) (int, error)
}

type Handler interface {
	Handle(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) Handle(w ResponseWriter, r *Request) {
	f(w, r)
}

type Step func(Handler) Handler

type Pipe struct {
	steps []Step
}

func NewPipe(steps ...Step) Pipe {
	return Pipe{append(([]Step)(nil), steps...)}
}

func (p Pipe) Then(h Handler) Handler {
	for i := range p.steps {
		h = p.steps[len(p.steps)-1-i](h)
	}
	return h
}

func HTTPStep(method string, url *url.URL, client *http.Client) Step {
	return func(h Handler) Handler {
		return &httpHandler{
			handler: h,
			method:  method,
			url:     url,
			client:  client,
		}
	}
}

type httpHandler struct {
	handler Handler
	method  string
	url     *url.URL
	client  *http.Client
}

func (h *httpHandler) Handle(w ResponseWriter, r *Request) {
	res, err := h.client.Do(&http.Request{
		URL:    h.url,
		Method: h.method,
		Body:   r.Body,
	})
	if err != nil {
		// TODO write error
		log.Println(err)
	}
	// TODO handle status code
	r.Body = res.Body
	h.handler.Handle(w, r)
}
