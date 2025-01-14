package main

import "net"

type User struct {
	Name    string
	Address string
	C       chan string
	conn    net.Conn
}

func NewUser(conn net.Conn) *User {
	remoteAddr := conn.RemoteAddr().String()
	user := &User{
		Name:    remoteAddr,
		Address: remoteAddr,
		C:       make(chan string),
		conn:    conn,
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
