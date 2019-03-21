package grpc

import (
	"context"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/honestbee/Zen/errs"
)

func logUnaryInterceptor(logger *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var remoteAddr string
		if p, ok := peer.FromContext(ctx); ok {
			remoteAddr = p.Addr.String()
		}

		logger.Info().Fields(map[string]interface{}{
			"from": remoteAddr,
			"path": info.FullMethod,
		}).Msgf("receiving data")

		resp, err := handler(ctx, req)

		if err != nil {
			er, ok := err.(*errs.Error)
			if !ok {
				er = errs.NewErr(errs.ServerInternalErrorCode, err)
			}
			if er.InternalErr != nil {
				err = status.Error(er.GRPCStatus, er.OutputErr)

				logger.Info().Fields(map[string]interface{}{
					"from":  remoteAddr,
					"path":  info.FullMethod,
					"error": er.Error(),
				}).Msgf("grpc unary error occurred")
			}
		}

		return resp, err
	}
}
