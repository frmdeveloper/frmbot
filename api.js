__path = process.cwd()
const {
   WAConnection,
   MessageType,
   Presence,
   MessageOptions,
   Mimetype,
   WALocationMessage,
   WA_MESSAGE_STUB_TYPES,
   ReconnectMode,
   ProxyAgent,
   GroupSettingChange,
   ChatModification,
   waChatKey,
   mentionedJid,
   processTime,
   WA_DEFAULT_EPHEMERAL
} = require("@adiwajshing/baileys")
const frm = require('./whatsapp/message.js')
const conn = require('./whatsapp/connect')
const client = conn.client
var express = require('express');
var router = express.Router();
const moment = require("moment-timezone") 
const fs = require("fs") 
const fetch = require('node-fetch')
const axios = require('axios')
const chalk = require('chalk')
var brainly = require('brainly-scraper');
const { spawn, exec } = require("child_process")
const { wait, simih, getBuffer, h2k, generateMessageID, getGroupAdmins, getRandom, banner, start, info, success, close } = require('./lib/functions')
const { fetchJson, uploadImages } = require('./lib/fetcher')
const { bgcolor, color } = require('./lib/color')

linkapp = 'http://frmdev.repl.co/refresh'
axios.get(linkapp)
	.then((a) => {
		console.log(a.data.result)
	})

function kyun(seconds){
  function pad(s){
    return (s < 10 ? '0' : '') + s;
  }
  var hours = Math.floor(seconds / (60*60));
  var minutes = Math.floor(seconds % (60*60) / 60);
  var seconds = Math.floor(seconds % 60);

  //return pad(hours) + ':' + pad(minutes) + ':' + pad(seconds)
  return `${pad(hours)} Jam ${pad(minutes)} Menit ${pad(seconds)} Detik`
}

router.get('/heleh', (req, res) => {
	res.json({result:'heleh terdeteksi'})
	})
router.get('/refresh', (req, res) => {
	res.json({result:`Ada yang membuka ${linkapp}`})
	setTimeout( () => {
	axios.get(linkapp)
	.then((a) => {
		console.log(a.data.result)
	})
	}, 10000)
})
router.get('/japriwa', (req, res) => {
	q = req.query.q
	untuk = req.query.untuk
	if (!untuk) return res.json({result:`silahkan tambahkan parameter untuk`})
	if (untuk.length == 0) return res.json({result:`UNTUK SIAPA ?`})
	if (!untuk.startsWith('0')) return res.json({result:`Gunakan kode negara tanpa diawali +`})
	if (!q) return res.json({result:`silahkan tambahkan parameter q`})
	if (q.length == 0) return res.json({result:`pesan kosong`})
	res.json({result:`mengirim ke ${untuk.split('@')[0]}\n*isi pesan:* ${q}`})
	client.sendMessage(`${untuk}@s.whatsapp.net`, `*[ FRM BOT ]\n*${q}`, MessageType.text)
	.catch(e => {
		res.json({result:'ERROR'})
		})
	})
router.get('/login', (req, res) => {
    res.sendFile(__path + '/login.html')
	})
router.get('/loginn', async(req, res) => {
    res.send('yee')
	})
router.get('/', (req, res) => {
	res.send('kosong')
	})
	
router.get('/c', async (req, res, next) => {
			try {
			q = req.query.q
			apikeyku = req.query.apikey
			command = req.query.cmd
			const redirect = (url) => {
				res.redirect(url)
			}
			const download = (link) => {
				res.download(link)
			} 
			const reply = (teks) => {
				res.json({result:teks})
				.catch(e => {
res.json({result:'ERROR'})
})
			}
			const send = (link) => {
				res.send(link)
			}
			const sendfile = (filenya) => {
				res.sendFile(filenya)
			}
	
switch(command) {
		case 'wa':
    		alihkan('http://wa.me/62895803265350')
    		break
    	case 'login':
    		
    		break
    	case 'shadow':
        		case 'cup':
                case 'cup1':
                case 'romance':
                case 'smoke':
                case 'burnpaper':
                case 'lovemessage':
                case 'undergrass':
                case 'love':
                case 'coffe':
                case 'woodheart':
                case 'flowerheart':
                case 'woodenboard':
                case 'summer3d':
                case 'wolfmetal':
                case 'nature3d':
                case 'underwater':
                case 'golderrose':
                case 'summernature':
                case 'letterleaves':
                case 'glowingneon':
                case 'fallleaves':
                case 'flamming':
                case 'harrypotter':
                case 'carvedwood':
                    lolimg = await getBuffer(`http://api.lolhuman.xyz/api/photooxy1/${command}?apikey=muzharzain&text=${q}`)
                    await fs.writeFileSync(`./sampah/${command}.jpg`, lolimg)
                    await sendfile(__path + `/sampah/${command}.jpg`)
                    break
				case 'ytv':
					ytv(args[0])
					.then((res) => {
					const { dl_link, thumb, title, filesizeF, filesize } = res
					redirect(dl_link)
					})
					break
    	default:
    	reply(`request *${command}* \n tidak ditemukan`)
	}
	} catch (e) {
		res.json({result:`terjadi kesalahan \n${e}`})
	}
})

module.exports = router