package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)


 


func WriteJSON(w http.ResponseWriter,statusCode int,body interface{})error{
w.Header().Add("Content-Type","application/json")
w.WriteHeader(statusCode)
if body!=nil{
if err:=json.NewEncoder(w).Encode(body);err!=nil{
	return err
}
}
return nil
}


type apiFunc func(w http.ResponseWriter, r *http.Request)error


func CreateHttpHandlerfunc(f apiFunc)http.HandlerFunc{
return func(w http.ResponseWriter, r*http.Request){
	if err:=f(w,r);err!=nil{
		if jsonErr:=WriteJSON(w,http.StatusInternalServerError,"Internal Server Error");jsonErr!=nil{
    log.Fatal(jsonErr)
		}
	 log.Fatal(err)
	}
}
} 


type APIServer struct{
	listenAddr string
  	
}


func NewApiServer()*APIServer{
return &APIServer{listenAddr:":8080"}
}



func (s *APIServer) run()error{
	r:=mux.NewRouter()
	r.Handle("/AdRequest",CreateHttpHandlerfunc(s.bookSpot)).Methods("GET")
	log.Printf("server running at port%s",s.listenAddr)
  if err:=http.ListenAndServe(s.listenAddr,r);err!=nil{
		return err
	}
	return nil
}



type AdObject struct{
	AdId uuid.UUID  `json:"adid"`
	Bidprice int    `json:"bidprice"` 
}


func (s *APIServer) bookSpot(w http.ResponseWriter,r *http.Request)error{ 
//generates an  int between 1 and 10 both incl
randomInt:=rand.Intn(10)+1
//send the Add request if the randomInt is not 3 or 5 or 7 
if randomInt!=3{
newAdObject:=AdObject{AdId: uuid.New(),Bidprice: (rand.Intn(100)+1)}
if err:=WriteJSON(w,http.StatusOK,newAdObject);err!=nil{
	return err
}}else{
	if err:=WriteJSON(w,http.StatusNoContent,nil);err!=nil{
	return err
}
}
return nil
}


