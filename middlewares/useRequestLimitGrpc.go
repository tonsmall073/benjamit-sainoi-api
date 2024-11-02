package middlewares

import (
	"bjm/utils"
	"context"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	requestCount = make(map[string]int)
	timerMap     = make(map[string]*time.Timer)
)

func limitRequestGrpc(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	defaultMax := 20
	defaultExpiration := time.Duration(30)

	if convInt, convErr := strconv.Atoi(os.Getenv("REQUEST_SEND_LIMIT_MAX")); convErr == nil && convInt > 0 {
		defaultMax = convInt
	}
	if convInt, convErr := strconv.Atoi(os.Getenv("REQUEST_SEND_DELAY_LIMIT_SECONDS")); convErr == nil && convInt > 0 {
		defaultExpiration = time.Duration(convInt)
	}

	ip, ok := getClientIP(ctx)
	if !ok {
		ip = "localhost"
		// return nil, status.Error(codes.Internal, "could not get client IP")
	}

	_, isExist := requestCount[ip]
	if !isExist {
		requestCount[ip] = 0
	}

	if requestCount[ip] >= defaultMax {
		errorResponse := utils.ErrorResponseModel{MessageDesc: "request limit exceeded", StatusCode: 8}
		return errorResponse, status.Error(codes.ResourceExhausted, "request limit exceeded")
	} else {
		requestCount[ip]++
	}

	resetTimer(ip, defaultExpiration)

	return handler(ctx, req)
}

func resetTimer(ip string, duration time.Duration) {
	if timer, exists := timerMap[ip]; exists {
		timer.Stop()
	}

	timerMap[ip] = time.AfterFunc(duration*time.Second, func() {
		delete(requestCount, ip)
		delete(timerMap, ip)
	})
}

func getClientIP(ctx context.Context) (string, bool) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if addr := md["remote_addr"]; len(addr) > 0 {
			return addr[0], true
		}
	}
	return "", false
}

func UseLimitRequestGrpc() grpc.UnaryServerInterceptor {
	return limitRequestGrpc
}
