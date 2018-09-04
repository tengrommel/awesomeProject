package process2

import (
	"awesomeProject/net/chatroom/common/message"
	"awesomeProject/net/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	// 字段？
	Conn net.Conn
}

// 编写一个函数serverProcessLogin函数，专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代码...
	// 1、先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	// 先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 再声明一个LoginResMes
	var loginResMes message.LoginResMes
	// 如果用户的id=100，密码=123456，认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500 // 500 状态码 表示该用户不存在
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}
	// 3 将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	// 4.将data赋值给resMes
	resMes.Data = string(data)
	// 5.对resMes 进行序列化，准备序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	// 6.发送data，我们将其封装到writePkg
	// 因为使用了分层的模式（mvc），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
