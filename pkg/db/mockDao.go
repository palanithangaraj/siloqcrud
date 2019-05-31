package db

import (
	"siloqcrud/pkg/model"
	 "fmt"
        log "github.com/sirupsen/logrus"
)

type MockDao struct {

}

//InitializeMock creates a Mock Dao instance that can be used in unit tests
//where the interface is used for DI
func InitializeMock() *MockDao {
	return &MockDao{}
}

//Just mimics the DB layer
func (dao *MockDao) ReadJoke(name *model.Name, joke *model.Joke) string {
        if name == nil || joke == nil {
                log.Fatalf("Error - reading date")
        }

        s:=fmt.Sprintf("%s %s's %s", name.Name, name.Surname, joke.Value.Joke)
        //TODO: handle db crud here as appropriate
        log.Info(s)
        return fmt.Sprintf(s)
}

