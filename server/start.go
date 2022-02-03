package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Server struct {}

func NewServer() *http.Server {
	router := mux.NewRouter()
	s := Server{}

	router.HandleFunc("/", s.home).Methods(http.MethodGet)
	router.HandleFunc("/far", s.farFromHome).Methods(http.MethodGet)

	fmt.Println("Your server is up!!!")
	return &http.Server{
		Addr: ":8094",
		Handler: router,
	}
}

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	SayHelloGo(ctx, w)

}



func (s *Server) farFromHome(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 9 * time.Second)
	defer cancel()

	SayByeGo(ctx, w)
}




func SayHelloGo(ctx context.Context, w http.ResponseWriter) {

	select {
	case <-time.After(6 * time.Second):
		fmt.Fprintf(w, "Hello, %s!", "Go")
		defer log.Printf("Did the job!")
	case <-	ctx.Done():
		log.Print(ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}


func SayByeGo(ctx context.Context, w http.ResponseWriter) {

	deadline, _ := ctx.Deadline()

	select {
	case <-time.After(8 * time.Second):
		fmt.Fprintf(w, "Bye %s!", "Go")
		defer log.Printf("Did the job before %v!", deadline)
	case <-ctx.Done():
		log.Print(ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
