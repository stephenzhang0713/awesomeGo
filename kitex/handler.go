package main

import (
	"context"
	"kitex/kitex_gen/api"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return &api.Response{Message: req.Message}, nil
}

// VisitOneway implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) VisitOneway(ctx context.Context, req *api.Request) (err error) {
	// TODO: Your code here...
	return nil
}
