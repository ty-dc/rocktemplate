// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package grpcManager

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spidernet-io/rocktemplate/api/v1/grpcService"
)

func (s *grpcClientManager) SendRequestForExecRequest(ctx context.Context, serverAddress []string, request *grpcService.ExecRequestMsg) (*grpcService.ExecResponseMsg, error) {

	if e := s.clientDial(ctx, serverAddress); e != nil {
		return nil, errors.Errorf("failed to dial, error=%v", e)
	}
	defer s.client.Close()

	c := grpcService.NewCmdServiceClient(s.client)

	if r, err := c.ExecRemoteCmd(ctx, request); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}
