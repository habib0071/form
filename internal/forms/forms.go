package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}


func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form {
		data,
		errors(map[string][]string{}),
	}
} 


func (f *Form) Has(feild string, r *http.Request) bool {
	x := r.Form.Get(feild)
	if x == "" {
		return false
	}
	return true
}