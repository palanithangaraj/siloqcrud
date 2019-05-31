package db

import (
	"siloqcrud/pkg/model"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type DataStore interface {
	ReadJoke(name *model.Name, joke *model.Joke) string
}

//Receiver and handler for all database operations.
type Dao struct {
	//TODO: add session context
        //databaseName   string
        collectionName string
}

func init() {
	//TODO: Init db connections, session here.
}

//Just mimics the DB layer
func (dao *Dao) ReadJoke(name *model.Name, joke *model.Joke) string {
	if name == nil || joke == nil {
		log.Fatalf("Error - reading date")
	}

	s:=fmt.Sprintf("%s %s's %s", name.Name, name.Surname, joke.Value.Joke)
	//TODO: handle db crud here as appropriate
	log.Info(s)
	return fmt.Sprintf(s)
}
