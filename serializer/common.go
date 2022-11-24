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

// 列表序列化器

type ListResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Total  int64       `json:"total"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
