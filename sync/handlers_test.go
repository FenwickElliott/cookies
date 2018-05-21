package sync

import (
	"net/http/httptest"
	"testing"

	"gopkg.in/mgo.v2"
)

var testServiceCookie string

func init() {
	service = Service{
		Name: "testService",
	}
	session, err := mgo.Dial("127.0.0.1")
	check(err)
	service.c = session.DB(service.Name).C("test")
	service.c.RemoveAll(nil)
	check(err)
}

func TestIn(t *testing.T) {
	req := httptest.NewRequest("HEAD", "/in?partner=testPartner&cookie=abc123", nil)
	w := httptest.NewRecorder()

	service.in(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("in did not return 200 status code")
	}

	for _, c := range resp.Cookies() {
		if c.Name == "testServiceID" {
			testServiceCookie = c.Value
		}
	}
	if testServiceCookie == "" {
		t.Errorf("no service cookie set")
	}
}
