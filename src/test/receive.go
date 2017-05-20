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

func mq_receive_wait(MQ_NAME string) {
	/*Creating AMQP Message Queue Connection*/
	l_mq_chan_ptr, err := MQ_pkg.MQ_create_connection_fn(IP, PORT, CTAG, USERNAME, PASSWORD, VHOST)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}
	/*Declaring Message Queue Exchange*/
	err = MQ_pkg.MQ_exchangedeclare_fn(l_mq_chan_ptr, EXCHANGE_NAME, EXCHANGE_TYPE, MQ_DURABLE, MQ_NOWAIT)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}
	/*Declaring Queue*/
	l_queue_s, err := MQ_pkg.MQ_queue_declare_fn(l_mq_chan_ptr, Dest_Name, MQ_DURABLE, MQ_NOWAIT)
	if err != nil {
		fmt.Println("ERROR: MQ_pkg.MQ_create_connection_fn()" + "\n")
		return
	}
	/*Declaring Message Queue with Exchange*/
	err = MQ_pkg.MQ_queue_bind_fn(l_mq_chan_ptr, &l_queue_s, EXCHANGE_NAME, BINDING_KEY, MQ_NOWAIT)
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
func main() {
	mq_receive_wait(MQ_NAME)
}
