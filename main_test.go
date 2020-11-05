package main

import (
	//"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

// Test Cases

func Test_HomePage(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/")
	if err != nil {
		// handle `err`
	}
	defer res.Body.Close()

	if err != nil {
	   t.Errorf("Expected nil, received %s", err.Error())
	}
	
	resData, err := ioutil.ReadAll(res.Body)
    if err != nil {
        //log.Fatal(err)
    }
	resString := string(resData)
	
	expect := "Welcome to the HomePage!"

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, expect, resString)
 }

func Test_ReturnAllArticle(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/all", returnAllArticles)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Return All Article", func(t *testing.T) {
	   res, err := http.Get(ts.URL + "/all")
	   if err != nil {
		// handle `err`
	}
	   defer res.Body.Close()

	   if err != nil {
		  t.Errorf("Expected nil, received %s", err.Error())
	    }

	   if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	 	}

		/*resData, err := ioutil.ReadAll(res.Body)
		if err != nil {
		}*/

	var article []Article
	
	decodeErr := json.NewDecoder(res.Body).Decode(&article)
	if decodeErr != nil {
		//if error do something
	}

		expect := []Article{
			Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
			Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		}

		assert.Equal(t, expect, article)
		
	})
 }

func Test_ReturnSingleArticle(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/article/{id}", returnAllArticles)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Return Article 1", func(t *testing.T) {
	   res, err := http.Get(ts.URL + "/article/1")
	   if err != nil {
		// handle `err`
		}
	   if err != nil {
		  t.Errorf("Expected nil, received %s", err.Error())
	    }
	   if res.StatusCode != http.StatusOK {
		  t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	    }
	})

	t.Run("Return Article 3", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/article/3")
		if err != nil {
		   t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
		   t.Errorf("Expected %d, received %d", http.StatusNotFound, res.StatusCode)
		}
	 })
 }