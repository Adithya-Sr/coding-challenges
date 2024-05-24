package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuctionApi(t *testing.T){
s:=NewApiServer()
testPlacementId:=5
//setting up test server
testServer:=httptest.NewServer(CreateHttpHandlerfunc(s.addPlacement))
defer testServer.Close()
resp,err:=http.Get(fmt.Sprintf("%s/AdPlacement?placementId=%d",testServer.URL,testPlacementId))
fmt.Printf("%s/AdPlacement?placementId=%d\n",testServer.URL,testPlacementId)
if err!=nil{
t.Fatalf("couldn't send request:%v",err)
}
if resp.StatusCode==200{
if resp.ContentLength==0{
	t.Fatalf("response body not found")
}
defer resp.Body.Close()
body:= make(map[string]interface{},2)
err=json.NewDecoder(resp.Body).Decode(&body)
if err!=nil{
	t.Fatalf("couldn't decode response body%v",err)
}
if body["placementId"].(float64)!=float64(testPlacementId){
t.Errorf("sent and received placementIds dont match,expected:%d,received:%d",testPlacementId,body["placementId"])
}
}else{
	t.Fatalf("unexpected statuscode received:%d",resp.StatusCode)
}
}