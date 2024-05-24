package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)




func TestBidder(t *testing.T){
s:=NewApiServer()
testServer:=httptest.NewServer(CreateHttpHandlerfunc(s.bookSpot))
defer testServer.Close()
resp,err:=http.Get(fmt.Sprintf("%s/AdRequest",testServer.URL))
if err!=nil{
	t.Fatalf("couldn't send request,%v",err)
}
fmt.Println(resp.StatusCode)
defer resp.Body.Close()
if resp.StatusCode==204{
	if resp.ContentLength!=0{
		t.Error("response body expected to be empty for status:204")
	}
}else if resp.StatusCode==200{
	if resp.ContentLength==0{
		t.Errorf("expected response body for status:200")
	}
body:=AdObject{}
if err:=json.NewDecoder(resp.Body).Decode(&body);err!=nil{
	t.Fatalf("error in decoding response:%v",err)
}
fmt.Println(body)
if body.AdId==uuid.Nil{
	t.Errorf("expected valid AdId, received:%v",body.AdId)
}
if body.Bidprice<1{
	t.Errorf("expected bidprice to be positive,received:%d",body.Bidprice)
}
}else{
	t.Fatalf("unexpected status code received:%d",resp.StatusCode)
}
}