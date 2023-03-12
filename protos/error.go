package protos

var (
	InvalidRequestParameters = &RspInfo{Code: 20000, Msg: "Invalid Request Parameters"}
)

var (
	UnknownError = &RspInfo{Code: 90000, Msg: "Unknown Error"}
)
