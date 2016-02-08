package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func add(x int , y int) int {
	return  x+y
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s &s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

//func main() {
//	r:= mux.NewRouter()
//	r.HandleFunc("/request", handleRequest)
//
//	var x float64 = 20.0 //static type declaration
//
//	y := 42 //dynamic type declaration
//
//	var a, b, c = 3, 4, "foo"  //type interference a=3 b=4 c= foo
//	var g *int
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Printf("a is of type %T\n", a)
//	fmt.Printf("b is of type %T\n", b)
//	fmt.Printf("c is of type %T\n", c)
//
//	fmt.Println(x)
//	fmt.Println(y)
//	fmt.Printf("x is of type %T\n", x)
//	fmt.Printf("y is of type %T\n", y)
//	fmt.Println(10!=20)
//	fmt.Println(&a)
//	g=&b
//	fmt.Println(g)
//
////	http.HandleFunc("/",handleRequest)
////	http.HandleFunc("/world",handleRequest)
//
//	log.Fatal(http.ListenAndServe(":8484",nil))
////	var x float64;0
////	x = 5.32322
////	fmt.Println(x);
////	fmt.Printf("The x of type %T\n",x)
////	fmt.Println(math.Pi)
////	fmt.Println("The time is %g and have a good day" , time.Now())
////	fmt.Println("Function Output is ",add(45,25))
//}
//
//func  handleRequest(w http.ResponseWriter  ,r *http.Request)  {
//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//}
//

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", Log(http.DefaultServeMux)))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}