package handler

import (
	"encoding/json"
	"net/http"

	"io/ioutil"

	"log"

	proto "github.com/radekdymacz/tca/srv/name/proto"
	context "golang.org/x/net/context"
)

//Namer struct implements handler interface
type Namer struct{}

//Random name from http://uinames.com/api/
func (n *Namer) Random(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	// fetch random mane
	//ures, err := http.Get("http://uinames.com/api/")
	client := &http.Client{}
	ureq, err := http.NewRequest("GET", "http://uinames.com/api/", nil)
	if err != nil {
		return err
	}
	ureq.Header.Set("Content-Type", "application/json")
	ures, err := client.Do(ureq)

	if err != nil {
		return err
	}
	//defer freeing up the client
	defer ures.Body.Close()
	body, err := ioutil.ReadAll(ures.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, rsp)
	if err != nil {

		return err
	}
	log.Printf("Response: %s", rsp.String())
	return nil
}
