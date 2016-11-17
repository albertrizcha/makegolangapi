package Handle

import (
	"log"
	"net/http"

	"github.com/albertrizcha/makegolangapi/Func"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Func.HomePage)
	myRouter.HandleFunc("/all", Func.ReturnAllArticles)
	myRouter.HandleFunc("/key/{key}", Func.ReturnOneArticle)
	myRouter.HandleFunc("/key/{key}/{var1}/{var2}/", Func.ReturnOneArticles)
	myRouter.HandleFunc("/testJson", Func.TestJSON).Methods("POST")
	myRouter.HandleFunc("/json", Func.TestJson)

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
