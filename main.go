package main

import (
	"fmt"
	"github.com/poovarasansivan/Golang/config"
    "iapi/routes"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	
)

func main(){
	config.ConnectDB();
	defer config.Database.Close()
    router := mux.NewRouter()
	router.HandleFunc("/", routes.Sample).Methods("POST")
   
}