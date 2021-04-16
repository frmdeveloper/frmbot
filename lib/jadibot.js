let { WAConnection: _WAConnection, MessageType } = require('@adiwajshing/baileys')
let WAConnection = require('./lib/simple').WAConnection(_WAConnection)
let qrcode = require('qrcode')

if (global.conns instanceof Array) console.log()// for (let i of global.conns) global.conns[i] && global.conns[i].user ? global.conns[i].close().then(() => delete global.conns[id] && global.conns.splice(i, 1)).catch(global.client.logger.error) : delete global.conns[i] && global.conns.splice(i, 1)
else global.conns = []

let handler  = async (mek, { client, args, prefix, command }) => {
  let parent = args[0] && args[0] == 'plz' ? client : global.client
  let auth = false
  if ((args[0] && args[0] == 'plz') || global.client.user.jid == client.user.jid) {
    let id = global.conns.length
    let client = new WAConnection()
    if (args[0] && args[0].length > 200) {
      let json = Buffer.from(args[0], 'base64').toString('utf-8')
      // global.client.sendMessage(m.isGroup ? m.sender : m.chat, json, m)
      let obj = JSON.parse(json)
      await client.loadAuthInfo(obj)
      auth = true
    }
    client.on('qr', async qr => {
    	exec(`qrencode -o jadibot_${sender.split("@")[0]}.png ${qr}`)
      setTimeout(() => {
		kodeqrjadibot = fs.readFileSync(`./jadibot_${sender.split("@")[0]}.png`)
		teksjadibot = `Scan QR ini untuk jadi bot sementara\n\n1. Klik titik tiga di pojok kanan atas\n2. Ketuk WhatsApp Web\n3. Scan QR ini \nQR Expired dalam 20 detik`
    	client.sendMessage(from, kodeqrjadibot, image, {quoted: mek, caption: teksjadibot})
    }, 2000)
    })
    client.once('connection-validated', user => {
      client.sendMessage(from, 'Berhasil tersambung dengan WhatsApp - mu.\n*NOTE: Ini cuma numpang*\n' + JSON.stringify(user, null, 2), m)
    })
    client.on('message-new', global.client.handler)
    client.regenerateQRIntervalMs = null
    client.connect().then(async ({user}) => {
      if (auth) return
      await client.sendMessage(user.jid, `Kamu bisa login tanpa qr dengan pesan dibawah ini. untuk mendapatkan kode lengkapnya, silahkan kirim *${prefix}getcode* untuk mendapatkan kode yang akurat`, MessageType.text)
      client.sendMessage(user.jid, `${prefix}${command} ${Buffer.from(JSON.stringify(client.base64EncodedAuthInfo())).toString('base64')}`, MessageType.text)
    })
    client.on('close', client.logger.info)
    global.conns.push(client)
  } else throw 'Tidak bisa membuat bot didalam bot!\n\nhttps://wa.me/' + global.client.user.jid.split`@`[0] + '?text=.jadibot'
}

module.exports = jadibot

