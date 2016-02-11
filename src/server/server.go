package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/bmizerany/pat"
	"time"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func add(x int , y int) int {
return  x+y
}

func Log(handler http.Handler) http.Handler {
log.Println(handler)
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
log.Printf("%s %s %s &s", r.RemoteAddr, r.Method, r.URL)
handler.ServeHTTP(w, r)
})
}

func main() {


router := pat.New()
router.Get("/",http.HandlerFunc(Index))
router.Get("todos",http.HandlerFunc(TodoIndex))
router.Get("/todos/:todoId",http.HandlerFunc(TodoShow))
log.Println("Listening...")

s := &http.Server{
Addr:           ":8080",
Handler:        Log(router),
ReadTimeout:    10 * time.Second,
WriteTimeout:   10 * time.Second,
MaxHeaderBytes: 1 << 20,
}
//	log.Fatal(s.ListenAndServe())

err := s.ListenAndServe()
if err != nil{
fmt.Println("Unable to create server")
}

}

func Index(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
vars := r.URL.Query()
id := vars.Get(":todoId")
fmt.Fprintln(w,id)
}
