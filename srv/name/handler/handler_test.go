package handler

import (
	"testing"

	proto "github.com/radekdymacz/tca/srv/name/proto"
	context "golang.org/x/net/context"
)

func TestNamer_Random(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.Request
		rsp *proto.Response
	}
	tests := []struct {
		name    string
		n       *Namer
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Namer{}
			if err := n.Random(tt.args.ctx, tt.args.req, tt.args.rsp); (err != nil) != tt.wantErr {
				t.Errorf("Namer.Random() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
