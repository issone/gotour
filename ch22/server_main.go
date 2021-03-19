package main

import (
	"go_t/ch22/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 基于 TCP 的 RPC
// func main() {
// 	rpc.RegisterName("MathService", new(server.MathService))
// 	l, e := net.Listen("tcp", ":1234")
// 	if e != nil {
// 		log.Fatal("listen error:", e)
// 	}
// 	// rpc.Accept(l)
// 	http.Serve(l, nil) //换成http的服务
// 	return
// }

// 基于 HTTP的RPC
// func main() {
// 	rpc.RegisterName("MathService", new(server.MathService))
// 	// rpc.HandleHTTP()
// 	l, e := net.Listen("tcp", ":1234")
// 	if e != nil {
// 		log.Fatal("listen error:", e)
// 	}
// 	// rpc.Accept(l)
// 	http.Serve(l, nil) //换成http的服务
// 	return
// }

//基于 TCP 的 JSON RPC
// func main() {
// 	rpc.RegisterName("MathService", new(server.MathService))
// 	l, e := net.Listen("tcp", ":1234")
// 	if e != nil {
// 		log.Fatal("listen error:", e)
// 	}
// 	for {
// 		conn, err := l.Accept()
// 		if err != nil {
// 			log.Println("jsonrpc.Serve: accept:", err.Error())
// 			return
// 		}
// 		//json rpc
// 		go jsonrpc.ServeConn(conn)
// 	}
// }

// 基于 HTTP的JSON RPC
func main() {

	rpc.RegisterName("MathService", new(server.MathService))

	//注册一个path，用于提供基于http的json rpc服务

	http.HandleFunc(rpc.DefaultRPCPath, func(rw http.ResponseWriter, r *http.Request) {

		// 通过Hijack方法劫持链接，然后转交给 jsonrpc 处理，这样就实现了基于 HTTP 协议的 JSON RPC 服务
		conn, _, err := rw.(http.Hijacker).Hijack()

		if err != nil {

			log.Print("rpc hijacking ", r.RemoteAddr, ": ", err.Error())

			return

		}

		var connected = "200 Connected to JSON RPC"

		io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")

		jsonrpc.ServeConn(conn)

	})

	l, e := net.Listen("tcp", ":1234")

	if e != nil {

		log.Fatal("listen error:", e)

	}

	http.Serve(l, nil) //换成http的服务

}
