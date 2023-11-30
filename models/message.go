package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderId int64
	TargetId int64
	Type     int
	Media    int
	Content  string
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// clientMap is a map of client id to Node
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// rwlock for clientMap
var rwLock sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	//validate token
	query := request.URL.Query()
	//token := query.Get("token")
	id := query.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	// msgType := query.Get("type")
	tid := query.Get("targetId")
	targetId, _ := strconv.ParseInt(tid, 10, 64)
	// context := query.Get("context")
	isValid := true
	conn, err := (&websocket.Upgrader{
		//check token
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println("upgrade failed, err:", err)
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//user relation
	// bind userid with node, and add lock
	rwLock.Lock()
	clientMap[userId] = node
	rwLock.Unlock()
	//send message
	go sendMsgToClient(node)
	//receive message
	go receiveMsgFromClient(node)
	sendMsg(targetId, []byte("hello, welcome to chat room"))
}

func sendMsgToClient(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("sendMsgToClient: ", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("send message failed, err:", err)
				return
			}
		}
	}
}

func receiveMsgFromClient(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("receive message failed, err:", err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws]<<<<<<<: ", string(data))
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpsend()
	go udprecv()
}

// comlete udp data send protocol
func udpsend() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	for {
		select {
		case data := <-udpsendChan:
			fmt.Println("sendMsgToClient: ", string(data))
			_, err := con.Write(data)
			if err != nil {
				fmt.Println("send message failed, err:", err)
				return
			}
		}
	}

}

// comlete udp data receive protocol
func udprecv() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		var buf [1024]byte
		n, err := con.Read(buf[:])
		if err != nil {
			fmt.Println("read message failed, err:", err)
			return
		}
		dispatch(buf[:n])
	}
}

// backend scheduling logic processing
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}
	switch msg.Type {
	case 1: // private chat
		sendMsg(msg.TargetId, data)
		// case 2: // group chat
		// 	sendGroupMsg()
		// case 3: // broadcast
		// 	sendAllMsg()
		// case 4:
	}
}

func sendMsg(userID int64, msg []byte) {
	rwLock.RLock() // read lock
	node, ok := clientMap[userID]
	rwLock.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
