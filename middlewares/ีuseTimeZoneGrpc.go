package middlewares

import (
	"context"
	"os"
	"time"

	"google.golang.org/grpc"
)

type contextKey string

const timeZoneKey contextKey = "timeZone"

func setTimeZoneGrpc(timeZone *time.Location) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		ctx = context.WithValue(ctx, timeZoneKey, timeZone)
		return handler(ctx, req)
	}
}

func UseTimeZoneGrpc() grpc.UnaryServerInterceptor {
	defaultStr := "Asia/Bangkok"

	if getTimeZone := os.Getenv("SERVER_TIME_ZONE"); getTimeZone != "" {
		defaultStr = getTimeZone
	}
	loc, err := time.LoadLocation(defaultStr)
	if err != nil {
		panic(err)
	}

	return setTimeZoneGrpc(loc)
}
