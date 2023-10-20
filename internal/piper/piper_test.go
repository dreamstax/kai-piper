package piper_test

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/dreamstax/kai-piper/internal/piper"
)

func TestPipe(t *testing.T) {
	p := piper.NewPipe(adder, adder, adder)
	h := p.Then(final())

	var i uint64 = 0
	data := []byte{}
	data = binary.BigEndian.AppendUint64(data, i)
	rdr := bytes.NewReader(data)
	rc := io.NopCloser(rdr)

	w := bytes.NewBuffer([]byte{})
	r := &piper.Request{
		Body: rc,
	}

	h.Handle(w, r)

	res, err := io.ReadAll(w)
	if err != nil {
		t.Fatal(err)
	}

	v := binary.BigEndian.Uint64(res)
	log.Println(v)
}

func adder(h piper.Handler) piper.Handler {
	return piper.HandlerFunc(func(w piper.ResponseWriter, r *piper.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// transform
		v := binary.BigEndian.Uint64(body)
		v++

		// prepare
		out := binary.BigEndian.AppendUint64([]byte{}, v)
		rdr := bytes.NewReader(out)
		rc := io.NopCloser(rdr)
		r.Body = rc

		h.Handle(w, r)
	})
}

// TODO  should this be a built in
func final() piper.Handler {
	return piper.HandlerFunc(func(w piper.ResponseWriter, r *piper.Request) {
		io.Copy(w, r.Body)
	})
}

func benchmarkFuncPipe(b *testing.B, h piper.Handler, w piper.ResponseWriter, r *piper.Request) {
	h.Handle(w, r)
}

func BenchmarkFuncPipe(b *testing.B) {
	p := piper.NewPipe(adder, adder, adder)
	h := p.Then(final())

	var i uint64 = 0
	data := []byte{}
	data = binary.BigEndian.AppendUint64(data, i)
	rdr := bytes.NewReader(data)
	rc := io.NopCloser(rdr)

	w := bytes.NewBuffer([]byte{})
	r := &piper.Request{
		Body: rc,
	}

	benchmarkFuncPipe(b, h, w, r)
}

func TestHTTPStep(t *testing.T) {
	thandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hai")
	})
	tServer := httptest.NewServer(thandler)
	defer tServer.Close()

	turl, err := url.Parse(tServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	httpStep := piper.HTTPStep(
		http.MethodPost,
		turl,
		http.DefaultClient,
	)

	p := piper.NewPipe(httpStep)
	h := p.Then(final())

	w := bytes.NewBuffer([]byte{})
	r := &piper.Request{
		Body: io.NopCloser(bytes.NewBuffer([]byte{})),
	}

	h.Handle(w, r)

	res, err := io.ReadAll(w)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(string(res))
}

type payload struct {
	Value int
}

func adderServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := &payload{}
		err := json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		p.Value++
		json.NewEncoder(w).Encode(p)
	}))
}

func TestHTTPAdder(t *testing.T) {
	s := adderServer()
	url, err := url.Parse(s.URL)
	if err != nil {
		t.Fatal(err)
	}

	httpStep := piper.HTTPStep(
		http.MethodPost,
		url,
		http.DefaultClient,
	)
	p := piper.NewPipe(httpStep, httpStep, httpStep, httpStep)
	h := p.Then(final())

	w := bytes.NewBuffer([]byte{})
	data, err := json.Marshal(&payload{0})
	if err != nil {
		t.Fatal(err)
	}
	r := &piper.Request{
		Body: io.NopCloser(bytes.NewBuffer(data)),
	}

	h.Handle(w, r)

	res, err := io.ReadAll(w)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(string(res))
}
