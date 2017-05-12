package handler

import (
	"log"

	client "github.com/micro/go-micro/client"
	api "github.com/micro/micro/api/proto"
	joke "github.com/radekdymacz/tca/srv/joke/proto"
	name "github.com/radekdymacz/tca/srv/name/proto"
	context "golang.org/x/net/context"
)

//Joke struct to pass client
type Joke struct {
	Client client.Client
}

//Get random name and joke
func (a *Joke) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received API.Joke request")
	//call name and joke service

	nameSrvClient := name.NewNameClient("go.micro.srv.name", a.Client)
	nameSrvRes, err := nameSrvClient.Random(ctx, &name.Request{})
	if err != nil {
		return err
	}
	//preapre req and call random joke service
	jokeSrvReq := joke.Request{
		FirstName: nameSrvRes.GetName(),
		LastName:  nameSrvRes.GetSurname(),
	}
	jokeSrvClient := joke.NewJokeClient("go.micro.srv.joke", a.Client)
	jokeSrvRes, err := jokeSrvClient.ChuckNorris(ctx, &jokeSrvReq)
	if err != nil {
		return err
	}
	joke := jokeSrvRes.GetValue()

	rsp.StatusCode = 200
	rsp.Body = joke.GetJoke()
	return nil
}
