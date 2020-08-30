package frontend

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler_Index(t *testing.T) {
	s := httptest.NewServer(Handler())
	defer s.Close()

	resp, err := http.Get(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	want, err := ioutil.ReadFile("web/frontend/index.html")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, respBody) {
		t.Errorf("invalid response body\nwant:\n%s\ngot:\n%s", want, respBody)
	}
}
