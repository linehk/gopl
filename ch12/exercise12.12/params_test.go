package params

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestParams(t *testing.T) {
	type data struct {
		Email string `validate:"email"`
		Phone string `validate:"phone"`
		Zip   int    `validate:"zip"`
	}
	tests := []struct {
		req  *http.Request
		want data
	}{
		{
			&http.Request{Form: url.Values{
				"email": []string{"sulinehk@gmail.com"},
				"phone": []string{"15778989527"},
				"zip":   []string{"546000"},
			}},
			data{
				"sulinehk@gmail.com",
				"15778989527",
				546000,
			},
		},
		{
			&http.Request{Form: url.Values{
				"email": []string{"aaaaaaaaa"},
				"phone": []string{"aaaaaaa"},
				"zip":   []string{"1111111"},
			}},
			data{
				"",
				"",
				0,
			},
		},
	}
	for i, tt := range tests {
		var got data
		if err := Unpack(tt.req, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
