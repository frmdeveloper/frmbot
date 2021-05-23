/*
HAI NAMAKU AQULZZ
YAH DISINI AKU SEBAGAI PEMULA MAU MENCOBA MEMBUAT BOT KU SENDIRI
YANG PASTINYA BANYAK COPY PASTE
OKE TERIMA KASIH
*/
const { WAConnection: _WAConnection, MessageType } = require("@adiwajshing/baileys")
const fs = require('fs')
const { color } = require('../lib/color')
const { uploadimg, wait, simih, getBuffer, h2k, generateMessageID, getGroupAdmins, getRandom, banner, start, info, success, close, emojiStrip, banner2, processTime, bitly, shortlink } = require('../lib/functions')
const simple = require('../lib/simple.js')
const WAConnection = simple.WAConnection(_WAConnection)

const client = new WAConnection()
exports.client = client

exports.connect = async() => {
    let authofile = './FRMbotLOGIN.json'
	client.logger.level = 'warn'
	client.on('qr', () => {
	console.log(color('[','white'), color('!','red'), color(']','white'), color(' SCAN KODE QR DIATAS, PAKAI WHATSAPP'))
	})
	fs.existsSync(authofile) && client.loadAuthInfo(authofile)
	client.on('connecting', () => {
		console.log('Menghubungkan...')
	})
	client.on('open', () => {
		console.log('Terhubung')
		fs.writeFileSync(authofile, JSON.stringify(client.base64EncodedAuthInfo(), null, '\t'))
		client.sendMessage(client.user.jid, JSON.stringify(client.base64EncodedAuthInfo(), null, '\t'), MessageType.text)
		client.sendMessage(`${ownerNumber}`, `BOT BERHASIL DIAKTIFKAN`, MessageType.text)
	})
	client.connect({timeoutMs: 30*1000})
    return client
}