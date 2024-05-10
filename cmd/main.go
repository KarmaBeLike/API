package main

import (
	"fmt"
	"log"
	"net/http"

	"api/internal/database/postgres"
	"api/internal/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

func SignIn(ctx *gin.Context) { // sign-in

	// проверить пароль на его захэшированный пароль из бд
}

func signUp(ctx *gin.Context) {
	var users models.Users
	if err := ctx.BindJSON(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}
	// перед тем как сохранить пасс захещировать hashedpass

	err := postgres.InsertUsers(users.Login, //hashedpass)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	// mux := http.NewServeMux() // поменять на джин
	// mux.HandleFunc("/", home)
	// mux.HandleFunc("/sign-in", signIn)
	// mux.HandleFunc("/sign-up", signUp)
	// SetConnection()

	router := gin.Default()

	router.POST("/sign-in", SignIn) // localhost:8080/sign-up
	/*{
		"login":"",
		"password":""
	}
	*/

	db, err := postgres.SetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = postgres.CreateTable(db)
	if err != nil {
		fmt.Println("1")
		log.Panic(err)
	}

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("2")
		log.Println(err)
		log.Fatal(err)

	}
	log.Println("Starting server on 8080")
}
