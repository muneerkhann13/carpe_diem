package main

import (
	"communication/MQ_pkg"
	"fmt"
)

var (
	MQ_NAME       = "mq_payback"
	Dest_Name     = "mq_payback"
	IP            = "localhost"
	PORT          = 5672
	CTAG          = "mq_user"
	USERNAME      = "guest"
	PASSWORD      = "guest"
	MQ_DURABLE    = true
	MQ_NOWAIT     = false
	MQ_AUTOACK    = true
	VHOST         = ""
	EXCHANGE_NAME = "user"
	BINDING_KEY   = "user"
	EXPIRY_TIME   = "90000"
	EXCHANGE_TYPE = "direct"
)

func send(message []byte) {
	l_mq_chan_ptr, err := MQ_pkg.MQ_create_connection_fn(IP, PORT, CTAG, USERNAME, PASSWORD, VHOST)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}

	err = MQ_pkg.MQ_sendmessage_fn(l_mq_chan_ptr, EXCHANGE_NAME, BINDING_KEY, EXPIRY_TIME, Dest_Name, message)
}

func main() {

	send([]byte("Hello"))
}
