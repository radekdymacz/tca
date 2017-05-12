package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"io/ioutil"

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
	//construct url

	url := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=%s", url.QueryEscape(req.GetFirstName()), url.QueryEscape(req.GetLastName()), url.QueryEscape(req.GetLimitTo()))

	// calling the service
	//jokeRes, err := http.Get(url)
	client := &http.Client{}
	ureq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	ureq.Header.Set("Accept", "text/html")
	jokeRes, err := client.Do(ureq)
	if err != nil {
		return err
	}
	//free up client
	defer jokeRes.Body.Close()
	body, err := ioutil.ReadAll(jokeRes.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, rsp)
	if err != nil {
		return errors.New("Marshalling err")
	}

	log.Printf("Response: %s", rsp.String())

	return nil
}
