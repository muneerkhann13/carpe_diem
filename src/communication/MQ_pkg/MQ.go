package MQ_pkg

import (
	"fmt"
	"strings"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	channel_send *amqp.Channel
	tag          string
	done         chan error
}

func MQ_create_connection_fn(p_mq_ip_str string, p_mq_port_str int, p_mq_ctag_str, p_mq_username_str,
	p_mq_password_str, p_vhost string) (mq_chan_ptr *Consumer, err error) {
	mq_chan_ptr = &Consumer{
		conn:         nil,
		channel:      nil,
		channel_send: nil,
		tag:          p_mq_ctag_str,
		done:         make(chan error),
	}
	amqpURI := fmt.Sprint("amqp://", p_mq_username_str, ":", p_mq_password_str, "@", p_mq_ip_str, ":", p_mq_port_str, "/", p_vhost) //amqp://guest:guest@localhost:5672/
	fmt.Println("NxTxN : Channel SEND *********** ", amqpURI)
	mq_chan_ptr.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return
	}
	mq_chan_ptr.channel, err = mq_chan_ptr.conn.Channel()
	if err != nil {
		return
	}

	mq_chan_ptr.channel_send, err = mq_chan_ptr.conn.Channel()
	if err != nil {
		return
	}
	return
}

func MQ_sendmessage_fn(p_mq_chan_ptr *Consumer, p_exc_name_str, p_bindingkey_str, p_expiry_time, p_src_instance_str string, message []byte) (err error) {
	if err = p_mq_chan_ptr.channel_send.Publish(
		p_exc_name_str,   // publish to an exchange
		p_bindingkey_str, // routing to 0 or more queues
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            message,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,
			Expiration:      p_expiry_time, // 0-9
			AppId:           p_src_instance_str,
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)

	}
	return
}

func MQ_exchangedeclare_fn(p_mq_chan_ptr *Consumer, p_exc_name_str, p_exc_type_str string, p_mq_durable_str, p_mq_nowait_str bool) (err error) {
	err = p_mq_chan_ptr.channel.ExchangeDeclare(
		p_exc_name_str,   // name of the exchange
		p_exc_type_str,   // type
		p_mq_durable_str, // durable
		false,            // delete when complete
		false,            // internal
		p_mq_nowait_str,  // noWait
		nil,              // arguments
	)
	if err != nil {
		return
	}
	return
}

func MQ_queue_declare_fn(p_mq_chan_ptr *Consumer, p_queue_name_str string, p_mq_durable_str, p_mq_nowait_str bool) (queue_s amqp.Queue, err error) {
	queue_s, err = p_mq_chan_ptr.channel.QueueDeclare(
		p_queue_name_str, // name of the queue
		p_mq_durable_str, // durable
		false,            // delete when usused
		false,            // exclusive
		p_mq_nowait_str,  // noWait
		nil,              // arguments
	)
	if err != nil {
		return
	}
	return
}

func MQ_queue_bind_fn(p_mq_chan_ptr *Consumer, p_queue_s *amqp.Queue, p_exc_name_str, p_bindingkey_str string, p_mq_nowait bool) (err error) {
	err = p_mq_chan_ptr.channel.QueueBind(
		p_queue_s.Name,   // name of the queue
		p_bindingkey_str, // bindingKey
		p_exc_name_str,   // sourceExchange
		p_mq_nowait,      // noWait
		nil,              // arguments
	)
	if err != nil {
		return
	}
	return
}

func MQ_queue_consume_fn(p_mq_chan_ptr *Consumer, p_queue_s *amqp.Queue, p_mq_autoack_str, p_mq_nowait_str bool) (
	deliveries_s <-chan amqp.Delivery, err error) {
	deliveries_s, err = p_mq_chan_ptr.channel.Consume(
		p_queue_s.Name, //queue
		p_mq_chan_ptr.tag,
		p_mq_autoack_str, //Auto Ack
		false,            //exclusive
		false,            //noLocal
		p_mq_nowait_str,  //noWait
		nil)
	if err != nil {
		return
	}
	return
}

func MQ_shutdown_fn(p_mq_chan_ptr *Consumer) (err error) {
	if err := p_mq_chan_ptr.channel.Cancel(p_mq_chan_ptr.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}
	if err := p_mq_chan_ptr.channel_send.Cancel(p_mq_chan_ptr.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := p_mq_chan_ptr.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}
	return
}
func CheckNotifyClose(p_mq_chan_ptr *Consumer) {
	for {
		ERR1 := fmt.Sprintln("closing:", <-p_mq_chan_ptr.conn.NotifyClose(make(chan *amqp.Error)))
		if strings.Contains(ERR1, "closing: E") {
			fmt.Println(ERR1)
			return

		}
	}
}
func MQ_queue_check_fn(p_mq_chan_ptr *Consumer, q_name string) (conumser_count int, err error) {
	var state amqp.Queue
	state, err = p_mq_chan_ptr.channel.QueueInspect(q_name)
	conumser_count = state.Consumers
	return
}
