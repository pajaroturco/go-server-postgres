package main

import (
	"log"
	"net/http"

	"GoWebserverPostgres/router"
	"GoWebserverPostgres/services"
	"GoWebserverPostgres/utils"
)

func main() {
	log.Println("In Main App")

	var dbconn = utils.ConnectDB()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Listening on Port 9090")
	log.Fatal(http.ListenAndServe(":9090", &appRouter))
}
