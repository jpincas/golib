package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func GetFormValue(r *http.Request, key string) string {
	r.ParseForm()
	return r.Form.Get(key)
}

func GetFormValues(r *http.Request, key string) []string {
	r.ParseForm()
	return r.Form[key]
}

func GetFormValueBool(r *http.Request, key string) bool {
	r.ParseForm()
	boolS := strings.ToLower(r.Form.Get(key))
	if boolS == "true" ||
		boolS == "t" ||
		boolS == "yes" ||
		boolS == "y" {
		return true
	}

	return false
}

func TryDecodeResponseBody(r *http.Response, target interface{}) ([]byte, error) {
	if r.Body == nil {
		return []byte{}, errors.New("empty body request")
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return b, err
	}

	if b == nil {
		return b, errors.New("empty body request")
	}

	if err := json.Unmarshal(b, target); err != nil {
		return b, err
	}

	return b, nil
}

func TryDecodeRequestBody(r *http.Request, target interface{}) ([]byte, error) {
	if r.Body == nil {
		return []byte{}, errors.New("empty body request")
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return b, err
	}

	if b == nil {
		return b, errors.New("empty body request")
	}

	if err := json.Unmarshal(b, target); err != nil {
		return b, err
	}

	return b, nil
}

func TryDecodeBody(w http.ResponseWriter, r *http.Request, target interface{}) ([]byte, error) {
	return TryDecodeRequestBody(r, target)
}

func TryDecodeBodyCopy(w http.ResponseWriter, r *http.Request, target interface{}) ([]byte, error) {
	if r.Body == nil {
		return []byte{}, errors.New("empty body request")
	}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return b, err
	}

	if b == nil {
		return b, errors.New("empty body request")
	}

	if err := json.Unmarshal(b, target); err != nil {
		return b, err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	return b, nil
}

// A super simple reverse proxy, which just routes request to one endpoint
// irrespective of the path
func SimpleReverseProxy(target *url.URL) *httputil.ReverseProxy {
	director := func(r *http.Request) {
		r.URL = target
		r.Host = target.Host

		// Update headers to allow for SSL redirection
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		if _, ok := r.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			r.Header.Set("User-Agent", "")
		}
	}

	return &httputil.ReverseProxy{Director: director}
}

func OptionalDecode(w http.ResponseWriter, r *http.Request, target interface{}) error {
	if r.Body == nil {
		return nil
	}

	_, err := TryDecodeBody(w, r, target)
	return err
}

func GetURLParam(r *http.Request, s string) string {
	return chi.URLParam(r, s)
}

func MustGetURLParam(r *http.Request, s string) (string, error) {
	param := GetURLParam(r, s)

	var err error
	if param == "" {
		err = fmt.Errorf("the url paramater %s was not present", s)
	}

	return param, err
}

func MustGetURLIDParam(r *http.Request, s string) (primitive.ObjectID, error) {
	param := GetURLParam(r, s)

	if param == "" {
		return primitive.ObjectID{}, fmt.Errorf("the url paramater %s was not present", s)
	}

	eid, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		return eid, err
	}

	return eid, nil
}
