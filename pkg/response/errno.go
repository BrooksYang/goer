package response

type Errno struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
