package models

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/http"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderId uint
	TargetId uint
	Type     string
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

func chat(writer http.ResponseWriter, request *http.Request) {
	//validate token
	query := request.URL.Query()
	//token := query.Get("token")
	id := query.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	msgType := query.Get("type")
	targetId := query.Get("targetId")
	context := query.Get("context")
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

	//receive message
}
