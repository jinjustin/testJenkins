package main

import (
	//"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"io/ioutil"
)

// Test Cases

func TestAPI(t *testing.T) {
	t.Run("HomePage", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
		 t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(homePage)
		handler.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("Return All Articles", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/all", nil)
		if err != nil {
		 t.Error(err)
		}


		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(returnAllArticles)
		handler.ServeHTTP(resp, req)
		data, err := ioutil.ReadAll(resp.Body)
		var value []Article

		if err == nil && data != nil {
   				json.Unmarshal(data, &value)
			}

		expected := []Article{
			Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
			Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		}
		b, err := json.Marshal(expected)
		if err != nil{
			//
		}

		assert.Equal(t,b,resp.Body)
	})
}



	