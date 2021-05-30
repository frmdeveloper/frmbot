__path = process.cwd()
const {
   WAConnection: _WAConnection,
   MessageType,
   Presence,
   MessageOptions,
   Mimetype,
   MimetypeMap,
   WALocationMessage,
   WA_MESSAGE_STUB_TYPES,
   ReconnectMode,
   ProxyAgent,
   GroupSettingChange,
   ChatModification,
   waChatKey,
   mentionedJid,
   WA_DEFAULT_EPHEMERAL
} = require("@adiwajshing/baileys")
var express = require('express');
var router = express.Router();
const simple = require('./lib/simple.js')
const WAConnection = simple.WAConnection(_WAConnection)
const moment = require("moment-timezone") 
const fs = require("fs") 
const crypto = require('crypto')
const axios = require('axios')
const imageToBase64 = require('image-to-base64')
const WSF = require('wa-sticker-formatter')
const { bgcolor } = require('./lib/color')
const { fetchJson, uploadImages } = require('./lib/fetcher')
const { recognize } = require('./lib/ocr')
const tesseract = require("node-tesseract-ocr")
const { virtex } = require('./src/virtex')
const { virtex2 } = require('./src/virtex2')
const { cara } = require('./src/cara')
const { spawn, exec } = require("child_process")
const { uploadimg, wait, simih, getBuffer, h2k, generateMessageID, getGroupAdmins, getRandom, banner, start, info, success, close, emojiStrip, banner2, processTime, bitly, shortlink } = require('./lib/functions')
const { uptotele, uptonaufal } = require('./lib/uploadimage')
const { servers, yta, ytv } = require('./lib/y2mate')
const tiktod = require('tiktok-scraper')
const brainly = require('brainly-scraper')
const translate = require('translation-google')
const ffmpeg = require('fluent-ffmpeg')
const cd = 4.32e+7
const { removeBackgroundFromImageFile } = require('remove.bg')
const { ind } = require('./language')
const yts = require('yt-search')
const os = require('os')
const cfonts = require('cfonts')
const cheerio = require('cheerio')
const request = require('request')
const kagApi = require('@kagchi/kag-api')
const lolis = require('lolis.life')
const loli = new lolis()
const google = require('google-it')
const fetch = require('node-fetch')
const { EmojiAPI } = require("emoji-api");
const emoji = new EmojiAPI()
const imgbb = require('imgbb-uploader')
const qrlogo = require('branded-qr-code')
penting = JSON.parse(fs.readFileSync('./assets/penting.json'))
const frm = require('./whatsapp/message.js')
const conn = require('./whatsapp/connect')

nomowner = '62895803265350' //pakai kode negara, contoh: 62895803265350
ownerNumber = [`${nomowner}@s.whatsapp.net`]
botName = 'FRM BOT'
devName = 'FRM Developer'
ownerName = 'Fauzan Rifki Maulana'
LolKey = 'juanlol291002' //lolhuman.herokuapp.com atau juanlol291002 atau erdwpehub28
ZeksKey = 'caliph_71' //zeks.xyz
BarBarKey = 'IDxO1TFYnKADlX4pxcHa' //mhankbarbars.tech
VhtearKey = 'ZidanGanzz' //api.vhtear.com
TobzKey = 'Z4sxB1r91MFrgnK3sObn' //tobz.herokuapp.com
XteamKey = '9ccd5c3c92359b79' //api.xteam.xyz
shizukakey = 'istmeiky633' 
imgbbkey = "f4fde56c72298d6d92ce5133024cbba8"
keyrmbg = '6yWvBTgxbkW7LL8fA8ahiQXE'

linkapp = '/refresh'
axios.get(linkapp)

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

router.get('/:encoded_id', function(req, res){
	shortnya = req.params.encoded_id
	res.send(`Param ${shortnya} tidak ditemukan`)
	})
router.get('/heleh', (req, res) => {
	res.json({result:'heleh terdeteksi'})
	})
router.get('/refresh', (req, res) => {
	res.json({result:`Ada yang membuka ${linkapp}`})
	setTimeout( () => {
	axios.get(linkapp)
	axios.get('/heleh')
	}, 10000)
})
router.get('/japriwa', (req, res) => {
	q = req.query.q
	nomor = req.query.nomor
	if (!nomor) return res.json({result:`silahkan tambahkan parameter nomor`})
	if (nomor.length == 0) return res.json({result:`UNTUK SIAPA ?`})
	if (nomor.startsWith("0")) return res.json({result:`pakai kode negara. contoh: 628xxx`})
	if (!nomor.match(/^[0-9]+$/)) return res.json({result:`Nomor tujuan harus angka`})
	if (!q) return res.json({result:`silahkan tambahkan parameter q`})
	if (q.length == 0) return res.json({result:`pesan kosong`})
	res.json({result:`Mengirim pesan ke ${nomor.split('@')[0]}`})
	client.sendMessage(`${nomor}@s.whatsapp.net`, `*[ FRM BOT ]*\n\n${q}`, MessageType.text)
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
			const sendfile = (filenya, formatnya) => {
				res.type(formatnya)
				res.sendFile(filenya)
			}
	
switch(command) {
		case 'wa':
    		redirect('http://wa.me/62895803265350')
    		break
    	case 'getip':
    		reply(req.ip)
    		break
    	case 'eval':
    		if (!q) return res.send(`parameter q kosong`)
    		try {
    		res.send(require('util').format(await eval(`;(async () => { ${q} })()`)))
    		} catch (e) {
    		res.send(e)
    		}
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
                    await sendfile(__path + `/sampah/${command}.jpg`, 'jpg')
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