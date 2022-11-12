package grpcManager

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spidernet-io/rocktemplate/api/v1/grpcService"
)

func (s *grpcClientManager) SendRequestForExecRequest(ctx context.Context, request *grpcService.ExecRequestMsg) (*grpcService.ExecResponseMsg, error) {
	if s.client == nil {
		return nil, errors.Errorf("please dial first")
	}
	c := grpcService.NewCmdServiceClient(s.client)

	if r, err := c.ExecRemoteCmd(ctx, request); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}
