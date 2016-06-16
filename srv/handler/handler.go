package handler

import (
	"github.com/Rakanixu/flag/srv/elastic"
	proto "github.com/Rakanixu/flag/srv/proto/flag"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

// Flag struct
type Flag struct{}

// Create srv handler
func (f *Flag) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {

	if len(req.Key) <= 0 || len(req.Description) <= 0 {
		return errors.BadRequest("go.micro.srv.flag.Flag.Create", "Fields required")
	}

	if err := elastic.Create(req); err != nil {
		return errors.InternalServerError("go.micro.srv.flag.Flag.Create", err.Error())
	}

	return nil
}

// Read srv handler
func (f *Flag) Read(ctx context.Context, req *proto.ReadRequest, rsp *proto.ReadResponse) error {
	if len(req.Key) <= 0 {
		return errors.BadRequest("go.micro.srv.flag.Flag.Read", "Flag key required")
	}

	flag, err := elastic.Read(req.Key)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.flag.Flag.Read", err.Error())
	}

	rsp.Key = flag.Key
	rsp.Description = flag.Description
	rsp.Value = flag.Value

	return nil
}

// Flip srv handler
func (f *Flag) Flip(ctx context.Context, req *proto.FlipRequest, rsp *proto.FlipResponse) error {
	if len(req.Key) <= 0 {
		return errors.BadRequest("go.micro.srv.flag.Flag.Flip", "Flag key required")
	}

	if err := elastic.Flip(req.Key); err != nil {
		return errors.InternalServerError("go.micro.srv.Flag.Flip", err.Error())
	}

	return nil
}

// Delete srv handler
func (f *Flag) Delete(ctx context.Context, req *proto.DeleteRequest, rsp *proto.DeleteResponse) error {
	if len(req.Key) <= 0 {
		return errors.BadRequest("go.micro.srv.flag.Flag.Delete", "Flag key required")
	}

	if err := elastic.Delete(req.Key); err != nil {
		return errors.InternalServerError("go.micro.srv.Flag.Delete", err.Error())
	}

	return nil
}

// List srv handler
func (f *Flag) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	flags, err := elastic.List()
	if err != nil {
		return errors.InternalServerError("go.micro.srv.Flag.List", err.Error())
	}

	rsp.Result = flags.Result

	return nil
}
