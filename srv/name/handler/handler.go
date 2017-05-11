package handler

import (
	"log"

	proto "github.com/radekdymacz/tca/srv/name/proto"
	context "golang.org/x/net/context"
)

//Joker struct implements handler interface
type Namer struct{}

//ChuckNorris fetches random joke from http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]
func (n *Namer) Random(ctx context.Context, req *proto.Request, rsp *proto.Response) error {

	rsp.Name = "test"

	log.Printf("Value: %v", rsp.String())
	return nil
}
