package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	//客户端连接
	conn net.Conn
}

//创建一个用户的方法
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	// 启动监听当前user channel消息的go程
	go user.ListenMessage()
	return user
}

//监听当前user channel 的go方法
//有消息直接发送给客户段
func (this *User) ListenMessage() {
	for {

		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
