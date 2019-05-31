package db

import (
	"siloqcrud/pkg/model"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type DataStore interface {
	ReadJoke(name *model.Name, joke *model.Joke) string
}

type Dao struct {

}

//Just mimics the DB layer
func (dao *Dao) ReadJoke(name *model.Name, joke *model.Joke) string {
	if name == nil || joke == nil {
		log.Fatalf("Error - reading date")
	}

	s:=fmt.Sprintf("%s %s's %s", name.Name, name.Surname, joke.Value.Joke)
	log.Info(s)
	return fmt.Sprintf(s)
}
