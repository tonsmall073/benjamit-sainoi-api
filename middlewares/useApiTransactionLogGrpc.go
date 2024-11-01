package middlewares

import (
	con "bjm/db/benjamit"
	model "bjm/db/benjamit/models"
	"bjm/utils"
	"bjm/utils/enums"
	"context"
	"encoding/json"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func logMiddlewareGrpc(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("[INFO] Request from %s", p.Addr)
	}
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("[ERROR] Response : %v", err)
	}

	contentType := ""
	origin := ""
	headers := make(map[string][]string)
	path := info.FullMethod
	method := info.FullMethod
	requestBody, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error marshalling request to JSON: %v", err)
		requestBody = []byte("Error converting to JSON")
	}

	responseBody, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error marshalling response to JSON: %v", err)
		responseBody = []byte("Error converting to JSON")
	}

	var statusCode codes.Code
	if err != nil {
		statusCode = status.Code(err)
	} else {
		statusCode = codes.OK
	}

	fullMethod := info.FullMethod
	parts := strings.Split(fullMethod, "/")
	if len(parts) == 3 {
		path = "/" + parts[1] // service path
		method = parts[2]     // method name
	} else if len(parts) == 4 {
		path = "/" + parts[1] + "/" + parts[2] // service path
		method = parts[3]                      // method name
	} else {
		path = fullMethod
		method = fullMethod
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		contentTypeList := md.Get("content-type")
		if len(contentTypeList) > 0 {
			contentType = contentTypeList[0]
		}
		originList := md.Get("origin")
		if len(originList) > 0 {
			origin = originList[0]
		}
	}

	go func() {
		for key, values := range md {
			headers[key] = values
		}
		headersJson, headersJsonErr := json.Marshal(headers)
		if headersJsonErr != nil {
			log.Printf("[ERROR] logging json marshal error: %v\n", headersJsonErr)
		}

		responseLog := model.ApiTransactionLog{
			Path:          path,
			Method:        method,
			ContentType:   contentType,
			StatusCode:    int(statusCode),
			ResponseBody:  string(responseBody),
			RequestBody:   string(requestBody),
			Headers:       string(headersJson),
			Origin:        origin,
			InterfaceType: enums.GRPC,
		}
		db, dbErr := con.Connect()
		defer con.ConnectClose(db)
		if dbErr != nil {
			log.Printf("[ERROR] failed to connect to database: %v\n", dbErr)
		} else {
			if err := db.Create(&responseLog).Error; err != nil {
				log.Printf("[ERROR] logging recording errors: %v\n", err)
			} else {
				log.Printf("[INFO] logging request for path: %s\n", path)
			}
		}
	}()

	return resp, err
}

func UseApiTransactionLogGrpc() grpc.ServerOption {
	cleanupOnce.Do(func() {
		go utils.LogCleanupTask(enums.GRPC)
	})

	return grpc.UnaryInterceptor(logMiddlewareGrpc)
}
