package links

import (
	"log"

	"github.com/glyphack/go-graphql-hackernews/graph/internal/users"
	database "github.com/glyphack/go-graphql-hackernews/internal/pkg/db/mysql"
)

// #1
type Link struct {
	ID   string
	Text string
	Done bool
	User *users.User
}

// #2
func (link Link) Save() int64 {
	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}
