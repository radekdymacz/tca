package handler

import (
	"testing"

	client "github.com/micro/go-micro/client"
	api "github.com/micro/micro/api/proto"
	context "golang.org/x/net/context"
)

func TestJoke_Get(t *testing.T) {
	type fields struct {
		Client client.Client
	}
	type args struct {
		ctx context.Context
		req *api.Request
		rsp *api.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Joke{
				Client: tt.fields.Client,
			}
			if err := a.Get(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Joke.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
