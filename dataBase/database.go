package dataBase

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"scouting_app/secrets"
)

type Firestore *firestore.Client

type DB struct {
	Db       *firestore.Client
	Year     int
	Code     string
	UserName string
	Schedule map[string]int
}

func CreateDB(year int, code string, userName string) *DB {
	db := &DB{Year: year, Code: code, UserName: userName}
	ctx := context.Background()
	sa := option.WithCredentialsJSON(*secrets.GetAdminKey())
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	db.Db, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
