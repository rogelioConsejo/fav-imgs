package server

import (
	"errors"
	"net/http"
	"strings"
)

func checkForId(r *http.Request) (err error) {
	if !r.URL.Query().Has(idKey) {
		err = errors.New("id not found")
	}
	return err
}

func getId(r *http.Request) (id string, err error) {
	id = r.URL.Query().Get(idKey)
	if strings.Trim(id, " ") == "" {
		err = errors.New("id parameter incorrect")
	}
	return id, err
}

func checkForUrl(r *http.Request) (err error) {
	if !r.URL.Query().Has(urlKey) {
		err = errors.New("url not found")
	}
	return err
}

func checkForTitle(r *http.Request) (err error) {
	if !r.URL.Query().Has(titleKey) {
		err = errors.New("title not found")
	}
	return err
}

const idKey = "id"
const titleKey = "title"
const urlKey = "url"

const root = "http://localhost:8080"
