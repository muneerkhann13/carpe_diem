package main

import (
	"communication/MQ_pkg"
	"fmt"
)

const (
	MQ_NAME     = "sys"
	MQ_IP       = "localhost"
	MQ_PORT     = 5672
	MQ_DURABLE  = true
	MQ_AUTOACK  = true
	MQ_NOWAIT   = false
	MQ_CTAG     = "ifainSubSystemTag"
	MQ_USERNAME = "guest"
	MQ_PASSWORD = "guest"
)

func main() {
	l_mq_chan_ptr, err := MQ_pkg.MQ_create_connection_fn(MQ_NAME, MQ_IP, MQ_PORT, MQ_CTAG, MQ_USERNAME, MQ_PASSWORD)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}

	err = MQ_pkg.MQ_sendmessage_fn(l_mq_chan_ptr, MQ_NAME, []byte("HELLO Nirbhay !!!!"))
}
