package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("bye world")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "ooh ooops", http.StatusBadRequest)
		return
	}
	b.l.Println(fmt.Sprintf("serving %s", d))

	fmt.Fprintf(rw, "bye %s \n", d)
}
