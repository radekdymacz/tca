// Code generated by protoc-gen-go.
// source: name.proto
// DO NOT EDIT!

package name

import (
	"reflect"
	"testing"

	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

func TestRequest_Reset(t *testing.T) {
	tests := []struct {
		name string
		m    *Request
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Request{}
			m.Reset()
		})
	}
}

func TestRequest_String(t *testing.T) {
	tests := []struct {
		name string
		m    *Request
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Request{}
			if got := m.String(); got != tt.want {
				t.Errorf("Request.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_ProtoMessage(t *testing.T) {
	tests := []struct {
		name string
		r    *Request
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{}
			r.ProtoMessage()
		})
	}
}

func TestRequest_Descriptor(t *testing.T) {
	tests := []struct {
		name  string
		r     *Request
		want  []byte
		want1 []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{}
			got, got1 := r.Descriptor()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.Descriptor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Request.Descriptor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestResponse_Reset(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			m.Reset()
		})
	}
}

func TestResponse_String(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Response.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_ProtoMessage(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			r.ProtoMessage()
		})
	}
}

func TestResponse_Descriptor(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
		want1  []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			got, got1 := r.Descriptor()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Descriptor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Response.Descriptor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestResponse_GetName(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			if got := m.GetName(); got != tt.want {
				t.Errorf("Response.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetSurname(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			if got := m.GetSurname(); got != tt.want {
				t.Errorf("Response.GetSurname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetGender(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			if got := m.GetGender(); got != tt.want {
				t.Errorf("Response.GetGender() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetRegion(t *testing.T) {
	type fields struct {
		Name    string
		Surname string
		Gender  string
		Region  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Response{
				Name:    tt.fields.Name,
				Surname: tt.fields.Surname,
				Gender:  tt.fields.Gender,
				Region:  tt.fields.Region,
			}
			if got := m.GetRegion(); got != tt.want {
				t.Errorf("Response.GetRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_publisher_Publish(t *testing.T) {
	type fields struct {
		c     client.Client
		topic string
	}
	type args struct {
		ctx  context.Context
		msg  interface{}
		opts []client.PublishOption
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
			p := &publisher{
				c:     tt.fields.c,
				topic: tt.fields.topic,
			}
			if err := p.Publish(tt.args.ctx, tt.args.msg, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("publisher.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPublisher(t *testing.T) {
	type args struct {
		topic string
		c     client.Client
	}
	tests := []struct {
		name string
		args args
		want Publisher
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPublisher(tt.args.topic, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPublisher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterSubscriber(t *testing.T) {
	type args struct {
		topic string
		s     server.Server
		h     interface{}
		opts  []server.SubscriberOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterSubscriber(tt.args.topic, tt.args.s, tt.args.h, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("RegisterSubscriber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewNameClient(t *testing.T) {
	type args struct {
		serviceName string
		c           client.Client
	}
	tests := []struct {
		name string
		args args
		want NameClient
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNameClient(tt.args.serviceName, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNameClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nameClient_Random(t *testing.T) {
	type fields struct {
		c           client.Client
		serviceName string
	}
	type args struct {
		ctx  context.Context
		in   *Request
		opts []client.CallOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &nameClient{
				c:           tt.fields.c,
				serviceName: tt.fields.serviceName,
			}
			got, err := c.Random(tt.args.ctx, tt.args.in, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("nameClient.Random() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nameClient.Random() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterNameHandler(t *testing.T) {
	type args struct {
		s    server.Server
		hdlr NameHandler
		opts []server.HandlerOption
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterNameHandler(tt.args.s, tt.args.hdlr, tt.args.opts...)
		})
	}
}

func TestName_Random(t *testing.T) {
	type fields struct {
		NameHandler NameHandler
	}
	type args struct {
		ctx context.Context
		in  *Request
		out *Response
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
			h := &Name{
				NameHandler: tt.fields.NameHandler,
			}
			if err := h.Random(tt.args.ctx, tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Name.Random() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
