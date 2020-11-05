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
var Articles := []Article



func homePage(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Welcome to the HomePage!"))
    //fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    /*w.Header().Set("Content-Type", "application/json")

    jData, err := json.Marshal(Articles)
        if err != nil {
            log.Println(err)
        }

    w.Write(jData)*/

    //json.NewEncoder(w).Encode(Articles)

    

    jData, err := json.Marshal([]Article{
        Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    })
    if err != nil {
        log.Println(err)
    }

    fmt.Println(string(jData))
    w.Write([]byte(string(jData)))
}

func returnOne(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(string(1)))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
        if article.ID == key {
            jData, err := json.Marshal(article)
                if err != nil {
                    log.Println(err)
                }
            w.Header().Set("Content-Type", "application/json")
            w.Write(jData)
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
    myRouter.HandleFunc("/one",returnOne)
    
    
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