package db

import (
        "testing"
	"siloqcrud/pkg/model"
)

func TestReadJoke(t *testing.T) {
        dao := InitializeMock()
	expected := ""
	name := model.Name{}
        joke := model.Joke{}
	joke.Value  = model.JokeItem{}
	joke.Value.Joke = ""
	name.Name = ""
	name.Surname = ""

	received := dao.ReadJoke(&name, &joke)
        if expected == received {
                t.Errorf("expected %v, got %v", expected, received)
        }
}
