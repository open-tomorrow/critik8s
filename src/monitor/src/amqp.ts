import { AMQPWebSocketClient } from '@cloudamqp/amqp-client'
import { fromEvent } from 'rxjs';
import { Subject } from 'rxjs';

export class AMQPService {
  amqpCli: AMQPWebSocketClient
  topic: string
  routingKey: string
  messagesSubject = new Subject<any>()

  constructor(host: string,
              port: string,
              usr: string,
              psw: string,
              topic: string,
              routingKey: string) {
    const url = "ws://" + host + ":" + port
    this.topic = topic
    this.routingKey = routingKey
    this.amqpCli = new AMQPWebSocketClient(url, "/", usr, psw)
  }

  async subscribe(){
    try {
      const conn = await this.amqpCli.connect()
      const ch = await conn.channel()

      await ch.exchangeDeclare(
        this.topic, // exchange name
        "topic", // type
      )

      const criticQ = await ch.queueDeclare();
      await ch.queueBind(criticQ.name, this.topic, this.routingKey);

      ch.basicConsume(criticQ.name, {noAck: true}, (v: any) => {
        const msg = new TextDecoder().decode(v.body);
        console.log(v);
        this.messagesSubject.next(msg)
      })

    } catch (err) {
      console.error("AMQPService::subscribe", err);
    }
  }
}
