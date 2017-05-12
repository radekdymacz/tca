package handler

import (
	"testing"

	proto "github.com/radekdymacz/tca/srv/joke/proto"
	context "golang.org/x/net/context"
)

func TestJoker_ChuckNorris(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.Request
		rsp *proto.Response
	}
	tests := []struct {
		name    string
		j       *Joker
		args    args
		wantErr bool
	}{
		{
			"empty-request",
			nil,
			*&args{
				context.Background(),
				&proto.Request{},
				&proto.Response{},
			},
			false,
		},
		{
			"name-only-request",
			nil,
			*&args{
				context.Background(),
				&proto.Request{
					FirstName: "Radek",
				},
				&proto.Response{},
			},
			false,
		},
		{
			"lastname-only-request",
			nil,
			*&args{
				context.Background(),
				&proto.Request{
					LastName: "Dymacz",
				},
				&proto.Response{},
			},
			false,
		},
		{
			"limitto-only-request",
			nil,
			*&args{
				context.Background(),
				&proto.Request{
					LimitTo: "nerdy",
				},
				&proto.Response{},
			},
			false,
		},
		{
			"full-request",
			nil,
			*&args{
				context.Background(),
				&proto.Request{
					FirstName: "Radek",
					LastName:  "Dymacz",
					LimitTo:   "nerdy",
				},
				&proto.Response{},
			},
			false,
		},
		{
			"unicode-request",
			nil,
			*&args{
				context.Background(),
				&proto.Request{
					FirstName: "世界",
					LastName:  "Dymacz",
					LimitTo:   "nerdy",
				},
				&proto.Response{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Joker{}
			if err := j.ChuckNorris(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {

				t.Errorf("Joker.ChuckNorris() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
