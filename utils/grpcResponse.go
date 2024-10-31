package utils

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcResponseJson[T any](body *T, statusCode int) (*T, error) {
	var res *ErrorResponseModel
	if statusCode < 0 {
		res = &ErrorResponseModel{
			MessageDesc: GrpcStatusCodes[13],
			StatusCode:  13,
		}
		marshel, marshelErr := json.Marshal(res)
		if marshelErr != nil {
			return (*T)(body), status.Error(codes.Code(13), marshelErr.Error())
		}
		if err := json.Unmarshal(marshel, body); err != nil {
			return (*T)(body), status.Error(codes.Code(13), err.Error())
		}
		return (*T)(body), status.Error(codes.Code(13), GrpcStatusCodes[13])
	}
	return body, status.Error(codes.Code(statusCode), GrpcStatusCodes[statusCode])
}

func GrpcResponseErrorJson[T any](body *T, messageDesc string, statusCode int) (*T, error) {
	var res *ErrorResponseModel
	if statusCode < 0 {
		res = &ErrorResponseModel{
			MessageDesc: GrpcStatusCodes[13],
			StatusCode:  13,
		}
		marshel, marshelErr := json.Marshal(res)
		if marshelErr != nil {
			return (*T)(body), status.Error(codes.Code(13), marshelErr.Error())
		}
		if err := json.Unmarshal(marshel, body); err != nil {
			return (*T)(body), status.Error(codes.Code(13), err.Error())
		}
		return (*T)(body), status.Error(codes.Code(13), GrpcStatusCodes[13])
	}

	res = &ErrorResponseModel{
		MessageDesc: messageDesc,
		StatusCode:  statusCode,
	}
	marshel, marshelErr := json.Marshal(res)
	if marshelErr != nil {
		return (*T)(body), status.Error(codes.Code(13), marshelErr.Error())
	}
	if err := json.Unmarshal(marshel, body); err != nil {
		return (*T)(body), status.Error(codes.Code(13), err.Error())
	}
	return (*T)(body), status.Error(codes.Code(statusCode), messageDesc)
}
