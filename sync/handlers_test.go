package sync

import (
	"encoding/json"
	"io/ioutil"
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
		t.Errorf("/in did not return 200 status code")
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

func TestPrint(t *testing.T) {
	req := httptest.NewRequest("GET", "/print", nil)
	w := httptest.NewRecorder()

	service.print(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Error("/print did not return 200 status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	body = body[1:74]

	var res map[string]string
	err = json.Unmarshal(body, &res)
	check(err)

	if res["_id"] != testServiceCookie {
		t.Errorf("returned service cookie didn't match")
	}
	if res["testPartner"] != "abc123" {
		t.Errorf("returned partner cookie didn't match")
	}
}
