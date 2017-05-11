package handler

import (
	"log"

	proto "github.com/radekdymacz/tca/srv/joke/proto"
	context "golang.org/x/net/context"
)

//Joker struct implements handler interface
type Joker struct{}

//ChuckNorris fetches random joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]
func (j *Joker) ChuckNorris(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	//Set defaults if request empty TODO move it out of the ndlaer into helper func
	if req.FirstName == "" {
		req.FirstName = "John"
	}
	if req.LastName == "" {
		req.LastName = "Doe"
	}
	if req.LimitTo == "" {
		req.LimitTo = "nerdy"
	}

	//TODO remove test data and implement calling the service

	s := make([]string, 1)
	s[0] = "test"
	rsp.Type = "test"
	rsp.Value = &proto.ChuckNorrisJoke{
		Id:         1,
		Joke:       "hello",
		Categories: s,
	}

	log.Printf("Value: %v", rsp.String())
	return nil
}
