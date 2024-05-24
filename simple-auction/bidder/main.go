package main

import(
//"github.com/joho/godotenv"
"log"
)




func main(){
/* 	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  } */
server:=NewApiServer()
if err:=server.run();err!=nil{
	log.Fatal(err)
}
	
}