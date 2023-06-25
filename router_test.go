package router

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_GET(t *testing.T) {
	testcases := []struct {
		name     string
		endpoint string
		handler  http.Handler
	}{
		{
			name:     "/のエンドポイントにハンドラを追加する",
			endpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/helloのエンドポイントにハンドラを追加する",
			endpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hogeのエンドポイントにハンドラを追加する",
			endpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
		{
			name:     "/hoge/fugaのエンドポイントにハンドラを追加する",
			endpoint: "/hoge/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			}),
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			r.GET(testcase.endpoint, testcase.handler)
			handler := r.Search(http.MethodGet, testcase.endpoint)
			// 関数のポインタを比較する
			if reflect.ValueOf(handler).Pointer() != reflect.ValueOf(testcase.handler).Pointer() {
				t.Errorf("ハンドラが異なります\nexpected: %v\nactual:%v", testcase.handler, handler)
			}
		})
	}
}
