package main

import "net"

type User struct {
	Name    string
	Address string
	C       chan string
	conn    net.Conn
	server  *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	remoteAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    remoteAddr,
		Address: remoteAddr,
		C:       make(chan string),
		conn:    conn,
		server:  server,
	}
	//启动监听
	go user.ListenMessage()
	return user
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

func (this *User) Online() {
	this.server.onlineUserMapLock.Lock()
	this.server.OnlineUserMap[this.Name] = this
	this.server.onlineUserMapLock.Unlock()

	//广播
	this.server.BroadCast(this, "have been online")
}

func (this *User) Offline() {
	delete(this.server.OnlineUserMap, this.Name)

	this.server.BroadCast(this, "offline")
}

func (this *User) Domessage(msg string) {
	if msg == "who" {
		//查询当前在线用户
		this.server.onlineUserMapLock.Lock()
		for _, e := range this.server.OnlineUserMap {
			this.SendMessageToSelf(e.Name + "is online")
		}
		this.server.onlineUserMapLock.Unlock()

	} else {
		this.server.BroadCast(this, msg)
	}

}

// 给当前用户发送消息
func (this *User) SendMessageToSelf(msg string) {
	this.conn.Write([]byte(msg))
}
