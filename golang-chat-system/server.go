package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//user 的map表
	OnlineMap map[string]*User
	//map锁
	mapLock sync.RWMutex
	//message广播管道
	Message chan string
}

// 创建server方法
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听Message广播消息channel的go程,一旦有消息就发送给全部user
func (this *Server) ListenMessage() {
	for {
		//不断的从Message取数据
		msg := <-this.Message
		//将message发送给全部user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	//把消息放到管道中
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	fmt.Println("连接建立成功")
	//用户上线了进行广播
	//1.将用户加入onlinemap中
	user := NewUser(conn)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	// 成功后调用广播
	this.BroadCast(user, "上线了")

	// 当前handler阻塞
	select {}
}

//启动服务器方法
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听message的goroutine
	go this.ListenMessage()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			// 失败
			fmt.Println("listener accept err:", err)
			continue
		}
		//成功
		//do handler
		go this.Handler(conn)
	}

}
