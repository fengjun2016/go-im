package handler

import (
	"encoding/json"
	"fmt"
	"go-im/app/model"
	"go-im/app/service"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"gopkg.in/fatih/set.v0"
)

//本核心在于形成userid 和 Node 的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行
	DataQueue chan []byte //此管道 用来收发消息
	GroupSets set.Interface
}

//定义命令行格式
const (
	CmdSingleMsg = 10
	CmdRoomMsg   = 11
	CmdHeart     = 0
)

type Message struct {
	Id      string `json:"id,omitempty" form:"id"`           //消息ID
	Userid  string `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   string `json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"`   //其他和数字相关的
}

//userid 和 node 的 映射关系
var clientMap map[string]*Node = make(map[string]*Node, 0)

//读写锁
var rwlocker sync.RWMutex

//消息存储 持久化服务
var messageService service.MessageService

var msgData chan Message

func init() {
	//消息存储缓冲通道 初始化
	msgData = make(chan Message, 10000)
}

//后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		logrus.Println(err.Error())
		return
	}

	switch msg.Cmd {
	case CmdSingleMsg:
		sendMsg(msg.Dstid, data)
	case CmdRoomMsg:
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.Dstid) {
				v.DataQueue <- data
			}
		}
	case CmdHeart:
		//检查 客户端的心跳
	}
}

//添加新的群ID到用户的groupset中
func AddGroupId(userId, gid string) {
	//取得node
	rwlocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	rwlocker.Unlock()
}

//实现聊天的功能
func Chat(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	userId := query.Get("id")
	token := query.Get("token")

	//校验token 是否合法
	isleagal := checkToken(userId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(req *http.Request) bool {
			return isleagal
		},
	}).Upgrade(rw, req, nil)

	//检查是否协议升级成功
	if err != nil {
		logrus.Println("upgrade error ", err.Error())
		return
	}

	//获得wensocket 连接 conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	//获取用户全部群id
	comIds, _ := contactService.SearchComunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}

	//由于这个map是全局共享的 而且 go并发 map是不安全的 所以需要对其上锁
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()

	//开启协程处理发送逻辑
	go sendproc(node)

	//开启协程完成接收逻辑
	go recvproc(node)

	//心跳包检测
	//启动一个协程，每隔1s向客户端发送一次心跳消息
	go heartbeat(node)

	//处理消息 持久化 存储
	go loadToDb()

	sendMsg(userId, []byte("welcome!"))
}

//校验token 是否合法
func checkToken(userId, token string) bool {
	user, err := userService.Find(userId)

	return user.Token == token && err == nil
}

//接收逻辑
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			logrus.Println(err.Error())
			return
		}

		dispatch(data)
		//todo 对 data进一步处理
		fmt.Printf("recv<=%s", data)

		//离线消息 持久化存储 不保证
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			logrus.Println(err.Error())
			return
		}

		msgData <- msg
	}
}

//发送逻辑
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				logrus.Println(err.Error())
				return
			}
		}
	}
}

//向每个客户端发送心跳包
func heartbeat(node *Node) {
	for {
		node.DataQueue <- []byte("heartbeat")
		time.Sleep(1 * time.Second)
	}
}

//发送消息, 发送到消息的管道
func sendMsg(userId string, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

//消息的持久化存储
func loadToDb() {
	var timeLocal, _ = time.LoadLocation("Asia/Chongqing")
	for {
		select {
		case msg := <-msgData:
			msgModel := model.Message{}
			msgModel.Userid = msg.Userid
			msgModel.Cmd = msg.Cmd
			msgModel.Dstid = msg.Dstid
			msgModel.Media = msg.Media
			msgModel.Content = msg.Content
			msgModel.Pic = msg.Pic
			msgModel.Url = msg.Url
			msgModel.Memo = msg.Memo
			msgModel.Amount = msg.Amount
			msgModel.Createat = time.Now().In(timeLocal)

			messageService.LoadToDb(msgModel)
		}
	}
}
