package main

import (
	"Communication/MQ_pkg"
	"fmt"
)

const (
	MQ_NAME     = "mq_ifain"
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

	err = MQ_pkg.MQ_exchangedeclare_fn(l_mq_chan_ptr, MQ_NAME)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}

	l_queue_s, err := MQ_pkg.MQ_queue_declare_fn(l_mq_chan_ptr, MQ_DURABLE, MQ_NOWAIT)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}

	err = MQ_pkg.MQ_queue_bind_fn(l_mq_chan_ptr, &l_queue_s, MQ_NAME)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}

	l_deliveries_s, err := MQ_pkg.MQ_queue_consume_fn(l_mq_chan_ptr, &l_queue_s, MQ_AUTOACK, MQ_NOWAIT)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}

	fmt.Println("Waiting For Message :")

	for d := range l_deliveries_s {
		fmt.Println(string(d.Body))
	}
}
