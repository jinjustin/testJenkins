package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)
// Article is ...
type Article struct {
	ID string `json:"ID"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// Articles is ...
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
        if article.ID == key {
            json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.     
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

// Existing code from above
func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}",deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}",returnSingleArticle)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.ID == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Articles = []Article{
        Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}