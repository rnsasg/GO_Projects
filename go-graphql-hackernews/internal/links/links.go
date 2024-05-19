package links

import (
	database "command-line-arguments/Users/kumarro/go/src/Go_Projects/go-graphql-hackernews/internal/pkg/db/mysql/mysql.go"
	"log"

	database "github.com/rnsasg/Go_Projects/go-graphql-hackernews/internal/pkg/db/mysql"
	"github.com/rnsasg/Go_Projects/go-graphql-hackernews/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}
