package errno

import "goer/pkg/response"

/*
|--------------------------------------------------------------------------
| 错误码设计
|--------------------------------------------------------------------------
|
| 5位数字
| 服务级别：1 位数表示： 1业务错误，4请求错误，5系统错误
| 服务模块：2 位数表示
| 具体错误：2 位数表示
|
*/

var (
	OK = &response.Errno{Code: 0, Msg: "OK"}

	// Request
	IllegalRequest   = &response.Errno{Code: 40000, Msg: "Illegal Request"}
	IllegalParameter = &response.Errno{Code: 40400, Msg: "Illegal Parameter"}
	InvalidToken     = &response.Errno{Code: 40401, Msg: "Invalid Token"}
	NotFound         = &response.Errno{Code: 40404, Msg: "Not found"}
	ValidationError  = &response.Errno{Code: 40422, Msg: "Validation Error"}
	TooManyRequest   = &response.Errno{Code: 40429, Msg: "Too many request"}

	// System
	InternalServerError = &response.Errno{Code: 50001, Msg: "Server Error"}
	NetworkCongested    = &response.Errno{Code: 50002, Msg: "Network Congested"}

	// User
	AccountExists         = &response.Errno{Code: 10101, Msg: "Account exists"}
	AccountNotFound       = &response.Errno{Code: 10102, Msg: "Account not found"}
	InvalidAccount        = &response.Errno{Code: 10103, Msg: "Invalid account"}
	AccountLocked         = &response.Errno{Code: 10104, Msg: "Your account has been locked"}
	InvalidReferrer       = &response.Errno{Code: 10105, Msg: "Invalid Referrer"}
	InvalidPassword       = &response.Errno{Code: 10106, Msg: "Invalid password"}
	InvalidGoogleCode     = &response.Errno{Code: 10107, Msg: "Invalid google code"}
	NeedPayPassword       = &response.Errno{Code: 10108, Msg: "Need pay password"}
	InvalidPayPassword    = &response.Errno{Code: 10109, Msg: "Invalid pay password"}
	AlreadyHasPayPassword = &response.Errno{Code: 10110, Msg: "Already has pay password"}
)
