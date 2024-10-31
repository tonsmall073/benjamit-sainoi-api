package utils

var GrpcStatusCodes = map[int]string{
	0:  "ok",
	1:  "canceled",
	2:  "unknown",
	3:  "invalid argument",
	4:  "deadline exceeded",
	5:  "not found",
	6:  "already exists",
	7:  "permission denied",
	8:  "resource exhausted",
	9:  "failed precondition",
	10: "aborted",
	11: "out of range",
	12: "unimplemented",
	13: "internal",
	14: "unavailable",
	15: "data loss",
	16: "unauthenticated",
	17: "unauthorized",
}
