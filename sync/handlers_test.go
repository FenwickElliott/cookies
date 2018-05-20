package sync

import (
	"fmt"
	"net/http/httptest"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func init() {
	service = Service{
		Name: "testService",
	}
	session, err := mgo.Dial("127.0.0.1")
	check(err)
	service.c = session.DB(service.Name).C("test")
	service.c.RemoveAll(nil)
}

func TestIn(t *testing.T) {
	req := httptest.NewRequest("HEAD", "/in", nil)
	w := httptest.NewRecorder()

	service.in(w, req)
	resp := w.Result()

	fmt.Println(resp)
}
