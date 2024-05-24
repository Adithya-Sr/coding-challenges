package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	//"os"
	"strconv"
	"time"
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
//os.Getenv("PORT") doesnt work in go-alpine image
return &APIServer{listenAddr:":3000"}
}



func (s *APIServer) run()error{
	r:=mux.NewRouter()
	r.Handle("/AdPlacement",CreateHttpHandlerfunc(s.addPlacement)).Methods("GET")
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
 

func (s *APIServer) addPlacement(w http.ResponseWriter, r *http.Request)error{
placementId,err:=strconv.Atoi(r.URL.Query().Get("placementId"))
if err!=nil{
	if err:=WriteJSON(w,http.StatusBadRequest,"bad request");err!=nil{
		return err
	}
	return nil
}
client := http.Client{
Timeout: time.Millisecond*200,  }
bidderRespChan:=make(chan interface{},10)
bidderResp:=make([]interface{},10)
highestBid:=AdObject{Bidprice: 0}
for i:=0;i<10;i++{
go getBids(client,bidderRespChan)
}
for i:=0;i<10;i++{
bidderResp[i]=<-bidderRespChan
log.Println(bidderResp[i])
if val,ok:=bidderResp[i].(AdObject);ok{
	if highestBid.Bidprice<val.Bidprice{
    highestBid=val
	}
}
}

WriteJSON(w,http.StatusOK,map[string]interface{}{"bid":highestBid,"placementId":placementId,})
return nil
}




func getBids(client http.Client,respChan chan interface{}){
resp,err:=client.Get("http://localhost:8080/AdRequest")
	if err!=nil{
		if netErr,ok:=err.(net.Error);ok && netErr.Timeout(){
			respChan<-fmt.Errorf("network timeout:%v",netErr)
			return
		}else{
    respChan<-err
		return}}
if resp.ContentLength == 0 {
        respChan <- fmt.Sprintln("empty response body")
        return
}
var decodedResp AdObject
defer resp.Body.Close()
if err := json.NewDecoder(resp.Body).Decode(&decodedResp);err!=nil{
	respChan<-err
	return
}
respChan<-decodedResp
}



