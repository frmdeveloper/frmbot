/*
HAI NAMAKU AQULZZ
YAH DISINI AKU SEBAGAI PEMULA MAU MENCOBA MEMBUAT BOT KU SENDIRI
YANG PASTINYA BANYAK COPY PASTE
OKE TERIMA KASIH
*/
const { WAConnection, MessageType } = require("@adiwajshing/baileys")
const fs = require('fs')
const { color } = require('../lib/color')

const client = new WAConnection()
exports.client = client

exports.connect = async() => {
    let authofile = './FRMbotLOGIN.json'
	const client = new WAConnection()
client.logger.level = 'warn'
console.log(banner.string)

	client.on('qr', () => {
	console.log(color('[','white'), color('!','red'), color(']','white'), color(' SCAN KODE QR DIATAS, PAKAI WHATSAPP'))
	})
	fs.existsSync('./FRMbotLOGIN.json') && client.loadAuthInfo('./FRMbotLOGIN.json')
	client.on('connecting', () => {
		console.log('Menghubungkan...')
	})
	client.on('open', () => {
		console.log('Terhubung')
		fs.writeFileSync('./FRMbotLOGIN.json', JSON.stringify(client.base64EncodedAuthInfo(), null, '\t'))
		client.sendMessage(client.user.jid, JSON.stringify(client.base64EncodedAuthInfo(), null, '\t'), MessageType.text)
		client.sendMessage(`${ownerNumber}`, `BOT BERHASIL DIAKTIFKAN`, MessageType.text)
	})
	client.connect({timeoutMs: 30*1000})
    return client
}