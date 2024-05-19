package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type Credentials struct {
	Password string `json:"password",db:"password"`
	Username string `json:"username",db:"username"`
}

// type Credentials struct {
// 	Password string `json:"password"`
// 	Username string `json:"username"`
// }

func Signup(w http.ResponseWriter, req *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(req.Body).Decode(creds)
	fmt.Println(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("insert into users values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(creds.Username, string(hashPassword))
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!", id)
}

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:admin123@/pwdauth")
	if err != nil {
		panic(err)
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {

	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Bad Request 1 ")
		return
	}

	stmt, err := db.Prepare("select password from users where username=?")
	if err != nil {
		log.Fatal(err)
	}

	storedCreds := &Credentials{}
	err = stmt.QueryRow(&creds.Username).Scan(&storedCreds.Password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(storedCreds.Password)
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("StatusUnauthorized")
		return
	}

	log.Println("Successfully able to login")
}

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	InitDB()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
