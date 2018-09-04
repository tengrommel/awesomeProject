package message

// 确定一些消息类型

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"data"` // 消息的类型
}

// 定义两个消息...后面需要再增加
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户的id
	UserPwd  string `json:"userPwd"`  // 用户的密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态码 500 表示该用户未注册 200 表示登陆成功
	Error string `json:"error"` // 返回错误信息
}

type RegisterMessage struct {
	//...
}
