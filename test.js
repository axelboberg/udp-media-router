const dgram = require('dgram')
const client = dgram.createSocket('udp4')
const server = dgram.createSocket('udp4')

let lastBuf

function wait (ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

class Client {
  constructor (host, port) {
    this.host = host
    this.port = port
    this.ready = true
  }

  sendIfReady (msg) {
    if (!this.ready) return
    this.ready = false
    client.send(msg, 0, msg.length, this.port, this.host, () => {
      this.ready = true
    })
  }
}

const CLIENTS = [
  new Client('127.0.0.1', 3002),
  new Client('127.0.0.1', 3003),
  new Client('127.0.0.1', 3004),
  new Client('127.0.0.1', 3005)
]

server.on('message', msg => {
  for (let client of CLIENTS) {
    client.sendIfReady(msg)
  }
})

server.bind(3001, '127.0.0.1')
