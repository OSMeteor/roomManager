package roomManager

type nodeMessage struct {
	messageType int         `json:"type"` //消息类型，类型为NODE_MESSAGE_TYPE组
	rommId      string      `json:"room_id"`
	body        interface{} `json:"body"` //消息体
}

const (
	_                                     = iota
	NODE_MESSAGE_TYPE_IN_HALL             //添加节点入大厅
	NODE_MESSAGE_TYPE_CLOSE_ROOM          //关闭节点服务
	NODE_MESSAGE_TYPE_CHANGE_ROOM         //节点房间变更
	NODE_MESSAGE_TYPE_SEND_MESSAGE        //节点发送消息
	NODE_MESSAGE_TYPE_CLEAN_ROOM          //清理房间垃圾节点
	NODE_MESSAGE_TYPE_FILL_USERID         //完善用户ID
	NODE_MESSAGE_TYPE_RELOAD_BANNED_IP    //重载IP地址黑名单
	NODE_MESSAGE_TYPE_RELOAD_BANNED_USER  //重载用户ID黑名单
	NODE_MESSAGE_TYPE_RELOAD_BANNED_WORDS //重载敏感词列表
)
