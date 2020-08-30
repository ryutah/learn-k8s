package backend

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler(t *testing.T) {
	s := httptest.NewServer(Handler())
	defer s.Close()

	reqPayload := payload{
		Name: "hogehoge",
		Age:  123,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(reqPayload)

	req, _ := http.NewRequest(http.MethodPost, s.URL, buf)
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var respPaylod payload
	if err := json.Unmarshal(respBody, &respPaylod); err != nil {
		t.Log(string(respBody))
		t.Fatal(err)
	}

	if want, got := resp.StatusCode, http.StatusCreated; want != got {
		t.Errorf("invalid status code\n  want: %v, got: %v", want, got)
	}
	if want, got := reqPayload, respPaylod; !reflect.DeepEqual(want, got) {
		t.Errorf("invalid response body\n  want: %v\n  got: %v", want, got)
	}
}

func TestHandler_BadRequest(t *testing.T) {
	s := httptest.NewServer(Handler())
	defer s.Close()

	req, _ := http.NewRequest(http.MethodPost, s.URL, nil)
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var respPaylod errorResponse
	if err := json.Unmarshal(respBody, &respPaylod); err != nil {
		t.Log(string(respBody))
		t.Fatal(err)
	}

	if want, got := resp.StatusCode, http.StatusBadRequest; want != got {
		t.Errorf("invalid status code\n  want: %v, got: %v", want, got)
	}
	if respPaylod.Message == "" {
		t.Error("response body is empty")
	}
}
