package process

import (
	"awesomeProject/net/chatroom/client/utils"
	"awesomeProject/net/chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

// 发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 1 创建一个message.Message
	var mes message.Message
	mes.Type = message.SmsMesType
	// 2、创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content // 内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	// 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sendGrop , err=", err.Error())
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes ", err.Error())
		return
	}
	//将mes发送给服务
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendGroup err=", err)
		return
	}
	return
}
