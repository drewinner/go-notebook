package chap03

import "net/http"

/**
*	提供https服务
 */
func Demo01() {
	server := http.Server{
		Addr:    "",
		Handler: nil,
	}
	//cert.pem是SSL证书
	//key.pem 是服务器私钥
	_ = server.ListenAndServeTLS("cert.pem", "key.pem")
}
