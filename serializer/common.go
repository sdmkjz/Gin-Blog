package serializer

// 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// Token序列化结构
type TokenData struct {
	User  string `json:"user"`
	Token string `json:"token"`
}
