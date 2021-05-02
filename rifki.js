/*
]=====> FAUZAN RIFKI MAULANA <=====[ ]=====> FRM DEVELOPER <=====[ ]=====> https://github.com/frmdeveloper/frmbot <=====[
*/
var express = require('express'),
    cors = require('cors'),
    secure = require('ssl-express-www');
const PORT = process.env.PORT || 8080 || 5000 || 3000
var { color } = require('./lib/color')
var apirouter = require('./api.js')
var app = express()
app.enable('trust proxy');
app.set("json spaces",2)
app.use(cors())
app.use(secure)
app.use(express.static("public"))
app.use('/', apirouter)
app.listen(PORT, () => {
    console.log(color("Server running on port " + PORT,'green'))
})

const dropboxV2Api = require('dropbox-v2-api')
const dropbox = dropboxV2Api.authenticate({
    token: '9ewnN6HaE5EAAAAAAAAAARRF-AjmOCUg7bC10gxrFJDoGlgTz1R8zspH0-yOoh73'
});

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
const qrcode = require("qrcode-terminal") 
const moment = require("moment-timezone") 
const fs = require("fs") 
const crypto = require('crypto')
const axios = require('axios')
const WSF = require('wa-sticker-formatter')
const { bgcolor } = require('./lib/color')
const { fetchJson, uploadImages } = require('./lib/fetcher')
const { recognize } = require('./lib/ocr')
const tesseract = require("node-tesseract-ocr")
const { virtex } = require('./src/virtex')
const { virtex2 } = require('./src/virtex2')
const { cara } = require('./src/cara')
const { spawn, exec } = require("child_process")
const { wait, simih, getBuffer, h2k, generateMessageID, getGroupAdmins, getRandom, banner, start, info, success, close } = require('./lib/functions')
const tiktod = require('tiktok-scraper')
const brainly = require('brainly-scraper')
const ffmpeg = require('fluent-ffmpeg')
const cd = 4.32e+7
const { removeBackgroundFromImageFile } = require('remove.bg')
const { ind } = require('./language')
const { yta, ytv } = require('./lib/ytdl')
const os = require('os')
const cheerio = require('cheerio')
const request = require('request')
const kagApi = require('@kagchi/kag-api')
const lolis = require('lolis.life')
const loli = new lolis()
const google = require('google-it')
const canvas = require('canvacord')
const fetch = require('node-fetch')
const { EmojiAPI } = require("emoji-api");
const emoji = new EmojiAPI()
const imgbb = require('imgbb-uploader')
const translate = require('@vitalets/google-translate-api')
tanda = '*───❉ FRM BOT ❉──*'
head1 = '*◪ ❀'
head2 = '❀*'
gaya1 = '║'
gaya2 = '╠☞'
gaya3 = '╰═─⊱'
katasandi = 'FRMbot'
nomerewa = ["0@s.whatsapp.net"]
masaaktif = '1'
jam1hari = '24'
monosp = '```'
prefix = '.'
gantiprefix = ''
blocked = []  
batre = [] 
limitawal = '70'
memberlimit = '3'
nggoroboguru = ''
nggopln = ''
fromnggoroboguru = '626262@s.whatsapp.net'
fromnggopln = '626262@s.whatsapp.net'
nomereroboguru = '6281578150000@s.whatsapp.net'
nomerepln = '628122123123@s.whatsapp.net'
tagstatus = 'status@broadcast'
cr = '_bot wa_'
pesansibuk = 'SEDANG SIBUK'
nomersibuk = ''
nomerwesdaftar = '626262@s.whatsapp.net'
statusbot = true
statuson = false
dibanned = '3'
sisabaterai = `belum diketahui`
hematdaya = `belum diketahui`
dicas = `belum diketahui`
sampah = ''
/*
]=====> INFO-INFO <=====[
*/
nomorkartu = ''
pulsakartu = ''
kuotakartu = ''
nomowner = '62895803265350' //pakai kode negara, contoh: 62895803265350
ownerNumber = [`${nomowner}@s.whatsapp.net`]
botName = 'FRM BOT'
devName = 'FRM Developer'
ownerName = 'Fauzan Rifki Maulana'
LolKey = 'muzharzain' //lolhuman.herokuapp.com atau juanlol291002 atau erdwpehub28
ZeksKey = 'apivinz' //zeks.xyz
BarBarKey = 'IDxO1TFYnKADlX4pxcHa' //mhankbarbars.tech
VhtearKey = '291002juan' //api.vhtear.com
TobzKey = 'Z4sxB1r91MFrgnK3sObn' //tobz.herokuapp.com
XteamKey = '9ccd5c3c92359b79' //api.xteam.xyz
shizukakey = 'istmeiky633' 
imgbbkey = "f4fde56c72298d6d92ce5133024cbba8"
/*
]=====> Hmmm <=====[
*/

/*       
]=====> FILE LUAR <=====[
*/
const liststiker = JSON.parse(fs.readFileSync('./sticker/liststiker.json'))
const listaudio = JSON.parse(fs.readFileSync('./audio/listaudio.json'))
const _jadibot = JSON.parse(fs.readFileSync('./database/user/datajadibot.json'))
const promo = JSON.parse(fs.readFileSync('./database/bot/promo.json'))
const _afk = JSON.parse(fs.readFileSync('./database/user/afk.json'))
const omongelek = JSON.parse(fs.readFileSync('./database/bot/omongelek.json'))
const _leveling = JSON.parse(fs.readFileSync('./database/group/leveling.json'))
const _level = JSON.parse(fs.readFileSync('./database/user/level.json'))
const _registered = JSON.parse(fs.readFileSync('./database/bot/registered.json'))
const welkom = JSON.parse(fs.readFileSync('./database/bot/welkom.json'))
const nsfw = JSON.parse(fs.readFileSync('./database/bot/nsfw.json'))
const samih = JSON.parse(fs.readFileSync('./database/bot/simi.json'))
const event = JSON.parse(fs.readFileSync('./database/bot/event.json'))
const _limit = JSON.parse(fs.readFileSync('./database/user/limit.json'))
const uang = JSON.parse(fs.readFileSync('./database/user/uang.json'))
ban = JSON.parse(fs.readFileSync('./database/user/banned.json'))

/*       
]=====> FILE MENU DILUAR <=====[
*/
const { help } = require('./lib/help')
const { simple } = require('./database/menu/simple')
const { gabut } = require('./database/menu/gabut')
const { groupm } = require('./database/menu/group')
const { download } = require('./database/menu/download')
const { dompet } = require('./database/menu/dompet')

const { random } = require('./database/menu/random')
const { other } = require('./database/menu/other')
const { owb } = require('./database/menu/owb')
const { maker } = require('./database/menu/maker')
const { sound } = require('./database/menu/sound')


/******** FIURE BOT IKI ********/
const makermenu2 = `
${head1} PEMBUATAN ${head2}
${gaya2} ${prefix}shadow
${gaya2} ${prefix}cup
${gaya2} ${prefix}cup1
${gaya2} ${prefix}romance
${gaya2} ${prefix}smoke
${gaya2} ${prefix}burnpaper
${gaya2} ${prefix}lovemessage
${gaya2} ${prefix}undergrass
${gaya2} ${prefix}love
${gaya2} ${prefix}coffe
${gaya2} ${prefix}woodheart
${gaya2} ${prefix}flowerheart
${gaya2} ${prefix}woodenboard
${gaya2} ${prefix}summer3d
${gaya2} ${prefix}wolfmetal
${gaya2} ${prefix}nature3d
${gaya2} ${prefix}underwater
${gaya2} ${prefix}golderrose
${gaya2} ${prefix}summernature
${gaya2} ${prefix}letterleaves
${gaya2} ${prefix}glowingneon
${gaya2} ${prefix}fallleaves
${gaya2} ${prefix}flamming
${gaya2} ${prefix}harrypotter
${gaya2} ${prefix}carvedwood
${gaya2} ${prefix}wetglass
${gaya2} ${prefix}multicolor3d
${gaya2} ${prefix}watercolor
${gaya2} ${prefix}luxurygold
${gaya2} ${prefix}galaxywallpaper
${gaya2} ${prefix}lighttext
${gaya2} ${prefix}beautifulflower
${gaya2} ${prefix}puppycute
${gaya2} ${prefix}royaltext
${gaya2} ${prefix}heartshaped
${gaya2} ${prefix}birthdaycake
${gaya2} ${prefix}galaxystyle
${gaya2} ${prefix}hologram3d
${gaya2} ${prefix}glossychrome
${gaya2} ${prefix}greenbush
${gaya2} ${prefix}metallogo
${gaya2} ${prefix}noeltext
${gaya2} ${prefix}glittergold
${gaya2} ${prefix}textcake
${gaya2} ${prefix}starsnight
${gaya2} ${prefix}wooden3d
${gaya2} ${prefix}textbyname
${gaya2} ${prefix}writegalacy
${gaya2} ${prefix}galaxybat
${gaya2} ${prefix}snow3d
${gaya2} ${prefix}birthdayday
${gaya2} ${prefix}goldplaybutton
${gaya2} ${prefix}silverplaybutton
${gaya2} ${prefix}freefire
${gaya2} ${prefix}darkneon ~teks~
${gaya2} ${prefix}candlemug ~teks~
${gaya2} ${prefix}lovemsg ~teks~
${gaya2} ${prefix}mugflower ~teks~
${gaya2} ${prefix}narutobanner ~teks~
${gaya2} ${prefix}paperonglass ~teks~
${gaya2} ${prefix}romancetext ~teks~
${gaya2} ${prefix}shadowtext ~teks~
${gaya2} ${prefix}coffecup ~teks~
${gaya2} ${prefix}coffecup2 ~teks~
${gaya2} ${prefix}glowingneon ~teks~
${gaya2} ${prefix}underwater ~teks~
${gaya2} ${prefix}hpotter ~teks~
${gaya2} ${prefix}woodblock ~teks~
${gaya1}
${gaya3}
*NB:* setelah teks diatas, silahkan masukkan teks anda
*contoh*
${prefix}cup ${botName}`

const tersimpan = `
${head1} SAVE ${head2}
${gaya2} ${prefix}getaudio ~nama nya~
${gaya2} ${prefix}getstiker ~nama nya~
${gaya2} ${prefix}saveaudio (tag audio)
${gaya2} ${prefix}savestiker (tag stiker)
${gaya1}
${gaya3}`

const donasi = `
*「 KIRIM PULSA 」*
*•* indosat
081615901727
*•* three
0895803265350

*══════════*
*══════════*

*「 KIRIM UANG 」*
*•* Aplikasi Dana
https://link.dana.id/qr/3jstu95e`

const edukasimenu = `
${head1} BELAJAR ${head2}
${gaya2} ${prefix}acakquran
${gaya2} ${prefix}apakah ~pertanyaan~
${gaya2} ${prefix}asupan
${gaya2} ${prefix}berita
${gaya2} ${prefix}bikinquote ~teks~ & ~namamu~
${gaya2} ${prefix}bisakah ~pertanyaan~
${gaya2} ${prefix}brainly ~soal~
${gaya2} ${prefix}bucin
${gaya2} ${prefix}cantikcek (geser gambar)
${gaya2} ${prefix}caklontong
${gaya2} ${prefix}chord ~judul~
${gaya2} ${prefix}dadu
${gaya2} ${prefix}dadu2
${gaya2} ${prefix}dadu3
${gaya2} ${prefix}dare
${gaya2} ${prefix}dare2
${gaya2} ${prefix}faktaunik
${gaya2} ${prefix}family100
${gaya2} ${prefix}gay
${gaya2} ${prefix}gantengcek (geser gambar)
${gaya2} ${prefix}google ~teks~
${gaya2} ${prefix}hobby
${gaya2} ${prefix}jamindo
${gaya2} ${prefix}jadwalsholat ~kode daerah~
${gaya2} ${prefix}jadwaltv ~channel~
${gaya2} ${prefix}kapankah ~pertanyaan~
${gaya2} ${prefix}katailham
${gaya2} ${prefix}lirik ~judul~
${gaya2} ${prefix}nulis ~teks~
${gaya2} ${prefix}nulis2 ~teks~
${gaya2} ${prefix}ramalhp ~628xx~
${gaya2} ${prefix}resep ~namanya~
${gaya2} ${prefix}sisahari
${gaya2} ${prefix}pantun
${gaya2} ${prefix}quotes
${gaya2} ${prefix}quotes2
${gaya2} ${prefix}quotes3
${gaya2} ${prefix}rate
${gaya2} ${prefix}tebakgambar
${gaya2} ${prefix}timer ~total~ ~satuan~
${gaya2} ${prefix}totalhuruf (tag pesan)
${gaya2} ${prefix}toxic
${gaya2} ${prefix}truth
${gaya2} ${prefix}tulis ~teks|Nama|Kelas~
${gaya2} ${prefix}watak
${gaya2} ${prefix}wiki ~teks~
${gaya2} ${prefix}wikien ~teks~
${gaya1}
${gaya3}
NB: Yang dicoret harus diganti
NB: Tanda kurung itu info
    tidak perlu diketik lagi`

const makermenu = `
${head1} MAKER MENU ${head2}
${gaya2} ${prefix}attp ~teks~
${gaya2} ${prefix}apiteks ~teks~
${gaya2} ${prefix}banner
${gaya2} ${prefix}bass
${gaya2} ${prefix}bitly ~link panjang~
${gaya2} ${prefix}blood ~teks~
${gaya2} ${prefix}cloudtext ~teks~
${gaya2} ${prefix}cglitch ~teks~ & ~teks~
${gaya2} ${prefix}cml ~teks~ & ~teks~
${gaya2} ${prefix}cphlogo ~teks~ & ~teks~
${gaya2} ${prefix}cpubg ~teks~ & ~teks~
${gaya2} ${prefix}dropwater ~teks~
${gaya2} ${prefix}emoji ~teks~
${gaya2} ${prefix}fast (geser vn)
${gaya2} ${prefix}ffbaner ~teks~
${gaya2} ${prefix}firework ~teks~
${gaya2} ${prefix}gemboktext ~teks~
${gaya2} ${prefix}gemuk (geser vn)
${gaya2} ${prefix}glitchtext ~teks~
${gaya2} ${prefix}greenneon ~teks~
${gaya2} ${prefix}halloweentext ~teks~
${gaya2} ${prefix}halah (geser pesan teks)
${gaya2} ${prefix}hilih (geser pesan teks)
${gaya2} ${prefix}huluh (geser pesan teks)
${gaya2} ${prefix}heleh (geser pesan teks)
${gaya2} ${prefix}holoh (geser pesan teks)
${gaya2} ${prefix}katakan ~teks~
${gaya2} ${prefix}lava ~teks~
${gaya2} ${prefix}lovemake ~teks~
${gaya2} ${prefix}metaldark ~teks~
${gaya2} ${prefix}neon ~teks~
${gaya2} ${prefix}neontext ~teks~
${gaya2} ${prefix}ninjalogo ~teks~
${gaya2} ${prefix}pornhub ~teks~ & ~teks~
${gaya2} ${prefix}metalteks ~teks~
${gaya2} ${prefix}qrcode ~teks~
${gaya2} ${prefix}raindrop (tag foto)
${gaya2} ${prefix}sandwrite ~teks~
${gaya2} ${prefix}slow (geser vn)
${gaya2} ${prefix}sticker (geser foto)
${gaya2} ${prefix}silktext ~teks~
${gaya2} ${prefix}sumery ~teks~
${gaya2} ${prefix}summer ~teks~
${gaya2} ${prefix}tupai (geser vn)
${gaya2} ${prefix}tahta ~teks~
${gaya2} ${prefix}textlight ~teks~
${gaya2} ${prefix}toimg (tag stiker)
${gaya2} ${prefix}tomp3 (tag video)
${gaya2} ${prefix}tourl (tag foto)
${gaya2} ${prefix}triggered (tag foto)
${gaya2} ${prefix}ttp ~teks~
${gaya2} ${prefix}tts ~kodebhs~  ~teks~
${gaya1}
${gaya3}
NB: Yang dicoret harus diganti
NB: Tanda kurung itu info
    tidak perlu diketik lagi
NB: *geser = tag*`

const downloader = `
${head1} DOWNLOAD ${head2}
${gaya2} ${prefix}fb ~link video fb~
${gaya2} ${prefix}pinterest ~teks~
${gaya2} ${prefix}tiktokstalk ~username~
${gaya2} ${prefix}happymod ~nama app~
${gaya2} ${prefix}igstalk ~username~
${gaya2} ${prefix}ig ~link ig~
${gaya2} ${prefix}moddroid ~nama app~
${gaya2} ${prefix}ytmp4 ~link yt~
${gaya2} ${prefix}ytmp3 ~link yt~
${gaya2} ${prefix}tiktok ~link tiktok~
${gaya2} ${prefix}joox ~judul~
${gaya2} ${prefix}play ~judul~
${gaya2} ${prefix}playvideo ~judul~
${gaya2} ${prefix}pptiktok ~username~
${gaya2} ${prefix}snack ~link snackvideo~
${gaya2} ${prefix}soundcloud ~link~
${gaya2} ${prefix}stalkig ~username~
${gaya2} ${prefix}wp ~nama pemandangan~
${gaya1}
${gaya3}
NB: Yang dicoret harus diganti
NB: Tanda kurung itu info
    tidak perlu diketik lagi`

const cekmenu = `
╭═─⊱ ❰ *CHECK* ❱ ⊰─═
${gaya1} *❀ Check User ${head2}
${gaya2} *123#
${gaya2} ${prefix}profile
${gaya2}
${gaya1} *❀ Bot ${head2}
${gaya2} p
${gaya2} tes
${gaya2} ${prefix}banlist
${gaya2} ${prefix}chatmu
${gaya2} ${prefix}frmgrup
${gaya2} ${prefix}hapus (geser pesanku)
${gaya2} ${prefix}info
${gaya2} ${prefix}lb
${gaya2} ${prefix}leaderboard
${gaya2} ${prefix}listuser
${gaya2} ${prefix}makasih
${gaya2} ${prefix}mutual
${gaya2} ${prefix}next
${gaya2} ${prefix}teswaktu
${gaya2} ${prefix}ping
${gaya2} ${prefix}quoted ~code~
${gaya2} ${prefix}save ~namamu~
${gaya2} ${prefix}sisahari
${gaya2} ${prefix}thanks
${gaya1}
${gaya3}
NB: Yang dicoret harus diganti
NB: Tanda kurung itu info
    tidak perlu diketik lagi`

const wibumenu = `
${head1} ACAK GAMBAR ${head2}
${gaya2} ${prefix}1cak
${gaya2} ${prefix}akira
${gaya2} ${prefix}anjing
${gaya2} ${prefix}blowjob
${gaya2} ${prefix}boruto
${gaya2} ${prefix}cium
${gaya2} ${prefix}hentai
${gaya2} ${prefix}husbu
${gaya2} ${prefix}kiss
${gaya2} ${prefix}kpop
${gaya2} ${prefix}kurumi
${gaya2} ${prefix}loli
${gaya2} ${prefix}loli2
${gaya2} ${prefix}meme
${gaya2} ${prefix}memeindo
${gaya2} ${prefix}miku
${gaya2} ${prefix}minato
${gaya2} ${prefix}nangis
${gaya2} ${prefix}neko
${gaya2} ${prefix}nekonime
${gaya2} ${prefix}pokemon
${gaya2} ${prefix}naruto
${gaya2} ${prefix}hinata
${gaya2} ${prefix}itori
${gaya2} ${prefix}peluk
${gaya2} ${prefix}ranime
${gaya2} ${prefix}randomanime
${gaya2} ${prefix}randomcry
${gaya2} ${prefix}randomhentai
${gaya2} ${prefix}randomhentong
${gaya2} ${prefix}rize
${gaya2} ${prefix}sasuke
${gaya2} ${prefix}sakura
${gaya2} ${prefix}slap
${gaya2} ${prefix}tampar
${gaya2} ${prefix}trap
${gaya2} ${prefix}waifu
${gaya2} ${prefix}wait
${gaya2} ${prefix}wibu
${gaya1}
${gaya3}`

const grupmenu = `
${head1} GROUP MENU ${head2}
${gaya2} ${prefix}add ~62xxx~
${gaya2} ${prefix}demote ~@tag~
${gaya2} ${prefix}edotense ~@tag~
${gaya2} ${prefix}event on
${gaya2} ${prefix}event off
${gaya2} ${prefix}fitnah ~@tag & pesannya & pesanbot~
${gaya2} ${prefix}grup buka
${gaya2} ${prefix}grup tutup
${gaya2} ${prefix}hedsot ~@tag~
${gaya2} ${prefix}hidetag ~pesanmu~
${gaya2} ${prefix}hidetag5 ~pesanmu~
${gaya2} ${prefix}kick ~@tag~
${gaya2} ${prefix}kickfast ~@tag~
${gaya2} ${prefix}kickme
${gaya2} ${prefix}leaderboard
${gaya2} ${prefix}leveling on
${gaya2} ${prefix}leveling off
${gaya2} ${prefix}linkgrup
${gaya2} ${prefix}listadmin
${gaya2} ${prefix}listonline
${gaya2} ${prefix}notifgrup on
${gaya2} ${prefix}notifgrup off
${gaya2} ${prefix}nsfw on
${gaya2} ${prefix}nsfw off
${gaya2} ${prefix}pengumuman ~teks~
${gaya2} ${prefix}peringatan ~teks~
${gaya2} ${prefix}promote
${gaya2} ${prefix}setname ~nama grup~
${gaya2} ${prefix}setdesc ~desk grup~
${gaya2} ${prefix}tagall
${gaya2} ${prefix}tagme
${gaya2} ${prefix}ubah.ikon
${gaya1}
${gaya3}`

const ownermenu = `
${head1} OWNER MENU ${head2}
${gaya2} ${prefix}anggotagrup
${gaya2} ${prefix}ban
${gaya2} ${prefix}bc
${gaya2} ${prefix}bcgc
${gaya2} ${prefix}block
${gaya2} ${prefix}bot off
${gaya2} ${prefix}bot on
${gaya2} ${prefix}bunuhbot
${gaya2} ${prefix}clearall
${gaya2} ${prefix}clearbc
${gaya2} ${prefix}clone
${gaya2} ${prefix}edit gaya1 ~${gaya1}~
${gaya2} ${prefix}edit gaya2 ~${gaya2}~
${gaya2} ${prefix}edit gaya3 ~${gaya3}~
${gaya2} ${prefix}edit head1 ~${head1}~
${gaya2} ${prefix}edit head2 ~${head2}~
${gaya2} ${prefix}edit pp
${gaya2} ${prefix}edit prefix
${gaya2} ${prefix}edit reply
${gaya2} ${prefix}eval ~cmdnya~
${gaya2} ${prefix}hapuschat ~codechat~
${gaya2} ${prefix}kickall
${gaya2} ${prefix}leave
${gaya2} ${prefix}off ~62xx@g.us~
${gaya2} ${prefix}on ~62xx@g.us~
${gaya2} ${prefix}owneronly off
${gaya2} ${prefix}owneronly on
${gaya2} ${prefix}reboot
${gaya2} ${prefix}restart
${gaya2} ${prefix}run ~code~
${gaya2} ${prefix}shutdown
${gaya2} ${prefix}sibuk on
${gaya2} ${prefix}sibuk off
${gaya2} ${prefix}ubahpp
${gaya2} ${prefix}unban
${gaya2} ${prefix}unblock
${gaya2} >
${gaya2} $
${gaya1}
${gaya3}`
/****** FIURE BOT IKI ******/

/*
]=====> FUNGSI <=====[
*/
const sleep = async (ms) => {
	return new Promise(resolve => setTimeout(resolve, ms))
}

const addAfkUser = (userId, time, reason, _dir) => {
    const obj = { id: userId, time: time, reason: reason }
    _dir.push(obj)
    fs.writeFileSync('./database/user/afk.json', JSON.stringify(_dir))
}

const checkAfkUser = (userId, _dir) => {
    let status = false
    Object.keys(_dir).forEach((i) => {
        if (_dir[i].id === userId) {
            status = true
        }
    })
    return status
}

const getAfkReason = (userId, _dir) => {
    let position = null
    Object.keys(_dir).forEach((i) => {
        if (_dir[i].id === userId) {
            position = i
        }
    })
    if (position !== null) {
        return _dir[position].reason
    }
}

const getAfkTime = (userId, _dir) => {
    let position = null
    Object.keys(_dir).forEach((i) => {
        if (_dir[i].id === userId) {
            position = i
        }
    })
    if (position !== null) {
        return _dir[position].time
    }
}

const getAfkId = (userId, _dir) => {
    let position = null
    Object.keys(_dir).forEach((i) => {
        if (_dir[i].id === userId) {
            position = i
        }
    })
    if (position !== null) {
        return _dir[position].id
    }
}

const getAfkPosition = (userId, _dir) => {
    let position = null
    Object.keys(_dir).forEach((i) => {
        if (_dir[i].id === userId) {
            position = i
        }
    })
    return position
}

const getLevelingXp = (sender) => {
            let position = false
            Object.keys(_level).forEach((i) => {
                if (_level[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                return _level[position].xp
            }
        }

        const getLevelingLevel = (sender) => {
            let position = false
            Object.keys(_level).forEach((i) => {
                if (_level[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                return _level[position].level
            }
        }

        const getLevelingId = (sender) => {
            let position = false
            Object.keys(_level).forEach((i) => {
                if (_level[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                return _level[position].id
            }
        }

        const addLevelingXp = (sender, amount) => {
            let position = false
            Object.keys(_level).forEach((i) => {
                if (_level[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                _level[position].xp += amount
                fs.writeFileSync('./database/user/level.json', JSON.stringify(_level))
            }
        }

        const addLevelingLevel = (sender, amount) => {
            let position = false
            Object.keys(_level).forEach((i) => {
                if (_level[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                _level[position].level += amount
                fs.writeFileSync('./database/user/level.json', JSON.stringify(_level))
            }
        }

        const addLevelingId = (sender) => {
            const obj = {id: sender, xp: 1, level: 1}
            _level.push(obj)
            fs.writeFileSync('./database/user/level.json', JSON.stringify(_level))
        }
             
         const getRegisteredRandomId = () => {
            return _registered[Math.floor(Math.random() * _registered.length)].id
        }

        const addRegisteredUser = (userid, sender, age, time, serials) => {
            const obj = { id: userid, name: sender, age: age, time: time, serial: serials }
            _registered.push(obj)
            fs.writeFileSync('./database/bot/registered.json', JSON.stringify(_registered))
        }

        const createSerial = (size) => {
            return crypto.randomBytes(size).toString('hex').slice(0, size)
        }

        const checkRegisteredUser = (sender) => {
            let status = false
            Object.keys(_registered).forEach((i) => {
                if (_registered[i].id === sender) {
                    status = true
                }
            })
            return status
        }
        
        const cekWesDaftar = (nomerwesdaftar) => {
            let status = false
            Object.keys(_registered).forEach((i) => {
                if (_registered[i].id === nomerwesdaftar) {
                    status = true
                }
            })
            return status
        }
        
        const addATM = (sender) => {
        	const obj = {id: sender, uang : 0}
            uang.push(obj)
            fs.writeFileSync('./database/user/uang.json', JSON.stringify(uang))
        }
        
        const addKoinUser = (sender, amount) => {
            let position = false
            Object.keys(uang).forEach((i) => {
                if (uang[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                uang[position].uang += amount
                fs.writeFileSync('./database/user/uang.json', JSON.stringify(uang))
            }
        }
        
        const checkATMuser = (sender) => {
        	let position = false
            Object.keys(uang).forEach((i) => {
                if (uang[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                return uang[position].uang
            }
        }
        
        const bayarLimit = (sender, amount) => {
        	let position = false
            Object.keys(_limit).forEach((i) => {
                if (_limit[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                _limit[position].limit -= amount
                fs.writeFileSync('./database/user/limit.json', JSON.stringify(_limit))
            }
        }
        
        const umureuser = (sender) => {
        	let position = false
            Object.keys(_registered).forEach((i) => {
                if (_registered[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                return _registered[position].age
            }
        }
        
        const namaneuser = (sender) => {
        	let position = false
            Object.keys(_registered).forEach((i) => {
                if (_registered[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                return _registered[position].name
            }
        }
        
        const stikget = (anunestik) => {
        	let position = false
            Object.keys(liststiker).forEach((i) => {
                if (liststiker[i].nomorstik === anunestik) {
                    position = i
                }
            })
            if (position !== false) {
                return liststiker[position].namastik
            }
        }
        	
        const confirmATM = (sender, amount) => {
        	let position = false
            Object.keys(uang).forEach((i) => {
                if (uang[i].id === sender) {
                    position = i
                }
            })
            if (position !== false) {
                uang[position].uang -= amount
                fs.writeFileSync('./database/user/uang.json', JSON.stringify(uang))
            }
        }
        
         const limitAdd = (sender) => {
             let position = false
            Object.keys(_limit).forEach((i) => {
                if (_limit[i].id == sender) {
                    position = i
                }
            })
            if (position !== false) {
                _limit[position].limit += 1
                fs.writeFileSync('./database/user/limit.json', JSON.stringify(_limit))
            }
        }
             
        
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
/*
]=====> SCAN QR <=====[
*/

const client = new WAConnection()
client.logger.level = 'warn'
console.log(banner.string)
   client.on('qr', () => {
	console.log(color('[','white'), color('!','red'), color(']','white'), color(' SCAN KODE QR DIATAS, PAKAI WHATSAPP'))
})

	client.on('credentials-updated', () => {
		fs.writeFileSync('./FRMbotLOGIN.json', JSON.stringify(client.base64EncodedAuthInfo(), null, '\t'))
		console.log('ingfokan cuyy...')
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

client.on('group-participants-update', async (anu) => {
			const mdata = await client.groupMetadata(anu.jid)
			if (anu.action == 'add' && anu.participants[0].includes(client.user.jid)) return client.sendMessage(mdata.id, `Halo semua, saya member baru.\n\nJangan memasukan saya ke grup *${mdata.subject}*\nJika belum izin kepada wa.me/${nomowner}\n\nbot akan keluar`, MessageType.text, {contextInfo: {"mentionedJid": [anu.participants[0]]}, quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(mdata.id ? { remoteJid: mdata.id } : {}) }, message: { conversation: `PERINGATAN` }}})
		if (!welkom.includes(anu.jid)) return
		try {
			console.log(anu)
			if (anu.action == 'add') {
				num = anu.participants[0]
				try {
					ppimg = await client.getProfilePicture(`${anu.participants[0].split('@')[0]}@c.us`)
				} catch {
					ppimg = 'https://i0.wp.com/www.gambarunik.id/wp-content/uploads/2019/06/Top-Gambar-Foto-Profil-Kosong-Lucu-Tergokil-.jpg'
				}
				teks = `*[ SELAMAT DATANG ]*\nKamu masuk di Grup *${mdata.subject}* \n___________________________\n@${num.split('@')[0]} \nMohon untuk follow instagramku *>* \ninstagram.com/frm_developer`
				let buff = await getBuffer(ppimg)
				client.sendMessage(mdata.id, teks, MessageType.text, {contextInfo: {"mentionedJid": [num]}, quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(mdata.id ? { remoteJid: mdata.id } : {}) }, message: { conversation: `_Notifikasi grup_` }}})
				client.sendMessage(mdata.id, ind.intro(), MessageType.text, {quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(mdata.id ? { remoteJid: mdata.id } : {}) }, message: { conversation: `salin ini!  Usahakan jangan ada teks yang dihapus` }}})
			} else if (anu.action == 'remove') {
				num = anu.participants[0]
				try {
					ppimg = await client.getProfilePicture(`${num.split('@')[0]}@c.us`)
				} catch {
					ppimg = 'https://i0.wp.com/www.gambarunik.id/wp-content/uploads/2019/06/Top-Gambar-Foto-Profil-Kosong-Lucu-Tergokil-.jpg'
				}
				teks = `SELAMAT TINGGAL... @${num.split('@')[0]}👋* \n_Haduh haduh, ngapain saya ngirim ini. Dia kan tidak akan tahu_\n\nYang lain, mohon untuk follow instagramku *>* \ninstagram.com/frm_developer`
				let buff = await getBuffer(ppimg)
				client.sendMessage(mdata.id, teks, MessageType.text, {contextInfo: {"mentionedJid": [num]}, quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(mdata.id ? { remoteJid: mdata.id } : {}) }, message: { conversation: `_Notifikasi grup_` }}})
			} else if (anu.action == 'promote') {
				num = anu.participants[0]
				try {
					ppimg = await client.getProfilePicture(`${num.split('@')[0]}@c.us`)
				} catch {
					ppimg = 'https://i0.wp.com/www.gambarunik.id/wp-content/uploads/2019/06/Top-Gambar-Foto-Profil-Kosong-Lucu-Tergokil-.jpg'
				}
				teks = `SELAMAT UNTUK ADMIN BARU... @${num.split('@')[0]}👋* \n\n*SAYA JADIKAN ADMIN DONG, PLIIIIS*\njika sudah, saya jadikan admin lagi. Biar kebal\n\nMohon untuk follow instagramku *>* \ninstagram.com/frm_developer`
				let buff = await getBuffer(ppimg)
				client.sendMessage(mdata.id, teks, MessageType.text, {contextInfo: {"mentionedJid": [num]}, quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(mdata.id ? { remoteJid: mdata.id } : {}) }, message: { conversation: `_Notifikasi grup_` }}})
			} else if (anu.action == 'demote') {
				num = anu.participants[0]
				try {
					ppimg = await client.getProfilePicture(`${num.split('@')[0]}@c.us`)
				} catch {
					ppimg = 'https://i0.wp.com/www.gambarunik.id/wp-content/uploads/2019/06/Top-Gambar-Foto-Profil-Kosong-Lucu-Tergokil-.jpg'
				}
				teks = `SAYA TURUT BERDUKA UNTUK ADMIN YANG DIPECAT... @${num.split('@')[0]}👋*\n\nMohon untuk follow instagramku *>* \ninstagram.com/frm_developer`
				let buff = await getBuffer(ppimg)
				client.sendMessage(mdata.id, teks, MessageType.text, {contextInfo: {"mentionedJid": [num]}, quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(mdata.id ? { remoteJid: mdata.id } : {}) }, message: { conversation: `_Notifikasi grup_` }}})
			}
		} catch (e) {
			console.log('Error : %s', color(e, 'red'))
		}
	})
	client.on('CB:Blocklist', json => {
		if (blocked.length > 2) return
	    for (let i of json[1].blocklist) {
	    	blocked.push(i.replace('c.us','s.whatsapp.net'))
	    }
	})
	
	client.on('CB:action,,call', async json => {
    const callerId = json[2][0][1].from;
    console.log("call dari "+ callerId)
        await client.sendMessage(callerId, `Jangan melakukan panggilan suara dengan bot, silahkan lakukan pamggilan suara dengan wa.me/${nomowner}\nOKE`, MessageType.text)
        await client.blockUser(callerId, "add")
})
	
	client.on('CB:action,,battery', json => {
		global.batteryLevelStr = json[2][0][1].value
		global.batterylevel = parseInt(batteryLevelStr)
		baterai = batterylevel
		batre = []
		statusbatre = json[2][0][1]
		batre.push(statusbatre)
        if (json[2][0][1].live == 'true') charging = true
        if (json[2][0][1].live == 'false') charging = false
        console.log(json[2][0][1])
		sisabaterai = `${json[2][0][1].value}%`
		hematdaya = json[2][0][1].powersave
		dicas = json[2][0][1].live
	})

	client.on('chat-update', async (mek) => {
		try {
			if (!mek.hasNewMessage) return
			mek = JSON.parse(JSON.stringify(mek)).messages[0]
			if (!mek.message) return
			mek.message = (Object.keys(mek.message)[0] === 'ephemeralMessage') ? mek.message.ephemeralMessage.message : mek.message
			if (mek.key && mek.key.remoteJid == 'status@broadcast') return
			if (mek.key.fromMe) return
			global.prefix
			global.blocked
			const me = client.user
			const content = JSON.stringify(mek.message)
			const from = mek.key.remoteJid
			const type = Object.keys(mek.message)[0]
			const { text, extendedText, contact, location, liveLocation, image, video, sticker, document, audio, product } = MessageType
            body = (type === 'conversation' && mek.message.conversation) ? mek.message.conversation : (type == 'imageMessage') && mek.message.imageMessage.caption ? mek.message.imageMessage.caption : (type == 'videoMessage') && mek.message.videoMessage.caption ? mek.message.videoMessage.caption : (type == 'extendedTextMessage') && mek.message.extendedTextMessage.text ? mek.message.extendedTextMessage.text : ''
			budy = (type === 'conversation') ? mek.message.conversation : (type === 'extendedTextMessage') ? mek.message.extendedTextMessage.text : ''
			if (body === null || body === undefined) return console.log('TIDAK ADA TEKS')
			const time = moment.tz('Asia/Jakarta').format('DD/MM HH:mm:ss')
			const timi = moment.tz('Asia/Jakarta').add(30, 'days').calendar();
			const timu = moment.tz('Asia/Jakarta').add(20, 'days').calendar();
			const command = body.slice(0).trim().split(/ +/).shift().toLowerCase()
			const args = body.trim().split(/ +/).slice(1)
			const cilik = body.slice(0).toLowerCase()
			const isCmd = body.startsWith(prefix)
			const tescuk = ["0@s.whatsapp.net"]
			const isGroup = mek.key.remoteJid.endsWith('@g.us')
			const q = args.join(' ')
			const qcilik = args.join(' ').toLowerCase()
			const botNumber = client.user.jid
			const sender = isGroup ? mek.participant : mek.key.remoteJid
			pushname = client.contacts[sender] != undefined ? client.contacts[sender].vname || client.contacts[sender].notify : undefined
			const speed = require('performance-now')
			const totalchat = await client.chats.all()
			const d = new Date
			const locale = 'id'
			const hari = d.toLocaleDateString(locale, { weekday: 'long' })
			const jamtok = moment.tz('Asia/Jakarta').format('HH')
			const menittok = moment.tz('Asia/Jakarta').format('mm')
			const tanggaltok = moment.tz('Asia/Jakarta').format('DD')
			const jam = moment.tz('Asia/Jakarta').format('HH:mm:ss')
			const jammenit = moment.tz('Asia/Jakarta').format('HH:mm')
			const tanggal = moment.tz('Asia/Jakarta').format('DD/MM/YYYY')
			const gmt = new Date(0).getTime() - new Date('1 January 1970').getTime()
			pushname2 = client.contacts[sender] != undefined ? client.contacts[sender].vname || client.contacts[sender].notify : undefined
			const groupMetadata = isGroup ? await client.groupMetadata(from) : ''
			const groupName = isGroup ? groupMetadata.subject : ''
			const groupId = isGroup ? groupMetadata.jid : ''
			const groupMembers = isGroup ? groupMetadata.participants : ''
			const groupDesc = isGroup ? groupMetadata.desc : ''
            const groupAdmins = isGroup ? getGroupAdmins(groupMembers) : ''

/*
]=====> Rifki ID <=====[
*/
            const isEventon = isGroup ? event.includes(from) : false
            const isRegistered = checkRegisteredUser(sender)
            const isUser = cekWesDaftar(nomerwesdaftar)
            const isBotGroupAdmins = groupAdmins.includes(botNumber) || false
            const isLevelingOn = isGroup ? _leveling.includes(from) : false
			const isGroupAdmins = groupAdmins.includes(sender) || false
			const isWelkom = isGroup ? welkom.includes(from) : false
			const isNsfw = isGroup ? nsfw.includes(from) : false
			const isSimi = isGroup ? samih.includes(from) : false
			const isOwner = ownerNumber.includes(sender)
			const isRoboGuru = nomereroboguru.includes(sender)
			const isNggoRoboguru = nggoroboguru.includes(sender)
			const isBanned = ban.includes(sender)
			const isNggoPln = nggopln.includes(sender)
			const isPln = nomerepln.includes(sender)
			const isImage = type === 'imageMessage'
			const isGanggu = sender.includes(sender)
			const isPromo = promo.includes(sender)
			const isAfkOn = checkAfkUser(sender, _afk)
			const isOnOff = statuson
			const isUrl = (url) => {
			    return url.match(new RegExp(/https?:\/\/(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_+.~#?&/=]*)/, 'gi'))
			}
			const reply = (teks) => {
				client.updatePresence(from, Presence.composing)
				client.sendMessage(from, teks, text, {quoted:mek})
			}
			const hapus = (dihapus) => {
			client.deleteMessage(from,  dihapus.key )
			}
			const sendMess = (hehe, teks) => {
				client.updatePresence(from, Presence.composing)
				client.sendMessage(hehe, teks, text)
			}
			const mentions = (teks, memberr, id) => {
				(id == null || id == undefined || id == false) ? client.sendMessage(from, teks.trim(), extendedText, {contextInfo: {"mentionedJid": memberr}}) : client.sendMessage(from, teks.trim(), extendedText, {quoted: mek, contextInfo: {"mentionedJid": memberr}})
			}
			const sendImage = (teks) => {
			client.updatePresence(from, Presence.composing)
		    client.sendMessage(from, teks, image, {quoted:mek})
		    }
		    const costum = (pesan, tipe, target, target2) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(from, pesan, tipe, {quoted: { key: { fromMe: false, participant: `${target}`, ...(from ? { remoteJid: from } : {}) }, message: { conversation: `${target2}` }}})
			}
			const fitnah = (ditag, pesannya, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(from, pesanku, text, {quoted: { key: { fromMe: false, participant: ditag, ...(from ? { remoteJid: 'status@broadcast' } : {}) }, message: { conversation: pesannya }}})
			}
			const tuduh = (ditag, pesannya, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(from, pesanku, text, {quoted: { key: { fromMe: false, participant: ditag, ...(from ? { remoteJid: from } : {}) }, message: { conversation: pesannya }}})
			}
			const tagpalsu = (ditag, pesannya, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(from, pesanku, text, {quoted: { key: { fromMe: false, participant: `${ditag}@s.whatsapp.net`, ...(from ? { remoteJid: from } : {}) }, message: { conversation: pesannya }}})
			}
			const faketag = (untuk, ditag, pesannya, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(untuk, pesanku, text, {quoted: { key: { fromMe: false, participant: ditag, ...(from ? { remoteJid: untuk } : {}) }, message: { conversation: pesannya }}})
			}
			const fitnah2 = (untuk, ditag, pesannya, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(untuk, pesanku, text, {quoted: { key: { fromMe: false, participant: ditag, ...(from ? { remoteJid: 'status@broadcast' } : {}) }, message: { conversation: pesannya }}})
			}
			const fitnah3 = (untuk, ditag, pesannya, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(untuk, pesanku, text, {quoted: { key: { fromMe: false, participant: ditag, ...(from ? { remoteJid: '6281615901727-1612997366@g.us' } : {}) }, message: { conversation: pesannya }}})
			}
			const fitnahgmbr = (ditag, pesannya, gambare, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(from, pesanku, gambare, image, {quoted: { key: { fromMe: false, participant: ditag, ...(from ? { remoteJid: 'status@broadcast' } : {}) }, message: { conversation: pesannya }}})
			}
			const fitnahgmbr2 = (untuk, ditag, pesannya, gambare, pesanku) => {
			client.updatePresence(from, Presence.composing)
			client.sendMessage(untuk, pesanku, gambare, image, {quoted: { key: { fromMe: false, participant: ditag, ...(from ? { remoteJid: 'status@broadcast' } : {}) }, message: { conversation: pesannya }}})
			}
		    const sendPtt = (teks) => {
			client.updatePresence(from, Presence.composing)
		    client.sendMessage(from, audio, mp3, {quoted:mek})
		    }
			const diteruskan = (teruskan) => {
				client.updatePresence(from, Presence.composing)
				client.sendMessage(from, teruskan, text, {quoted : mek, contextInfo: { forwardingScore: 1000, isForwarded: true}})
			}
			const stiker = async(dadistiker) => {
				stickermk = new WSF.Sticker(`${dadistiker}`, { crop: true, animated: false, pack: `${pushname}`, author: 'By FRM BOT' })
				await stickermk.build()
				stcBuffr = await stickermk.get()
				client.updatePresence(from, Presence.composing)
				client.sendMessage(from, stcBuffr, sticker, {quoted:mek}).catch((err) => reply('error'))
			}
			const stikergif = async(bahanstiker) => {
				stickermk = new WSF.Sticker(`${bahanstiker}`, { crop: true, animated: true, pack: `${pushname}`, author: 'By FRM BOT' })
				await stickermk.build()
				stcBuffr = await stickermk.get()
				client.updatePresence(from, Presence.composing)
				client.sendMessage(from, stcBuffr, sticker, {quoted:mek}).catch((err) => reply('error'))
			}
			const fakethumb = (gmbrnya, captionnya, teksnya) => {
				client.updatePresence(from, Presence.composing)
            	client.sendMessage(from, teksnya, image, {thumbnail:gmbrnya,quoted:mek,caption:captionnya})
        	}
			const sendMediaURL = async(to, url, text="", mids=[]) =>{
                if(mids.length > 0){
                    text = normalizeMention(to, text, mids)
                }
                const fn = Date.now() / 10000;
                const filename = fn.toString()
                let mime = ""
                var download = function (uri, filename, callback) {
                    request.head(uri, function (err, res, body) {
                        mime = res.headers['content-type']
                        request(uri).pipe(fs.createWriteStream(filename)).on('close', callback);
                    });
                };
                download(url, filename, async function () {
                    console.log('done');
                    let media = fs.readFileSync(filename)
                    let type = mime.split("/")[0]+"Message"
                    if(mime === "image/gif"){
                        type = MessageType.video
                        mime = Mimetype.gif
                    }
                    if(mime.split("/")[0] === "audio"){
                        mime = Mimetype.mp4Audio
                    }
                    selfb.sendMessage(to, media, type, { quoted: mek, mimetype: mime, caption: text,contextInfo: {"mentionedJid": mids}})
                    
                    fs.unlinkSync(filename)
                });
            }  
			const fakestatus = (teks) => {
				client.updatePresence(from, Presence.composing)
            	client.sendMessage(from, teks, text, {
                quoted: {
                    key: {
                        fromMe: false,
                        participant: `${sender}`, ...(from ? { remoteJid: "status@broadcast" } : {})
                    },
                    message: {
                        "imageMessage": {
                            "url": "https://mmg.whatsapp.net/d/f/At0x7ZdIvuicfjlf9oWS6A3AR9XPh0P-hZIVPLsI70nM.enc",
                            "mimetype": "image/jpeg",
                            "caption": `${body}`,
                            "fileSha256": "+Ia+Dwib70Y1CWRMAP9QLJKjIJt54fKycOfB2OEZbTU=",
                            "fileLength": "28777",
                            "height": 1080,
                            "width": 1079,
                            "mediaKey": "vXmRR7ZUeDWjXy5iQk17TrowBzuwRya0errAFnXxbGc=",
                            "fileEncSha256": "sR9D2RS5JSifw49HeBADguI23fWDz1aZu4faWG/CyRY=",
                            "directPath": "/v/t62.7118-24/21427642_840952686474581_572788076332761430_n.enc?oh=3f57c1ba2fcab95f2c0bb475d72720ba&oe=602F3D69",
                            "mediaKeyTimestamp": "1610993486",
                            "jpegThumbnail": fs.readFileSync('./fauzan.rifki.m/quoted.png'),
                            "scansSidecar": "1W0XhfaAcDwc7xh1R8lca6Qg/1bB4naFCSngM2LKO2NoP5RI7K+zLw=="
                        }
                    }
                }
            })
        }
        	const statuswa = {
                    key: {
                        fromMe: false,
                        participant: `${sender}`, ...(from ? { remoteJid: "status@broadcast" } : {})
                    },
                    message: {
                        "imageMessage": {
                            "url": "https://mmg.whatsapp.net/d/f/At0x7ZdIvuicfjlf9oWS6A3AR9XPh0P-hZIVPLsI70nM.enc",
                            "mimetype": "image/jpeg",
                            "caption": `${body}`,
                            "fileSha256": "+Ia+Dwib70Y1CWRMAP9QLJKjIJt54fKycOfB2OEZbTU=",
                            "fileLength": "28777",
                            "height": 1080,
                            "width": 1079,
                            "mediaKey": "vXmRR7ZUeDWjXy5iQk17TrowBzuwRya0errAFnXxbGc=",
                            "fileEncSha256": "sR9D2RS5JSifw49HeBADguI23fWDz1aZu4faWG/CyRY=",
                            "directPath": "/v/t62.7118-24/21427642_840952686474581_572788076332761430_n.enc?oh=3f57c1ba2fcab95f2c0bb475d72720ba&oe=602F3D69",
                            "mediaKeyTimestamp": "1610993486",
                            "jpegThumbnail": fs.readFileSync('./fauzan.rifki.m/quoted.png'),
                            "scansSidecar": "1W0XhfaAcDwc7xh1R8lca6Qg/1bB4naFCSngM2LKO2NoP5RI7K+zLw=="
                        }
                    }
                }
		
			/***************** ngganti prefix ********/
				  if (body.startsWith(`.`)) {
                  	prefix = '.'
                  }
                  if (body.startsWith(`!`)) {
                  	prefix = '!'
                  }
                  if (body.startsWith(`#`)) {
                  	prefix = '#'
                  }
                  if (body.startsWith(`/`)) {
                  	prefix = '/'
                  }
                  if (body.startsWith(`~`)) {
                  	prefix = '~'
                  }
                  if (body.startsWith(`%`)) {
                  	prefix = '%'
                  }
                  if (body.startsWith(`-`)) {
                  	prefix = '-'
                  }
                  if (body.startsWith(`+`)) {
                  	prefix = '+'
                  }
                  if (body.startsWith(`=`)) {
                  	prefix = '='
                  }
                  if (body.startsWith(`_`)) {
                  	prefix = '_'
                  }
            /***************** ngganti prefix ********/
            
            // TAMBAHAN SAAT BOT OFF / ON
            if (command.includes(`${prefix}bot`) && qcilik.includes(`on`) && isOwner) {
				if (isOnOff) return reply('SUDAH ON')
				statuson = true
				reply('BERHASIL MENYALAKAN')
			}
			if (command.includes(`${prefix}bot`) && qcilik.includes(`off`) && isOwner) {
				if (!isOnOff) return reply('SUDAH OFF')
				statuson = false
				reply('BERHASIL MEMATIKAN')
			}
			if (command.includes(`${prefix}owneronly`) && qcilik.includes(`on`) && isOwner) {
				if (statusbot) return reply('MODE OWNER SAJA SUDAH AKTIF')
				statusbot = true
				reply('MODE OWNER SAJA AKTIF')
			}
			if (command.includes(`${prefix}owneronly`) && qcilik.includes(`off`) && isOwner) {
				if (!statusbot) return reply('MODE OWNER SAJA SUDAH MATI')
				statusbot = false
				reply('MODE OWNER SAJA MATI')
			}
			if (budy.includes(`🌿🌿🌿🌿🌿`) && budy.endsWith(`🍃🍃🍃🍃🍃`)) {
				intro0 = `${body.split(`🌿🌿🌿🌿🌿`)[1]}`
				intro1 = `${intro0.split(`🍃🍃🍃🍃🍃`)[0]}`
				kosong = ''
				introne = `*🤝 PERKENALAN DITERIMA 🤝*\n${intro1}\n*🤝 TERIMAKASIH ??*\n\n_🌱 bot ini ramah lingkungan 🌱_\n_🌱 tidak mengandung zat nuklir 🌱_`
				reply(`${introne}`)
			}
			if (budy.includes(`🌿🌿🌿🌿🌿`)) {
				reply(`WIDIH`)
			}
if (!isOnOff) return
			//pesan tambahan
				if (cilik === `tes` || cilik === `woy` || cilik === `bot`) {
					client.updatePresence(from, Presence.composing)
					client.sendMessage(from, `Maaf, ada perlu apa`, text, {quoted: statuswa})
				  }
				if (cilik === `p` || cilik === `${prefix}p` || cilik === `🅿️`)  {
				  client.updatePresence(from, Presence.composing)
				  reply(`السَّلاَمُ عَلَيْكُمْ وَرَحْمَةُ اللهِ وَبَرَكَاتُهُ`)
				  reply(`hai`)
				  }
                if (budy.includes(`@${me.jid.split('@')[0]}`)) {
                  	client.updatePresence(from, Presence.composing)
                  	client.updatePresence(from, Presence.available)
                  	ditagwong = fs.readFileSync(`./fauzan.rifki.m/ditagwong.webp`)
            		client.sendMessage(from, ditagwong, sticker, {quoted: mek})
                  }
				if (cilik.includes(`ripki`) || cilik.includes(`rifki`) || cilik.includes(`fauzan`)) {
					client.updatePresence(from, Presence.available)
                    const ip =['Ada yang aneh','Ya ada apa','Ada perlu apa','kenapa?','Kamu lagi ngapain?','Disana enak','Apa bisa saya bantu','Halo','Apa kabar','Assalamualaikum','Waalaikumsalam','Kamu siapa','Halo','Hey','Aku siapa','Saya siapa hayoo?','sekarang jam berapa','Kamu siapa','Rumahmu dimana','Mau ngapain','selamat pagi','Selamat siang','Selamat sore','Selamat malam','Selamat tidur','Selamat Whatsapp an','Kamu sekarang lagi apa','Ini siapa ya?','1+1=2','Namaku siapa?']
					const ki = ip[Math.floor(Math.random() * ip.length)]
                  ripki = reply(`Nama ownerku terdeteksi\n${ki}`)
                  client.deleteMessage(from, ripki.key)
                  }
			const listbahasa = `*List kode Bahasa*\n
*Code       Bahasa*
${monosp} sq        Albanian
 ar        Arabic
 hy        Armenian
 ca        Catalan
 zh        Chinese
 zh-cn     Chinese (China)
 zh-tw     Chinese (Taiwan)
 zh-yue    Chinese (Cantonese)
 hr        Croatian
 cs        Czech
 da        Danish
 nl        Dutch
 en        English
 en-au     English (Australia)
 en-uk     English (United Kingdom)
 en-us     English (United States)
 eo        Esperanto
 fi        Finnish
 fr        French
 de        German
 el        Greek
 ht        Haitian Creole
 hi        Hindi
 hu        Hungarian
 is        Icelandic
 id        Indonesian
 it        Italian
 ja        Japanese
 ko        Korean
 la        Latin
 lv        Latvian
 mk        Macedonian
 no        Norwegian
 pl        Polish
 pt        Portuguese
 pt-br     Portuguese (Brazil)
 ro        Romanian
 ru        Russian
 sr        Serbian
 sk        Slovak
 es        Spanish
 es-es     Spanish (Spain)
 es-us     Spanish (United States)
 sw        Swahili
 sv        Swedish
 ta        Tamil
 th        Thai
 tr        Turkish
 vi        Vietnamese
 cy        Welsh
      ${monosp}`
      
      
/*
]=====> LEVELING <=====[
*/
            if (isGroup && isRegistered && isLevelingOn) {
            const currentLevel = getLevelingLevel(sender)
            const checkId = getLevelingId(sender)
            try {
                if (currentLevel === undefined && checkId === undefined) addLevelingId(sender)
                const amountXp = Math.floor(Math.random() * 10) + 500
                const requiredXp = 5000 * (Math.pow(2, currentLevel) - 1)
                const getLevel = getLevelingLevel(sender)
                addLevelingXp(sender, amountXp)
                if (requiredXp <= getLevelingXp(sender)) {
                    addLevelingLevel(sender, 1)
                    bayarLimit(sender, 3)
                    await reply(ind.levelup(namaneuser, sender, getLevelingXp,  getLevel, getLevelingLevel))
                }
            } catch (err) {
                console.error(err)
            }
        }
        
        	if (!isGroup && isRegistered) {
            const currentLevel = getLevelingLevel(sender)
            const checkId = getLevelingId(sender)
            try {
                if (currentLevel === undefined && checkId === undefined) addLevelingId(sender)
                const amountXp = Math.floor(Math.random() * 10) + 500
                const requiredXp = 5000 * (Math.pow(2, currentLevel) - 1)
                const getLevel = getLevelingLevel(sender)
                addLevelingXp(sender, amountXp)
                if (requiredXp <= getLevelingXp(sender)) {
                    addLevelingLevel(sender, 1)
                    bayarLimit(sender, 3)
                    await reply(ind.levelup(namaneuser, sender, getLevelingXp,  getLevel, getLevelingLevel))
                }
            } catch (err) {
                console.error(err)
            }
        }
/*
]=====> CHECK LIMIT BY LANN ID <=====[
*/
          const checkLimit = (sender) => {
          	let found = false
                    for (let lmt of _limit) {
                        if (lmt.id === sender) {
                            let limitCounts = limitawal - lmt.limit
                            if (limitCounts <= 0) return client.sendMessage(from,`Limit anda sudah habis\n\n_Note : limit bisa di dapatkan dengan cara ${prefix}buylimit dan naik level atau besok anda kami beri ${limitawal} pesan_`, text,{ quoted: mek})
                            reply(ind.limitcount(limitCounts))
                            found = true
                        }
                    }
                    if (found === false) {
                        let obj = { id: sender, limit: 0 }
                        _limit.push(obj)
                        fs.writeFileSync('./database/user/limit.json', JSON.stringify(_limit))
                        reply(ind.limitcount(limitCounts))
                    }
				}
			const ceklimit = (sender) => {
          	let cari = false
                    for (let lmit of _limit) {
                        if (lmit.id === sender) {
                            let sisalimit = limitawal - lmit.limit
                            return sisalimit
                            cari = true
                        }
                    }
                    if (cari === false) {
                        let objl = { id: sender, limit: 0 }
                        _limit.push(objl)
                        fs.writeFileSync('./database/user/limit.json', JSON.stringify(_limit))
                    }
				}
			
			
				
/*
]=====> LIMITED BY LANN ID <=====[
*/
           const isLimit = (sender) =>{ 
		      let position = false
              for (let i of _limit) {
              if (i.id === sender) {
              	let limits = i.limit
              if (limits >= limitawal ) {
              	  position = true
                    reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    return true
              } else {
              	_limit
                  position = true
                  return false
               }
             }
           }
           if (position === false) {
           	const obj = { id: sender, limit: 0 }
                _limit.push(obj)
                fs.writeFileSync('./database/user/limit.json',JSON.stringify(_limit))
           return false
       }
     }

        
            if (isGroup) {
					try {
						const getmemex = groupMembers.length	
					    if (getmemex < memberlimit) {
						reply(`maaf member group belum memenuhi syarat. minimal member group adalah ${memberlimit}`)
						setTimeout( () => {
 	                           client.groupLeave(from) 
 					   	}, 5000)
								setTimeout( () => {
								client.updatePresence(from, Presence.composing)
								reply("1detik")
							}, 4000)
								setTimeout( () => {
								client.updatePresence(from, Presence.composing)
								reply("2detik")
							}, 3000)
								setTimeout( () => {
								client.updatePresence(from, Presence.composing)
								reply("3detik")
							}, 2000)
								setTimeout( () => {
								client.updatePresence(from, Presence.composing)
								reply("4detik")
							}, 1000)
								setTimeout( () => {
								client.updatePresence(from, Presence.composing)
								reply("5detik")
							}, 0)
					    }
		       } catch (err) { console.error(err)  }
 	       }
 
			
            if (checkAfkUser(sender, _afk) && !isCmd) {
                _afk.splice(getAfkPosition(sender, _afk), 1)
                fs.writeFileSync('./database/user/afk.json', JSON.stringify(_afk))
                reply(ind.afkDone(namaneuser(sender)))
            }

			colors = ['red','white','black','blue','yellow','green']
			const isMedia = (type === 'imageMessage' || type === 'videoMessage')
			const isQuotedImage = type === 'extendedTextMessage' && content.includes('imageMessage')
			const isQuotedVideo = type === 'extendedTextMessage' && content.includes('videoMessage')
			const isQuotedSticker = type === 'extendedTextMessage' && content.includes('stickerMessage')
			const isQuotedAudio = type === 'extendedTextMessage' && content.includes('audioMessage')
			//private chat message
			if (!isGroup && isCmd) console.log(color(`\n━━━━━━━━━━━━━━━━━━━━\n${sender.split('@')[0]} (${pushname2})`), (`\n${body}\n               ${jammenit}\n`))
			if (!isGroup && !isCmd) console.log(color(`\n━━━━━━━━━━━━━━━━━━━━\n${sender.split('@')[0]} (${pushname2})`), (`\n${body}\n               ${jammenit}\n`))
			
			//group message
			if (isCmd && isGroup) console.log(color(`\n━━━━━━━━━━━━━━━━━━━━\n${sender.split('@')[0]} (${pushname2})`), color(`•> ${groupName}`), (`\n${body}\n               ${jammenit}\n`))
			if (!isCmd && isGroup) console.log(color(`\n━━━━━━━━━━━━━━━━━━━━\n${sender.split('@')[0]} (${pushname2})`), color(`•> ${groupName}`), (`\n${body}\n               ${jammenit}\n`))

if (isBanned) return
switch(command) {
				case prefix+'shadow':
                case prefix+'cup':
                case prefix+'cup1':
                case prefix+'romance':
                case prefix+'smoke':
                case prefix+'burnpaper':
                case prefix+'lovemessage':
                case prefix+'undergrass':
                case prefix+'love':
                case prefix+'coffe':
                case prefix+'woodheart':
                case prefix+'flowerheart':
                case prefix+'woodenboard':
                case prefix+'summer3d':
                case prefix+'wolfmetal':
                case prefix+'nature3d':
                case prefix+'underwater':
                case prefix+'golderrose':
                case prefix+'summernature':
                case prefix+'letterleaves':
                case prefix+'glowingneon':
                case prefix+'fallleaves':
                case prefix+'flamming':
                case prefix+'harrypotter':
                case prefix+'carvedwood':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    if (args.length == 0) return reply('Teksnya mana um')
                    txt = args.join(" ")
                    lolimg = await getBuffer(`http://api.lolhuman.xyz/api/photooxy1/${body.slice(1).split(' ')[0]}?apikey=${LolKey}&text=${txt}`)
                    client.updatePresence(from, Presence.composing)
                    client.sendMessage(from, lolimg, image, {quoted:mek, caption: `By ${botName}`})
                    break
                case prefix+'wetglass':
                case prefix+'multicolor3d':
                case prefix+'watercolor':
                case prefix+'luxurygold':
                case prefix+'galaxywallpaper':
                case prefix+'lighttext':
                case prefix+'beautifulflower':
                case prefix+'puppycute':
                case prefix+'royaltext':
                case prefix+'heartshaped':
                case prefix+'birthdaycake':
                case prefix+'galaxystyle':
                case prefix+'hologram3d':
                case prefix+'glossychrome':
                case prefix+'greenbush':
                case prefix+'metallogo':
                case prefix+'noeltext':
                case prefix+'glittergold':
                case prefix+'textcake':
                case prefix+'starsnight':
                case prefix+'wooden3d':
                case prefix+'textbyname':
                case prefix+'writegalacy':
                case prefix+'galaxybat':
                case prefix+'snow3d':
                case prefix+'birthdayday':
                case prefix+'goldplaybutton':
                case prefix+'silverplaybutton':
                case prefix+'freefire':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    if (args.length == 0) return reply('Teksnya mana um')
                    txt = args.join(" ")
                    lolimg = await getBuffer(`http://api.lolhuman.xyz/api/ephoto1/${body.slice(1).split(' ')[0]}?apikey=${LolKey}&text=${txt}`)
                    client.updatePresence(from, Presence.composing)
                    client.sendMessage(from, lolimg, image, {quoted:mek, caption: `By ${botName}`})
                    break
                case prefix+'darkneon':
                case prefix+'candlemug':
                case prefix+'lovemsg':
                case prefix+'mugflower':
                case prefix+'narutobanner':
                case prefix+'paperonglass':
                case prefix+'romancetext':
                case prefix+'shadowtext':
                case prefix+'coffecup':
                case prefix+'coffecup2':
                case prefix+'glowingneon':
                case prefix+'underwater':
                case prefix+'hpotter':
                case prefix+'woodblock':
                	if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    if (args.length == 0) return reply('Teksnya mana um')
                    txt = args.join(" ")
                    reply(ind.wait())
                    buffer = await getBuffer(`https://videfikri.com/api/textmaker/${body.slice(1).split(' ')[0]}/?text=${txt}`)
                    client.sendMessage(from, buffer, image, {caption: 'Nih kak.. *Jangan lupa subscribe DappaUhuy*', quoted: mek})
                    break
			case prefix+'buat':
				const grouup = await client.groupCreate (`${q.split('&')[1]}`, [`${args[0]}@s.whatsapp.net`, `${args[1]}@s.whatsapp.net`])
reply (`Membuat grup`)
client.sendMessage(grouup.gid, "halo", text) // say hello to everyone on the group
			break
				case prefix+'join':
					if (args.length < 1) return reply(`Mohon berikan tautan tndangan grup`)
					let linkgrup = `${body.split('whatsapp.com/')[1]}`
					let islink = q.match(/(https:\/\/chat.whatsapp.com)/gi)
					if (!islink) return reply('Maaf link group-nya salah! ')
						response = await client.acceptInvite (linkgrup)
						reply("Bergabung ke: " + response.gid)
					break							
/*
]=====> SIMPLE MENU <=====[
*/			  case prefix+'stiker': 
				case prefix+'sticker':
				
				    if (!isRegistered) return reply(ind.noregis())
				    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
						encmedia = isQuotedImage ? JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo : mek
						media = await client.downloadAndSaveMediaMessage(encmedia)
						stiker('./undefined.jpeg')
					} else if ((isMedia && mek.message.videoMessage.seconds < 11 || isQuotedVideo && mek.message.extendedTextMessage.contextInfo.quotedMessage.videoMessage.seconds < 11) && args.length == 0) {
						const encmedia = isQuotedVideo ? JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo : mek
						const media = await client.downloadAndSaveMediaMessage(encmedia)
						stikergif('./undefined.jpeg')
						}
					await limitAdd(sender)
						break
			case prefix+'nulis2':
			case prefix+'tulis2':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply('Teksnya mana kak? Contoh : ${prefix}nulis1 Rifki baik hati')
				reply('「❗」WAIT BRO GUE NULIS DUMLU YAKAN')
				buff = await getBuffer(`https://api.zeks.xyz/api/nulis?text=${q}&apikey=${ZeksKey}`)
				client.sendMessage(from, buff, image, {quoted: mek, caption: 'Lebih baik nulis sendiri ya kak :*'})
				await limitAdd(sender)
				break
				case prefix+'blackpink':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`「❗」Contoh : ${prefix}blackpink Rifki`)
					pink = body.slice(11)
					reply('「❗」Hah Blekping :v')
					lol = await getBuffer(`https://api.zeks.xyz/api/logobp?text=${pink}&apikey=${ZeksKey}`)
					client.sendMessage(from, lol, image, {quoted: mek})
					await limitAdd(sender)
					break
				case prefix+'ninjalogo':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				var gh = body.slice(11)
				var nin = gh.split("&")[0];
				var ja = gh.split("&")[1];
				if (args.length < 1) return reply('「❗」Contoh : ${prefix}ninjalogo Rifki & Gans')
				reply(ind.wait())
				buffer = await getBuffer(`https://api.xteam.xyz/textpro/ninjalogo?text=${nin}&text2=${ja}&APIKEY=${XteamKey}`)
				client.sendMessage(from, buffer, image, {quoted: mek})
				await limitAdd(sender)
				break
			case prefix+'coffetext':
                  if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`「❗」Contoh : ${prefix}blackpink Rifki`)
					coff = body.slice(11)
					mhe = await fetchJson(`https://api.shizukaa.xyz/api/coffie?apikey=${shizukakey}&text=${coff}`)
					reply(ind.wait())
					atu = await getBuffer(mhe.result.url)
					client.sendMessage(from, atu, image, {quoted: mek})
					await limitAdd(sender)
					break				
		case prefix+'halloweentext':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
				if (args.length < 1) return reply(ind.wrongf())
				ween = body.slice(15)
				if (ween.length > 10) return reply('Teksnya kepanjangan, maksimal 9 karakter')
				reply(ind.wait())
				buffer = await getBuffer(`https://api.xteam.xyz/textpro/helloweenfire?text=${ween}&APIKEY=${XteamKey}`)
		    client.sendMessage(from, buffer, image, {quoted: mek})
		    await limitAdd(sender)	
		    break
				case prefix+'pornhub':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				var gh = body.slice(9)
				var porn = gh.split("&")[0];
				var hub = gh.split("&")[1];
				if (args.length < 1) return reply(`#pornhub FRM & bot\n*itu contohnya*`)
				reply(ind.wait())
				buffer = await getBuffer(`https://api.xteam.xyz/textpro/ph?text=${porn}&text2=${hub}&APIKEY=${XteamKey}`)
				client.sendMessage(from, buffer, image, {quoted: mek})
				await limitAdd(sender)
				break
//GEMBOK TEXT
                case prefix+'gemboktext':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					var gh = body.slice(12)
					var gem = gh.split("&")[0];
					var bok = gh.split("&")[1];
					if (args.length < 1) return reply('[❗] Contoh : ${prefix}gemboktext 11 01 2021 & Rifki dan Nadia')
					reply(ind.wait())
					buffer = await getBuffer(`https://api.vhtear.com/padlock?text1=${gem}&text2=${bok}&apikey=${VhtearKey}`)
					client.sendMessage(from, buffer, image, {quoted: mek})
					await limitAdd(sender)
					break
                case prefix+'glitchtext':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					var gh = body.slice(12)
					var gli = gh.split("&")[0];
					var tch = gh.split("&")[1];
					if (args.length < 1) return reply('[❗] Contoh : ${prefix}glitchtext Rifki & Gans')
					reply(ind.wait())
					buffer = await getBuffer(`https://api.xteam.xyz/textpro/glitch?text=${gli}&text2=${tch}&APIKEY=${XteamKey}`)
					client.sendMessage(from, buffer, image, {quoted: mek})
					await limitAdd(sender)
					break
					case prefix+'ttperr':
					if (!isRegistered) return reply(ind.noregis())
				    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('yang mau dijadiin text sticker apaan, titit kah?')
					ranp = getRandom('.png')
					rano = getRandom('.webp')
					teks = body.slice(4).trim()
					anu = await fetchJson(`https://tobz-api.herokuapp.com/api/ttp?text=${body.slice(5)}&apikey=${TobzKey}`, {method: 'get'})
					if (anu.error) return reply(anu.error)
					exec(`wget ${anu.result} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=20 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						client.sendMessage(from, fs.readFileSync(rano), sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
                        await limitAdd(sender)
					break
				case prefix+'toimg':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (!isQuotedSticker) return reply('Reply atau Tag sticker yang mau dijadiin gambar kak >_<')
					reply(ind.wait())
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.png')
					exec(`ffmpeg -i ${media} ${ran}`, (err) => {
						fs.unlinkSync(media)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(ran)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: 'nih kak [(^.^)]'})
						fs.unlinkSync(ran)
					})
					await limitAdd(sender)
					break
                case prefix+'bikinquote':
                
                if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                var gh = body.slice(12)
					var quote = gh.split("&")[0];
					var wm = gh.split("&")[1];
					 pref = `yang mau dijadiin quote apaan, titit?\n\ncontoh : ${prefix}bikinquote Hidup hancur jika tidak utuh & By. FRM Developer`
					if (args.length < 1) return reply(pref)
					reply(ind.wait())
					bikinquote = await fetchJson(`https://terhambar.com/aw/qts/?kata=${quote}&author=${wm}&tipe=random`, {method: 'get'})
					buffer = await getBuffer(bikinquote.result)
					client.sendMessage(from, buffer, image, {caption: 'Nih kak >_<', quoted: mek})
					await limitAdd(sender)
					break
                   case prefix+'stalkig':
                   case prefix+'igstalk':
                   if (!isRegistered) return reply(ind.noregis())
					if (isBanned) return reply(`Maaf, nomor kamu tidak dapat menggunakan bot ini\nSilahkan mohon kepada bosku / ownerku`)
                   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                     teks = body.slice(9)
                     igstalk = await fetchJson(`https://api.vhtear.com/igprofile?query=${teks}&apikey=${VhtearKey}`, {method: 'get'})
                     reply('「❗」Sabar Lagi Stalking IG nya kak')
                     buffer = await getBuffer(igstalk.result.picture)
                     hasil = `*Username?* : _${igstalk.result.username}_ \n *Nama??* : _${igstalk.result.full_name}_ \n *Jumlah Follower??﹦?* : _${igstalk.result.follower}_ \n *Jumlah Following?* : _${igstalk.result.follow}_ \n *Jumlah Post?* : _${igstalk.result.post_count}_ \n *Biografi?? :* _${igstalk.result.biography}`
                    client.sendMessage(from, buffer, image, {quoted: mek, caption: hasil})
                    await limitAdd(sender)
			       break
		case prefix+'silktext':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
				if (args.length < 1) return reply(ind.wrongf())
				silk = body.slice(10)
				if (silk.length > 7) return reply('Teksnya kepanjangan, maksimal 6 karakter')
				reply(ind.wait())
				buffer = await getBuffer(`https://api.vhtear.com/silktext?text=${silk}&apikey=${VhtearKey}`)
		    client.sendMessage(from, buffer, image, {quoted: mek})
		    await limitAdd(sender)	
		    break								
/*
]=====> GABUTZ MENU <=====[
*/
				case `${prefix}bisakah`:
				case `${prefix}bisa`:
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply(`Yang di tanyakan apa ?\n*KETIK* ${command} pertanyaanmu\n*CONTOH* ${command} saya hidup ?`)
					bisakah = body.slice(1)
					const bisa =['Jangan tanya saya','Tanya orang lain aja','Ngapain tanya saya','Tidak akan bisa','Bisa','Tidak Bisa']
					const keh = bisa[Math.floor(Math.random() * bisa.length)]
					reply('Pertanyaan : *'+bisakah+'*\n\nJawaban : '+ keh)
					await limitAdd(sender)
					break
				case `${prefix}kapankah`:
				case `${prefix}kapan`:
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply(`Yang di tanyakan apa ?\n*KETIK* ${command} pertanyaanmu\n*CONTOH* ${command} saya hidup ?`)
					kapankah = body.slice(1)
					const kapan =['Jangan tanya saya','Tanya orang lain aja','Ngapain tanya saya','Tidak akan pernah','Kalok udah kiamat','Besok','Lusa','Tadi','4 Hari Lagi','5 Hari Lagi','6 Hari Lagi','1 Minggu Lagi','2 Minggu Lagi','3 Minggu Lagi','1 Bulan Lagi','2 Bulan Lagi','3 Bulan Lagi','4 Bulan Lagi','5 Bulan Lagi','6 Bulan Lagi','1 Tahun Lagi','2 Tahun Lagi','3 Tahun Lagi','4 Tahun Lagi','5 Tahun Lagi','6 Tahun Lagi','1 Abad lagi','3 Hari Lagi']
					const koh = kapan[Math.floor(Math.random() * kapan.length)]
					reply('Pertanyaan : *'+kapankah+'*\n\nJawaban : '+ koh)
					await limitAdd(sender)
					break
           case `${prefix}apakah`:
           case `${prefix}apa`:
           if (!isRegistered) return reply(ind.noregis())
		   if (isBanned) return reply(ind.diban())
           if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
           if (args.length < 1) return reply(`Yang di tanyakan apa ?\n*KETIK* ${command} pertanyaanmu\n*CONTOH* ${command} saya hidup ?`)
					apakah = body.slice(1)
					const apa =['Jangan tanya saya','Tanya orang lain aja','Ngapain tanya saya','Nggak','Iya','Tidak','Bisa Jadi']
					const kah = apa[Math.floor(Math.random() * apa.length)]
					reply('Pertanyaan : *'+apakah+'*\n\nJawaban : '+ kah)
					await limitAdd(sender)
					break
			case `${prefix}rate`:
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply(`Yang di tanyakan apa ?\n*KETIK* ${command} pertanyaanmu\n*CONTOH* ${command} hidupku`)
					rate = body.slice(1)
					const ra =['4','9','17','28','34','48','59','62','74','83','97','100','29','94','75','82','41','39']
					const te = ra[Math.floor(Math.random() * ra.length)]
					reply('Pertanyaan : *'+rate+'*\n\nJawaban : '+ te+'%')
					await limitAdd(sender)
					break
			case prefix+'cantikcek':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isMedia && !isQuotedImage) return reply(`kirim gambar dengan teks (caption) *${command}*\n*ATAU*\ngeser gambar (tag) lalu balas ${command}`)
					cantik = body.slice(1)
					const can =['1','2','3','4','5','6','7','8','9','10','11','12','13','14','15','16','17','18','19','20','21','22','23','24','25','26','27','28','29','30','31','32','33','34','35','36','37','38','39','40','41','42','43','44','45','46','47','48','49','50','51','52','53','54','55','56','57','58','59','60','61','62','63','64','65','66','67','68','69','70','71','72','73','74','75','76','77','78','79','80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','96','97','98','99','100']
					const tik = can[Math.floor(Math.random() * can.length)]
					reply('Pertanyaan : *'+cantik+'*\n\nJawaban : '+ tik+'%')
					await limitAdd(sender)
					break
			case prefix+'gantengcek':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isMedia && !isQuotedImage) return reply(`kirim gambar dengan teks (caption) *${command}*\n*ATAU*\ngeser gambar (tag) lalu balas ${command}`)
					ganteng = body.slice(1)
					const gan =['1','2','3','4','5','6','7','8','9','10','11','12','13','14','15','16','17','18','19','20','21','22','23','24','25','26','27','28','29','30','31','32','33','34','35','36','37','38','39','40','41','42','43','44','45','46','47','48','49','50','51','52','53','54','55','56','57','58','59','60','61','62','63','64','65','66','67','68','69','70','71','72','73','74','75','76','77','78','79','80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','96','97','98','99','100']
					const teng = gan[Math.floor(Math.random() * gan.length)]
					reply('Pertanyaan : *'+ganteng+'*\n\nJawaban : '+ teng+'%')
					await limitAdd(sender)
					break
		case prefix+'watak':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Yang di tanyakan apa ?\n*KETIK* ${command} pertanyaanmu\n*CONTOH* ${command} hidupku`)
					watak = body.slice(1)
					const wa =['Penyayang','Pemurah','Pemarah','Pemaaf','Penurut','Baik','Baperan','Baik Hati','penyabar','UwU','top deh, pokoknya','Suka Membantu']
					const tak = wa[Math.floor(Math.random() * wa.length)]
					reply('Pertanyaan : *'+watak+'*\n\nJawaban : '+ tak)
					await limitAdd(sender)
				        break
		case prefix+'hobby':
                    if (!isRegistered) return reply(ind.noregis())
                   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                   if (args.length < 1) return reply(`Yang di tanyakan apa ?\n*KETIK* ${command} pertanyaanmu\n*CONTOH* ${command} Rifki`)
					hobby = body.slice(1)
					const hob =['Memasak','Membantu Atok','Mabar','Nobar','Sosmedtan','Membantu Orang lain','Nonton Anime','Nonton Drakor','Naik Motor','Nyanyi','Menari','Bertumbuk','Menggambar','Foto fotoan Ga jelas','Maen Game','Berbicara Sendiri']
					const by = hob[Math.floor(Math.random() * hob.length)]
					reply('Pertanyaan : *'+hobby+'*\n\nJawaban : '+ by)
					await limitAdd(sender)
					break
			case `${prefix}dadu2`:
                    if (!isRegistered) return reply(ind.noregis())
                   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					dadu = body.slice(1)
					daduu = ['https://i.ibb.co/Bw42zpY/jogodedados-128px-2.gif','https://i.ibb.co/njdfrHT/jogodedados-128px-1.gif','https://i.ibb.co/BBcyPp2/jogodedados-128px-3.gif','https://i.ibb.co/YhhDbX5/jogodedados-128px-4.gif','https://i.ibb.co/qFTd1K1/jogodedados-128px-6.gif','https://i.ibb.co/9g8ns1b/jogodedados-128px-5.gif']
					daduuu = daduu[Math.floor(Math.random() * daduu.length)]
					dadua = await getBuffer(daduuu)
					client.sendMessage(from, dadua, sticker, {quoted: mek})
            		await limitAdd(sender)
					break
			case `${prefix}dadu`:
                    if (!isRegistered) return reply(ind.noregis())
                   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					dadu = body.slice(1)
					daduu =['1','2','3','4','5','6']
					daduuu = daduu[Math.floor(Math.random() * daduu.length)]
					stiker(`./fauzan.rifki.m/dadu${daduuu}.webp`)
            		await limitAdd(sender)
					break
			case prefix+'dadu3':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				anu9 = await fetchJson(`https://leyscoders-api.herokuapp.com/api/dadu?apikey=demo`, {method:'get'})
    			stiker(anu9.result)
    			await limitAdd(sender)       	
				break
           case prefix+'seberapagay':
           
           if (!isRegistered) return reply(ind.noregis())
           if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					gay = body.slice(13)
		   seberapagay = await fetchJson(`https://arugaz.herokuapp.com/api/howgay`, {method: 'get'})
		   hasil = `Nih Liat Data Gay Si ${gay}\n\n\nPersentase Gay : ${seberapagay.persen}%\nAlert!!! : ${seberapagay.desc}`
		   reply(hasil)
		   await limitAdd(sender)
					break
                case prefix+'nangis':
                case prefix+'cry':
                
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					cry = await fetchJson(`https://waifu.pics/api/sfw/cry`, {method: 'get'})
					reply('SABAR NGAB')
					if (cry.error) return reply(cry.error)
					exec(`wget ${cry.url} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
					case prefix+'cium':
					
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					cium = await fetchJson(`https://waifu.pics/api/sfw/kiss`, {method: 'get'})
					reply('Mwahhh')
					if (cium.error) return reply(cium.error)
					exec(`wget ${cium.url} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
					case prefix+'peluk':
					
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					reply('Peyukkkk')
					peluk = await fetchJson(`https://waifu.pics/api/sfw/hug`, {method: 'get'})
					if (peluk.error) return reply(peluk.error)
					exec(`wget ${peluk.url} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						auu = fs.readFileSync(ranp)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					break
                case prefix+'truth':
                
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					const trut =['Pernah suka sama siapa aja? berapa lama?','Kalau boleh atau kalau mau, di gc/luar gc siapa yang akan kamu jadikan sahabat?(boleh beda/sma jenis)','apa ketakutan terbesar kamu?','pernah suka sama orang dan merasa orang itu suka sama kamu juga?','Siapa nama mantan pacar teman mu yang pernah kamu sukai diam diam?','pernah gak nyuri uang nyokap atau bokap? Alesanya?','hal yang bikin seneng pas lu lagi sedih apa','pernah cinta bertepuk sebelah tangan? kalo pernah sama siapa? rasanya gimana brou?','pernah jadi selingkuhan orang?','hal yang paling ditakutin','siapa orang yang paling berpengaruh kepada kehidupanmu','hal membanggakan apa yang kamu dapatkan di tahun ini','siapa orang yang bisa membuatmu sange','siapa orang yang pernah buatmu sange','(bgi yg muslim) pernah ga solat seharian?','Siapa yang paling mendekati tipe pasangan idealmu di sini','suka mabar(main bareng)sama siapa?','pernah nolak orang? alasannya kenapa?','Sebutkan kejadian yang bikin kamu sakit hati yang masih di inget','pencapaian yang udah didapet apa aja ditahun ini?','kebiasaan terburuk lo pas di sekolah apa?']
					const ttrth = trut[Math.floor(Math.random() * trut.length)]
					truteh = await getBuffer(`https://i.ibb.co/305yt26/bf84f20635dedd5dde31e7e5b6983ae9.jpg`)
					client.sendMessage(from, truteh, image, { caption: '*Truth*\n\n'+ ttrth, quoted: mek })
					await limitAdd(sender)
					break
                case prefix+'dare':
                	
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))                
					const dare =['Kirim pesan ke mantan kamu dan bilang "aku masih suka sama kamu','telfon crush/pacar sekarang dan ss ke pemain','pap ke salah satu anggota grup','Bilang "KAMU CANTIK BANGET NGGAK BOHONG" ke cowo','ss recent call whatsapp','drop emot 🤥 setiap ngetik di gc/pc selama 1 hari','kirim voice note bilang can i call u baby?','drop kutipan lagu/quote, terus tag member yang cocok buat kutipan itu','pake foto sule sampe 3 hari','ketik pake bahasa daerah 24 jam','ganti nama menjadi "gue anak lucinta luna" selama 5 jam','chat ke kontak wa urutan sesuai %batre kamu, terus bilang ke dia "i lucky to hv you','prank chat mantan dan bilang " i love u, pgn balikan','record voice baca surah al-kautsar','bilang "i hv crush on you, mau jadi pacarku gak?" ke lawan jenis yang terakhir bgt kamu chat (serah di wa/tele), tunggu dia bales, kalo udah ss drop ke sini','sebutkan tipe pacar mu!','snap/post foto pacar/crush','teriak gajelas lalu kirim pake vn kesini','pap mukamu lalu kirim ke salah satu temanmu','kirim fotomu dengan caption, aku anak pungut','teriak pake kata kasar sambil vn trus kirim kesini','teriak " anjimm gabutt anjimmm " di depan rumah mu','ganti nama jadi " BOWO " selama 24 jam','Pura pura kerasukan, contoh : kerasukan maung, kerasukan belalang, kerasukan kulkas, dll']
					const der = dare[Math.floor(Math.random() * dare.length)]
					tod = await getBuffer(`https://i.ibb.co/305yt26/bf84f20635dedd5dde31e7e5b6983ae9.jpg`)
					client.sendMessage(from, tod, image, { quoted: mek, caption: '*Dare*\n\n'+ der })
					await limitAdd(sender)
					break
				case prefix+'dare2':
                	
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					dare2 = await fetchJson(`https://xptnbotapinew.herokuapp.com/?dare&apikey=xptn`, {method: 'get'})
					reply(anu.desc)
					await limitAdd(sender)
					break
                  case prefix+'timer':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))          
				if (args.length < 1) return reply(`#timer 5 detik\nitu contoh nya`)      
				if (args[1]=="detik") {var timer = args[0]+"000"
				} else if (args[1]=="menit") {var timer = args[0]+"0000"
				} else if (args[1]=="jam") {var timer = args[0]+"00000"
				} else {return reply("*pilih:*\ndetik\nmenit\njam")}
				setTimeout( () => {
				reply("Waktu habis")
				}, timer)
				await limitAdd(sender)
				break						   
/*
]=====> MENU GRUP <=====[
*/		 
				case `${prefix}welcome`:
				case `${prefix}notifgrup`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isGroup) return reply(ind.groupo())
					if (args.length < 1) return reply('Boo :')
					if (args[0] === 'on') {
						if (isWelkom) return reply('*SUDAH AKTIF* !!!')
						welkom.push(from)
						fs.writeFileSync('./database/bot/welkom.json', JSON.stringify(welkom))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'off') {
						welkom.splice(from, 1)
						fs.writeFileSync('./database/bot/welkom.json', JSON.stringify(welkom))
						reply('BERHASIL MEMATIKAN')
					} else if (args[0] === 'enable') {
						if (isWelkom) return reply('*SUDAH AKTIF* !!!')
						welkom.push(from)
						fs.writeFileSync('./database/bot/welkom.json', JSON.stringify(welkom))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'disable') {
						welkom.splice(from, 1)
						fs.writeFileSync('./database/bot/welkom.json', JSON.stringify(welkom))
						reply('BERHASIL MEMATIKAN')
					} else {
						reply(ind.satukos())
					}
					await limitAdd(sender)
					break
                 case `${prefix}event`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isGroup) return reply(ind.groupo())
					if (args.length < 1) return reply('Boo :')
					if (args[0] === 'on') {
						if (isEventon) return reply('*SUDAH AKTIF* !!!')
						event.push(from)
						fs.writeFileSync('./database/bot/event.json', JSON.stringify(event))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'off') {
						event.splice(from, 1)
						fs.writeFileSync('./database/bot/event.json', JSON.stringify(event))
						reply('BERHASIL MEMATIKAN')
					} else if (args[0] === 'enable') {
						if (isEventon) return reply('*SUDAH AKTIF* !!!')
						event.push(from)
						fs.writeFileSync('./database/bot/event.json', JSON.stringify(event))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'disable') {
						event.splice(from, 1)
						fs.writeFileSync('./database/bot/event.json', JSON.stringify(event))
						reply('BERHASIL MEMATIKAN')
					} else {
						reply(ind.satukos())
					}
					await limitAdd(sender)
					break
                case prefix+'leveling':
                if (!isGroup) return reply(ind.groupo())
                if (args.length < 1) return reply('Ekhemm >_<')
                if (args[0] === 'enable') {
                    if (isLevelingOn) return reply('*fitur level sudah aktif sebelum nya*')
                    _leveling.push(from)
                    fs.writeFileSync('./database/group/leveling.json', JSON.stringify(_leveling))
                     reply(ind.lvlon())
                } else if (args[0] === 'disable') {
                	if (!isLevelingOn) return reply('*fitur level sudah mati sebelum nya*')
                    _leveling.splice(from, 1)
                    fs.writeFileSync('./database/group/leveling.json', JSON.stringify(_leveling))
                     reply(ind.lvloff())
                } else if (args[0] === 'on') {
                    if (isLevelingOn) return reply('*fitur level sudah aktif sebelum nya*')
                    _leveling.push(from)
                    fs.writeFileSync('./database/group/leveling.json', JSON.stringify(_leveling))
                     reply(ind.lvlon())
                } else if (args[0] === 'off') {
                	if (!isLevelingOn) return reply('*fitur level sudah mati sebelum nya*')
                    _leveling.splice(from, 1)
                    fs.writeFileSync('./database/group/leveling.json', JSON.stringify(_leveling))
                     reply(ind.lvloff())
				} else {
                    reply(ind.satukos())
                }
					break
				case `${prefix}simi`:
				case `${prefix}simih`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isGroup) return reply(ind.groupo())
					if (args.length < 1) return reply('Boo :')
					if (args[0] === 'on') {
						if (isSimi) return reply('SUDAH ON')
						samih.push(from)
						fs.writeFileSync('./database/bot/simi.json', JSON.stringify(samih))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'off') {
						if (!isSimi) return reply('SUDAH OFF')
						samih.splice(from, 1)
						fs.writeFileSync('./database/bot/simi.json', JSON.stringify(samih))
						reply('BERHASIL MEMATIKAN')
					} else if (args[0] === 'enable') {
						if (isSimi) return reply('SUDAH ON')
						samih.push(from)
						fs.writeFileSync('./database/bot/simi.json', JSON.stringify(samih))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'disable') {
						if (!isSimi) return reply('SUDAH OFF')
						samih.splice(from, 1)
						fs.writeFileSync('./database/bot/simi.json', JSON.stringify(samih))
						reply('BERHASIL MEMATIKAN')
					} else {
						reply(ind.satukos())
					}
					await limitAdd(sender)
					break
				case `${prefix}nsfw`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isGroup) return reply(ind.groupo())
					if (args.length < 1) return reply('Boo :')
					if (args[0] === 'on') {
						if (isNsfw) return reply(' *sudah aktif*  !!')
						nsfw.push(from)
						fs.writeFileSync('./database/bot/nsfw.json', JSON.stringify(nsfw))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'off') {
						nsfw.splice(from, 1)
						fs.writeFileSync('./database/bot/nsfw.json', JSON.stringify(nsfw))
						reply('BERHASIL MEMATIKAN')
					} else if (args[0] === 'enable') {
						if (isNsfw) return reply(' *sudah aktif*  !!')
						nsfw.push(from)
						fs.writeFileSync('./database/bot/nsfw.json', JSON.stringify(nsfw))
						reply('BERHASIL MENYALAKAN')
					} else if (args[0] === 'disable') {
						nsfw.splice(from, 1)
						fs.writeFileSync('./database/bot/nsfw.json', JSON.stringify(nsfw))
						reply('BERHASIL MEMATIKAN')
					} else {
						reply(ind.satukos())
					}
					await limitAdd(sender)
					break
				case prefix+'admin':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())
					teks = `*DAFTAR ATASAN GROUP* _${groupMetadata.subject}_\n*TOTAL* : ${groupAdmins.length}\n\n`
					no = 0
					for (let admon of groupAdmins) {
						no += 1
						teks += `[${no.toString()}] @${admon.split('@')[0]}\n`
					}
					mentions(teks, groupAdmins, true)
					break
				case prefix+'grup':
				case prefix+'group':
										
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (args[0] === 'buka') {
					    reply(`*BERHASIL MEMBUKA GROUP*`)
						client.groupSettingChange(from, GroupSettingChange.messageSend, false)
					} else if (args[0] === 'tutup') {
						reply(`*BERHASIL MENUTUP GROUP*`)
						client.groupSettingChange(from, GroupSettingChange.messageSend, true)
					}
				case prefix+'open':
				case prefix+'buka':
										
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					reply(`*BERHASIL MEMBUKA GROUP*`)
					client.groupSettingChange(from, GroupSettingChange.messageSend, false)
					break
				case prefix+'close':
				case prefix+'tutup':
										
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					reply(`*BERHASIL MENUTUP GROUP*`)
					client.groupSettingChange(from, GroupSettingChange.messageSend, true)
					break
				case prefix+'add':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (args.length < 1) return reply(`Salah\n\nGini contoh nya\n${prefix}add ${me.jid.split("@")[0]}`)
					hapusspasi = body.replace(' ', '')
					hapusmin = hapusspasi.replace('-', '')
					hapustambah = hapusmin.replace('+', '')
					hasiladd = hapustambah.replace('add', 'add ')
					reply(hasiladd)
					try {
						num = `${hapustambah.split('add')[1]}@s.whatsapp.net`
						client.groupAdd(from, [num])
						reply(`Menambahkan ${hapustambah.split('add')[1]}`)
					} catch (e) {
						console.log('Error :', e)
						reply('Anjim yang mau di add di private, dahlah :)')
					}
					break
                case prefix+'hidetag':
					                
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isGroup) return reply(ind.groupo())
					var value = body.slice(9)
					var group = await client.groupMetadata(from)
					var member = group['participants']
					var mem = []
					member.map( async adm => {
					mem.push(adm.id.replace('c.us', 's.whatsapp.net'))
					})
					options = {
					text: value,
					contextInfo: { mentionedJid: mem },
					quoted: mek
					}
					client.sendMessage(from, options, text)
					await limitAdd(sender)
					break
			case prefix+'me':
			case prefix+'profil':
			case prefix+'profile':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				levele = getLevelingLevel(sender)
                xpne = getLevelingXp(sender)
                if (levele === undefined && xpne === undefined) return reply(`*Aktifkan level, ketik*\n${prefix}leveling on`)
                butuhxp = 5000 * (Math.pow(2, levele) - 1)
				try {
                	ppneme = await client.getProfilePicture(`${sender.split('@')[0]}@c.us`)
                } catch {
					ppneme = fs.readFileSync(`./fauzan.rifki.m/ppkosong.png`)
                }
					menya = `*TAG:* @${sender.split("@")[0]}
*Nama:* ${pushname}
*Nama terdaftar:* ${namaneuser(sender)}
*Nomor:* ${sender.split("@")[0]}
*Limit:* ${ceklimit(sender)} pesan
*Saldo:* Rp ${checkATMuser(sender)}
*Level:* ${levele}
*XP:* ${xpne}/${butuhxp}`
					client.sendMessage(from, ppneme, image, {quoted:mek, caption: `${menya}`, contextInfo: { mentionedJid: [sender] }})
	kontakme = 'BEGIN:VCARD\n' 
            + `VERSION:3.0\n` 
            + `FN:${args.join(' ')}\n` 
            + `ORG: minta di save;\n` 
            + `TEL;type=CELL;type=VOICE;waid=${sender.split("@")[0]}:+${sender.split("@")[0]}\n` 
            + `END:VCARD` 
            			client.sendMessage(from, {displayname: "Jeff", vcard: kontakme}, MessageType.contact, { quoted: mek })
				break
			case prefix+'level':
                                
                if (!isRegistered) return reply(ind.noregis())
                const userLevel = getLevelingLevel(sender)
                const userXp = getLevelingXp(sender)
                if (userLevel === undefined && userXp === undefined) return reply(ind.lvlnul())
                const requiredXp = 5000 * (Math.pow(2, userLevel) - 1)
                resul = `*♡ LEVEL ♡*\n➸ NAMA: ${namaneuser(sender)}\n➸ NAMA AKUN: ${pushname}\n➸ NOMOR: wa.me/${sender.split("@")[0]}\n➸ XP: ${userXp}/${requiredXp}\n➸ LEVEL: ${userLevel}`
                reply(resul)
                try {
                	pepe = await client.getProfilePicture(`${sender.split('@')[0]}@c.us`)
                } catch {
					auuu = 'https://i0.wp.com/www.gambarunik.id/wp-content/uploads/2019/06/Top-Gambar-Foto-Profil-Kosong-Lucu-Tergokil-.jpg'
					pepe = fs.readFileSync(`./fauzan.rifki.m/ppkosong.png`)
                }
                const rank = new canvas.Rank()
                    .setAvatar(pepe)
                    .setLevel(userLevel)
                    .setLevelColor('#000000', '#000000')
                    .setRank(Number('0'))
                    .setCurrentXP(userXp, '#000000')
                    .setOverlay('#000000', 100, false)
                    .setRequiredXP(requiredXp, '#000000')
                    .setProgressBar('#000000', 'COLOR')
                    .setBackground('COLOR', '#ffffff')
                    .setUsername(namaneuser(sender), '#000000')
                    .setDiscriminator(sender.substring(6, 10))
                rank.build()
                    .then(async (buffer) => {
                        canvas.write(buffer, `${sender}_card.png`)
                        levele = fs.readFileSync(`./${sender}_card.png`)
                        client.sendMessage(from, levele, image, {caption: `Silahkan Follow instagramku dengan cara klik link ini\nhttps://instagram.com/frm_developer`, quoted:mek})
                        fs.unlinkSync(`${sender}_card.png`)
                    })
                    .catch(async (err) => {
                        console.error(err)
                        reply('Error!')
                    })
            break
                 case prefix+'linkgrup':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))                
				    if (!isGroup) return reply(ind.groupo())
				    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				    if (!isBotGroupAdmins) return reply(ind.badmin())
				    linkgc = await client.groupInviteCode (from)
				    yeh = `https://chat.whatsapp.com/${linkgc}\n\nlink Group *${groupName}*`
				    reply(yeh)
			        await limitAdd(sender)
					break
				case prefix+'tagall':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))		
					if (!isGroup) return reply(ind.groupo())
					members_id = []
					teks = (args.length > 1) ? body.slice(8).trim() : ''
					teks += '\n\n'
					for (let mem of groupMembers) {
						teks += `➸ @${mem.jid.split('@')[0]}\n`
						members_id.push(mem.jid)
					}
					client.updatePresence(from, Presence.composing)
					mentions(teks, members_id, true)
					break
           case prefix+'setname':
                if (!isRegistered) return reply(ind.noregis())           
                if (!isGroup) return reply(ind.groupo())
				if (!isBotGroupAdmins) return reply(ind.badmin())
                client.groupUpdateSubject(from, `${body.slice(9)}`)
                reply('⟪ SUKSES ⟫ Mengubah Nama Grup')
					break
                case prefix+'setdesc':
                if (!isRegistered) return reply(ind.noregis())                
                if (!isGroup) return reply(ind.groupo())
				if (!isBotGroupAdmins) return reply(ind.badmin())
                client.groupUpdateDescription(from, `${body.slice(9)}`)
                reply('⟪ SUKSES ⟫ Mengubah Desk Grup')
					break
           case prefix+'demote':
                if (!isRegistered) return reply(ind.noregis())           
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return reply(`Minta contoh ?\n${command} @62xxxx`)
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (mentioned.length > 1) {
						teks = ''
						for (let _ of mentioned) {
							teks += `*jabatan kamu di copot*🏃 :\n`
							teks += `@_.split('@')[0]`
						}
						mentions(teks, mentioned, true)
						client.groupDemoteAdmin(from, mentioned)
					} else {
						mentions(`Yahh @${mentioned[0].split('@')[0]} Jabatan kamu sebagai leluhur di grup telah di copot🏃`, mentioned, true)
						client.groupDemoteAdmin(from, mentioned)
					}
					break
				case prefix+'promote':
                if (!isRegistered) return reply(ind.noregis())				
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return reply(`Minta contoh ?\n${command} @62xxxx`)
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (mentioned.length > 1) {
						teks = ''
						for (let _ of mentioned) {
							teks += `Yeee🥳 Kamu naik jabatan >_< :\n`
							teks += `@_.split('@')[0]`
						}
						mentions(teks, mentioned, true)
						client.groupMakeAdmin(from, mentioned)
					} else {
						mentions(`Selamat?? @${mentioned[0].split('@')[0]} *anda naik menjadi admin group* >_<`, mentioned, true)
						client.groupMakeAdmin(from, mentioned)
					}
					break
				case prefix+'hedsot':
                if (!isRegistered) return reply(ind.noregis())				
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return reply(`Minta contoh ?\n${command} @62xxxx`)
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (mentioned.length > 1) {
						teks = 'Bismillah Hedsot >_< :\n'
						for (let _ of mentioned) {
							teks += `@${_.split('@')[0]}\n`
						}
						mentions(teks, mentioned, true)
						client.groupRemove(from, mentioned)
						mentions(teks, mentioned, true)
						client.groupAdd(from, [num])
					} else {
						mentions(`Berhasil Meng hedsot pala nya  : @${mentioned[0].split('@')[0]}`, mentioned, true)
						client.groupRemove(from, mentioned)
						}
					break
					case prefix+'hidetag5':
                if (!isRegistered) return reply(ind.noregis())				
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					var value = body.slice(9)
					var group = await client.groupMetadata(from)
					var member = group['participants']
					var mem = []
					member.map( async adm => {
					mem.push(adm.id.replace('c.us', 's.whatsapp.net'))
					})
					options = {
					text: value,
					contextInfo: { mentionedJid: mem },
					quoted: mek
					}
					client.sendMessage(from, options, text)
	                .then(() => {client.sendMessage(from, options, text)})
	                .then(() => {client.sendMessage(from, options, text)})
	                .then(() => {client.sendMessage(from, options, text)})
	                .then(() => {client.sendMessage(from, options, text)})
					break
                 case prefix+'fitnah':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())                 
				if (args.length < 1) return reply(`Gini kak : ${prefix}fitnah @tag&pesan&balasanbot\n\nContoh : ${prefix}fitnah @628xxx&hai&hai juga`)
				mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
				var gh = body.slice(8)
					var replace = gh.split("&")[0];
					var target = gh.split("&")[1];
					var bot = gh.split("&")[2];
					tuduh(`${mentioned}`, `${target}`, `${bot}`)
					break
				case prefix+'tuduh':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())                 
					if (args.length < 1) return reply(`Gini kak : ${prefix}tuduh 628xxx&pesan&balasanbot\n\nContoh : ${prefix}tuduh ${me.jid.split('@')[0]}&hai&hai juga`)
					var hi = body.slice(7)
					var replace1 = hi.split("&")[0];
					var target1 = hi.split("&")[1];
					var bot1 = hi.split("&")[2];
					tagpalsu(`${replace1}`, `${target1}`, `${bot1}`)
					break
				case prefix+'pengumuman':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())                 
					if (args.length < 1) return reply(`#pengumuman jangan rame ya\n\nitu contoh nya`)
					var value = q
					var group = await client.groupMetadata(from)
					var member = group['participants']
					var mem = []
					member.map( async adm => {
					mem.push(adm.id.replace('c.us', 's.whatsapp.net'))
					})
					var hidetage = {
					text: value,
					contextInfo: { mentionedJid: mem },
					quoted: mek
					}
					tuduh(`${nomerewa}`, `📢 _pengumuman_ 📢`, hidetage)
					await limitAdd(sender)
					break
					break
				case prefix+'peringatan':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())                 
					if (args.length < 1) return reply(`#peringatan jangan spam, nanti akan saya kick\n\nitu contoh nya`)
					tuduh(`${nomerewa}`, `⚠️ _peringatan_ ⚠️`, `🔪 *=>* ${q}`)
					break
				case prefix+'tuduh':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())                 
					if (args.length < 1) return reply(`Gini kak : ${prefix}fitnah 62xxx&pesan&balasanbot\n\nContoh : ${prefix}fitnah 628xxx&hai&hai juga`)
					var ftn = body.slice(7)
					var sengditag = ftn.split("&")[0];
					var pesane = ftn.split("&")[1];
					var pesanku = ftn.split("&")[2];
					tuduh(`${sengditag}@s.whatsapp.net`, `${pesane}`, `${pesanku}`)
					break
/*
]=====> DOWNLOAD MENU <=====[
*/
				case prefix+'yutubdl':
					if (args.length < 1) return reply('Urlnya mana um?')
                  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if(!isUrl(args[0]) && !args[0].includes('youtu')) return reply('URL NYA TIDAK VALID KAK')				
		yutubdl = await fetchJson(`https://api.vhtear.com/ytdl?link=${args[0]}&apikey=${VhtearKey}`, {method: 'get'})
					if (yutubdl.error) return reply(yutubdl.error)
					teks = `*➸ JUDUL* : ${yutubdl.result.title}\n\n*[WAIT] Proses Dumlu Yakan*`
					thumb = await getBuffer(yutubdl.result.imgUrl)
					client.sendMessage(from, thumb, image, {quoted: mek, caption: teks})
					buffer = await getBuffer(yutubdl.result.UrlVideo)
					client.sendMessage(from, buffer, video, {mimetype: 'video/mp4', quoted: mek})
                    await limitAdd(sender)
					break
				case prefix+'tiktod':
					if (args.length < 1) return reply('Urlnya mana um?')
                  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					tiktod = await fetchJson(`https://api.zeks.xyz/api/tiktok?url=${args[0]}&apikey=${ZeksKey}`,)
					reply('[WAIT] Proses Dumlu Yakan')
					rmln = await getBuffer(tiktod.result.result.server_1)
					client.sendMessage(from, rmln, video, {mimetype: 'video/mp4', caption: `*By.* ${tiktod.result.username}`, quoted: mek})
					await limitAdd(sender)
					break
                     case prefix+'play2':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal)) 
                reply(ind.wait())
                play2 = await fetchJson(`https://api.vhtear.com/ytmp3?query=${body.slice(6)}&apikey=${VhtearKey}`)
               if (play2.error) return reply(play2.error)
                 infomp3 = `*「❗」Lagu Ditemukan*\n➸ Judul : ${play2.result.title}\n➸ Durasi : ${play2.result.duration}\n➸ Size : ${play2.result.size}\n\n*[WAIT] Proses Dumlu Yakan*`
                buffer = await getBuffer(play2.result.image)
                lagu = await getBuffer(play2.result.mp3)
                client.sendMessage(from, buffer, image, {quoted: mek, caption: infomp3})
                client.updatePresence(from, Presence.recording)
                client.sendMessage(from, lagu, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
                break				
/*
]=====> LIMIT MENU <=====[
*/
				case prefix+'limit':
				                  
				   if (!isRegistered) return reply(ind.noregis())
				   checkLimit(sender)
					break
				case prefix+'atm':
                  				
				if (!isRegistered) return reply(ind.noregis())
				const kantong = checkATMuser(sender)
				reply(ind.uangkau(pushname, sender, kantong))
				break
				case prefix+'buylimit':
				case `${prefix}belikuota limit`:
				case `${prefix}buy`:
				case `${prefix}buylimit`:
                if (!isRegistered) return reply(ind.noregis())
				if (args.length < 1) return reply(`* SELAMAT DATANG *\nSelamat datang\n\n\n*Harga*\nRp 1000 = 1 pesan\n\n*Cara Beli:*\n${prefix}buylimit ~jumlah limitnya~\n*Contoh:*\n${prefix}buylimit 5`)
				payout = body.slice(10)
				koinPerlimit = 1000
				total = koinPerlimit * payout
				if ( checkATMuser(sender) <= total) return reply(`Maaf, saldo anda tidak mencukupi untuk melakukan pembelian ini. Silahkan ketik *${prefix}kerja* untuk mengisi saldo anda\nTerima kasih`)
				if ( checkATMuser(sender) >= total ) {
					confirmATM(sender, total)
					bayarLimit(sender, payout)
					await reply(`*⟪ PEMBAYARAN BERHASIL ⟫*\n\n➸ pengirim : FRM BOT\n➸ penerima : ${namaneuser(sender)}\n➸ nominal pembelian : ${payout} pesan \n➸ harga kuota limit : ${koinPerlimit}/pesan\n➸ sisa saldo : Rp ${checkATMuser(sender)}\n\nproses berhasil dengan SN\n${monosp}${createSerial(15)}${monosp}`)
				} 
				break
/*
]=====> RANDOM MENU <=====[
*/
                case prefix+'pokemon':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                   anu = await fetchJson(`https://api.fdci.se/rep.php?gambar=pokemon`, {method: 'get'})
					reply(ind.wait())
					var n = JSON.parse(JSON.stringify(anu));
					var nimek =  n[Math.floor(Math.random() * n.length)];
					pok = await getBuffer(nimek)
					client.sendMessage(from, pok, image, { quoted: mek })
					await limitAdd(sender)
					break
                case prefix+'anjing':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                   anu = await fetchJson(`https://api.fdci.se/rep.php?gambar=anjing`, {method: 'get'})
					reply(ind.wait())
					var n = JSON.parse(JSON.stringify(anu));
					var nimek =  n[Math.floor(Math.random() * n.length)];
					pok = await getBuffer(nimek)
					client.sendMessage(from, pok, image, { quoted: mek })
					await limitAdd(sender)
					break
                case prefix+'blowjob':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					blowjob = await fetchJson('https://waifu.pics/api/nsfw/blowjob', {method: 'get'})
					if (blowjob.error) return reply(blowjob.error)
					exec(`wget ${blowjob.url} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
			case prefix+'nekonime':
			case prefix+'neko':
                if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						neko = await fetchJson(`https://waifu.pics/api/nsfw/neko`, {method: 'get'})
						neko1 = await getBuffer(neko.url)
						client.sendMessage(from, neko1, image, {quoted: mek, caption: 'Nih nekonime mu >_<'})
					await limitAdd(sender)
					break
                case prefix+'kpop':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                                        reply(ind.wait())
                                        kpop = await fetchJson(`https://tobz-api.herokuapp.com/api/randomkpop?apikey=${TobzKey}`, {method: 'get'})
                                        if (kpop.error) return reply(kpop.error)
                                        buffer = await getBuffer(kpop.result)
                                        client.sendMessage(from, buffer, image, {quoted: mek, caption: tanda})
                                        await limitAdd(sender)
                                        break
                case prefix+'husbu':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                   if (!isGroup) return reply(ind.groupo())
						res = await fetchJson(`https://tobz-api.herokuapp.com/api/husbu?apikey=${TobzKey}`)
						buffer = await getBuffer(res.image)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: '>_<'})
					await limitAdd(sender)
					break
			case prefix+'chord':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                anu = await fetchJson(`https://tobz-api.herokuapp.com/api/chord?q=${body.slice(7)}&apikey=${TobzKey}`)
                reply(anu.result)
                await limitAdd(sender)
                break
			case prefix+'loli':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(ind.wait())
					loline = await fetchJson(`https://api.vhtear.com/randomloli&apikey=${VhtearKey}`, {method: 'get'})
					buffer = await getBuffer(loline.result.result)
					client.sendMessage(from, buffer, image, {quoted: mek})
					await limitAdd(sender)
					break					
                case prefix+'randomhentong':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					gatauda = body.slice(15)
					reply(ind.wait())
					randomhentong = await fetchJson(`https://api.vhtear.com/randomhentai?apikey=${VhtearKey}`, {method: 'get'})
					buffer = await getBuffer(randomhentong.result.url)
					client.sendMessage(from, buffer, image, {quoted: mek, caption: `nih hentong mu`})
					await limitAdd(sender)
					break					
					case prefix+'wibu':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						data = await fetchJson(`https://api.vhtear.com/randomwibu&apikey=${VhtearKey}`)
						buffer = await getBuffer(data.result.foto)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: '>_<'})
					await limitAdd(sender)
					break
                case prefix+'darkjokes':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))   
				 data = fs.readFileSync('./src/darkjokes.js');
                 jsonData = JSON.parse(data);
                 randIndex = Math.floor(Math.random() * jsonData.length);
                 randKey = jsonData[randIndex];
                hasil = await getBuffer(randKey.result)
                sendImage(hasil, mek, '*GELAP BOS :V*')
				break										
/*
]=====> OTHER MENU <=====[
*/				
            	case prefix+'mining':
                      if (!isRegistered) return reply(ind.noregis())
                      if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                      if (!isEventon) return reply(`maaf ${pushname} event mining tidak di aktifkan sama owner Rifki`)
                      if (isOwner) {
                      const one = 999999999
                      addLevelingXp(sender, one)
                      addLevelingLevel(sender, 99)
                      reply(`karena Rifki baik Bot memberikan ${one}Xp >_<`)
                      }else{
                      const mining = Math.ceil(Math.random() * 10000)
                      addLevelingXp(sender, mining)
                      await reply(`*selamat* ${pushname} kamu mendapatkan *${mining}Xp*\n\nmining hanya digunakan untuk menambah XP`)
                      }
                    await limitAdd(sender)
					break
                case prefix+'moddroid':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			data = await fetchJson(`https://tobz-api.herokuapp.com/api/moddroid?q=${body.slice(10)}&apikey=${TobzKey}`)
			hepi = data.result[0] 
			teks = `*➸ Nama*: ${data.result[0].title}\n*➸ publisher*: ${hepi.publisher}\n*➸ mod info:* ${hepi.mod_info}\n*➸ size*: ${hepi.size}\n*➸ latest version*: ${hepi.latest_version}\n*➸ genre*: ${hepi.genre}\n*➸ link:* ${hepi.link}\n*➸ download*: ${hepi.download}`
			buffer = await getBuffer(hepi.image)
			client.sendMessage(from, buffer, image, {quoted: mek, caption: `${teks}`})
			await limitAdd(sender)
			break
		case prefix+'lirik':
			if (!isRegistered) return reply(ind.noregis())
			if (isBanned) return reply(ind.diban())
			if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			anu = await fetchJson(`https://tobz-api.herokuapp.com/api/lirik?q=${body.slice(7)}&apikey=${TobzKey}`)
			thum = await getBuffer(anu.result.thumb)
			teks = `*「 LAGU DI TEMUKAN 」*\n\n*Judul* : ${anu.result.judul}\n*Album* : ${anu.result.album}\n*public in* : ${anu.result.dipublikasi}\n*Lyrics* : ${anu.result.lirik}`
			client.sendMessage(from, thum, image, { quoted : mek, caption: teks })
			await limitAdd(sender)
			break
	case prefix+'minecraft.terbaru':
		request.get({
		headers: {'content-type' : 'application/x-www-form-urlencoded'},
        url:     'https://feedback.minecraft.net/hc/en-us/sections/360001186971-Release-Changelogs',
      },function(error, response, body){
          let $ = cheerio.load(body);
          var y = $.html().split('class="article-list-link">')[1];
          var t = y.split('</a>')[0];
      reply(t)
      })
      break
case prefix+'nama':
case prefix+'artinama':
	
	if (!isRegistered) return reply(ind.noregis())
	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
    request.get({
        headers: {'content-type' : 'application/x-www-form-urlencoded'},
        url:     `https://www.primbon.com/arti_nama.php?nama1=${q}&proses=+Submit%21+`,
      },function(error, response, body){
          let $ = cheerio.load(body);
          var y = $.html().split('arti:')[1];
          var t = y.split('method="get">')[1];
          var f = y.replace(t ," ");
          var x = f.replace(/<br\s*[\/]?>/gi, "\n");
          var h  = x.replace(/<[^>]*>?/gm, '');
      console.log(""+ h);
      reply(h)
      })
      await limitAdd(sender) 
			break
			case prefix+'happymod':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			data = await fetchJson(`https://tobz-api.herokuapp.com/api/happymod?q=${body.slice(10)}&apikey=${TobzKey}`)
			hupo = data.result[0] 
			teks = `*➸ Nama*: ${data.result[0].title}\n*➸ version*: ${hupo.version}\n*➸ size:* ${hupo.size}\n*➸ root*: ${hupo.root}\n*➸ purchase*: ${hupo.price}\n*➸ link*: ${hupo.link}\n*➸ download*: ${hupo.download}`
			buffer = await getBuffer(hupo.image)
			client.sendMessage(from, buffer, image, {quoted: mek, caption: `${teks}`})
			await limitAdd(sender)
			break
            case prefix+'bitly':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
               client.updatePresence(from, Presence.composing) 
                data = await fetchJson(`https://tobz-api.herokuapp.com/api/bitly?url=${args[0]}&apikey=${TobzKey}`)
                hasil = `link panjang : ${args[0]}\n\nlink pendek : ${data.result}`
                reply(hasil)
                await limitAdd(sender)
                break
					case prefix+'pinterest':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					client.updatePresence(from, Presence.composing) 
					data = await fetchJson(`https://api.vhtear.com/pinterest?query=${q}&apikey=${VhtearKey}`, {method: 'get'})
					reply(ind.wait())
					var pinterest = JSON.parse(JSON.stringify(data.result));
					var hasilpinterest =  pinterest[Math.floor(Math.random() * pinterest.length)];
					gmbrhasilpinterest = await getBuffer(hasilpinterest)
					client.sendMessage(from, gmbrhasilpinterest, image, { quoted: mek, caption: `*⟪ PINTEREST ⟫*`})
					var pinterest2 = JSON.parse(JSON.stringify(data.result));
					var hasilpinterest2 =  pinterest2[Math.floor(Math.random() * pinterest2.length)];
					gmbrhasilpinterest2 = await getBuffer(hasilpinterest2)
					client.sendMessage(from, gmbrhasilpinterest2, image, { quoted: mek, caption: `*⟪ PINTEREST ⟫*`})
					await limitAdd(sender)
					break 
					case prefix+'resepmasakan':
					case prefix+'resep':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                   resep = await fetchJson(`https://mnazria.herokuapp.com/api/resep?key=${q}`, {method: 'get'})
                   if (resep.error) return reply(resep.error)
                   buff = await getBuffer(resep.thumb_item)
                   hasil = `*➸ title* \n ${resep.title} *➸ item_name* \n ${resep.item_name} *➸ ingredient* \n${resep.ingredient} *➸ step* \n${resep.step}`
                   client.sendMessage(from, buff, image, {quoted: mek, caption: hasil})
                   await limitAdd(sender)
					break
                case prefix+'beritahoax':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    client.updatePresence(from, Presence.composing) 
					data = await fetchJson(`https://docs-jojo.herokuapp.com/api/infohoax`, {method: 'get'})
					teks = '♡───────────♡\n'
					for (let i of data.result) {
						teks += `*➸ Gambar* : ${i.image}\n*➸ Title* : ${i.title}\n*➸ link* : ${i.link}\n*➸ tag* : ${i.tag}\n♡───────────♡\n`
					}
					reply(teks.trim())
					await limitAdd(sender)
					break
					case prefix+'brainly':
	                  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    brien = body.slice(9)
					brainly(`${brien}`).then(res => {
					teks = '♡───────────♡\n'
					for (let Y of res.data) {
						teks += `\n*「 BRAINLY 」*\n\n*➸ Pertanyaan:* ${Y.pertanyaan}\n\n*➸ Jawaban:* ${Y.jawaban[0].text}\n♡───────────♡\n`
					}
					reply(teks)
                        console.log(res)
                    })
					await limitAdd(sender)
					break
                case prefix+'virtex':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(`*INILAH VIRTEX TERBERAT DI DUNIA*\n\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏tapi boong`)
					lagutapi = fs.readFileSync('./fauzan.rifki.m/tapiboong.m4a')
					client.updatePresence(from, Presence.recording)
					client.sendMessage(from, lagutapi, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
					await limitAdd(sender)
					break
                case prefix+'virtex2':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (!isGroup) return reply(ind.groupo())
					await costum(virtex2(pushname, prefix, botName, ownerName, getLevelingLevel, sender, _registered), text, tescuk, cr)
					await limitAdd(sender)
					break
				case prefix+'mutual':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                if (isGroup) return  reply( 'TIDAK BISA DI GRUP KAK')
                anug = getRegisteredRandomId(_registered).replace('@s.whatsapp.net','')
                await reply('Mencari Pasangan >_<')
                await reply(`wa.me/${anug}`)
                await reply( `Pasangan Ditemukan: 🐊\n*${prefix}next* — Temukan Pasangan Baru`)
                await limitAdd(sender)
            break
            case prefix+'next':  
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                if (isGroup) return  reply( 'TIDAK BISA DI GRUP KAK')
                anug = getRegisteredRandomId(_registered).replace('@s.whatsapp.net','')
                await reply('Mencari Pasangan >_<')
                await reply(`wa.me/${anug}`)
                await reply( `Pasangan Ditemukan: 🐊\n*${prefix}next* — Temukan Pasangan Baru`)
                await limitAdd(sender)
                break
/*
]=====> MAKER MENU <=====[
*/
				case prefix+'raindrop':
if (!isRegistered) return reply(ind.noregis())
if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
  ted = isQuotedImage ? JSON.parse(JSON.stringify(mek).replace('quotedM', 'm')).message.extendedTextMessage.contextInfo: mek
  reply(ind.wait())
  owgi = await client.downloadAndSaveMediaMessage(ted)
  tels = body.slice(7)
  raindrop = await imgbb(imgbbkey, owgi)
  hehre = await getBuffer(`https://videfikri.com/api/textmaker/raindrop/?urlgbr=${raindrop.display_url}`)
 client.sendMessage(from, hehre, image, {quoted:mek})
} else {
  reply(`Tag gambar nya kak`)
}
break
							case prefix+'trigger':
							case prefix+'triggered':
				          	case prefix+'tg':
                                        if (!isRegistered) return reply(ind.noregis())
										if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                                         if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
                                         ger = isQuotedImage ? JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo : mek
                                         owgi = await client.downloadAndSaveMediaMessage(ger)
                                         anu = = await imgbb(imgbbkey, owgi)
                                        teks = `${anu.display_url}`
                                        ranp = getRandom('.gif')
                                        rano = getRandom('.webp')
                                        anu1 = `https://some-random-api.ml/canvas/triggered?avatar=${teks}`
                                         exec(`wget ${anu1} -O ${ranp} && ffmpeg -i ${ranp} ${rano}`, (error, stdout, stderr) => {
                                                nobg = fs.readFileSync(rano)
                                                 client.sendMessage(from, nobg, sticker, {quoted: mek})
                                                fs.unlinkSync(rano)
                                        })
                                    
                                             } else {
                                                 reply('Gunakan foto!')
                                          }
                                             break
				case prefix+'imgtourl':
				case prefix+'tourl':
				case prefix+'tolink':
					if (!isRegistered) return reply(ind.noregis())
if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
  ted = isQuotedImage ? JSON.parse(JSON.stringify(mek).replace('quotedM', 'm')).message.extendedTextMessage.contextInfo: mek
  owgi = await client.downloadAndSaveMediaMessage(ted)
  tels = body.slice(7)
  anu = await imgbb(imgbbkey, owgi)
  gambare = await getBuffer(anu.display_url)
  client.sendMessage(from, gambare, image, {quoted: mek, caption: `${anu.data.display_url}`})
} else {
  reply('tag gambar/foto')
}
break
				case prefix+'lovemake':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Teksnya mana um')
					love = body.slice(10)
					if (love.length > 12) return reply('Teksnya kepanjangan, maksimal 9 karakter')
					reply(ind.wait())
					bufferxcz = await getBuffer(`https://api.vhtear.com/lovemessagetext?text=${love}&apikey=${VhtearKey}`, {method: 'get'})
					client.sendMessage(from, bufferxcz, image, {quoted: mek, caption: ' '+love})
					await limitAdd(sender)
					break
				case prefix+'metalteks':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Teksnya mana um')
					love = body.slice(10)
					if (love.length > 12) return reply('Teksnya kepanjangan, maksimal 9 karakter')
					reply(ind.wait())
					bufferxcz = await getBuffer(`https://api.vhtear.com/metal_maker?text=${love}&apikey=${VhtearKey}`, {method: 'get'})
					client.sendMessage(from, bufferxcz, image, {quoted: mek, caption: ' '+love})
					await limitAdd(sender)
					break
				case prefix+'apiteks':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Teksnya mana um')
					love = body.slice(9)
					if (love.length > 12) return reply('Teksnya kepanjangan, maksimal 9 karakter')
					reply(ind.wait())
					bufferxcz = await getBuffer(`https://api.vhtear.com/fire_maker?text=${love}&apikey=${VhtearKey}`, {method: 'get'})
					client.sendMessage(from, bufferxcz, image, {quoted: mek, caption: ' '+love})
					await limitAdd(sender)
					break
				case prefix+'ffbaner':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Teksnya mana um')
					love = body.slice(9)
					if (love.length > 12) return reply('Teksnya kepanjangan, maksimal 9 karakter')
					reply(ind.wait())
					bufferxcz = await getBuffer(`https://api.vhtear.com/bannerff?title=${love}&text=by. frm&apikey=${VhtearKey}`, {method: 'get'})
					client.sendMessage(from, bufferxcz, image, {quoted: mek, caption: ' '+love})
					await limitAdd(sender)
					break
				case prefix+'ramalhp':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('teks nya mana om')
					ramalhp = await fetchJson(`https://api.vhtear.com/nomerhoki?no=${q}&apikey=${VhtearKey}`, {method: 'get'})
					reply(ramalhp.result.hasil)
					await limitAdd(sender)
					break
				case `${prefix}banner`:
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					tipelist = ['console','block','simpleBlock','simple','3d','simple3d','chrome','huge','shade','slick','grid','pallet','tiny']
					if (args.length < 1) return reply(`${tanda}\n*${prefix}banner ~font&teks~*`)
					if (!tipelist.includes(args[0])) return reply(`${tanda}\nFont salah\n*Fontnya:*\nconsole, block, simpleBlock, simple, 3d, simple3d, chrome, huge, shade, slick, grid, pallet, tiny`)
					gh = body.slice(8)
					gl1 = gh.split(" ")[1];
					anu = cfonts.render((`${gl1}`), {
					font: `${args[0]}`,
					color: 'candy',
					align: 'center',
					lineHeight: 1
					});
					reply(`*MIRINGKAN HP MU*\n${monosp}${anu.string}${monosp}`)
					await limitAdd(sender)
					break
case prefix+'hartatahta':
case prefix+'tahta':

if (!isRegistered) return reply(ind.noregis())
if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
if (args.length < 1) return reply('「❗」Contoh : ${prefix}hartatahta bagi duit')
har = body.slice(12)
reply('「❗」Hirti Tihti Tai Anjg :v')
tahta = await getBuffer(`https://api.vhtear.com/hartatahta?text=${q}&apikey=${VhtearKey}`)
client.sendMessage(from, tahta, image, {quoted: mek})
await limitAdd(sender)
break
case prefix+'cphlogo':
    if (!isRegistered) return reply(ind.noregis())
	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
	gh = `${body.slice(9)}`
	gbl1 = gh.split("&")[0];
	gbl2 = gh.split("&")[1];
	if (args.length < 1) return reply('Teksnya mana gan?')
	buffer = await getBuffer(`https://api.vhtear.com/pornlogo?text1=${gbl1}&text2=${gbl2}&apikey=${VhtearKey}`, {method: 'get'})
	client.sendMessage(from, buffer, image, {quoted: mek})
	await limitAdd(sender) 
	break
case prefix+'cglitch':
	if (!isRegistered) return reply(ind.noregis())
	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
    if (args.length < 1) return reply(`*CONTOH:*\n${prefix}cglitch Fauzan & Rifki`)
    hm = `${body.slice(8)}`
	text1 = hm.split("&")[0];
    text2 = hm.split("&")[1];                    
    glitch = await getBuffer(`https://api.vhtear.com/glitchtext?text1=${text1}&text2=${text2}&apikey=${VhtearKey}`, {method: 'get'})
    client.sendMessage(from, glitch, image, {quoted: mek, caption: 'nih gan'})
	await limitAdd(sender) 
	break 
case prefix+'cml':
				if (!isRegistered) return reply(ind.noregis())
	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply(`${command} Fauzan & Rifki`)
                     if (args.length > 10) return reply('karakter minimal 10')
					cml = `${body.slice(5)}`
					cml1 = cml.split("&")[0];
					cml2 = cml.split("&")[1];
					buffer = await getBuffer(`https://api.vhtear.com/logoml?hero=${cml1}&text=${cml2}&apikey=${VthearKey}`, {method: 'get'})
					client.sendMessage(from, buffer, image, {quoted: mek})
					await limitAdd(sender) 
					break
				case prefix+'cpubg':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`${command} Fauzan & Rifki`)
                     if (args.length > 10) return reply('karakter minimal 10')
					cpubg = `${body.slice(7)}`
					cpubg1 = cpubg.split("&")[0];
					cpubg2 = cpubg.split("&")[1];
					anu = await fetchJson(`http://lolhuman.herokuapp.com/api/photooxy2/pubg?apikey=${LolKey}&text1=${cpubg1}&text2=${cpubg2}`, {method: 'get'})
					cpubg = await getBuffer(anu.result)
					client.sendMessage(from, cpubg, image, {quoted: mek})
					await limitAdd(sender) 
					break
case prefix+'cloudtext':

if (!isRegistered) return reply(ind.noregis())
if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
if (args.length < 1) return reply('「❗」Contoh : ${prefix}cloudtext Rifki')
cloud = body.slice(11)
reply('「❗」Bentar Bro Terbang dumlu yakan')
buffer = await getBuffer(`https://api.xteam.xyz/textpro/cloudtext?text=${cloud}&APIKEY=${XteamKey}`)
client.sendMessage(from, buffer, image, {quoted: mek})
await limitAdd(sender)
break
case prefix+'firework2':
	const browser = await puppeteer.launch({args: ['--no-sandbox']});
      const page = await browser.newPage();
        await page.goto('https://textpro.me/firework-sparkle-text-effect-930.html');
        await page.type('#text-0', text);
        await page.click('[name="submit"]')
        await page.waitForNavigation()
        const bodyHandle = await page.$('body');
        const html = await page.evaluate(body => body.innerHTML, bodyHandle);
        await bodyHandle.dispose();
        const $ = cheerio.load(html)
        const result = 'https://textpro.me' + $('#content-wrapper > section > div > div.col-md-9 > div:nth-child(4) > div > img').attr('src')
        await browser.close();
        client.sendMessage(from, result, text)
	break
	
	
				case prefix+'fast':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.mp3')
					exec(`ffmpeg -i ${media} -filter:a "atempo=1.63asetrate=44100" ${ran}`, (err, stderr, stdout) => {
					fs.unlinkSync(media)
					if (err) return reply('Error!')
					hah = fs.readFileSync(ran)
					client.updatePresence(from, Presence.recording)
					client.sendMessage(from, hah, audio, {mimetype: 'audio/mp4', ptt:true, quoted: mek})
					fs.unlinkSync(ran)
					})
					break 
				case prefix+'slow':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.mp3')
					exec(`ffmpeg -i ${media} -filter:a "atempo=0.7,asetrate=44100" ${ran}`, (err, stderr, stdout) => {
						fs.unlinkSync(media)
						if (err) return reply('Error!')
						hah = fs.readFileSync(ran)
						client.updatePresence(from, Presence.recording)
						client.sendMessage(from, hah, audio, {mimetype: 'audio/mp4', ptt:true, quoted: mek})
						fs.unlinkSync(ran)
					})
				break
				case prefix+'tupai':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.mp3')
					exec(`ffmpeg -i ${media} -filter:a "atempo=0.5,asetrate=65100" ${ran}`, (err, stderr, stdout) => {
						fs.unlinkSync(media)
						if (err) return reply('Error!')
						hah = fs.readFileSync(ran)
						client.updatePresence(from, Presence.recording)
						client.sendMessage(from, hah, audio, {mimetype: 'audio/mp4', ptt:true, quoted: mek})
						fs.unlinkSync(ran)
					})
				break
				case prefix+'gemuk':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.mp3')
					exec(`ffmpeg -i ${media} -filter:a "atempo=1.6,asetrate=22100" ${ran}`, (err, stderr, stdout) => {
						fs.unlinkSync(media)
						if (err) return reply('Error!')
						hah = fs.readFileSync(ran)
						client.updatePresence(from, Presence.recording)
						client.sendMessage(from, hah, audio, {mimetype: 'audio/mp4', ptt:true, quoted: mek})
						fs.unlinkSync(ran)
					})
				break
				case prefix+'bass':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.mp3')
					exec(`ffmpeg -i ${media} -af equalizer=f=94:width_type=o:width=2:g=30 ${ran}`, (err, stderr, stdout) => {
						fs.unlinkSync(media)
						if (err) return reply('Error!')
						hah = fs.readFileSync(ran)
						client.updatePresence(from, Presence.recording)
						client.sendMessage(from, hah, audio, {mimetype: 'audio/mp4', ptt:true, quoted: mek})
						fs.unlinkSync(ran)
					})
					break
		case prefix+'img2url':
			if (!isRegistered) return reply(ind.noregis())
			if (isBanned) return reply(ind.diban())
			if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
			reply(ind.wait())
            encmediasek = isQuotedImage ? JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo : mek
            mediane = await  client.downloadAndSaveMediaMessage(encmediasek)
            linkgmb = await uploadImages(mediane, `${sender.split('@')[0]}_img`)
                     reply(linkgmb)
                }
               
            await limitAdd(sender) 	
            break
            case prefix+'getaudio':
            case prefix+'getvn':
            	if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply('Judulnya mana')
				if (!listaudio.includes(qcilik)) return reply(`Kami tidak memiliki audio yang bernama ${q}`)
				client.updatePresence(from, Presence.recording)
				getaudio = fs.readFileSync(`./audio/${qcilik}.mp3`)
				client.sendMessage(from, getaudio, MessageType.audio, {quoted: mek, mimetype: Mimetype.mp4Audio, ptt:true})
				await limitAdd(sender)
				break
			case prefix+'terus':
			 const messages = await client.loadConversation (`${sender}`, 1)
const message = messages[0]
await conn.forwardMessage (`${from}@s.whatsapp.net`, message)
break
			case prefix+'getstiker':
			case prefix+'getsticker':
            	if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply('Namanya apa ?')
				if (!liststiker.includes(q)) return reply(`Kami tidak memiliki stiker yang bernama ${q}`)
				getstikere = fs.readFileSync(`./sticker/${q}.webp`)
				client.sendMessage(from, getstikere, sticker, {quoted: mek})
            	await limitAdd(sender)
				break
				case prefix+'savesticker':
				case prefix+'savestiker':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			    	if (!isQuotedSticker) return reply('Reply stiker nya')
					svst = q
					if (!svst) return reply('Nama sticker nya apa?')
					if (liststiker.includes(q)) return reply(`sudah ada`)
					boij = JSON.parse(JSON.stringify(mek).replace('quotedM', 'm')).message.extendedTextMessage.contextInfo
					delb = await client.downloadMediaMessage(boij)
						fs.writeFileSync(`./sticker/${svst}.webp`, delb)
            			liststiker.push(q)
            			fs.writeFileSync('./sticker/liststiker.json', JSON.stringify(liststiker))
					reply(`Berhasil menyimpan sticker!`)
					await limitAdd(sender)
					break
				case prefix+'saveaudio':
				case prefix+'savemp3':
				case prefix+'savevn':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isQuotedAudio) return reply('Tag / geser audio nya om')
					gsh = qcilik
					if (!gsh) return reply('Nama file nya apa?')
					if (listaudio.includes(qcilik)) return reply(`sudah ada`)
					uyw = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						gx = await client.downloadMediaMessage(uyw)
						fs.writeFileSync(`./sampah/download.mp3`, gx)
						exec(`ffmpeg -i ./sampah/download.mp3 ./audio/${qcilik}.mp3`, (error, stdout, stderr) => {
  if (error) {
  	fs.writeFileSync(`./audio/${qcilik}.mp3`, gx)
  listaudio.push(qcilik)
  fs.writeFileSync('./audio/listaudio.json', JSON.stringify(listaudio))
    reply(`Menyimpan MP3`);
    return;
  }
  listaudio.push(qcilik)
  fs.writeFileSync('./audio/listaudio.json', JSON.stringify(listaudio))
  reply('Mengonversi File Opus menjadi Mp3')
});
					await limitAdd(sender)
					break
				case 'tesvn':
					uyw = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					gx = await client.downloadMediaMessage(uyw)
					fs.writeFileSync(`./`, gx)
					reply('Berhasil menyimpan audio!')
					await limitAdd(sender)
					break
				case prefix+'quoted':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`ups`)
					client.sendMessage(from, `ini hasilnya`, text, {quoted:args.join(' ')})
					await limitAdd(sender)
					break
				case 'apatu':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ditag = JSON.parse(JSON.stringify(mek).replace('quotedM', 'm')).message.extendedTextMessage.contextInfo
					if (ditag.message.conversation) return reply(ditag.message.conversation)
					
					downloadm = await client.downloadMediaMessage(ditag)
					break
				case prefix+'fordward':
				case prefix+'diteruskan':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`ups`)
	   			client.sendMessage(from, `${q}`, MessageType.text, {contextInfo: { forwardingScore: 600, isForwarded: true }}) 
				break
/*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  
]=====> OWNER MENU <=====[
*/
				case '$':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					exec(`${q}`, (error, stdout, stderr) => {
					if (error) {
    				reply(`*exec error:*\nmiringkan hpmu\n\n${monosp}${error}${monosp}`);
    				return;
					}
					reply(`*stdout:*\nmiringkan hpmu\n\n${monosp}${stdout}${monosp}`);
					reply(`*stderr:*\nmiringkan hpmu\n\n${monosp}${stderr}${monosp}`);
					});
				break
				case 'shutdown':
            		if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
	        		diteruskan(`Bye`)
                	await sleep(5000)
					return reply(JSON.stringify(eval(process.exit())))
					break
				case prefix+'on':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					if (args.length < 1) return reply(`Code chat nya mana`)
					if (onoffnya.includes(q)) return reply('SUDAH MENYALA')
					onoffnya.push(q)
					fs.writeFileSync('./database/bot/onoff.json', JSON.stringify(onoffnya))
					reply('BERHASIL MENYALAKAN')
					break
				case prefix+'off':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					if (args.length < 1) return reply(`Code chat nya mana`)
					if (!onoffnya.includes(q)) return reply('TIDAK ADA DI DALAM DAFTAR BOT ON')
					onoffnya.splice(q, 1)
					fs.writeFileSync('./database/bot/onoff.json', JSON.stringify(onoffnya))
					reply('BERHASIL MEMATIKAN')
					break
				case prefix+'listboton':
					listboton = `*BOT AKTIF DI*\n\n`
					for (let listbotonn of onoffnya) {
						listboton += `*•* ${listbotonn}\n`
					}
						listboton += `\n*SELESAI*`
					reply(listboton.trim())
					reply(`*CONTOH CARA MEMATIKAN*\n${prefix}off 62xx@g.us\n\n*CONTOH CARA MENYALAKAN*\n${prefix}on 62xx@g.us`)
					await limitAdd(sender)
					break
				case prefix+'run':
				case '>':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					try{
                	sy = args.join(' ')
                	return eval(sy)
                	} catch (e) {
                	reply(`${monosp}${e}${monosp}`)
                	}
                	break
				case prefix+'eval':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
                	if (!q) return reply(ind.wrongf())
                	try {
         	           let evaled = await eval(q)
         	           if (typeof evaled !== 'string') evaled = require('util').inspect(evaled)
          	          await reply(evaled)
       	         } catch (err) {
        	            console.error(err)
          	          await reply(`Error!\n\n${err}`)
  	   	       }
        	    break 
				case prefix+'anggotagrup':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					memberlimit = args[0]
					reply(`Sekarang Bot bisa dimasukan ke grup jika anggotanya lebih dari ${memberlimit}`)
					await limitAdd(sender)
					break
				case `${prefix}bc`: 
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					if (args.length < 1) return reply('pesan tidak ada')
						teksbc = []
						teksbc.push(args.join(' '))
						for (let i = 0; i < igdl.result.length; i++) {
                    	dibc = _registered[i].id
						client.sendMessage(dibc, teksbc, text, {quoted: { key: { fromMe: false, participant: `${me.jid}`, ...(from ? { remoteJid: `status@broadcast` } : {}) }, message: { conversation: `BOARDCAST` }}})
						}
					await limitAdd(sender)
					break
				case prefix+'bcgc':
				    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					if (args.length < 1) return reply('Teksnya mana bosku >_<')
					anu = await groupMembers
					nom = mek.participant
					if (isMedia && !mek.message.videoMessage || isQuotedImage) {
						const encmedia = isQuotedImage ? JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo : mek
						buff = await client.downloadMediaMessage(encmedia)
						for (let _ of anu) {
							client.sendMessage(_.jid, buff, image, {caption: `*「 BC GROUP 」*\n\n➸ Dari Grup : ${groupName}\n➸ Pengirim : wa.me/${(sender.split('@')[0])}\n➸ Pesan : ${body.slice(6)}`})
						}
						reply('*「 SUKSES BOSKU 」*')
					} else {
						for (let _ of anu) {
							sendMess(_.jid, `*「 BC GROUP 」*\n\n➸ Dari Grup : ${groupName}\n➸ Pengirim : wa.me/${(sender.split('@')[0])}\n➸ Pesan : ${body.slice(6)}`)
						}
						reply('*「 SUKSES BOSKU 」*')
					}
					await limitAdd(sender)
					break
				case prefix+'clearall':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					anu = await client.chats.all()
					client.setMaxListeners(25)
					for (let _ of anu) {
						client.deleteChat(_.jid)
					}
					reply(ind.clears())
					await limitAdd(sender)
					break
				case prefix+'hapuschat':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					client.deleteChat(args[0])
					reply(`Menghapus chat dengan ${args[0]}`)
					break
				case prefix+'block':
					client.updatePresence(from, Presence.composing) 
					client.chatRead (from)
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isGroup) return reply(ind.groupo())
					if (!isOwner) return reply(ind.ownerb())
					client.blockUser (`${body.slice(7)}@c.us`, "add")
					client.sendMessage(from, `perintah Diterima, memblokir ${body.slice(7)}@c.us`, text)
					await limitAdd(sender)
					break
				case prefix+'unblock':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isGroup) return reply(ind.groupo())
					if (!isOwner) return reply(ind.ownerb())
				    client.blockUser (`${body.slice(9)}@c.us`, "remove")
					client.sendMessage(from, `Perintah Diterima, membuka ${body.slice(9)}@c.us`, text)
					await limitAdd(sender)
					break   				
				case prefix+'edit':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
				    client.updatePresence(from, Presence.composing) 
				if (args[0] === 'pp') {
					enmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(enmedia)
					await client.updateProfilePicture(botNumber, media)
					reply('Makasih profil barunya bosku😗')
				} else if (args[0] === 'reply') {
						cr = body.slice(10)
						reply(`reply berhasil di ubah menjadi : ${cr}`)
				} else if (args[0] === 'prefix') {
					prefix = args[0]
					reply(`*「 SUKSES 」* Prefix jadi ➸ : ${prefix}`)
				} else if (args[0] === 'head1') {
					head1 = q
					reply(`sukses mengganti, head kiri menjadi ➸ : ${head1}`)
				} else if (args[0] === 'head2') {
					head2 = q
					reply(`sukses mengganti, head kanan menjadi ➸ : ${head2}`)
				} else if (args[0] === 'gaya1') {
					gaya1 = q
					reply(`sukses mengganti, gaya pinggir menjadi ➸ : ${gaya1}`)
				} else if (args[0] === 'gaya2') {
					gaya2 = q
					reply(`sukses mengganti, gaya pinggir menjadi ➸ : ${gaya2}`)
				} else if (args[0] === 'gaya3') {
					gaya3 = q
					reply(`sukses mengganti, gaya batas bawah menjadi ➸ : ${gaya3}`)
				}
					await limitAdd(sender)
					break
				case prefix+'clone':
					if (!isRegistered) return reply(ind.noregis())
					if (args.length < 1) return reply(`ups`)
					if (!isGroup) return reply(ind.groupo())
					if (!isOwner) return reply(ind.ownerg())
					if (args.length < 1) return reply(' *TAG YANG MAU DI CLONE!!!* ')
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return reply('Tag cvk')
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid[0]
					let { jid, id, notify } = groupMembers.find(x => x.jid === mentioned)
					try {
						pp = await client.getProfilePicture(id)
						buffer = await getBuffer(pp)
						client.updateProfilePicture(botNumber, buffer)
						mentions(`Foto profile Berhasil di perbarui menggunakan foto profile @${id.split('@')[0]}`, [jid], true)
					} catch (e) {
						reply(ind.stikga())
					}
					await limitAdd(sender)
					break
                case prefix+'ban':
					if (!isRegistered) return reply(ind.noregis())
					if (args.length < 1) return reply(`ups`)
					bnnd = body.slice(6)
					ban.push(`${bnnd}@s.whatsapp.net`)
					fs.writeFileSync('./database/user/banned.json', JSON.stringify(ban))
					reply(`Nomor wa.me/${bnnd} telah dibanned !`)
					await limitAdd(sender)
					break
				case prefix+'banlist':
					client.updatePresence(from, Presence.composing) 
 
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					teks = 'This is list of ban number :\n'
					for (let benn of ban) {
						teks += `~> @${benn.split('@')[0]}\n`
					}
					teks += `Total : ${ban.length}`
					client.sendMessage(from, teks.trim(), extendedText, {quoted: mek, contextInfo: {"mentionedJid": ban}})
					await limitAdd(sender)
					break
				case prefix+'unban':
					if (!isRegistered) return reply(ind.noregis())
					if (args.length < 1) return reply(`ups`)
					bnnd = body.slice(8)
					ban.splice(`${bnnd}@s.whatsapp.net`, 1)
					fs.writeFileSync('./database/user/banned.json', JSON.stringify(ban))
					reply(`Nomor wa.me/${bnnd} telah di unban!`)
					await limitAdd(sender)
					break
				case `${prefix}sibuk`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					var sibuk = body.slice(14)
					var sibuk1 = sibuk.split(" ")[1]
					if (args[0] === 'on') {
						nomersibuk = ''
						pesansibuk = `${sibuk1}`
						reply('on')
					} else if (args[0] === 'off') {
						nomersibuk = `${nomerewa}`
						pesansibuk = `${sibuk1}`
						reply(`off`)
					}
					await limitAdd(sender)
					break
case prefix+'iri':
client.updatePresence(from, Presence.recording)
const irimp3 = fs.readFileSync('./assets/iri.mp3');
client.sendMessage(from, irimp3, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'pale':
client.updatePresence(from, Presence.recording)
const pa = fs.readFileSync('assets/pale.mp3')
client.sendMessage(from, pa, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'sound':
client.updatePresence(from, Presence.recording)
const soun = fs.readFileSync('assets/sound.mp3')
client.sendMessage(from, soun, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break 
case prefix+'sound1':
client.updatePresence(from, Presence.recording)
satu = fs.readFileSync('./assets/sound1.mp3');
client.sendMessage(from, satu, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'sound2':
client.updatePresence(from, Presence.recording)
dua = fs.readFileSync('./assets/sound2.mp3');
client.sendMessage(from, dua, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'sound3':
client.updatePresence(from, Presence.recording)
tiga = fs.readFileSync('./assets/sound3.mp3');
client.sendMessage(from, tiga, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'sound4':
client.updatePresence(from, Presence.recording)
empat = fs.readFileSync('./assets/sound4.mp3');
client.sendMessage(from, empat, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'sound5':
client.updatePresence(from, Presence.recording)
lima = fs.readFileSync('./assets/sound5.mp3');
client.sendMessage(from, lima, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'sound6':
client.updatePresence(from, Presence.recording)
enam = fs.readFileSync('./assets/sound6.mp3');
client.sendMessage(from, enam, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break
case prefix+'sound7':
client.updatePresence(from, Presence.recording)
tujuh = fs.readFileSync('./assets/sound7.mp3');
client.sendMessage(from, tujuh, MessageType.audio, {quoted: mek, mimetype: 'audio/mp4', ptt:true})
break														
/*
]=====> TQTO <=====[
> Rifki ID
> REVOER ID
> ARIS ID
> NADIA CANS
> NAZWA
> VHTEAR
> TOBZ
> MHANKBARBAR
> All Creator Bot WhatsApp
*/



//dari saya sebelumnya
				case prefix+'jadiboterror':
let client1 = new WAConnection()
client1.on('qr', qr => {
   qrcode.generate(qr, { small: true })
   exec(`qrencode -o ./sampah/jadibot_${qr}.png ${body.slice(8)}`)
   reply('membuat kode qr')
   setTimeout( () => {
					qrcodene = fs.readFileSync(`./sampah/jadibot_${qr}.png`)
					}, 1000) // 1000 = 1detik,
					setTimeout( () => {
					client.sendMessage(from, qrcodene, image, {quoted: mek, caption: 'ni'})
					}, 1500) // 1000 = 1detik,
   console.log(color('[','white'),color('','red'),color(']','white'),color('      ^\nSCAN QR CODE','white'),color('BOT','red'),color('By','white'),color('Fauzan Rifki Maulana','yellow'))
})

client1.on('credentials-updated', () => {
	const authInfo = client.base64EncodedAuthInfo()
   console.log(`tersambung`)
   reply(`tersambung`)
   fs.writeFileSync(`./database/user/jadibot@${sender}.json`, JSON.stringify(authInfo, null, '\t'))
})
fs.existsSync('./database/user/jadibot@${sender}.json') && client1.loadAuthInfo('./database/user/jadibot@${sender}.json')
client1.connect();
const buffqr = await getBuffer(`https://api.qrserver.com/v1/create-qr-code/?size=500x500&data=${qr}`)
client.sendMessage(from, buffqr, image, {quoted: mek, caption: `Scan sebelum kadaluarsa\n${monosp}${qr}${monosp}`})
    				break
    			case prefix+'qrcode':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					const tex = encodeURIComponent(body.slice(8))
					if (!tex) return reply(`${prefix}qrcode teksnya`)
					const buff = await getBuffer(`https://api.qrserver.com/v1/create-qr-code/?size=500x500&data=${body.slice(4)}`)
					stiker(`https://api.qrserver.com/v1/create-qr-code/?size=500x500&data=${body.slice(4)}`)
					client.sendMessage(from, buff, image, {quoted: mek})
					await limitAdd(sender)
					break
				case prefix+'help':
				case prefix+'menu':
                    if (!isRegistered) return reply(ind.noregis())
					reqXp  = 5000 * (Math.pow(2, getLevelingLevel(sender)) - 1)
					uangku = checkATMuser(sender)
					menuweton = ['Pahing', 'Pon','Wage','Kliwon','Legi']
					menuneweton = menuweton[Math.floor(((d * 1) + gmt) / 84600000) % 5]
					fitnah(`${nomerewa}`, `${tanda}`, `╭═─⊱ ❰ *INFO INFO INFO* ❱ ⊰─═
${gaya2} *Hari:* ${hari} ${menuneweton}
${gaya2} *Tanggal:* ${tanggal}
${gaya2} *jam:* ${jam} WIB
${gaya2} *Nama:* ${namaneuser(sender)}
${gaya2} *Umurmu:* ${umureuser(sender)} tahun
${gaya2} *Nomer:* ${sender.split("@")[0]}
${gaya2} *XP:* ${getLevelingXp(sender)}
${gaya2} *Level:* ${getLevelingLevel(sender)}
${gaya3}  ⸨ BOT nya FRM ⸩  ⊰─═╯\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${donasi}`)
reply(`*SILAHKAN PILIH BENTUK MENUNYA*\n\n${prefix}allmenu\n     _menu dikirim satu_\n${prefix}allmenu2\n     _menu dikirim terpisah_\n    _bisa meramaikam chat_`)
					break
				case `${prefix}allmenu2`:
				case `${prefix}allhelp2`:
				case `${prefix}listmenu2`:
				case `${prefix}listhelp2`:
				case `${prefix}semuamenu2`:
				case `${prefix}semua.menu2`:
                    if (!isRegistered) return reply(ind.noregis())
					reqXp  = 5000 * (Math.pow(2, getLevelingLevel(sender)) - 1)
					uangku = checkATMuser(sender)
					menuweton = ['Pahing', 'Pon','Wage','Kliwon','Legi']
					menuneweton = menuweton[Math.floor(((d * 1) + gmt) / 84600000) % 5]
					fitnah(`${nomerewa}`, `${tanda}`, `╭═─⊱ ❰ *INFO INFO INFO* ❱ ⊰─═
${gaya2} *Hari:* ${hari} ${menuneweton}
${gaya2} *Tanggal:* ${tanggal}
${gaya2} *jam:* ${jam} WIB
${gaya2} *Nama* : ${namaneuser(sender)}
${gaya2} *Umurmu* : ${umureuser(sender)} tahun
${gaya2} *Nomer* : ${sender.split("@")[0]}
${gaya2} *XP* : ${getLevelingXp(sender)}
${gaya2} *Level* : ${getLevelingLevel(sender)}
${gaya3}  ⸨ BOT nya FRM ⸩  ⊰─═╯\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${donasi}`)
reply(`TERSIMPAN\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${tersimpan}`)
reply(`MENU CEK\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${cekmenu}`)
reply(`MENU EDUKASI\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${edukasimenu}`)
reply(`ACAK GAMBAR\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${wibumenu}`)
reply(`DOWNLOAD\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${downloader}`)
reply(`MENU DI GRUP\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${grupmenu}`)
reply(`PEMBUATAN\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${makermenu}`)
reply(`PEMBUATAN 2\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${makermenu2}`)
					break
				case `${prefix}allmenu`:
				case `${prefix}allhelp`:
				case `${prefix}listmenu`:
				case `${prefix}listhelp`:
				case `${prefix}semuamenu`:
				case `${prefix}semua.menu`:
                    if (!isRegistered) return reply(ind.noregis())
					reqXp  = 5000 * (Math.pow(2, getLevelingLevel(sender)) - 1)
					uangku = checkATMuser(sender)
					menuweton = ['Pahing', 'Pon','Wage','Kliwon','Legi']
					menuneweton = menuweton[Math.floor(((d * 1) + gmt) / 84600000) % 5]
					fitnah(`${nomerewa}`, `${tanda}`, `╭═─⊱ ❰ *INFO INFO INFO* ❱ ⊰─═
${gaya2} *Hari:* ${hari} ${menuneweton}
${gaya2} *Tanggal:* ${tanggal}
${gaya2} *jam:* ${jam} WIB
${gaya2} *Nama* : ${namaneuser(sender)}
${gaya2} *Umurmu* : ${umureuser(sender)} tahun
${gaya2} *Nomer* : ${sender.split("@")[0]}
${gaya2} *XP* : ${getLevelingXp(sender)}
${gaya2} *Level* : ${getLevelingLevel(sender)}
${gaya3}  ⸨ BOT nya FRM ⸩  ⊰─═╯\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${donasi}`)
reply(`TERSIMPAN\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏${tersimpan}
MENU CEK\n${cekmenu}
MENU EDUKASI\n${edukasimenu}
ACAK GAMBAR\n${wibumenu}
DOWNLOAD\n$downloader}
MENU DI GRUP\n${grupmenu}
PEMBUATAN\n${makermenu}
PEMBUATAN 2\n${makermenu2}`)
					break
				case '*123#':
                    if (!isRegistered) return reply(ind.noregis())
					
					reply(`Kode USSD berjalan...`)
					setTimeout( () => {
					
					reply(`*${sender.split("@")[0]} Saldo Rp ${checkATMuser(sender)} s.d null*

${prefix}level
${prefix}limit
${prefix}mining
${prefix}belikuota limit ~jumlah~
${prefix}buylimit ~jumlah~
${prefix}kerja



*════════════*`)
					}, 2000) // 1000 = 1s,
					promo.push(sender)
						fs.writeFileSync('./database/bot/promo.json', JSON.stringify(promo))
					break
				case prefix+'bug':
				case prefix+'lapor':
				case prefix+'report':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Pesan mu mana ?`)
					reply(`Mengirim laporan bug ke wa.me/${nomowner}\n*Dengan pesan*\n${q}`)
					client.updatePresence(`${ownerNumber}`, Presence.composing)
					client.sendMessage(`${ownerNumber}`, `${q}`, text, {quoted: mek})
					client.updatePresence(`${ownerNumber}`, Presence.composing)
					client.sendMessage(`${ownerNumber}`, `untuk membalas silahkan gunakan pesan di bawah ini`, text, {quoted: mek})
					client.updatePresence(`${ownerNumber}`, Presence.composing)
					client.sendMessage(`${ownerNumber}`, `${prefix}balas ${sender} ${from}| ~pesanmu~`, text, {quoted: mek})
					break
				case prefix+'balas':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`${prefix}balas 62xxx@s.whatsapp.net pesanmu`)
					if (args[0].startsWith('08')) return reply(`${tanda}\nPakai 62 jangan pakai 08`)
					reply(`mengirim balasan ke ${args[1]}`)
					client.updatePresence(`${args[1]}`, Presence.composing)
					client.sendMessage(`${args[1]}`, `*Untuk ${args[0]}*\n${body.split("|")[1]}`, text, {quoted: { key: { fromMe: false, participant: `${args[0]}`, ...(from ? { remoteJid: `${args[1]}` } : {}) }, message: { conversation: `_aku ngebug guys_` }}})
					break
				case prefix+'sisa.hari':
				case prefix+'sisahari':
					sakjamberapamenit = '60'
					sakhariberapamenit = '1440'
					sisojam = jam1hari - jamtok
					sisomenit = sisojam * sakjamberapamenit
					sisonejamtambahmenit = sisomenit - menittok
					reply(`Sisa hari di indonesia bagian barat adalah \n${sisojam} jam *atau* ${sisonejamtambahmenit} menit`)
					break
				case prefix+'1k10':
                    if (!isRegistered) return reply(ind.noregis())
					if (!isPromo) return reply(`Maaf, kamu tidak dapat menggunakan promo ini. Silahkan coba lagi nanti`)
					if (args.length < 1) return reply(`kuota limit bot murah, Rp 1000 dapat 10 pesan untuk 1 hari\nuntuk membeli silahkan ketik 1k10 y`)
					if (args[0] === 'y') {
					if ( checkATMuser(sender) <= '1000') return reply(`maaf uang kamu belum mencukupi. silahkan kumpulkan dan beli nanti`)
						confirmATM(sender, '1000')
						bayarLimit(sender, '10')
						fitnah(`${nomerewa}`, `* PEMBELIAN kuota limit BERHASIL *`, `*PEMBELIAN kuota limit*\n\n\n*Penerima* : ${namaneuser(sender)}\n*nominal pembelian* : 5 pesan\n *harga kuota limit* : Rp 1000\n *Sisa saldo mu* : Rp ${checkATMuser(sender)}\n\n\n${createSerial(15)}`)
				  }
					promo.splice(sender, 1)
						fs.writeFileSync('./database/bot/promo.json', JSON.stringify(promo))
					break
				case prefix+'free5k':
                    if (!isRegistered) return reply(ind.noregis())
					if (!isPromo) return reply(`Maaf, kamu tidak dapat menggunakan promo ini. Silahkan coba lagi nanti`)
					if (args.length < 1) return reply(`saldo gratis untuk kamu, Rp 5000.\nMau? ketik free5k y`)
					if (args[0] === 'y') {
                      addKoinUser(sender, '5000')
                      dompetisi = checkATMuser(sender)
                      fitnah(`${nomerewa}`, `berhasil`, `Pembelian saldo Rp 5000 dengan biaya Rp 0 berhasil. Silahkan ketik *123# lalu kirim`)
                     }
                     promo.splice(sender, 1)
						fs.writeFileSync('./database/bot/promo.json', JSON.stringify(promo))
                     break
				case prefix+'listuser':
                    if (!isRegistered) return reply(ind.noregis())
					listuser = `*PENGGUNA BOT INI*\nJumlah: ${_registered.length}\nMạ.af kan saya\n\n`
					for (let listuser1 of _registered) {
						listuser += `*Nama:* ${listuser1.name}
*Umur:* ${listuser1.age}
*No. Wa:* ${listuser1.id.split("@")[0]}
*Waktu daftar:* ${listuser1.time}
*No. SN:* ${listuser1.serial}\n\n`
					}
					listuser += `\nBot by Fauzan Rifki Maulana\nMạ.af saya telah membocorkan`
					reply(listuser.trim())
					break
				case prefix+'listaudio':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					listaudionya = `*DAFTAR AUDIO YANG TERSIMPAN*\n\n`
					for (let listaudione of listaudio) {
						listaudionya += `*•* ${listaudione}\n`
					}
						listaudionya += `\n*SELESAI*`
					reply(listaudionya.trim())
					reply(`*CARA AMBIL*\nCukup ketik nama audio nya, lalu kirim\n\n*NB:* _jika anda belum terdaftar, maka kami tidak membalasnya_`)
					await limitAdd(sender)
					break
				case prefix+'liststicker':
				case prefix+'liststiker':
				case prefix+'getstik':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					liststikernya = `*DAFTAR STIKER YANG TERSIMPAN*\n\n`
					for (let liststikere of liststiker) {
						liststikernya += `*•* ${liststikere}\n`
					}
						liststikernya += `\n*SELESAI*`
					reply(liststikernya.trim())
					reply(`*CARA AMBIL*\nCukup ketik *${prefix}getstik ~nama stiker~*, lalu kirim\n\n*NB:* _perhatikan huruf besar kecil nya, harus sama_`)
					await limitAdd(sender)
					break
				case `*${body.slice(1).split("#")[0]}#`:
                    if (!isRegistered) return reply(ind.noregis())
					fitnah(`${sender}`, `${tanda}`, `Kode MMI *${body.slice(1).split("#")[0]}# ada masalah sambungan atau kode MMI tidak berlaku`)
					break
				case `${prefix}makermenu`:
                    if (!isRegistered) return reply(ind.noregis())
					reply(`${makermenu}`)
					break
				case `${prefix}downloader`:
                    if (!isRegistered) return reply(ind.noregis())
					reply(`${downloader}`)
					break
				case `${prefix}edukasi`:
				case `${prefix}education`:
				case `${prefix}belajar`:
                    if (!isRegistered) return reply(ind.noregis())
					reply(`${edukasimenu}`)
					break
				case `${prefix}cek`:
                    if (!isRegistered) return reply(ind.noregis())
					reply(`${cekmenu}`)
					break
				case `${prefix}wibumenu`:
                    if (!isRegistered) return reply(ind.noregis())
					reply(`${wibumenu}`)
					break
				case `${prefix}grupmenu`:
                    if (!isRegistered) return reply(ind.noregis())
					reply(`${grupmenu}`)
					break
				case `${prefix}ownermenu`:
				case `${prefix}menubosku`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isOwner) return reply(`Mạ.af. Anda Bukan Ownerku\n${ownermenu}`)
					reply(`${ownermenu}`)
					break
				case prefix+'frmgroup':
				case prefix+'groupfrm':
				case prefix+'grupfrm':
				case prefix+'frmgrup':
				case prefix+'grupbot':
				case prefix+'groupbot':
					if (isGroup) return reply(`Anda sudah berada di grupnya FRM BOT`)
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(`*LINK GRUP YANG SAYA PUNYA*`)
					fitnah(`${nomerewa}`, `_Grupnya frm bot_`, `https://chat.whatsapp.com/EXHQaghKyaI6BCJfIFMsKR`)
					fitnah(`${nomerewa}`, `_Grupnya frm bot_`, `https://chat.whatsapp.com/JKOzD8sU19tEnkjPSbTGXI`)
					fitnah(`${nomerewa}`, `_Grupnya frm bot_`, `https://chat.whatsapp.com/BTmVicnGmL3Ligggdmt8XZ`)
			        await limitAdd(sender)
					break
				case `${prefix}makasih`:
				case `${prefix}thank`:
				case `${prefix}thanks`:
                    if (!isRegistered) return reply(ind.noregis())
					fitnah(`${nomerewa}`, `${tanda}`, `
_~instagram.com/frm_developer~_
_~fb.com/fauzan.rifki.m~_
_~t.me/frm_developer~_

◪ ❀ *TERIMAKASIH* ❀
kami ucapkan makasih
untuk:

*SC ini dari*
*☞* https://github.com/affisjunianto/botwasapv3
*☞* https://github.com/Rifki666/babybot
*☞* dan ada potongan kecil dari sc lain

o==[]::::::>
><))):>
('_')

*REST API & WEB HELPER*
☞ api.fdci.se
☞ api.itsmeikyxsec404.xyz
☞ api.shizukaa.xyz
☞ api.vhtear.com
☞ api.xteam.xyz
☞ api.zeks.xyz
☞ chat.whatsapp.com
☞ docs-jojo.herokuapp.com
☞ drive.google.com
☞ freerestapi.herokuapp.com
☞ mnazria.herokuapp.com
☞ i.ibb.co
☞ itsmeikygans.my.id
☞ lolhuman.herokuapp.com
☞ kocakz.herokuapp.com
☞ st4rz.herokuapp.com
☞ terhambar.com
☞ tobz-api.herokuapp.com
☞ translate.google.com
☞ videfikri.com
☞ wa.me
☞ waifu.pics`)
			break
				case `${prefix}dompet`:
				case `${prefix}wallet`:
				case `${prefix}uang`:
				case `${prefix}money`:
                    if (!isRegistered) return reply(ind.noregis())
				cekdompet = checkATMuser(sender)
				fitnah(`${nomerewa}`, `ISI DOMPET MU`, `* dompet *\n*Nama* : ${pushname}\n*Nomer* : ${sender.split("@")[0]}\n*Uang* : Rp ${cekdompet}\n`)
				
				break
			case `${prefix}gift.kuota `:
				
                if (!isRegistered) return reply(ind.noregis())
				if (args.length < 1) return reply(`* SELAMAT DATANG *\nSelamat datang\n\nSilahkan ketik\n${prefix}gift.kuota limit ~nomor~  ~jumlah kuota limit~\n*CONTOH*\n${prefix}gift.kuota limit 62895803265350 7`)
				if (!isUser) return reply(`nomor tidak terdaftar`)
				payout = (`-${args[1]}`)
				const totale = args[1]
				nomerwesdaftar = `${args[0]}@s.whatsapp.net`
				if ( checkLimit(sender) >= totale) return reply(`maaf kuota limit kamu belum mencukupi. silahkan kumpulkan dan transfer nanti`)
				if ( checkLimit(sender) <= totale ) {
					limit(sender, totale)
					bayarLimit(args[0], payout)
					fitnah(`${nomerewa}`, `* TRANSFER kuota limit BERHASIL *`, `*TRANSFER kuota limit*\n\n\n*Penerima* : ${args[0]}\n*nominal transfer* : ${payout} pesan\n *Sisa kuota limit mu* : Rp ${checkLimit(sender)}\n\n\n${createSerial(15)}`)
				} 
				
				break
				case `${prefix}jualxp`:
				case `${prefix}salexp`:
                    if (!isRegistered) return reply(ind.noregis())
					if (args.length < 1) return reply(`SELAMAT DATANG\nSilahkan Jual XP disini\nKetik ${prefix}jualxp ~jumlah uang~\n\nBiaya Rp 1000 dapat 1 Limit.`)
					totaluang = args[0]
					const XpPerUang = 1000
					const totaljualxp = XpPerUang * totaluang
					if ( getLevelingXp(sender) <= totaljualxp) return reply(`maaf uang kamu belum mencukupi. silahkan kumpulkan dan beli nanti`)
					if ( getLevelingXp(sender) >= totaljualxp ) {
					addLevelingXp(sender, totaljualxp)
					addKoinUser(sender, totaluang)
					fitnah(`${nomerewa}`, `* PENJUALAN XP BERHASIL *`, `*PENJUALAN XP* \n\n\n*Penerima* : ${pushname}\n*nominal penjualan* : Rp ${totaluang} \n *harga xp* : Rp ${XpPerUang} per XP\n *Sisa uang mu* : Rp ${checkATMuser(sender)}\n\n\n_${createSerial(15)}_`)
				}
					
					break
				case `${prefix}kerja`:
                      
                	  if (!isRegistered) return reply(ind.noregis())
				      if (isBanned) return reply(ind.diban())
				if (args.length < 1) {
                      bayaran = 1000
                      addKoinUser(sender, bayaran)
                      dompetisi = checkATMuser(sender)
                      kerjane = {
							text: `*HASIL BEKERJA*\n\n*Penerima:* @${sender.split("@")[0]}\n*Gaji:* Rp ${bayaran}\n*Saldo anda:* Rp ${dompetisi}\n*No. SN:* ${monosp}${createSerial(15)}${monosp}`,
							contextInfo: { mentionedJid: [sender] }
						}
                      fitnah(`${nomerewa}`, `_berhasil_`, kerjane)
				} if (args.length > 0) {
					bayaran = 1000
					addKoinUser(sender, bayaran)
					dompetisi = checkATMuser(sender)
					mergawe = {
							text: `*HASIL BEKERJA*\n\n*Penerima:* @${sender.split("@")[0]}\n*Pekerjaan:* ${q}\n*Gaji:* Rp ${bayaran}\n*Saldo anda:* Rp ${dompetisi}\n*No. SN:* ${monosp}${createSerial(15)}${monosp}`,
							contextInfo: { mentionedJid: [sender] }
						}
                      fitnah(`${nomerewa}`, `_berhasil_`, mergawe)
				}
					break
				case `${prefix}10k`:
                      
                	  if (!isRegistered) return reply(ind.noregis())
				      if (isBanned) return reply(ind.diban())
					tipelist = [`${katasandi}`]
					if (!tipelist.includes(args[0])) return reply(`${tanda}\nKatasandi salah, silahkan beli. Hubungi wa.me/62895803265350`)
                      bayaran = 10000
                      addKoinUser(sender, bayaran)
                      dompetisi = checkATMuser(sender)
                      fitnah(`${nomerewa}`, `UHUI`, `*WOOW*\n\n\nRp ${bayaran} telah ditambahkan\nsaldomu sekarang Rp ${dompetisi}\n\n\n_${createSerial(15)}_`)
					break
				case `${prefix}addtaksopan`:
                    if (!isRegistered) return reply(ind.noregis())
					if (args.length < 1) return reply(`ketik\n${prefix}addtaksopan ~kata jeleknya~\nCONTOH\n${prefix}addtaksopan cok`)
					var gaksopan = args[0]
               	 var katane = body.slice(13+gaksopan.length)
					omongelek.push(katane)
						fs.writeFileSync('./database/bot/omongelek.json', JSON.stringify(omongelek))
						fitnah(nomerewa, `berhasil`, `Kata ${katane} telah ditambahkan`)
					break
				case prefix+'tf.saldo':
					reply(`Untuk mentransfer saldo, silahkan ketik\n${prefix}tf _nomornya_ _nominal_\n\nContoh\n${prefix}tf ${me.jid.split("@")[0]} 5000`)
					break
				case prefix+'tf':
                    if (!isRegistered) return reply(ind.noregis())
					if (isBanned) return reply(`Maaf, nomor kamu tidak dapat menggunakan bot ini\nSilahkan mohon kepada bosku / ownerku`)
					if (args.length < 1) return reply(`Untuk mentransfer saldo, silahkan ketik\n${prefix}tf _nomornya_ _nominal_\n\nContoh\n${prefix}tf ${me.jid.split("@")[0]} 5000`)
					if (args[0].startsWith('08')) return reply(`${tanda}\nPakai 62 jangan pakai 08`)
                var tujuan = args[0]
                var jumblah = body.slice(4+tujuan.length)
                if (checkATMuser(sender) < jumblah) return reply(`saldo mu tidak mencukupi untuk melakukan transfer`)
                tujuantf = `${tujuan.replace("@", '')}@s.whatsapp.net`
                nomerwesdaftar = `${tujuan}@s.whatsapp.net`
                if (!isUser) return reply(`Nomor ${args[0]} tidak terdaftar`)
                fee = 0.005 *  jumblah
                hasiltf = jumblah - fee
			if (isUser) {
                addKoinUser(`${args[0]}@s.whatsapp.net`, hasiltf)
                confirmATM(sender, jumblah)
                fitnah(`${nomerewa}`, `berhasil`, `Transfer saldo Rp ${jumblah} ke : +${nomerwesdaftar} berhasil\nDengan pajak : ${fee}`)
                fitnah2(`${nomerwesdaftar}`, `Dari ${sender.split("@")[0]}\nDari ${namaneuser(sender)}`, `*Nomor WhatsApp* mu telah di isi saldo senilai Rp ${jumblah} dengan SN ${createSerial(15)}`)
			}
					break
					case prefix+'attp':
					case prefix+'ttp':
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply('${prefix}ttp orang itu aneh\n\ncontohnya itu')
				let yosh = body.slice(6)
				stikergif(`https://api.xteam.xyz/attp?file&text=${encodeURIComponent(yosh)}`)
            	await limitAdd(sender)
           	 break
           case prefix+'ninjalogo':
				 if (!isRegistered) return reply(ind.noregis())
				 if (isBanned) return reply(ind.diban())
				 if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				 var gh = body.slice(11)
				 var nin = gh.split("&")[0];
				 var ja = gh.split("&")[1];
				 if (args.length < 1) return reply(`「❗」Contoh : ${prefix}ninjalogo Rifki & Gans`)
				 reply(ind.wait())
				 buffer = await getBuffer(`https://api.xteam.xyz/textpro/ninjalogo?text=${nin}&text2=${ja}&APIKEY=${XteamKey}`)
				 client.sendMessage(from, buffer, image, {quoted: mek})
				 await limitAdd(sender)
				 break
           	case prefix+'cloudtext':
 if (!isRegistered) return reply(ind.noregis())
				 	if (isBanned) return reply(ind.diban())
				 	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
 if (args.length < 1) return reply(`Contoh : ${prefix}cloudtext Rifki`)
 cloud = body.slice(11)
 reply('Bentar Bro terbang dulu yekan yahaha hayukk')
 buffer = await getBuffer(`https://api.xteam.xyz/textpro/cloudtext?text=${cloud}&APIKEY=${XteamKey}`)
 client.sendMessage(from, buffer, image, {quoted: mek})
 await limitAdd(sender)
 break
           		case prefix+'glitchtext':
				 	if (!isRegistered) return reply(ind.noregis())
				 	if (isBanned) return reply(ind.diban())
				 	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					 var gh = body.slice(12)
					 var gli = gh.split("&")[0];
					 var tch = gh.split("&")[1];
					 if (args.length < 1) return reply(`[❗] Contoh : ${prefix}glitchtext Rifki & Gans`)
					 reply(ind.wait())
					 buffer = await getBuffer(`https://api.xteam.xyz/textpro/glitch?text=${gli}&text2=${tch}&APIKEY=${XteamKey}`)
					 client.sendMessage(from, buffer, image, {quoted: mek})
					 await limitAdd(sender)
					 break
					case `${prefix}summer`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/summerysandwriting?text=${body.slice(8)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break
					case `${prefix}sandwrite`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/sandwriting?text=${body.slice(11)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break 
					case `${prefix}metaldark`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/metaldarkgold?text=${body.slice(11)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break 
					case `${prefix}dropwater`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/dropwater?text=${body.slice(11)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break 
					case `${prefix}greenneon`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					gneondl = await getBuffer(`https://api.xteam.xyz/textpro/greenneon?text=${body.slice(11)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, gneondl, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break 
				case prefix+'neon':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`「❗」Contoh : ${prefix}neontext Rifki`)
					naon = body.slice(6)
					reply('「❗」WAIT GANS')
					alu = await getBuffer(`https://api.xteam.xyz/textpro/neon?text=${naon}&APIKEY=${XteamKey}`)
					client.sendMessage(from, alu, image, {quoted: mek})
					break
				case prefix+'textlight':
                  if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply(ind.wrongf())
				ligh = body.slice(11)
				if (ligh.length > 10) return reply('Teksnya kepanjangan, maksimal 9 karakter')
				reply(ind.wait())
				lawak = await getBuffer(`https://api.zeks.xyz/api/tlight?text=${ligh}&apikey=apivinz`)
		    client.sendMessage(from, lawak, image, {quoted: mek})
		    await limitAdd(sender)
		    break
					case `${prefix}neontext`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/neontext?text=${body.slice(10)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break 
				case `${prefix}toxic`: //frm developer
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					const asune =['anjing','babi lu','anak anjing','udah tolol nub Lagi','muka lo kek monyet','udah jomblo sendirian lagi dirumah tolol','so so an mau punya pacar muka aja kek monyet lepass dari kandang','ganteng doang di toxic aja dibilang baperan','pantek kau','bangsat kau','ku entod kalian nangis kau','memek lu semua','lihat anak anjing lagi baca','ganteng doang jemput cewe dipanggang','kamu cantik beb bullshit anjing cowo buaya','anak dajjal','puki lu','anjing ngajak gelud','sama hantu takut cupu ngentod','cupu cupu aja gausah bacot','kontol lu semua','bocah lu semua kontol']
					const coknya = asune[Math.floor(Math.random() * asune.length)]
					reply(`${tanda}\n${coknya}`)
					await limitAdd(sender)
					break
					case `${prefix}sumery`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/summerysandwriting?text=${body.slice(8)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break
					case `${prefix}blood`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/bloodontheroastedglass?text=${body.slice(7)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break
					case `${prefix}firework`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					arugazzz = await getBuffer(`https://api.xteam.xyz/textpro/fireworksparkle?text=${body.slice(10)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, arugazzz, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break
					case `${prefix}lava`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(ind.wrongf())
					reply(ind.wait())
					aruga = await getBuffer(`https://api.xteam.xyz/textpro/lava?text=${body.slice(6)}&APIKEY=${XteamKey}`)
					client.sendMessage(from, aruga, image, {caption: 'Nih kak', quoted: mek})
					await limitAdd(sender)
					break
                case `${prefix}1cak`:
				    try {
					    
                    if (!isRegistered) return reply(ind.noregis())
						if (isBanned) return reply(ind.diban())
					    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						res = await fetchJson(`https://st4rz.herokuapp.com/api/1cak`, {method: 'get'})
						buffer = await getBuffer(res.result)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: 'ni anjim'})
					} catch (e) {
						console.log(`Error :`, color(e,'red'))
						reply(ind.wrongf())
					}
					await limitAdd(sender)
					break
                case prefix+'quotes':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					quotes = await fetchJson(`https://api.itsmeikyxsec404.xyz/quotesad?apikey=itsmeiky633`, {method: 'get'})
					reply(`${tanda}\n${quotes.result}`)
					await limitAdd(sender)
					break
				case prefix+'quotes3':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					quotes3 = await fetchJson(`https://tobz-api.herokuapp.com/api/randomquotes?apikey=${TobzKey}`, {method: 'get'})
					reply(`${tanda}\n${quotes3.quotes}`)
					await limitAdd(sender)
					break
				case prefix+'asupan':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					asupan = await fetchJson(`http://lolhuman.herokuapp.com/api/asupan?apikey=${LolKey}`, {method: 'get'})
					if (asupan.message) return reply(asupan.message)
					asupan1 = await getBuffer(asupan.result)
					client.sendMessage(from, asupan1, video, {mimetype: 'video/mp4', quoted: mek, caption: `${tanda}`})
					break
				case prefix+'katailham':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					katailham = await fetchJson(`https://api.shizukaa.xyz/api/bacotanilham?apikey=itsmeiky633`, {method: 'get'})
					reply(`${tanda}\n${katailham.result}`)
					await limitAdd(sender)
					break
				case prefix+'pantun':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					pantun = await fetchJson(`http://kocakz.herokuapp.com/api/random/text/pantun`, {method: 'get'})
					reply(`${pantun.result}`)
					await limitAdd(sender)
					break
				case prefix+'quotes2':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				 quotes = body.slice(1)
				 const quo =['Lebih baik mengerti sedikit daripada salah mengerti.','Hampir semua pria memang mampu bertahan menghadapi kesulitan. Namun, jika Anda ingin menguji karakter sejati pria, beri dia kekuasaan.','Bila tekad seseorang kuat dan teguh, Tuhan akan bergabung dalam usahanya.','Penderitaan adalah pelajaran.','Ilmu pengetahuan tanpa agama adalah pincang.','Hidup itu seperti sebuah sepeda, agar tetap seimbang kita harus tetap bergerak.','Perbedaan masa lalu, sekarang, dan masa depan tak lebih dari ilusi yang keras kepala.','Sebuah meja, sebuah kursi, semangkuk buah, dan sebuah biola; apa lagi yang dibutuhkan agar seseorang bisa merasa bahagia?','Belas kasihanlah terhadap sesama, bersikap keraslah terhadap diri sendiri.','Cara paling baik untuk menggerakkan diri Anda ialah memberi tugas kepada diri sendiri.','Kita tidak boleh kehilangan semangat. Semangat adalah stimulan terkuat untuk mencintai, berkreasi dan berkeinginan untuk hidup lebih lama.','Manusia akan bahagia selama ia memilih untuk bahagia.','Saya tidak berharap menjadi segalanya bagi setiap orang. Saya hanya ingin menjadi sesuatu untuk seseorang.','Apabila sempurna akal seseorang, maka sedikit perkataannya.','Bahagialah orang yang dapat menjadi tuan untuk dirinya, menjadi kusir untuk nafsunya dan menjadi kapten untuk bahtera hidupnya.','Sahabat yang jujur lebih besar harganya daripada harta benda yang diwarisi dari nenek moyang.','Yang paling melelahkan dalam hidup adalah menjadi orang yang tidak tulus.','Terbuka untuk Anda, begitulah Tuhan memberi kita jalan untuk berusaha. Jangan pernah berfikir jalan sudah tertutup.','Penundaan adalah kuburan dimana peluang dikuburkan.','Cinta bukan saling menatap mata, namun melihat ke arah yang sama bersama-sama.','Kita adalah apa yang kita kerjakan berulang kali. Dengan demikian, kecemerlangan bukan tindakan, tetapi kebiasaan.','Jangan pernah mencoba menjadikan putra atau putri Anda menjadi seperti Anda. Diri Anda hanya cukup satu saja.','Jika Anda bisa membuat orang lain tertawa, maka Anda akan mendapatkan semua cinta yang Anda inginkan.','Masalah akan datang cepat atau lambat. Jika masalah datang, sambut dengan sebaik mungkin. Semakin ramah Anda menyapanya, semakin cepat ia pergi.','Kita tak bisa melakukan apapun untuk mengubah masa lalu. Tapi apapun yang kita lakukan bisa mengubah masa depan.','Kesabaran adalah teman dari kebijaksanaan.','Orang-orang kreatif termotivasi oleh keinginan untuk maju, bukan oleh keinginan untuk mengalahkan orang lain.','Dimanapun engkau berada selalulah menjadi yang terbaik dan berikan yang terbaik dari yang bisa kita berikan.','Kebencian seperti halnya cinta, berkobar karena hal-hal kecil.','Anda tidak perlu harus berhasil pada kali pertama.','Satu jam yang intensif, jauh lebih baik dan menguntungkan daripada bertahun-tahun bermimpi dan merenung-renung.','Hal terbaik yang bisa Anda lakukan untuk orang lain bukanlah membagikan kekayaan Anda, tetapi membantu dia untuk memiliki kekayaannya sendiri.','Tidak ada jaminan keberhasilan, tetapi tidak berusaha adalah jaminan kegagalan.','Aku tidak tahu kunci sukses itu apa, tapi kunci menuju kegagalan adalah mencoba membuat semua orang senang.']
				 const tes = quo[Math.floor(Math.random() * quo.length)]
				 client.sendMessage(from, ''+tes+'\n\n_By : ⸸Rifki⸸Panutanque._', text, { quoted: mek })
				 await limitAdd(sender)
				 break
				case prefix+'mediafire':
                    if (!isRegistered) return reply(ind.noregis())
						if (isBanned) return reply(ind.diban())
					    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					mf = await fetchJson(`https://api-anoncybfakeplayer.herokuapp.com/mediafire?url=${q}`, {method: 'get'})
					reply(`Tunggu, sedang memproses file dengan ukuran ${mf.filesize}`)
					mediafiredl = await getBuffer(mf.result)
					client.sendMessage(from, mediafiredl, document, {quoted:mek})
					await limitAdd(sender)
					break
				case `${prefix}ytcari`:
				case `${prefix}ytsearch`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ytcari = await fetchJson(`https://api.vhtear.com/youtube?query=${q}&apikey=${VhtearKey}`, {method: 'get'})
					if (ytcari.error) return reply(ytcari.error)
					teks = '=================\n'
					for (let i of ytcari.result) {
						teks += `*Title* : ${i.title}\n*URL* : https://www.youtube.com/watch?v=${i.id}\n*Published* : ${i.publishTime}\n*Duration* : ${i.duration}\n*Views* : ${h2k(i.views)}\n=================\n`
					}
					reply(teks.trim())
					break
					case `${prefix}infonomor`:
                    if (!isRegistered) return reply(ind.noregis())
if (isBanned) return reply(`Maaf, nomor kamu tidak dapat menggunakan bot ini\nSilahkan mohon kepada bosku / ownerku`)
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    if (args.length < 1) return reply(`Masukan Nomor\nContoh : ${prefix}infonomor 0812345678`)
                data = await fetchJson(`https://docs-jojo.herokuapp.com/api/infonomor?no=${body.slice(11)}`)
                if (data.error) return reply(data.error)
                if (data.result) return reply(data.result)
                hasil = ` internasional : ${data.international}\n nomor : ${data.nomor}\n operator : ${data.op}`
                reply(hasil)
                await limitAdd(sender)
					break 
				case `${prefix}jadwaltv`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					hasiljadwaltv = await fetchJson(`https://api.zeks.xyz/api/jadwaltv?channel=${body.slice(10)}&apikey=${ZeksKey}`, {method: 'get'})
					if (hasiljadwaltv.listchannel) return reply(`${tanda}\n${hasiljadwaltv.listchannel}`)
					jadwaltv = `${hasiljadwaltv.result}`
					jadwaletv = (jadwaltv.trim())
					reply(jadwaletv.replace(',','\n'))
					await limitAdd(sender)
					break
				case prefix+'news':
				case prefix+'berita':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					data = await fetchJson(`https://api.zeks.xyz/api/tribunews?apikey=${ZeksKey}`, {method: 'get'})
					teks = '=================\n'
					for (let i of data.result) {
						teks += `*Judul* : ${i.title}\n*link* : ${i.link}\n*Keterangan* : ${i.ket}\n*Url* : ${i.url}\n=================\n`
					}
					reply(teks.trim())
					await limitAdd(sender)
					break
				case prefix+'spam':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(`*SILAHKAN PILIH SPAM NYA*\n\n${prefix}spamcall\n${prefix}spamemail\n${prefix}spamsms`)
					await limitAdd(sender)
					break
				case prefix+'spamcall':
				case prefix+'callspam':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Format salah, silahkan ketik \n${prefix}spamcall ~nomornya~\n*CONTOH*\n${prefix}spamcall ${me.jid.split("@")[0]}`)
						spamcall = await fetchJson(`https://videfikri.com/api/call/?nohp=${args[1]}`, {method: 'get'})
						reply(`${spamcall.result.nohp}\n${spamcall.result.logs}`)
						await limitAdd(sender)
				case prefix+'spamemail':
				case prefix+'emailspam':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Format salah, silahkan ketik \n${prefix}spamcall ~nomornya~\n*CONTOH*\n${prefix}spamcall ${me.jid.split("@")[0]}`)
						spamemail = await fetchJson(`https://videfikri.com/api/spamemail/?email=${args[1]}&subjek=Hallo&pesan=Silahkan bayar tagihan listrik Anda`, {method: 'get'})
						reply(spamemail.result.log_lengkap)
						await limitAdd(sender)
				case prefix+'spamsms':
				case prefix+'smsspam':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Format salah, silahkan ketik \n${prefix}spamsms ~nomornya~\n*CONTOH*\n${prefix}spamsms ${me.jid.split("@")[0]}`)
						spamsms = await fetchJson(`https://core.ktbs.io/v2/user/registration/otp/${args[1]}`, {method: 'get'})
                        spamsms1 = await fetchJson(`https://api.danacita.co.id/users/send_otp/?mobile_phone=${args[1]}`, {method: 'get'})
                        spamsms2 = await fetchJson(`https://account-api-v1.klikindomaret.com/api/PreRegistration/SendOTPSMS?NoHP=${args[1]}`, {method: 'get'})
						reply(`mengirim sms`)
						await limitAdd(sender)
					break
				case prefix+'faktaunik':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					fakta = await fetchJson(`https://videfikri.com/api/fakta`, {method: 'get'})
					reply(`${fakta.result.fakta}`)
					break
				case prefix+'wiki':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`mana teks nya`)
					wiki = await fetchJson(`https://videfikri.com/api/wiki/?query=${q}`, {method: 'get'})
					reply(`${wiki.result.judul}\nMenurut wikipedia\n${wiki.result.isi_konten}`)
					break
				case prefix+'wikien':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`mana teks nya`)
					wiki = await fetchJson(`https://videfikri.com/api/wikieng/?query=${q}`, {method: 'get'})
					reply(`${wiki.result.title}\naccording to wikipedia\n${wiki.result.desc}`)
					break
                case `${prefix}slap`:
                	
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    kapankah = body.slice(1)
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					const slap =['anjing','babi lu','anak anjing','udah tolol nub Lagi','muka lo kek monyet','udah jomblo sendirian lagi dirumah tolol','so so an mau punya pacar muka aja kek monyet lepass dari kandang','ganteng doang di toxic aja dibilang baperan','pantek kau','bangsat kau','ku entod kalian nangis kau','memek lu semua','lihat anak anjing lagi baca','ganteng doang jemput cewe dipanggang','kamu cantik beb bullshit anjing cowo buaya','anak dajjal','puki lu','anjing ngajak gelud','sama hantu takut cupu ngentod','cupu cupu aja gausah bacot','kontol lu semua','bocah lu semua kontol','3 Hari Lagi']
					const ple = slap[Math.floor(Math.random() * slap.length)]
					pod = fs.readFileSync('./fauzan.rifki.m/tampar.gif')
					client.sendMessage(from, pod, image, { quoted: mek, caption: tanda+'\n*Toxic*\n\n'+ ple })
					await limitAdd(sender)
					break
				case prefix+'jadian':
					if (!isRegistered) return reply(ind.noregis())
                	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isGroup) return reply(ind.groupo())
					jds = []
					const jdii = groupMembers
					const koss = groupMembers
					const akuu = jdii[Math.floor(Math.random() * jdii.length)]
					const diaa = koss[Math.floor(Math.random() * koss.length)]
					teks = `Ciee.. yang lagi jadian \n@${akuu.jid.split('@')[0]} ♥️ @${diaa.jid.split('@')[0]} `
					jds.push(akuu.jid)
					jds.push(diaa.jid)
					mentions(teks, jds, true)
					await limitAdd(sender)
					break
					case `${prefix}tampar`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					tampar = await fetchJson(`https://waifu.pics/api/sfw/slap`, {method: 'get'})
					tampare = await getBuffer(tampar.url)
					client.sendMessage(from, tampar, sticker, {quoted: mek})
					exec(`ffmpeg -i ${tampare} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(tampare)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
				case prefix+'trap':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					hmtrap = await fetchJson(`https://waifu.pics/api/nsfw/trap`, {method: 'get'})
					trap = await getBuffer(hmtrap.url)
					client.sendMessage(from, trap, image, {quoted: mek})
					await limitAdd(sender)
					break
                case `${prefix}beritahoax`:
                     
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    client.updatePresence(from, Presence.composing) 
					data = await fetchJson(`https://docs-jojo.herokuapp.com/api/infohoax`, {method: 'get'})
					teks = '=================\n'
					for (let i of data.result) {
						teks += `*Gambar* : ${i.image}\n*Title* : ${i.title}\n*link* : ${i.link}\n*tag* : ${i.tag}\n=================\n`
					}
					reply(teks.trim())
					await limitAdd(sender)
					break
				case prefix+'katafrm':
				case prefix+'katarifki':
				case prefix+'katafauzan':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
rifkiberkata = ["Takdir mati bisa di ubah dengan cara bunuh diri",
"android berasal dari linux yang dikembangkan oleh google",
"dia yang selalu menemaniku adalah teman dekat ku"]
					katanerifki = rifkiberkata[Math.floor(Math.random() * rifkiberkata.length)]
					reply(`${katanerifki} *-* kata Fαυzαη Rιƒкι MαυĻαηα`)
					break
				case prefix+'bucin':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					bucin = [
  "Aku memilih untuk sendiri, bukan karena menunggu yang sempurna, tetapi butuh yang tak pernah menyerah.",
  "Seorang yang single diciptakan bersama pasangan yang belum ditemukannya.",
  "Jomblo. Mungkin itu cara Tuhan untuk mengatakan 'Istirahatlah dari cinta yang salah'.",
  "Jomblo adalah anak muda yang mendahulukan pengembangan pribadinya untuk cinta yang lebih berkelas nantinya.",
  "Aku bukan mencari seseorang yang sempurna, tapi aku mencari orang yang menjadi sempurna berkat kelebihanku.",
  "Pacar orang adalah jodoh kita yang tertunda.",
  "Jomblo pasti berlalu. Semua ada saatnya, saat semua kesendirian menjadi sebuah kebersamaan dengannya kekasih halal. Bersabarlah.",
  "Romeo rela mati untuk juliet, Jack mati karena menyelamatkan Rose. Intinya, kalau tetap mau hidup, jadilah single.",
  "Aku mencari orang bukan dari kelebihannya tapi aku mencari orang dari ketulusan hatinya.",
  "Jodoh bukan sendal jepit, yang kerap tertukar. Jadi teruslah berada dalam perjuangan yang semestinya.",
  "Kalau kamu jadi senar gitar, aku nggak mau jadi gitarisnya. Karena aku nggak mau mutusin kamu.",
  "Bila mencintaimu adalah ilusi, maka izinkan aku berimajinasi selamanya.",
  "Sayang... Tugas aku hanya mencintaimu, bukan melawan takdir.",
  "Saat aku sedang bersamamu rasanya 1 jam hanya 1 detik, tetapi jika aku jauh darimu rasanya 1 hari menjadi 1 tahun.",
  "Kolak pisang tahu sumedang, walau jarak membentang cintaku takkan pernah hilang.",
  "Aku ingin menjadi satu-satunya, bukan salah satunya.",
  "Aku tidak bisa berjanji untuk menjadi yang baik. Tapi aku berjanji akan selalu mendampingi kamu.",
  "Kalau aku jadi wakil rakyat aku pasti gagal, gimana mau mikirin rakyat kalau yang selalu ada dipikiran aku hanyalah dirimu.",
  "Lihat kebunku, penuh dengan bunga. Lihat matamu, hatiku berbunga-bunga.",
  "Berjanjilah untuk terus bersamaku sekarang, esok, dan selamanya.",
  "Rindu tidak hanya muncul karena jarak yang terpisah. Tapi juga karena keinginan yang tidak terwujud.",
  "Kamu tidak akan pernah jauh dariku, kemanapun aku pergi kamu selalu ada, karena kamu selalu di hatiku, yang jauh hanya raga kita bukan hati kita.",
  "Aku tahu dalam setiap tatapanku, kita terhalang oleh jarak dan waktu. Tapi aku yakin kalau nanti kita pasti bisa bersatu.",
  "Merindukanmu tanpa pernah bertemu sama halnya dengan menciptakan lagu yang tak pernah ternyayikan.",
  "Ada kalanya jarak selalu menjadi penghalang antara aku sama kamu, namun tetap saja di hatiku kita selalu dekat.",
  "Jika hati ini tak mampu membendung segala kerinduan, apa daya tak ada yang bisa aku lakukan selain mendoakanmu.",
  "Mungkin di saat ini aku hanya bisa menahan kerinduan ini. Sampai tiba saatnya nanti aku bisa bertemu dan melepaskan kerinduan ini bersamamu.",
  "Melalui rasa rindu yang bergejolak dalam hati, di situ terkadang aku sangat membutuhkan dekap peluk kasih sayangmu.",
  "Dalam dinginnya malam, tak kuingat lagi; Berapa sering aku memikirkanmu juga merindukanmu.",
  "Merindukanmu itu seperti hujan yang datang tiba-tiba dan bertahan lama. Dan bahkan setelah hujan reda, rinduku masih terasa.",
  "Sejak mengenalmu bawaannya aku pengen belajar terus, belajar menjadi yang terbaik buat kamu.",
  "Tahu gak perbedaan pensi sama wajah kamu? Kalau pensil tulisannya bisa dihapus, tapi kalau wajah kamu gak akan ada yang bisa hapus dari pikiran aku.",
  "Bukan Ujian Nasional besok yang harus aku khawatirkan, tapi ujian hidup yang aku lalui setelah kamu meninggalkanku.",
  "Satu hal kebahagiaan di sekolah yang terus membuatku semangat adalah bisa melihat senyumanmu setiap hari.",
  "Kamu tahu gak perbedaanya kalau ke sekolah sama ke rumah kamu? Kalo ke sekolah pasti yang di bawa itu buku dan pulpen, tapi kalo ke rumah kamu, aku cukup membawa hati dan cinta.",
  "Aku gak sedih kok kalo besok hari senin, aku sedihnya kalau gak ketemu kamu.",
  "Momen cintaku tegak lurus dengan momen cintamu. Menjadikan cinta kita sebagai titik ekuilibrium yang sempurna.",
  "Aku rela ikut lomba lari keliling dunia, asalkan engkai yang menjadi garis finishnya.",
  "PR-ku adalah merindukanmu. Lebih kuat dari Matematika, lebih luas dari Fisika, lebih kerasa dari Biologi.",
  "Cintaku kepadamu itu bagaikan metabolisme, yang gak akan berhenti sampai mati.",
  "Kalau jelangkungnya kaya kamu, dateng aku jemput, pulang aku anter deh.",
  "Makan apapun aku suka asal sama kamu, termasuk makan ati.",
  "Cinta itu kaya hukuman mati. Kalau nggak ditembak, ya digantung.",
  "Mencintaimu itu kayak narkoba: sekali coba jadi candu, gak dicoba bikin penasaran, ditinggalin bikin sakaw.",
  "Gue paling suka ngemil karena ngemil itu enak. Apalagi ngemilikin kamu sepenuhnya...",
  "Dunia ini cuma milik kita berdua. Yang lainnya cuma ngontrak.",
  "Bagi aku, semua hari itu adalah hari Selasa. Selasa di Surga bila dekat denganmu...",
  "Bagaimana kalau kita berdua jadi komplotan penjahat? Aku curi hatimu dan kamu curi hatiku.",
  "Kamu itu seperti kopi yang aku seruput pagi ini. Pahit, tapi bikin nagih.",
  "Aku sering cemburu sama lipstikmu. Dia bisa nyium kamu tiap hari, dari pagi sampai malam.",
  "Hanya mendengar namamu saja sudah bisa membuatku tersenyum seperti orang bodoh.",
  "Aku tau teman wanitamu bukan hanya satu, dan menyukaimu pun bukan hanya aku.",
  "Semenjak aku berhenti berharap pada dirimu, aku jadi tidak semangat dalam segala hal..",
  "Denganmu, jatuh cinta adalah patah hati paling sengaja.",
  "Sangat sulit merasakan kebahagiaan hidup tanpa kehadiran kamu disisiku.",
  "Melalui rasa rindu yang bergejolak dalam hati, di situ terkadang aku sangat membutuhkan dekap peluk kasih sayangmu.",
  "Sendainya kamu tahu, sampai saat ini aku masih mencintaimu.",
  "Terkadang aku iri sama layangan..talinya putus saja masih dikejar kejar dan gak rela direbut orang lain...",
  "Aku tidak tahu apa itu cinta, sampai akhirnya aku bertemu denganmu. Tapi, saat itu juga aku tahu rasanya patah hati.",
  "Mengejar itu capek, tapi lebih capek lagi menunggu\nMenunggu kamu menyadari keberadaanku...",
  "Jangan berhenti mencinta hanya karena pernah terluka. Karena tak ada pelangi tanpa hujan, tak ada cinta sejati tanpa tangisan.",
  "Aku punya sejuta alasan unutk melupakanmu, tapi tak ada yang bisa memaksaku untuk berhenti mencintaimu.",
  "Terkadang seseorang terasa sangat bodoh hanya untuk mencintai seseorang.",
  "Kamu adalah patah hati terbaik yang gak pernah aku sesali.",
  "Bukannya tak pantas ditunggu, hanya saja sering memberi harapan palsu.",
  "Sebagian diriku merasa sakit, Mengingat dirinya yang sangat dekat, tapi tak tersentuh.",
  "Hal yang terbaik dalam mencintai seseorang adalah dengan diam-diam mendo akannya.",
  "Kuharap aku bisa menghilangkan perasaan ini secepat aku kehilanganmu.",
  "Demi cinta kita menipu diri sendiri. Berusaha kuat nyatanya jatuh secara tak terhormat.",
  "Anggaplah aku rumahmu, jika kamu pergi kamu mengerti kemana arah pulang. Menetaplah bila kamu mau dan pergilah jika kamu bosan...",
  "Aku bingung, apakah aku harus kecewa atu tidak? Jika aku kecewa, emang siapa diriku baginya?\n\nKalau aku tidak kecewa, tapi aku menunggu ucapannya.",
  "Rinduku seperti ranting yang tetap berdiri.Meski tak satupun lagi dedaunan yang menemani, sampai akhirnya mengering, patah, dan mati.",
  "Kurasa kita sekarang hanya dua orang asing yang memiliki kenangan yang sama.",
  "Buatlah aku bisa membencimu walau hanya beberapa menit, agar tidak terlalu berat untuk melupakanmu.",
  "Aku mencintaimu dengan segenap hatiku, tapi kau malah membagi perasaanmu dengan orang lain.",
  "Mencintaimu mungkin menghancurkanku, tapi entah bagaimana meninggalkanmu tidak memperbaikiku.",
  "Kamu adalah yang utama dan pertama dalam hidupku. Tapi, aku adalah yang kedua bagimu.",
  "Jika kita hanya bisa dipertemukan dalam mimpi, aku ingin tidur selamanya.",
  "Melihatmu bahagia adalah kebahagiaanku, walaupun bahagiamu tanpa bersamaku.",
  "Aku terkadang iri dengan sebuah benda. Tidak memiliki rasa namun selalu dibutuhkan. Berbeda dengan aku yang memiliki rasa, namun ditinggalkan dan diabaikan...",
  "Bagaimana mungkin aku berpindah jika hanya padamu hatiku bersinggah?",
  "Kenangan tentangmu sudah seperti rumah bagiku. Sehingga setiap kali pikiranku melayang, pasti ujung-ujungnya akan selalu kembali kepadamu.",
  "Kenapa tisue bermanfaat? Karena cinta tak pernah kemarau. - Sujiwo Tejo",
  "Kalau mencintaimu adalah kesalahan, yasudah, biar aku salah terus saja.",
  "Sejak kenal kamu, aku jadi pengen belajar terus deh. Belajar jadi yang terbaik buat kamu.",
  "Ada yang bertingkah bodoh hanya untuk melihatmu tersenyum. Dan dia merasa bahagia akan hal itu.",
  "Aku bukan orang baik, tapi akan belajar jadi yang terbaik untuk kamu.",
  "Kita tidak mati, tapi lukanya yang membuat kita tidak bisa berjalan seperti dulu lagi.",
  "keberadaanmu bagaikan secangkir kopi yang aku butuhkan setiap pagi, yang dapat mendorongku untuk tetap bersemangat menjalani hari.",
  "Aku mau banget ngasih dunia ke kamu. Tapi karena itu nggak mungkin, maka aku akan kasih hal yang paling penting dalam hidupku, yaitu duniaku.",
  "Mending sing humoris tapi manis, ketimbang sok romantis tapi akhire tragis.",
  "Ben akhire ora kecewa, dewe kudu ngerti kapan waktune berharap lan kapan kudu mandeg.",
  "Aku ki wong Jowo seng ora ngerti artine 'I Love U'. Tapi aku ngertine mek 'Aku tresno awakmu'.",
  "Ora perlu ayu lan sugihmu, aku cukup mok setiani wes seneng ra karuan.",
  "Cintaku nang awakmu iku koyok kamera, fokus nang awakmu tok liyane mah ngeblur.",
  "Saben dino kegowo ngimpi tapi ora biso nduweni.",
  "Ora ketemu koe 30 dino rasane koyo sewulan.",
  "Aku tanpamu bagaikan sego kucing ilang karete. Ambyar.",
  "Pengenku, Aku iso muter wektu. Supoyo aku iso nemokne kowe lewih gasik. Ben Lewih dowo wektuku kanggo urip bareng sliramu.",
  "Aku ora pernah ngerti opo kui tresno, kajaba sak bare ketemu karo sliramu.",
  "Cinta aa ka neng moal leungit-leungit sanajan aa geus kawin deui.",
  "Kasabaran kaula aya batasna, tapi cinta kaula ka anjeun henteu aya se epna.",
  "Kanyaah akang moal luntur najan make Bayclean.",
  "Kenangan endah keur babarengan jeung anjeun ek tuluy diinget-inget nepi ka poho.",
  "Kuring moal bakal tiasa hirup sorangan, butuh bantosan jalmi sejen.",
  "Nyaahna aa ka neg teh jiga tukang bank keur nagih hutang (hayoh mumuntil).",
  "Kasabaran urang aya batasna, tapi cinta urang ka maneh moal aya beakna.",
  "Hayang rasana kuring ngarangkai kabeh kata cinta anu aya di dunya ieu, terus bade ku kuring kumpulkeun, supaya anjeun nyaho gede pisan rasa cinta kuring ka anjeun.",
  "Tenang wae neng, ari cinta Akang mah sapertos tembang krispatih; Tak lekang oleh waktu.",
  "Abdi sanes jalmi nu sampurna pikeun anjeun, sareng sanes oge nu paling alus kanggo anjeun. Tapi nu pasti, abdi jalmi hiji-hijina nu terus emut ka anjeun.",
  "Cukup jaringan aja yang hilang, kamu jangan.",
  "Sering sih dibikin makan ati. Tapi menyadari kamu masih di sini bikin bahagia lagi.",
  "Musuhku adalah mereka yang ingin memilikimu juga.",
  "Banyak yang selalu ada, tapi kalo cuma kamu yang aku mau, gimana?",
  "Jam tidurku hancur dirusak rindu.",
  "Cukup China aja yang jauh, cinta kita jangan.",
  "Yang penting itu kebahagiaan kamu, aku sih gak penting..",
  "Cuma satu keinginanku, dicintai olehmu..",
  "Aku tanpamu bagaikan ambulans tanpa wiuw wiuw wiuw.",
  "Cukup antartika aja yang jauh. Antarkita jangan.",
  "-Kehilangan adalah rencana semesta agar manusia belajar bijaksana Kehilangan adalah rencana semesta agar manusia belajar bijaksana*\n\n\n*Quotes By:rdlvy",
"-bukan aku yg menyesal kehilangammu, tapi kmu yg akan menyesal kehilangan orang sebaik aku*\n\nQuotes By Unknow",
"-aku tak pernah menyesal mencintaimu tpi yg aku sesali adalah terlalu berharap kau adalah takdir nyatanya kau hanya hadir*\n\n*Quotes By unknow",
"-lepaskan apa yg bukan milikmu, ikhlaskan apa yg sudah tidak lagi ditempatnya, merelakan pertanyaan yg ga akan pernah ada jawabannya, menerima bahwa kadang ga semua orang  bisa bikin kamu bahagia*\n\n*Quotes By unknow",
"-jika memang komitmen, dia takkan pernah berubah kepadamu. semakin lama hubungan dan semakin tau keadaan bukan semakin bosan tetapi semakin sayang dan mau menerima kekurangan, bukan justru semakin hilang dan banyak alasan*\n\n*Quotes By unknow",
"-dan sekarang kita berjalan dijalan kita sendiri. saatnya saya mengarungi lautan, menghadapi pasang surut dan angin saya akan mencari alasan lain untuk hidup, kamu menghancurkanku tapi aku akan berusaha baik baik saja tanpamu*\n\n*Quotes By unknow",
"-aku tidak memaksamu bertahan, jika aku tak membahagiakan saja silahkan*\n\n*Quotes By unknow",
"-hatiku pernah mati berkali kali, dihujam oleh orang orang terbaik yg dulu pernah memupuk arti dan janji, lalu meninggalkan bila sudah tak berarti*\n\n*Quotes By unkow",
"-tetap tumbuh walau dunia membunuh, tetap melangkah meski jejak sudah kehilangan arah*\n\n*Quotes By unknow",
"-kenapa masih dia? karna sampai saat ini saya belum mencintai seseorng sekuat saya mencintai dia*\n\n*Quotes By NN",
"-aku bukan hanya takut kehilanganmu, tapi aku takut kamu berpura pura mencintaiku*\n\n*Quotes By ~NN",
"-patah hatinya ke satu orang mati rasanya ke semua orang*\n\n Quotes By unknown",
"-Sebaik apapun kamu, akan tetap salah dimata orng yg salah. tetapi jika seseorang benar mencintaimu, sebanyak apapun kekuranganmu akan tetap dianggap terbaik olehnya*\n\n*Quotes By ~NN",
"-khalil gibran pernah berkata jika kamu mencintai seseorang, biarkan ia pergi. kalau ia kembali berarti ia adalah milikmu kalau ia tidak kembali, maka ia memang tidak pernah menjadi milikmu",
"-kita jalani ini apa adanya. tak mengikat tak terikat, kita hidup bersama jika tidak sama sama kita terima*\n\n*Quotes By ~NN",
"-dia yg bilang takut ditinggalkan justru malah dia yg meninggalkan*\n\n*Quotes By ~NN",
"-kamu beruntung jika sudah menemukan seseorang yg tidak pernah bosan untuk mengerti segala sikapmu yg terkadang menyebalkan dan menjengkelkan*\n\n*Quotes By ~NN",
"-tidak ada yg bisa menjamin dia berubah meskipun kamu banyak berkorban. tidak ada yg dapat menjamin dia berhenti menyakiti meskipun kamu sudah menerima dan memaafkannya berkali kali. karena perubahan itu datangnya dari diri sendiri*\n\n*Quotes By NN",
"-anda kehilangan saya?, saya sudah lebih dulu kehilangan anda bukan saya yg berubah tapi anda*\n\n*Quotes By ~NN",
"-jangan karena kamu tau dia begitu mencintaimu kamu jadi seenaknya. jangan karna kamu tau dia selalu memaafkan kesalahanmu kamu terus berulah melakukan kesalahan yg sama. ingat jika dia sudah sampai di titik lelah kemudian memutuskan pergi mungkin untukmu dia tidak memberi kesempatan lagi. maka selagi memiliki jaga dengan baik dan sepenuh hati*\n\n*Quotes By ~NN",
"-terkadang meski kita sudah ikhlas melepaskan seseorang yg dicintai ada saat saat tertntu kita mulai mengingatnya lagi membayangkan kejadian yg membuat dada kita terasa sakit lagi. sebab yg namanya luka tidak pernah bisa sembuh dengan sempurna tetap akan ada bekas disana bagaimanapun kita menutupi dan mengobatinya*\n\n*Quotes By ~NN",
"-apa bukti bahwa dia mencintai kita dengan tulus? ketika kamu sudah mematahkan hatinya dia masih tetap berkeinginan untuk membahagiakanmu dan dia tidak akan mudah untuk pergi meninggalkanmu*\n\n*Quotes By ~NN",
"-semacam diberi harapan namun tidak diprioritaskan. semacam dipertahankan namun tidak diperhatikan*\n\n*Quotes By ~NN",
"-pertama wanita akan memperlakukanmu sebagaimna dia ingin diperlakukan. lalu wanita akan memperlakukanmu sebagaimana kamu memperlakukan mereka*\n\n*Quotes By ~NN",
"-hilangku tidak dicari hadirku tak dinanti pergiku tak ditahan kembaliku tak diharapkan. ternyata mencintai sendirian rasanya menyakitkan*\n\n*Quotes By ~NN",
"-perempuan pasti akan cemburu ketika melihat lelakinya dekat dengan perempuan lain meski dia temen sekalipun. munafik jika perempuan itu mengatakan dia baik² saja*\n\n*Quotes By ~NN",
"-bohong itu hak kamu dan tugasku hanya pura² diam*\n\n*Quotes By NN",
"-kalo bosen bilang jangan diem, biar aku pergi sendiri tanpa kamu suruh*\n\n*Quotes By NN",
"-aku tau bahwa kamu berpura² mencintaiku*\n\n*Quotes By NN"
]
					bucin1 = bucin[Math.floor(Math.random() * bucin.length)]
					reply(`${bucin1}`)
					break
				case prefix+'tomp3':
				case prefix+'kemp3':
				case prefix+'toaudio':
                	
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                	client.updatePresence(from, Presence.composing) 
					if (!isQuotedVideo) return reply('tag / geser videonya um ')
					reply(`tunggu...\nJika tidak dibalas, berarti error`)
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.mp4')
					exec(`ffmpeg -i ${media} ${ran}`, (err) => {
						fs.unlinkSync(media)
						if (err) return reply('Maaf, gagal pada saat mengkonversi video ke mp3 ')
						bufferlkj = fs.readFileSync(ran)
						client.updatePresence(from, Presence.recording)
						client.sendMessage(from, bufferlkj, audio, {mimetype: 'audio/mp4', quoted: mek})
						fs.unlinkSync(ran)
					})
					await limitAdd(sender)
					break
				case `${prefix}setppbot`:
				case `${prefix}setbotpp`:
				case `${prefix}nggopp`:
				case `${prefix}pakaipp`:
				case `${prefix}gantipp`:
				case `${prefix}ubahpp`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
				    client.updatePresence(from, Presence.composing) 
					if (!isQuotedImage) return reply(`Kirim gambar lalu geser gambar yang sudah dikirim lalu ketik ${prefix}ubahpp`)
					enmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(enmedia)
					await client.updateProfilePicture(botNumber, media)
					reply('Makasih profil barunya')
					await limitAdd(sender)
					break 
				case `${prefix}roboguru`:
						
                    if (!isRegistered) return reply(ind.noregis())
						if (isBanned) return reply(ind.diban())
						if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						if (isNggoRoboguru) return reply(`Untuk Menghindari Kerusakan Jawaban\nSilahkan Coba Beberapa detik lagi\n\n_pesan ini muncul karena kami melayani orang lain_`)
						if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
						encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						media = await client.downloadAndSaveMediaMessage(encmedia)
								kebotguru = fs.readFileSync(`./${media}`)
								roboguru2 = `${encmedia.message.imageMessage.caption}`
								client.sendMessage(`${nomereroboguru}`, kebotguru, image, {caption: `${roboguru2}`, quoted: { key: { fromMe: false, participant: `${sender}`, ...(from ? { remoteJid: tagstatus } : {}) }, message: { conversation: `Dari ${sender.split("@")[0]}` }}})
							} else {
								client.sendMessage(`${nomereroboguru}`, `${body.slice(10)}`, text, {quoted: { key: { fromMe: false, participant: `${sender}`, ...(from ? { remoteJid: tagstatus } : {}) }, message: { conversation: `Dari ${sender.split("@")[0]}` }}})
							}
								nggoroboguru = sender
								fromnggoroboguru = from
							setTimeout( () => {
								nggoroboguru = me.jid
								fromnggoroboguru = me.jid
							}, 10000) // 1000 = 1s,
							await limitAdd(sender)
							break
					case `${prefix}pln`:
						
                   	 if (!isRegistered) return reply(ind.noregis())
						if (isBanned) return reply(ind.diban())
						if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						if (isNggoPln) return reply(`Untuk Menghindari Kerusakan Pesan\nSilahkan Coba Beberapa detik lagi\n\n_pesan ini muncul karena kami melayani orang lain_`)
						if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
						encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						media = await client.downloadAndSaveMediaMessage(encmedia)
								kepln = fs.readFileSync(`./${media}`)
								pln2 = `${encmedia.message.imageMessage.caption}`
								client.sendMessage(`${nomerepln}`, kepln, image, {caption: `${pln2}`, quoted: { key: { fromMe: false, participant: `${sender}`, ...(from ? { remoteJid: from } : {}) }, message: { conversation: `Dari ${sender.split("@")[0]}` }}})
							} else {
								client.sendMessage(`${nomerepln}`, `${body.slice(5)}`, text, {quoted: { key: { fromMe: false, participant: `${sender}`, ...(from ? { remoteJid: from } : {}) }, message: { conversation: `Dari ${sender.split("@")[0]}` }}})
							}
								nggopln = sender
								fromnggopln = from
							setTimeout( () => {
								nggopln = me.jid
								fromnggopln = me.jid
							}, 10000) // 1000 = 1s,
							await limitAdd(sender)
							break
					case `${prefix}brainly`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    brien = body.slice(9)
					brainly(`${brien}`).then(res => {
					teks = '\n'
					for (let Y of res.data) {
						teks += `\n* _BRAINLY_ *\n\n* Pertanyaan:* ${Y.pertanyaan}\n\n* Jawaban:* ${Y.jawaban[0].text}\n\n`
					}
					client.sendMessage(from, teks, text, {quoted: mek, detectLinks: false})
                        console.log(res)
                    })
					await limitAdd(sender)
					break 
				case prefix+'buatstatus':
                    if (!isRegistered) return reply(ind.noregis())
				    if (!isOwner) return reply(ind.ownerb())
					client.sendMessage('status@broadcast', body.slice(12), text, {quoted: mek})
					fitnah('status@broadcast', 'sudah', ' ')
					break 
				case `${prefix}bcgc`:
                    if (!isRegistered) return reply(ind.noregis())
				     if (!isOwner) return reply(ind.ownerb())
					if (args.length < 1) return reply('.......')
					anu = await groupMembers
					nom = mek.participant
					if (isMedia && !mek.message.videoMessage || isQuotedImage) {
						encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						media = await client.downloadAndSaveMediaMessage(encmedia)
						for (let _ of anu) {
							client.sendMessage(_.jid, media, image, {caption: `* BC GROUP *\n\nDari Grup : ${groupName}\nPengirim : wa.me/${(sender.split('@')[0])}\nPesan : ${body.slice(6)}`})
						}
						reply('')
					} else {
						for (let _ of anu) {
							sendMess(_.jid, `* BC GROUP *\n\nDari Grup : ${groupName}\nPengirim : wa.me/${(sender.split('@')[0])}\nPesan : ${body.slice(6)}`)
						}
						reply('Sukses broadcast group')
					}
					break 
				case `${prefix}resep`:
                    if (!isRegistered) return reply(ind.noregis())
                   anu = await fetchJson(`https://mnazria.herokuapp.com/api/resep?key=${q}`, {method: 'get'})
                   if (anu.error) return reply(anu.error)
                   buff = await getBuffer(anu.thumb_item)
                   hasil = `*title* \n ${anu.title} *item_name* \n ${anu.item_name} *ingredient* \n${anu.ingredient} *step* \n${anu.step}`
                   client.sendMessage(from, buff, image, {quoted: mek, caption: hasil})
                   await limitAdd(sender)
					break 
				case `${prefix}play`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					client.updatePresence(from, Presence.composing)
					try {
					play1 = await fetchJson(`https://api.zeks.xyz/api/ytplaymp3?q=${body.slice(6)}&apikey=${ZeksKey}`, {method: 'get'})
					play2 = `*Judul:* ${play1.result.title}\n*Source:* ${play1.result.source}\n*size:* ${play1.result.size}\n\n_tunggu, ini akan memakan waktu lama_`
					playthumb = await getBuffer(play1.result.thumbnail)
					client.sendMessage(from, playthumb, image, {quoted: mek, caption: play2})
					playmp3 = await getBuffer(play1.result.url_audio)
					client.updatePresence(from, Presence.recording)
					client.sendMessage(from, playmp3, audio, {mimetype: 'audio/mpeg', filename: `${play1.result.title}.mp3`, quoted: mek})
					} catch (e) {
                        console.error('ERROR\nSilahkan ganti judul nya')
                        reply(`ERROR\nSilahkan ganti judul nya\n\n${e}`)
                    }
					break
				case `${prefix}playvideo`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length > 0) {
					playv1 = await fetchJson(`https://api.zeks.xyz/api/ytplaymp4?q=${body.slice(11)}&apikey=${ZeksKey}`, {method: 'get'})
					playv2 = `*Judul:* ${playv1.result.title}\n*Source:* ${playv1.result.source}\n*size:* ${playv1.result.size}\n\n_tunggu_`
					playvthumb = await getBuffer(playv1.result.thumbnail)
					client.sendMessage(from, playvthumb, image, {quoted: mek, caption: playv2})
					playmp4 = await getBuffer(playv1.result.url_video)
					client.sendMessage(from, playmp4, video, {mimetype: 'video/mp4', quoted: mek, filename: `${playv1.result.title}.mp4`, caption: `${playv1.result.title}\n${playv1.result.source}`})
					} else {
                        console.error('ERROR\nSilahkan ganti judul nya')
                        reply(`ERROR\nSilahkan ganti judul nya`)
                        }
					break
				case prefix+'emoji':
                    if (!isRegistered) return reply(ind.noregis())
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					emoji.get(args[0]).then(emoji => {
    				stiker(emoji.images[4].url)
					reply(emoji.images[4].url)
   				 })
					break
				case prefix+'emoji':
                    if (!isRegistered) return reply(ind.noregis())
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					stiker(`https://api.zeks.xyz/api/emoji-image?apikey=${ZeksKey}&emoji=${encodeURIComponent(args[0])}`)
					break
				case prefix+'memeindo':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					memein = await kagApi.memeindo()
					bufferll = await getBuffer(`https://imgur.com/${memein.hash}.jpg`)
					client.sendMessage(from, bufferll, image, {quoted: mek, caption: '!sticker'})
					await limitAdd(sender)
					break
				case prefix+'memeindo':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					memein = await fetchJson(`https://api.zeks.xyz/api/memeindo?apikey=benbenz`, {method: 'get'})
					buffermemein = await getBuffer(memein.result)
					client.sendMessage(from, buffermemein, image, {quoted: mek, caption: '!sticker'})
					await limitAdd(sender)
					break
				case `${prefix}infogrup`:
				case `${prefix}grupinfo`:
				case `${prefix}infogroup`:
				case `${prefix}groupinfo`:
                    if (!isRegistered) return reply(ind.noregis())
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                	if (!isGroup) return reply(ind.groupo())
                	client.updatePresence(from, Presence.composing)
					try {
						ppimg = await client.getProfilePicture(from)
					} catch {
						ppimg = 'https://drive.google.com/u/0/uc?id=1zGz2KQiZCEYBWhgIhoZnUKTqqA2-hWfi&export=download'
					}
					let buf = await getBuffer(ppimg)
					teks = (args.length > 1) ? body.slice(8).trim() : ''
					teks += `${tanda}\n*Nama grup :* ${groupName}\n*Deskripsi :* ${groupDesc}\n*Jumlah Admin :* ${groupAdmins.length}\n*Jumlah anggota :* ${groupMembers.length}\n\n*Admin:*`
					no = 0
					for (let admon of groupAdmins) {
						no += 1
						teks += `@${admon.split('@')[0]}\n`
						contextInfo: { mentionedJid: [admon] }
					}
					for (let mem of groupMembers) {
						teks += `** ${mem.jid.split('@')[0]}\n`
						anggota_id = []
						anggota_id.push
						contextInfo: { mentionedJid: [mem] }
					}
		        	client.sendMessage(from, buf, image, {quoted: mek, caption: `${teks}\n\n${anggota_id}`})
                	break
				case `${prefix}kickall`:
                    
                    if (!isRegistered) return reply(ind.noregis())
                    if (!isOwner) return reply(ind.ownerb())
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    if (!isBotGroupAdmins) return reply(ind.badmin())
			        members_id = []
					teks = (args.length > 1) ? body.slice(8).trim() : ''
					teks += '\n\n'
					for (let mem of groupMembers) {
						teks += `** ${mem.jid.split('@')[0]}\n`
						members_id.push(mem.jid)
					}
					mentions(teks, members_id, true)
					client.groupRemove(from, members_id)
					break 
				case `${prefix}hapus`:
				case `${prefix}delete`:
				case `${prefix}delet`:
				case `${prefix}del`:
				case `hapus`:
				case `delete`:
                    if (!isRegistered) return reply(ind.noregis())
                   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                   try {
					client.deleteMessage(from, { id: mek.message.extendedTextMessage.contextInfo.stanzaId, remoteJid: from, fromMe: true })
					await limitAdd(sender)
					} catch {
						reply(`Geser / tag pesan ku untuk menghapus nya`)
					}
					break
				case `${prefix}setreply`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isOwner) return reply(ind.ownerb())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    client.updatePresence(from, Presence.composing) 
					if (args.length < 1) return
					cr = body.slice(10)
					reply(`reply berhasil di ubah menjadi : ${cr}`)
					await limitAdd(sender)
					break 
				case `${prefix}grouplist`:
				case `${prefix}grupmu`:
                    if (!isRegistered) return reply(ind.noregis())
                   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					client.updatePresence(from, Presence.composing) 
					teks = `\`\`\`Ini adalah list group FRM BOT :\n\n\`\`\``
					no = 0
					for (let hehehe of groupId) {
						no += 1
						teks += `\`\`\`[${no.toString()}]\`\`\` @${hehehe.split('@')[0]}\n`
					}
					teks += `\n\`\`\`Total grup : ${groupId.length}\`\`\``
					client.sendMessage(from, teks.trim(), extendedText, {quoted: mek})
					await limitAdd(sender)
					break
			case `${prefix}daftar`:
                if (isRegistered) return  reply(ind.rediregis())
                if (!q.includes('|')) return  reply(`${prefix}daftar Rifki|16\n*itu contohnya*`)
                const namaUser = q.substring(0, q.indexOf('|') - 0)
                const umurUser = q.substring(q.lastIndexOf('|') + 1)
                const serialUser = createSerial(20)
                veri = sender 
                if (args[0].startsWith('|')) {
                	daftar = body.slice(8)
                	namaUser1 = daftar.split("|")[1];
               	 umurUser1 = daftar.split("|")[2];
                	serialUser1 = createSerial(20)
                    addRegisteredUser(sender, namaUser1, umurUser1, time, serialUser1)
                    await reply(ind.registered(namaUser1, umurUser1, serialUser1, time, sender))
                    addATM(sender)
                    addLevelingId(sender)
                    console.log(color('[REGISTER]'), color(time, 'yellow'), 'Name:', color(namaUser1, 'cyan'), 'Age:', color(umurUser1, 'cyan'), 'Serial:', color(serialUser1, 'cyan'))
                } else if (!args[0].startsWith('|')) {
                	forme = body.slice(8)
                	namaUser2 = forme.split("|")[0];
               	 umurUser2 = forme.split("|")[1];
                	serialUser2 = createSerial(20)
                    addRegisteredUser(sender, namaUser2, umurUser2, time, serialUser2)
                    await reply(ind.registered(namaUser2, umurUser2, serialUser2, time, sender))
                    addATM(sender)
                    addLevelingId(sender)
                    console.log(color('[REGISTER]'), color(time, 'yellow'), 'Name:', color(namaUser2, 'cyan'), 'Age:', color(umurUser2, 'cyan'), 'Serial:', color(serialUser2, 'cyan'))
                }
					break
			case prefix+'hilih':
			case prefix+'huluh':
			case prefix+'halah':
			case prefix+'heleh':
			case prefix+'holoh':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
				hlh = encmedia.message.conversation || encmedia.message.imageMessage.caption || encmedia.message.videoMessage.caption || encmedia.message.extendedTextMessage.text
				ter = command[2].toLowerCase()
				reply(`${hlh.replace(/[aiueo]/g, ter).replace(/[AIUEO]/g, ter.toUpperCase())}`)
				await limitAdd(sender)
				break
			case prefix+'totalhuruf':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
				ttlhrf = encmedia.message.conversation || encmedia.message.imageMessage.caption || encmedia.message.videoMessage.caption || encmedia.message.extendedTextMessage.text
				totalhrf = await fetchJson(`https://videfikri.com/api/jumlahhuruf/?query=${encodeURIComponent(ttlhrf)}`, {method: 'get'})
				reply(`_Jumlah karakter pada pesan tersebut aaaadalaaaaah_\n*${totalhrf.result.jumlah}* karakter`)
				await limitAdd(sender)
				break
			case prefix+'getstickererr':
			case prefix+'getstikererr':
                    if (!isRegistered) return reply(ind.noregis())
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    getstik1 = await fetchJson(`https://api.zeks.xyz/api/searchsticker?apikey=${ZeksKey}&q=${q}`, {method: 'get'})
                    for (let i = 0; i < getstik1.sticker.length; i++) {
                    ambilstikere = await getBuffer(getstik1.sticker[i])
                    exec(`cwebp -q 75 ambilstikere -o ./sampah/getstiker_${sender}.webp`)
					.then(() => {ambilstiker = fs.readFileSync(`./sampah/getstiker_${sender}.webp`)
                    client.sendMessage(from, ambilstiker, sticker, {quoted: mek})
                    })
                    }
					await limitAdd(sender)
                    break
          case prefix+'spam':
          	if (!isRegistered) return reply(ind.noregis())
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    for (let i = 0; i < body.slice(6); i++) {
					reply(`hmm`)
				}
				await limitAdd(sender)
                    break
          case `${prefix}speed`:
          case `${prefix}ping`:
          	
                    if (!isRegistered) return reply(ind.noregis())
                   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			timestamp = speed()
            latensi = speed() - timestamp
            reply(`*Kecepatan internet:* ${latensi.toFixed(4)} detik\n\nINFO: lebih kecil lebih cepat`)
            await limitAdd(sender)
					break
			case `${prefix}donasi`:
			case `${prefix}donate`:
			case `${prefix}menyumbang`:
                    if (!isRegistered) return reply(ind.noregis())
					reply(donasi)
					client.sendMessage(from, fs.readFileSync(`./fauzan.rifki.m/qrdanafrm.webp`), sticker, {quoted:mek})
					break
				case prefix+'runtime':
					if (!isRegistered) return reply(ind.noregis())
                    if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					client.updatePresence(from, Presence.composing) 
					uptime = process.uptime()
					reply(`Umur bot sekarang adalah\n*${kyun(uptime)}*`)
					await limitAdd(sender)
				break
				case `${prefix}info`:
                    if (!isRegistered) return reply(ind.noregis())
         	       if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					timestamp = speed()
                    latensi = speed() - timestamp
					uptime = process.uptime()
					weton = ['Pahing', 'Pon','Wage','Kliwon','Legi']
					hasilweton = weton[Math.floor(((d * 1) + gmt) / 84600000) % 5]
					infobotnya = fs.readFileSync('./fauzan.rifki.m/infobotnya.jpg')
					infonggocmd = fs.readFileSync('./fauzan.rifki.m/infonggocmd.jpg')
					infonggowa = fs.readFileSync('./fauzan.rifki.m/infonggowa.jpg')
//INFO BOTNYA
					teksinfobotnya = `⏰ ${jam} WIB
🗓️ ${hari} ${hasilweton}
╰> ${tanggal}

▬▭▬▭▬▭▬▭▬▭▬▭▬
*INFO BOT*

*❀ Nama akun wa bot:*
╰> ${me.name}
*❀ Pembuat:*
╰> Fauzan Rifki Maulana
*❀ Nomor Bot:*
╰> ${me.jid.split('@')[0]}
*❀ Bosku:*
╰> ${nomowner}
*❀ Prefix:*
╰> ${prefix}
*❀ Kontak terblokir:*
╰> ${blocked.length}
kontak terblokir kadang error
*❀ Total Chat:*
╰> ${totalchat.length}
*❀ Total pengguna:*
╰> ${_registered.length}
*❀ Lama bot aktif:*
╰> ${kyun(uptime)}
*❀ PP Bot:*
╰> ${me.imgUrl}

*Follow igku*
instagram.com/frm_developer
*Facebook*
fb.com/fauzan.rifki.m
*Telegram*
t.me/frm_developer
▬▭▬▭▬▭▬▭▬▭▬▭▬
*Dibanned ${dibanned} kali, sebab*
${prefix}kickall
${prefix}bc`

//INFO PERANGKAT WA
					teksinfonggowa = `${head1} Merek Perangkat: ${head2}
_${me.phone.device_manufacturer}_
*❀ Model Perangkat: ❀*
_${me.phone.device_model}_
*❀ Versi OS: ❀*
_${me.phone.os_version}_
*❀ Nomor Versi OS: ❀*
_${me.phone.os_build_number}_
*❀ MCC: ❀*
_${me.phone.mcc}_
*❀ MNC: ❀*
_${me.phone.mnc}_
*❀ Versi WhatsApp: ❀*
_${me.phone.wa_version}_
*❀ Sisa baterai: ❀*
_${sisabaterai}_
*❀ Penghemat Baterai: ❀*
_${hematdaya}_
*❀ Dicas: ❀*
_${dicas}_
*NB:* jika _true_ berarti ya
jika _false_ berarti tidak`
					
//INFO TERMINAL
client.sendMessage(from, infonggocmd, image, { quoted: mek, caption: `*❀ Kecepatan Internet:*
_${latensi.toFixed(4)} detik_
   _(lebih besar lebih lambat)_
*❀ Sistem Operasi:*
_${os.platform()}_
*❀ RAM*:
Dipakai Shell / Terminal:
_${(process.memoryUsage().heapUsed / 1024 / 1024).toFixed(2)} MB_
Total: _${Math.round(os.totalmem / 1024 / 1024)} MB_
*❀ MEM:*
Sisa: _${os.freemem()}_
Total: _${os.totalmem()}_
*❀ CPU:*
Model: _${os.cpus()[0].model}_
Speed: _${os.cpus()[0].speed}_` })
					client.sendMessage(from, infobotnya, image, { quoted: mek, caption: teksinfobotnya })
					client.sendMessage(from, infonggowa, image, { quoted: mek, caption: teksinfonggowa })
					console.log(me)
					await limitAdd(sender)
					break
				case `${prefix}blocklist`: 
				case `${prefix}listblock`:
                    if (!isRegistered) return reply(ind.noregis())
             	   if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					teks = 'KONTAK TERBLOKIR:\n'
					for (let block of blocked) {
						teks += ` @${block.split('@')[0]}\n`
					}
					teks += ` : ${blocked.length}`
					client.sendMessage(from, teks.trim(), extendedText, {quoted: mek, contextInfo: {"mentionedJid": blocked}})
					await limitAdd(sender)
					break
                case `${prefix}hidetag`:
            	if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isGroup) return reply(ind.groupo())
					var value = body.slice(9)
					var group = await client.groupMetadata(from)
					var member = group['participants']
					var mem = []
					member.map( async adm => {
					mem.push(adm.id.replace('c.us', 's.whatsapp.net'))
					})
					var hidetage = {
					text: value,
					contextInfo: { mentionedJid: mem },
					quoted: mek
					}
					client.sendMessage(from, hidetage, text)
					await limitAdd(sender)
					break
                case `${prefix}quotemaker`:
            	if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                var gh = body.slice(12)
					var quote = gh.split("|")[0];
					var wm = gh.split("|")[1];
					 pref = `Usage: \n${prefix}quotemaker teks|watermark\n\nEx :\n${prefix}quotemaker ini contoh|bicit`
					if (args.length < 1) return reply(pref)
					reply(ind.wait())
					anu = await fetchJson(`https://terhambar.com/aw/qts/?kata=${quote}&author=${wm}&tipe=random`, {method: 'get'})
					buffer = await getBuffer(anu.result)
					client.sendMessage(from, buffer, image, {caption: 'Nih anjim', quoted: mek})
					await limitAdd(sender)
					break				
				case `${prefix}ssweb`:
            	if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Urlnya mana om')
					reply(ind.wait())
					ssweb = await fetchJson(`https://api.zeks.xyz/api/ssweb?url=${body.slice(7)}&apikey=${ZeksKey}`)
					if (ssweb.message) return reply(ssweb.message)
					sswebhasil = await getBuffer(ssweb.result)
					client.sendMessage(from, sswebhasil, image, {quoted: mek, caption: `${tanda}\n${body.slice(7)}`})
					await limitAdd(sender)
					break
                case `${prefix}pokemon`:
            	if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.fdci.se/rep.php?gambar=pokemon`, {method: 'get'})
					reply(ind.wait())
					var n = JSON.parse(JSON.stringify(anu));
					var nimek =  n[Math.floor(Math.random() * n.length)];
					pok = await getBuffer(nimek)
					client.sendMessage(from, pok, image, { quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'gay':
                    if (!isRegistered) return reply(ind.noregis())
                	if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                	const deskgay =["Mungkin sudah ada 1-2 korban!","Berjiwa gay tetapi tidak membabi buta!","WOAKEOAWKOEKAW KABOOORRRRR!!! KALAU INI JANGANKAN BOOLMU, KNALPOT AJA DISODOK!","Jujur lo udah berapa banyak korban"]
                	const deskegay = deskgay[Math.floor(Math.random() * deskgay.length)]
                	const persengay =['1','2','3','4','5','6','7','8','9','10','11','12','13','14','15','16','17','18','19','20','21','22','23','24','25','26','27','28','29','30','31','32','33','34','35','36','37','38','39','40','41','42','43','44','45','46','47','48','49','50','51','52','53','54','55','56','57','58','59','60','61','62','63','64','65','66','67','68','69','70','71','72','73','74','75','76','77','78','79','80','81','82','83','84','85','86','87','88','89','90','91','92','93','94','95','96','97','98','99','100']
					const persenegay = persengay[Math.floor(Math.random() * persengay.length)]
      				reply(`*GAY LU*\n*Persentase* : ${persenegay}%\n${deskegay}`)
      				break
                case `${prefix}anjing`:
            	if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.fdci.se/rep.php?gambar=anjing`, {method: 'get'})
					reply(ind.wait())
					var n = JSON.parse(JSON.stringify(anu));
					var nimek =  n[Math.floor(Math.random() * n.length)];
					pok = await getBuffer(nimek)
					client.sendMessage(from, pok, image, { quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'acakquran':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					client.updatePresence(from, Presence.composing)
					acakqurane = await fetchJson(`https://api.zeks.xyz/api/randomquran`, {method: 'get'})
					reply(`${tanda}\n*Nama:* ${acakqurane.result.asma}/${acakqurane.result.nama}\n*Nomor surah:* ${acakqurane.result.nomor}\n*Arti:* ${acakqurane.result.arti}\n*keterangan:* ${acakqurane.result.keterangan}\n\n_tunggu, sendang mengirim suara orang baca Al Quran_`)
					suaraquran = await getBuffer(acakqurane.result.audio)
					client.updatePresence(from, Presence.recording)
					client.sendMessage(from, suaraquran, audio, {mimetype: 'audio/mpeg', filename: `${acakqurane.result.asma}.mp3`, quoted: mek})
					await limitAdd(sender)
					break
				case prefix+'ytmp4':
                case 'ytmp4':
                   case 'ytv':
			if (args.length === 0) return reply(`Kirim perintah *${prefix}ytmp4 [linkYt]*`)
			let isLinks2 = args[0].match(/(?:https?:\/{2})?(?:w{3}\.)?youtu(?:be)?\.(?:com|be)(?:\/watch\?v=|\/)([^\s&]+)/)
			if (!isLinks2) return reply(mess.error.link)
				try {
				reply(mess.wait)
				ytv(args[0])
				.then((res) => {
				const { dl_link, thumb, title, filesizeF, filesize } = res
				axios.get(`https://tinyurl.com/api-create.php?url=${dl_link}`)
				.then((a) => {
				if (Number(filesize) >= 40000) return sendMediaURL(from, thumb, `*YTMP 4!*\n\n*Title* : ${title}\n*Ext* : MP4\n*Filesize* : ${filesizeF}\n*Link* : ${a.data}\n\n_Untuk durasi lebih dari batas disajikan dalam mektuk link_`)
				const captionsYtmp4 = `*Data Berhasil Didapatkan!*\n\n*Title* : ${title}\n*Ext* : MP4\n*Size* : ${filesizeF}\n\n_Silahkan tunggu file media sedang dikirim mungkin butuh beberapa menit_`
				sendMediaURL(from, thumb, captionsYtmp4)
				sendMediaURL(from, dl_link).catch(() => reply(mess.error.link))
				})		
				})
				} catch (err) {
			    reply(mess.error.link) 
				}
				break
                case `${prefix}ytmp4`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Urlnya mana um?')
					if(!isUrl(args[0]) && !args[0].includes('youtu')) return reply('LINK RUSAK')
					ytmp4 = await fetchJson(`https://st4rz.herokuapp.com/api/ytv2?url=${body.slice(7)}`, {method: 'get'})
					dlytmp4 = await getBuffer(ytmp4.result)
					client.sendMessage(from, dlytmp4, video, {mimetype: 'video/mp4', quoted: mek, filename: `${ytmp4.title}`, caption: `${ytmp4.title}`})
					await limitAdd(sender)
					break
				case `${prefix}ytmp3`:
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Urlnya mana um?')
					if(!isUrl(args[0]) && !args[0].includes('youtu')) return reply(ind.wrogf)
					client.updatePresence(from, Presence.recording)
					ytmp3 = await fetchJson(`https://st4rz.herokuapp.com/api/yta2?url=${q}`, {method: 'get'})
					ytmp3dl = await getBuffer(ytmp3.result)
					client.sendMessage(from, ytmp3dl, audio, {mimetype: 'audio/mpeg', quoted: mek})
					await limitAdd(sender)
					break
				case `${prefix}fb`:
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Urlnya mana um?')
					fbdl = await fetchJson(`http://lolhuman.herokuapp.com/api/facebook?apikey=${LolKey}&url=${q}`, {method: 'get'})
					hasilfbdl = await getBuffer(fbdl.result[1].link)
					client.sendMessage(from, hasilfbdl, video, {mimetype: 'video/mp4', quoted: mek, caption: 'nih'})
					await limitAdd(sender)
					break
				case prefix+'fbdl':
					fbdl.getInfo("https://www.facebook.com/111683913906599/posts/195671178841205/")
    				.then(res => {
        			console.log(res)
    				});
    				break
				case prefix+'ig':
				case prefix+'igpost':
				case prefix+'igdl':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Urlnya mana um?')
						igdl = await fetchJson(`https://api.zeks.xyz/api/ig?url=${q}&apikey=${ZeksKey}`, {method: 'get'})
						formatig = igdl.result[0].type
						reply(`${igdl.caption}`)
					if (formatig.includes(`mp4`)) {
						for (let i = 0; i < igdl.result.length; i++) {
                    	hasiligdl = await getBuffer(igdl.result[i].url)
						client.sendMessage(from, hasiligdl, video, {mimetype: 'video/mp4', quoted: mek})
						}
					} else if (formatig.includes(`jpg`)) {
						for (let i = 0; i < igdl.result.length; i++) {
                    	hasiligdl = await getBuffer(igdl.result[i].url)
						client.sendMessage(from, hasiligdl, image, {quoted: mek})
						}
					}
						await limitAdd(sender)
					break
			case prefix+'snack':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply('Urlnya mana gan?')
				if (!isUrl(args[0]) && !args[0].includes('sck')) return reply(mess.error.Iv)
                anu = await fetchJson(`https://api-anoncybfakeplayer.herokuapp.com/sckdown?url=${args[0]}`, {method: 'get'})
               if (anu.error) return reply(anu.error)
                 sck = `「 *SNACK VIDEO DOWNLOADER* 」\n\n*• Format:* ${anu.format}\n*• Size:* ${anu.size}\n\n*TUNGGU SEBENTAR LAGI DIKIRIM MOHON JANGAN SPAM*`
                bufferddd = await getBuffer('https://raw.githubusercontent.com/azizae-official/aebot/main/src/glitchtext.png')
                 reply(mess.wait)
                buff = await getBuffer(anu.result)
                client.sendMessage(from, bufferddd, image, {quoted: mek, caption: sck})
                client.sendMessage(from, buff, video, {mimetype: 'video/mp4', filename: `${anu.format}.mp4`, quoted: mek})
                await limitAdd(sender) 
                break
			case prefix+'soundcloud':
				if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (args.length < 1) return reply('Urlnya mana um?')
				client.updatePresence(from, Presence.composing)
				soundcloud = await fetchJson(`https://api.zeks.xyz/api/soundcloud?apikey=${ZeksKey}&url=${q}`, {method: 'get'})
				thumbsound = await getBuffer(soundcloud.result.thumb)
				client.sendMessage(from, thumbsound, image, {mimetype: 'audio/mp4', quoted: mek, caption: `*Judul:* ${result.title}\n*Durasi:* ${result.duration}\n*Quality:* ${result.quality}`})
				client.updatePresence(from, Presence.recording)
				soundclouddl = await getBuffer(soundcloud.result.download)
				client.sendMessage(from, soundclouddl, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
				await limitAdd(sender)
				break
			case prefix+'tiktok':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Urlnya mana um?')
					reply(`tunggu...`)
					tiktok = await fetchJson(`https://api.vhtear.com/tiktokdl?link=${q}&apikey=${VhtearKey}`, {method: 'get'})
					if (tiktok.result.info) return reply(`${tiktok.result.info}\nAtau mungkin ada orang lain yang pernah download\nBot kan bukan saya doang`)
					hasiltiktok = await getBuffer(tiktok.result.video)
					client.sendMessage(from, hasiltiktok, video, {mimetype: 'video/mp4', quoted: mek, caption: `nih`})
					await limitAdd(sender)
					break
				case prefix+'wp':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('teks nya mana om')
					anwu = await fetchJson(`http://lolhuman.herokuapp.com/api/wallpaper?apikey=${LolKey}&query=${q}`, {method: 'get'})
					bufferx = await getBuffer(anwu.result)
					client.sendMessage(from, bufferx, image, {quoted: mek})
					break
				case prefix+'nulis':
					 if (args.length < 1) return reply('Yang mau di tulis titit kah?')
				 	if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					 reply('WAIT BRO GUE NULIS DULU YAKAN')
					 tulis = await getBuffer(`https://api.vhtear.com/write?text=${q}&apikey=${VhtearKey}`)
					 client.sendMessage(from, tulis, image, {quoted: mek})
										 await limitAdd(sender)
					 break
		 case prefix+'silktext':
				 if (!isRegistered) return reply(ind.noregis())
				 if (isBanned) return reply(ind.diban())
				 if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
				 if (args.length < 1) return reply(ind.wrongf())
				 silk = body.slice(10)
				 if (silk.length > 7) return reply('Teksnya kepanjangan, maksimal 6 karakter')
				 reply(ind.wait())
				 buffer = await getBuffer(`https://api.vhtear.com/silktext?text=${silk}&apikey=${VhtearKey}`)
			 client.sendMessage(from, buffer, image, {quoted: mek})
			 await limitAdd(sender)	
			 break	
			case prefix+'gemboktext':
				 if (!isRegistered) return reply(ind.noregis())
				 if (isBanned) return reply(ind.diban())
				 if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					 var gh = body.slice(12)
					 var gem = gh.split("&")[0];
					 var bok = gh.split("&")[1];
					 if (args.length < 1) return reply(`Contoh : ${prefix}gemboktext 11 01 2021 & Rifki dan Riska`)
					 reply(ind.wait())
					 buffer = await getBuffer(`https://api.vhtear.com/padlock?text1=${gem}&text2=${bok}&apikey=${VhtearKey}`)
					 client.sendMessage(from, buffer, image, {quoted: mek})
					 await limitAdd(sender)
					 break
				case `${prefix}yta`:
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Urlnya mana um?')
					if(!isUrl(args[0]) && !args[0].includes('youtu')) return reply(ind.stikga())
					function post(url, formdata) {
    console.log(Object.keys(formdata).map(key => `${key}=${encodeURIComponent(formdata[key])}`).join('&'))
    return fetch(url, {
        method: 'POST',
        headers: {
            accept: "*/*",
            'accept-language': "en-US,en;q=0.9",
            'content-type': "application/x-www-form-urlencoded; charset=UTF-8"
        },
        body: Object.keys(formdata).map(key => `${key}=${encodeURIComponent(formdata[key])}`).join('&')
    })
}
const ytIdRegex = /(?:http(?:s|):\/\/|)(?:(?:www\.|)youtube(?:\-nocookie|)\.com\/(?:watch\?.*(?:|\&)v=|embed\/|v\/)|youtu\.be\/)([-_0-9A-Za-z]{11})/
function ytv(url) {
    return new Promise((resolve, reject) => {
        if (ytIdRegex.test(url)) {
            let ytId = ytIdRegex.exec(url)
            url = 'https://youtu.be/' + ytId[1]
            post('https://www.y2mate.com/mates/id4/analyze/ajax', {
                url,
                q_auto: 0,
                ajax: 1
            })
                .then(res => res.json())
                .then(res => {
                    console.log('Scraping...')
                    document = (new JSDOM(res.result)).window.document
                    yaha = document.querySelectorAll('td')
                    filesize = yaha[yaha.length - 23].innerHTML
                    id = /var k__id = "(.*?)"/.exec(document.body.innerHTML) || ['', '']
                    thumb = document.querySelector('img').src
                    title = document.querySelector('b').innerHTML

                    post('https://www.y2mate.com/mates/id4/convert', {
                        type: 'youtube',
                        _id: id[1],
                        v_id: ytId[1],
                        ajax: '1',
                        token: '',
                        ftype: 'mp4',
                        fquality: 360
                    })
                        .then(res => res.json())
                        .then(res => {
                            let KB = parseFloat(filesize) * (1000 * /MB$/.test(filesize))
                            resolve({
                                dl_link: /<a.+?href="(.+?)"/.exec(res.result)[1],
                                thumb,
                                title,
                                filesizeF: filesize,
                                filesize: KB
                            })
                        }).catch(reject)
                }).catch(reject)
        } else reject('URL INVALID')
    })
}
					let { dl_link, thumb, title, filesize, filesizeF} = await ytv(args[0])
					client.sendMessage(from, dl_link, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
					await limitAdd(sender)
					break
                case `${prefix}text3d`:
            	if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
              	    if (args.length < 1) return reply('teksnya mana kak?')
                    teks = `${body.slice(8)}`
                    if (teks.length > 10) return reply('Teksnya kepanjangan, Maksimal 10 kalimat')
                    buff = await getBuffer(`https://docs-jojo.herokuapp.com/api/text3d?text=${teks}`, {method: 'get'})
                    client.sendMessage(from, buff, image, {quoted: mek, caption: `${teks}`})
			     	await limitAdd(sender)
					break
			    case `${prefix}fototiktok`:
				if (!isRegistered) return reply(ind.noregis())
			if (isBanned) return reply(ind.diban())
			if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    gatauda = body.slice(12)
                    anu = await fetchJson(`https://docs-jojo.herokuapp.com/api/tiktokpp?user=${gatauda}` , {method: 'get'})
			        buff = await getBuffer(anu.result)
                    reply(buff)
			        await limitAdd(sender)
					break
				case prefix+'tiktokdl':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					tiktokdl = await fetchJson(`http://docs-jojo.herokuapp.com/api/tiktok_nowm?url=${body.slice(10)}` , {method: 'get'})
					hasiltiktokdl = await getBuffer(tiktokdl.result.url)
					client.sendMessage(from, hasiltiktokdl, video, {mimetype: 'video/mp4', quoted: mek, filename: `${tiktokdl.result.title}.mp4`, caption: `*JUDUL:* ${tiktokdl.result.title}\n*DARI:* ${tiktokdl.result.from}`})
					await limitAdd(sender)
					break
				case prefix+'intro':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					client.updatePresence(from, Presence.composing) 
					client.sendMessage(from, ind.intro(), text, {quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(from ? { remoteJid: from } : {}) }, message: { conversation: `salin ini!  Usahakan jangan ada teks yang dihapus` }}})
					await limitAdd(sender)
					break
			    case `${prefix}map`:
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
               	 anu = await fetchJson(`https://mnazria.herokuapp.com/api/maps?search=${body.slice(5)}`, {method: 'get'})
                	buffer = await getBuffer(anu.gambar)
                	client.updatePresence(from, Presence.composing) 
                	client.sendMessage(from, buffer, image, {quoted: mek, caption: `${body.slice(5)}`})
					await limitAdd(sender)
					break
				case `${prefix}kbbi`:
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Apa yang mau dicari um?')
					try {
					kbbi = await fetchJson(`https://api.vhtear.com/kbbi?query=${q}&apikey=${VhtearKey}`, {method: 'get'})
					reply(`* KBBI *\n\n *Query* : ${q}\n\n *Result* : ${kbbi.data.result.hasil}\n\n_https://kbbi.kemdikbud.go.id/_`)
					await limitAdd(sender)
					} catch (err) {
                		reply(`KBBI Error : ${err}`)
           		 }
					break
		case prefix+'fotogoogle':
			argz = body.trim().split('|')
            if (argz.length >= 2) {
            const qwery = argz[1]
            const jum = argz[2]
            if(!qwery) return reply(`Kirim perintah *#googleimage [ |Query|Jumlah ]*, contoh = #googleimage |loli|3`)
            if(!jum) return reply(`Jumlah gambar diperlukan, contoh = #fotogoogle |loli|3`)
            if(jum >= 5) return reply('Jumlah terlalu banyak! Max 4')
            var gis = require('g-i-s');
            var opts = {
                searchTerm: qwery
                };
                gis(opts, logResults);
                    
                function logResults(error, results) {
                    if (error) {
                        reply('Maaf, Fitur Sedang Error')
                    } else {
                        const item = results.slice(0, jum)
                        item.forEach(async(res) => {
                        console.log(res)
                        const yurl = await urlShortener(res.url)
                        client.sendMessage(from, res.url, { quoted: mek, caption: ` Link : ${yurl}\n Image size : ${res.height} x ${res.width}`})  
                        await limitAdd(sender)
                        })
                    }
                }
            }
            break
	case prefix+'waktuindonesia':
	case prefix+'jamindo':
	case prefix+'waktuindo':
			
            if (!isRegistered) return reply(ind.noregis())
			if (isBanned) return reply(ind.diban())
            if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			reply(`Waktu Indonesia Barat: *${moment().utcOffset('+0700').format('HH:mm')}* WIB \nWaktu Indonesia Tengah: *${moment().utcOffset('+0800').format('HH:mm')}* WITA \nWaktu Indonesia Timur: *${moment().utcOffset('+0900').format('HH:mm')}* WIT`)
			break
	case prefix+'tinyurl':
			if (!isRegistered) return reply(ind.noregis())
			if (isBanned) return reply(ind.diban())
            if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
            tinyurl = await fetchJson(`https://tinyurl.com/api-create.php?url=${q}`, {method: 'get'})
            reply(JSON.stringify(tinyurl))
            break
	case prefix+'jadwalsholat':
			
            if (!isRegistered) return reply(ind.noregis())
			if (isBanned) return reply(ind.diban())
            if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
            tanggalsholat = moment.tz('Asia/Jakarta').format('YYYY-MM-DD')
            kodedaerah = await fetchJson(`https://api.banghasan.com/sholat/format/json/kota`, {method: 'get'})
					koddaerah = '=================\n'
					for (let kdedae of kodedaerah.kota) {
						koddaerah += `*Kode:* ${kdedae.id}\n*Nama:* ${kdedae.nama}\n`
					}
            jsholat = await fetchJson(`https://api.banghasan.com/sholat/format/json/jadwal/kota/${body.slice(14)}/tanggal/${tanggalsholat}`, {method: 'get'})
            if (jsholat.pesan) return reply(`MAAF, KODE DAERAH SALAH ATAU ADA YANG TIDAK LENGKAP\n\n${koddaerah.trim()}`)
            reply(`${tanda}\n*Kode kota:* ${query.kota}\n*Tanggal:* ${query.tanggal}\n*Imsak:* ${query.jadwal.data.imsak}\n*Subuh:* ${query.jadwal.data.subuh}\n*dhuha:* ${query.jadwal.data.dhuha}\n*Dhuhur:* ${query.jadwal.data.dzuhur}\n*Ashar:* ${query.jadwal.data.ashar}\n*Maghrib:* ${query.jadwal.data.maghrib}\n*Isya:* ${query.jadwal.data.isya}`)
            break
				case prefix+'moddroid':
                    if (!isRegistered) return reply(ind.noregis())
			if (isBanned) return reply(ind.diban())
			if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			data = await fetchJson(`https://tobz-api.herokuapp.com/api/moddroid?q=${body.slice(10)}&apikey=${TobzKey}`, {method: 'get'})
			hepi = data.result[0] 
			teks = `*Nama*: ${data.result[0].title}\n*publisher*: ${hepi.publisher}\n*mod info:* ${hepi.mod_info}\n*size*: ${hepi.size}\n*latest version*: ${hepi.latest_version}\n*genre*: ${hepi.genre}\n*link:* ${hepi.link}\n*download*: ${hepi.download}`
			buffer = await getBuffer(hepi.image)
			client.sendMessage(from, buffer, image, {quoted: mek, caption: `${teks}`})
			await limitAdd(sender)
			break
				case prefix+'happymod':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
			data = await fetchJson(`https://tobz-api.herokuapp.com/api/happymod?q=${body.slice(10)}&apikey=${TobzKey}`, {method: 'get'})
			hupo = data.result[0] 
			teks = `*Nama*: ${data.result[0].title}\n*version*: ${hupo.version}\n*size:* ${hupo.size}\n*root*: ${hupo.root}\n*purchase*: ${hupo.price}\n*link*: ${hupo.link}\n*download*: ${hupo.download}`
			buffer = await getBuffer(hupo.image)
			client.sendMessage(from, buffer, image, {quoted: mek, caption: `${teks}`})
			await limitAdd(sender)
			break
            case prefix+'bitly':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
               client.updatePresence(from, Presence.composing) 
                data = await fetchJson(`https://tobz-api.herokuapp.com/api/bitly?url=${args[0]}&apikey=${TobzKey}`, {method: 'get'})
                hasil = `link : ${args[0]}\n\nOutput : ${data.result}`
                reply(hasil)
                await limitAdd(sender)
                break
            case prefix+'nangis':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					anu = await fetchJson('https://tobz-api.herokuapp.com/api/cry?apikey=${TobzKey}', {method: 'get'})
					if (anu.error) return reply(anu.error)
					exec(`wget ${anu.result} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
				case prefix+'waifu':
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						res = await fetchJson(`https://waifu.pics/api/nsfw/waifu`, {method: 'get'})
						bufferm = await getBuffer(res.url)
						client.sendMessage(from, bufferm, image, {quoted: mek, caption: 'ezzzz'})
						await limitAdd(sender)
					break
				case prefix+'waifu2':
				
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(ind.wait())
					anu = await fetchJson(`https://tobz-api.herokuapp.com/api/waifu?apikey=${TobzKey}`, {method: 'get'})
					if (anu.error) return reply(anu.error)
					bufferss = await getBuffer(anu.image)
					waifu = `*${anu.desc}`
					client.sendMessage(from, bufferss, image, {quoted: mek, caption: waifu})
					break
				case prefix+'randomcry':
				case prefix+'cry':
				    try {
					
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					
						res = await fetchJson(`https://waifu.pics/api/sfw/cry`, {method: 'get'})
						bufferm = await getBuffer(res.url)
						client.sendMessage(from, bufferm, image, {quoted: mek, caption: 'ezzzz'})
						await limitAdd(sender)
					} catch (e) {
						console.log(`Error :`, color(e,'red'))
						sa = await getBuffer(`https://i.ibb.co/JcSjmNY/IMG-20210107-WA0052.jpg`)
						client.sendMessage(from, sa, image, {quoted: mek, caption: 'Error Kak!!'})
						reply(' *ERROR* ')
					}
					break
				case prefix+'randomhentai':
				case prefix+'hentai':
					
               	 if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				    try {
						res = await fetchJson(`https://tobz-api.herokuapp.com/api/hentai?apikey=${TobzKey}`, {method: 'get'})
						bufferxx = await getBuffer(res.result)
						client.sendMessage(from, bufferxx, image, {quoted: mek, caption: 'hentai teros'})
						await limitAdd(sender)
					} catch (e) {
						console.log(`Error :`, color(e,'red'))
						reply(' *ERROR* ')
					}
					break
                case prefix+'blowjoberror':
				
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					anu = await fetchJson('https://tobz-api.herokuapp.com/api/nsfwblowjob?apikey=${TobzKey}', {method: 'get'})
					if (anu.error) return reply(anu.error)
					exec(`wget ${anu.result} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
				case prefix+'kiss':
				case prefix+'cium':
				    try {
				
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						res = await fetchJson(`https://waifu.pics/api/sfw/kiss`, {method: 'get'})
						bufferv = await getBuffer(res.url)
						client.sendMessage(from, bufferv, image, {quoted: mek, caption: 'ezzzz'})
						await limitAdd(sender)
					} catch (e) {
						console.log(`Error :`, color(e,'red'))
						sa = await getBuffer(`https://i.ibb.co/JcSjmNY/IMG-20210107-WA0052.jpg`)
						client.sendMessage(from, sa, image, {quoted: mek, caption: 'Error Kak!!'})
						reply(' *ERROR* ')
					}
					break
					case prefix+'cium':
				
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					anu = await fetchJson('https://tobz-api.herokuapp.com/api/kiss?apikey=${TobzKey}', {method: 'get'})
					if (anu.error) return reply(anu.error)
					exec(`wget ${anu.result} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
					
					case prefix+'peluk':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					ranp = getRandom('.gif')
					rano = getRandom('.webp')
					anu = await fetchJson('https://tobz-api.herokuapp.com/api/hug?apikey=${TobzKey}', {method: 'get'})
					if (anu.error) return reply(anu.error)
					exec(`wget ${anu.result} -O ${ranp} && ffmpeg -i ${ranp} -vcodec libwebp -filter:v fps=fps=15 -lossless 1 -loop 0 -preset default -an -vsync 0 -s 512:512 ${rano}`, (err) => {
						fs.unlinkSync(ranp)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(rano)
						client.sendMessage(from, buffer, sticker, {quoted: mek})
						fs.unlinkSync(rano)
					})
					await limitAdd(sender)
					break
				case prefix+'qrcode1':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					exec(`qrencode -o ./sampah/qr_${sender.split("@")[0]}.png ${body.slice(8)}`)
					reply(`tunggu...`)
					setTimeout( () => {
					qrcodene = fs.readFileSync(`./sampah/qr_${sender.split("@")[0]}.png`)
					}, 1000) // 1000 = 1detik,
					setTimeout( () => {
					client.sendMessage(from, qrcodene, image, {quoted: mek, caption: 'ni'})
					}, 1500) // 1000 = 1detik,
					break
				case prefix+'translate':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`${prefix}translate kode bahasa  teks\n*CONTOH*\n${prefix}translate id i love you\n\n*TIDAK TAHU KODE BAHASA ?*\nketik${prefix}kodebhs\n*or* type ${prefix}codelang`)
                	const texto = q.substring(0, q.indexOf('|') - 1)
                	const languaget = q.substring(q.lastIndexOf('|') + 2)
                	translate(texto, {to: languaget}).then(res => {reply(res.text)})
           	 break
                case prefix+'husbu':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				    try {
						res = await fetchJson(`https://tobz-api.herokuapp.com/api/husbu?apikey=${TobzKey}`)
						buffer = await getBuffer(res.image)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: 'Ingat! Cintai husbumu'})
					} catch (e) {
						console.log(`Error :`, color(e,'red'))
						reply(' *ERROR* ')
					}
					await limitAdd(sender)
					break
                case prefix+'randomanime':
				case prefix+'ranime':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					gatauda = body.slice(8)
					reply(ind.wait())
					anu = await fetchJson(`https://tobz-api.herokuapp.com/api/randomanime?apikey=${TobzKey}`, {method: 'get'})
					buffer = await getBuffer(anu.result)
					client.sendMessage(from, buffer, image, {quoted: mek})
					await limitAdd(sender)
					break
			case `${prefix}nekonime`:
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				    try {
						res = await fetchJson(`https://tobz-api.herokuapp.com/api/nekonime?apikey=${TobzKey}`, {method: 'get'})
						buffer = await getBuffer(res.result)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: `${tanda}`})
						await limitAdd(sender)
					} catch (e) {
						console.log(`Error :`, color(e,'red'))
						reply(`${tanda}\n *ERROR* `)
					}
					break
			case prefix+'wibu':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				
					reply(ind.wait())
					anu = await fetchJson(`https://api.vhtear.com/randomwibu&apikey=${VhtearKey}`)
					if (anu.error) return reply(anu.error)
					bufferhh = await getBuffer(anu.result.foto)
					wibu = `  *nama* ${anu.result.nama}  *deskripsi* ${anu.result.deskripsi}`
					client.sendMessage(from, bufferhh, image, {quoted: mek, caption: wibu})
					await limitAdd(sender)
					break
			case prefix+'joox':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                data = await fetchJson(`https://tobz-api.herokuapp.com/api/joox?q=${body.slice(6)}&apikey=${TobzKey}`, {method: 'get'})
               if (data.error) return reply(data.error)
                 infomp3 = `*Lagu Ditemukan!!!*\nJudul : ${data.result.judul}\nAlbum : ${data.result.album}\nDipublikasi : ${data.result.dipublikasi}`
                buffer = await getBuffer(data.result.thumb)
                lagu = await getBuffer(data.result.mp3)
                client.updatePresence(from, Presence.composing)
                client.sendMessage(from, buffer, image, {quoted: mek, caption: infomp3})
                client.updatePresence(from, Presence.recording)
                client.sendMessage(from, lagu, audio, {mimetype: 'audio/mp4', filename: `${data.result.title}.mp3`, quoted: mek})
                await limitAdd(sender)
                break
			case prefix+'meme':
                
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					meme = await kagApi.memes()
					buffer = await getBuffer(`https://imgur.com/${meme.hash}.jpg`)
					client.sendMessage(from, buffer, image, {quoted: mek, caption: '.......'})
                                        await limitAdd(sender)
					break
			//animefoto
				case prefix+'naruto':
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=naruto&apikey=${VhtearKey}`, {method: 'get'})
					var naru = JSON.parse(JSON.stringify(anu.result));
					var to =  naru[Math.floor(Math.random() * naru.length)];
					nyew = await getBuffer(to)
					client.sendMessage(from, nyew, image, { caption: 'naruto!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'minato':
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=minato&apikey=${VhtearKey}`, {method: 'get'})
					var min = JSON.parse(JSON.stringify(anu.result));
					var ato =  min[Math.floor(Math.random() * min.length)];
					nyeq = await getBuffer(ato)
					client.sendMessage(from, nyeq, image, { caption: 'minato!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'boruto':
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=boruto&apikey=${VhtearKey}`, {method: 'get'})
					var bor = JSON.parse(JSON.stringify(anu.result));
					var uto =  bor[Math.floor(Math.random() * bor.length)];
					nyet = await getBuffer(uto)
					client.sendMessage(from, nyet, image, { caption: 'boruto!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'hinata':
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=hinata&apikey=${VhtearKey}`, {method: 'get'})
					var hina = JSON.parse(JSON.stringify(anu.result));
					var ta =  hina[Math.floor(Math.random() * hina.length)];
					nyei = await getBuffer(ta)
					client.sendMessage(from, nyei, image, { caption: 'hinata!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'sasuke':
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=sasuke&apikey=${VhtearKey}`, {method: 'get'})
					var sasu = JSON.parse(JSON.stringify(anu.result));
					var ke =  sasu[Math.floor(Math.random() * sasu.length)];
					nyeo = await getBuffer(ke)
					client.sendMessage(from, nyeo, image, { caption: 'sasuke!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'sakura':
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=sakura&apikey=${VhtearKey}`, {method: 'get'})
					var sak = JSON.parse(JSON.stringify(anu.result));
					var kura =  sak[Math.floor(Math.random() * sak.length)];
					nyep = await getBuffer(kura)
					client.sendMessage(from, nyep, image, { caption: 'sakura!!', quoted: mek })
					await limitAdd(sender)
					break
			case prefix+'loli':
				    try {
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					
						res = await fetchJson(`https://api.vhtear.com/randomloli&apikey=${VhtearKey}`, {method: 'get'})
						buffertt = await getBuffer(res.result.result)
						client.sendMessage(from, buffertt, image, {quoted: mek, caption: 'loli'})
						await limitAdd(sender)
					} catch (e) {
						console.log(`Error :`, color(e,'red'))
						sa = await getBuffer(`https://i.ibb.co/JcSjmNY/IMG-20210107-WA0052.jpg`)
						client.sendMessage(from, sa, image, {quoted: mek, caption: 'Error Kak!!'})
						reply(' *ERROR* ')
						}
						break
				case prefix+'loli2':
					
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=loli&apikey=${VhtearKey}`, {method: 'get'})
					var lol = JSON.parse(JSON.stringify(anu.result));
					var i2 =  lol[Math.floor(Math.random() * lol.length)];
					nyeee = await getBuffer(i2)
					client.sendMessage(from, nyeee, image, { caption: 'Oni chan baka!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'rize':
				
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=RizeKamishiro&apikey=${VhtearKey}`, {method: 'get'})
					var ri = JSON.parse(JSON.stringify(anu.result));
					var ze =  ri[Math.floor(Math.random() * ri.length)];
					nyeg = await getBuffer(ze)
					client.sendMessage(from, nyeg, image, { caption: 'rize chan!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'akira':
				
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=akiramado&apikey=${VhtearKey}`, {method: 'get'})
					var ak = JSON.parse(JSON.stringify(anu.result));
					var ara =  ak[Math.floor(Math.random() * ak.length)];
					nyeh = await getBuffer(ara)
					client.sendMessage(from, nyeh, image, { caption: 'akira chan!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'itori':
				
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=itori&apikey=${VhtearKey}`, {method: 'get'})
					var it = JSON.parse(JSON.stringify(anu.result));
					var ori =  it[Math.floor(Math.random() * it.length)];
					nyej = await getBuffer(ori)
					client.sendMessage(from, nyej, image, { caption: 'itori chan!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'kurumi':
				
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=kurumitokisakikawai&apikey=${VhtearKey}`, {method: 'get'})
					var kur = JSON.parse(JSON.stringify(anu.result));
					var imi =  kur[Math.floor(Math.random() * kur.length)];
					nyek = await getBuffer(imi)
					client.sendMessage(from, nyek, image, { caption: 'kurumi chan!!', quoted: mek })
					await limitAdd(sender)
					break
				case prefix+'miku':
				
            	    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/pinterest?query=Nakanomiku&apikey=${VhtearKey}`, {method: 'get'})
					var mi = JSON.parse(JSON.stringify(anu.result));
					var ku =  mi[Math.floor(Math.random() * mi.length)];
					nyel = await getBuffer(ku)
					client.sendMessage(from, nyel, image, { caption: 'miku chan!!', quoted: mek })
					await limitAdd(sender)
					break
            case prefix+'leaderboard':
				case prefix+'lb':
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				_level.sort((a, b) => (a.xp < b.xp) ? 1 : -1)
				uang.sort((a, b) => (a.uang < b.uang) ? 1 : -1)
                let leaderboardlvl = '-----[ *LEADERBOARD LEVEL* ]----\n\n'
                let leaderboarduang = '-----[ *LEADERBOARD UANG* ]----\n\n'
                let nom = 0
                try {
                    for (let i = 0; i < 10; i++) {
                        nom++
                        leaderboardlvl += `*[${nom}]* wa.me/${_level[i].id.replace('@s.whatsapp.net', '')}\n *XP*: ${_level[i].xp} *Level*: ${_level[i].level}\n`
                        leaderboarduang += `*[${nom}]* wa.me/${uang[i].id.replace('@s.whatsapp.net', '')}\n *Uang*: _Rp${uang[i].uang}_\n *Limit*: ${limitawal - _limit[i].limit}\n`
                    }
                    await reply(leaderboardlvl)
                    await reply(leaderboarduang)
                    await limitAdd(sender)
                } catch (err) {
                    console.error(err)
                    await reply(`minimal 10 user untuk bisa mengakses database`)
                }
				break
				case prefix+'mutual':
                
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                if (isGroup) return  reply( 'Command ini tidak bisa digunakan di dalam grup!')
                anug = getRegisteredRandomId(_registered).replace('@s.whatsapp.net','')
                await reply('Looking for a partner...')
                await reply(`wa.me/${anug}`)
                await reply( `Partner found: \n*${prefix}next* � find a new partner`)
            break
            case prefix+'next':
                
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                if (isGroup) return  reply( 'Command ini tidak bisa digunakan di dalam grup!')
                anug = getRegisteredRandomId(_registered).replace('@s.whatsapp.net','')
                await reply('Looking for a partner...')
                await reply(`wa.me/${anug}`)
                await reply( `Partner found: \n*${prefix}next* � find a new partner`)
            break
            case`${prefix}google`:
            
                const googleQuery = body.slice(8)
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                if(googleQuery == undefined || googleQuery == ' ') return reply(`*Hasil Pencarian : ${googleQuery}* tidak ditemukan`)
                google({ 'query': googleQuery }).then(results => {
                let vars = `_*Hasil Pencarian : ${googleQuery}*_\n`
                for (let i = 0; i < results.length; i++) {
                    vars +=  `\n\n*Judul* : ${results[i].title}\n\n*Deskripsi* : ${results[i].snippet}\n\n*Link* : ${results[i].link}\n\n`
                }
                    reply(vars)
                }).catch(e => {
                    console.log(e)
                    client.sendMessage(from, 'Google Error : ' + e);
                })
                await limitAdd(sender) 
                break 
            case prefix+'virtexasli':
            	
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				reply(`Halo ${pushname} / ${namaneuser(sender)}\nWaktu � 5 detik lagi_\nBersiap siaplah`)
				setTimeout( () => {
            	reply(ind.virtex(q))
            	fitnah(nomerewa, 'berhasil', `berhasil mengirim virtex dan menghapus virtex untuk bot`)
            	}, 5000) // 1000 = 1s,
            	break
			case prefix+'tulis': // BY MFARELS
                
                if (!isRegistered) return reply(ind.noregis())
                if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                if (args.length < 1) return reply(`Kirim Perintah ${command} Teks|Nama|Kelas-\n\n*Contoh :*\n${command} satu ditambah satu|Rifki|X`) // https://github.com/MFarelS/RajinNulis-BOT
                arg = body.trim().split('|') // INSTALL IMAGEMAGICK KALO MAU WORK
                const diNama = arg[1] // INSTALL, CENTANG KOLOM 1,2,3,5,6
                const diKelas = arg[2] // SUBSCRIBE MFARELS CH
                const diTulis = q.split('|')[0] // FOLLOW INSTAGRAM @mfarelsyahtiawan
                reply('menulis...') // NAMA, KELAS, WAKTU, BY ST4RZ
                const panjangKalimat = diTulis.replace(/(\S+\s*){1,10}/g, '$&\n')
                const panjangNama = diNama.replace(/(\S+\s*){1,10}/g, '$&\n')
                const panjangKelas = diKelas.replace(/(\S+\s*){1,10}/g, '$&\n')
                const panjangBaris = panjangKalimat.split('\n').slice(0, 30).join('\n')
                const panjangBarisNama = panjangNama.split('\n').slice(0, 30).join('\n')
                const panjangBarisKelas = panjangKelas.split('\n').slice(0, 30).join('\n')
                var months = ['- 1 -', '- 2 -', '- 3 -', '- 4 -', '- 5 -', '- 6 -', '- 7 -', '- 8 -', '- 9 -', '- 10 -', '- 11 -', '- 12 -'];
                var myDays = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu'];
                var date = new Date();
                var day = date.getDate();
                var month = date.getMonth();
                var thisDay = date.getDay(),
                    thisDay = myDays[thisDay];
                var yy = date.getYear();
                var year = (yy < 1000) ? yy + 1900 : yy;
                const waktunye = (day + ' ' + months[month] + ' ' + year)
                const harinye = (thisDay)
                spawn('convert', [
                    './mager/magernulis/magernulis1.jpg',
                    '-font',
                    './font/Zahraaa.ttf',
                    '-size',
                    '700x960',
                    '-pointsize',
                    '20',
                    '-interline-spacing',
                    '1',
                    '-annotate',
                    '+806+78',
                    harinye,
                    '-font',
                    './font/Zahraaa.ttf',
                    '-size',
                    '700x960',
                    '-pointsize',
                    '18',
                    '-interline-spacing',
                    '1',
                    '-annotate',
                    '+806+102',
                    waktunye,
                    '-font',
                    './font/Zahraaa.ttf',
                    '-size',
                    '700x960',
                    '-pointsize',
                    '18',
                    '-interline-spacing',
                    '1',
                    '-annotate',
                    '+360+100',
                    panjangBarisNama,
                    '-font',
                    './font/Zahraaa.ttf',
                    '-size',
                    '700x960',
                    '-pointsize',
                    '18',
                    '-interline-spacing',
                    '1',
                    '-annotate',
                    '+360+120',
                    panjangBarisKelas, 
                    '-font',
                    './font/Zahraaa.ttf',
                    '-size',
                    '700x960',
                    '-pointsize',
                    '20',
                    '-interline-spacing',
                    '-7.5',
                    '-annotate',
                    '+344+142',
                    panjangBaris,
                    `./mager/hasilnulis/nulis@${sender.split("@")[0]}.jpg`
                ])
                .on('error', () => reply('Error Bjeer, Keknya Scriptnya Lagi Error'))
                .on('exit', () => {
                	tulisanbuku = fs.readFileSync(`./mager/hasilnulis/nulis@${sender.split("@")[0]}.jpg`)
                	client.sendMessage(from, tulisanbuku, image, { quoted: mek, caption: `Sukses. Ditulis oleh bot`})
                })
            break // BY MFARELS
            case `kamu`:
            	if (args[0] === 'lumayan','ganteng','jelek','pinter','pintar','cantik','cowok','cewek') {
            		dahtahu = fs.readFileSync(`./fauzan.rifki.m/dahtahu.webp`)
            		client.sendMessage(from, dahtahu, sticker, {quoted: mek})
            	}
            	break
            case `zan`:
            case `rip`:
            case `rif`:
            	zan =['Ya ada apa','Ada perlu apa','kenapa?','Kamu lagi ngapain?','Disana enak','Apa bisa saya bantu','Halo','Apa kabar','Assalamualaikum','Waalaikumsalam','Kamu siapa','Halo','Hey','Aku siapa','Saya siapa hayoo?','sekarang jam berapa','Kamu siapa','Rumahmu dimana','Mau ngapain','selamat pagi','Selamat siang','Selamat sore','Selamat malam','Selamat tidur','Selamat Whatsapp an','Kamu sekarang lagi apa','Ini siapa ya?','1+1=2','Namaku siapa?']
				rif = zan[Math.floor(Math.random() * zan.length)]
                 hapuszan = reply(`${tanda}\nHalo *${pushname2}*\nNamaku terdeteksi (${command})\n\n${rif}`)
                 setTimeout( () => {
                 client.deleteMessage(from, { id: hapuszan })
                 }, 10000) // 1000 = 1s,
                 break
				case `ampunbangjago`:
				case `bangjago`:
				case `ampun`:
				case `${prefix}ampunbangjago`:
				case `${prefix}bangjago`:
				case `${prefix}ampun`:
				client.updatePresence(from, Presence.recording)
					lagubangjago = fs.readFileSync('./fauzan.rifki.m/bangjago.m4a')
					client.sendMessage(from, lagubangjago, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
					break
				case `welotka`:
				case `kang`:
				case `kangcopet`:
				case `bangkadada`:
				case `${prefix}welotka`:
				case `${prefix}kangcopet`:
				case `${prefix}bangkadada`:
				case `${prefix}kang`:
				client.updatePresence(from, Presence.recording)
					laguwelotka = fs.readFileSync('./fauzan.rifki.m/welot.mp3')
					client.sendMessage(from, laguwelotka, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
					break
				case `ladaladi`:
				case `${prefix}ladaladi`:
				client.updatePresence(from, Presence.recording)
					laguladaladi = fs.readFileSync('./fauzan.rifki.m/ladaladi.mp3')
					client.sendMessage(from, laguladaladi, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
					break
				
			// Nggo cek apikey
				case prefix+'zeks':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Ups, apikey nya mana`)
					try {
					zeks = await fetchJson(`https://api.zeks.xyz/api/pantun?apikey=${q}`, {method: 'get'})
					reply(`${tanda}\n\n*API:* api.zeks.xyz\n*ApiKey:* ${q}\n*Status ApiKey:* HIDUP\n\n${tanda}`)
					} catch {
					reply(`${tanda}\n\n*API:* api.zeks.xyz\n*ApiKey:* ${q}\n*Status ApiKey:* MATI GUYS\n\n${tanda}`)
					}
					await limitAdd(sender)
					break
				case prefix+'vhtear':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Ups, apikey nya mana`)
					try {
					vhtear = await fetchJson(`https://api.vhtear.com/random_pantun&apikey=${body.slice(8)}`, {method: 'get'})
					reply(`${tanda}\n\n*API:* api.vhtear.com\n*ApiKey:* ${body.slice(8)}\n*Status:* Hidup\n_${vhtear.result.pantun}_\n\n${tanda}`)
					} catch {
					reply(`${tanda}\n\n*API:* api.vhtear.com\n*ApiKey:* ${body.slice(8)}\n*Status:* Mati\n\n${tanda}`)
					}
					await limitAdd(sender)
					break
				case prefix+'tobz':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Ups, apikey nya mana`)
					tobz = await fetchJson(`https://tobz-api.herokuapp.com/api/wiki?q=Robot&apikey=${body.slice(6)}`, {method: 'get'})
					if (tobz.message) return reply(`${tanda}\n\n*API:* tobz-api.herokuapp.com\n*ApiKey:* ${body.slice(6)}\n*Status ApiKey:* ${tobz.message}\n\n${tanda}`)
					if (tobz.result) return reply(`${tanda}\n\n*API:* tobz-api.herokuapp.com\n*ApiKey:* ${body.slice(6)}\n*Status ApiKey:* HIDUP\n\n${tanda}`)
					await limitAdd(sender)
					break
				case prefix+'itech':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Ups, apikey nya mana`)
					itech = await fetchJson(`https://api.i-tech.id/tools/hilih?key=${body.slice(7)}&kata=kamu jangan gitu`, {method: 'get'})
					reply(`${tanda}\n\n*API:* api.i-tech.id\n*ApiKey:* ${body.slice(7)}\n*Status ApiKey:* ${itech.pesan}\n*NB* _jika muncul tulisan "Invalid API KEY" di status apikey\nBerarti ApiKey Tidak Aktif_\n\n${tanda}`)
					await limitAdd(sender)
					break
				case prefix+'itsmeiky':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Ups, apikey nya mana`)
					itsmeiky = await fetchJson(`https://api.itsmeikyxsec404.xyz/bacotandilan?apikey=${body.slice(10)}`, {method: 'get'})
					if (itsmeiky.msg) return reply(`${tanda}\n\n*ApiKey:* ${body.slice(10)}\n*Status ApiKey:* ${itsmeiky.msg}\n\n${tanda}`)
					reply(`${tanda}\n\n*API:* api.itsmeikyxsec404.xyz\n*ApiKey:* ${body.slice(10)}\n*Status ApiKey:* HIDUP\n\n${tanda}`)
					await limitAdd(sender)
					break
				case prefix+'xteam':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Ups, apikey nya mana`)
					apikeyxteam = await fetchJson(`https://api.xteam.xyz/cuaca?kota=surabaya&APIKEY=${body.slice(7)}`, {method: 'get'})
					cuaca = (`${apikeyxteam.message.cuaca}`)
					if (cuaca.includes(`undefined`)) return reply(`${tanda}\n\n*API:* api.xteam.xyz\n*ApiKey:* ${body.slice(7)}\n*Status ApiKey:* ${apikeyxteam.message}\n\n${tanda}`)
					reply(`${tanda}\n\n*ApiKey:* ${body.slice(7)}\n*Status ApiKey:* HIDUP\n${apikeyxteam.message.cuaca}\n\n${tanda}`)
					await limitAdd(sender)
					break
			// Nggo ngecek apikey wes entek
			
				case prefix+'dl':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					dlnya = await getBuffer(`${body.slice(4)}`)
						client.sendMessage(from, dlnya, image, {quoted: mek, caption: '.......'})
						.catch(err => {
						client.sendMessage(from, dlnya, video, {quoted: mek, caption: '.......'})
						}) 
						.else(err => {
						client.sendMessage(from, dlnya, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
						})
					break
				case `${prefix}ocr`: 
				case prefix+'jaditeks':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
						encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						media = await client.downloadAndSaveMediaMessage(encmedia)
						reply(ind.wait())
						await recognize(media, {lang: 'eng+ind', oem: 1, psm: 3})
							.then(teks => {
								reply(teks.trim())
								fs.unlinkSync(media)
								limitAdd(sender)
							})
							.catch(err => {
								reply(err.message)
								fs.unlinkSync(media)
							})
					} else {
						reply('Error, mencoba dengan metode lain')
						const config = {
  lang: "ind",
  oem: 1,
  psm: 3,
}

tesseract.recognize(media, config)
  .then(text => {
    reply(text)
  })
  .catch(error => {
    console.log(error.message)
  })
					}
					break
            case `${prefix}tagme`:
            case `${prefix}tagsaya`:
            	
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
            	tagme = {
					text: `Halo @${sender.split("@")[0]}\n\nItu kamu sudah di tag`,
					contextInfo: { mentionedJid: [sender] }
					}
					reply(tagme)
					break
				case `${prefix}stikererr`: 
				case `${prefix}stickererr`:
				    
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                    await limitAdd(sender)
					if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
						const encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						const media = await client.downloadAndSaveMediaMessage(encmedia)
						ran = getRandom('.webp')
						await ffmpeg(`./${media}`)
							.input(media)
							.on('start', function (cmd) {
								console.log(`Started : ${cmd}`)
							})
							.on('error', function (err) {
								console.log(`Error : ${err}`)
								fs.unlinkSync(media)
								reply(ind.stikga())
							})
							.on('end', function () {
								console.log('Finish')
								buff = fs.readFileSync(ran)
								client.sendMessage(from, buff, sticker, {quoted: mek})
								fs.unlinkSync(media)
								fs.unlinkSync(ran)
							})
							.addOutputOptions([`-vcodec`,`libwebp`,`-vf`,`scale='min(320,iw)':min'(320,ih)':force_original_aspect_ratio=decrease,fps=15, pad=320:320:-1:-1:color=white@0.0, split [a][b]; [a] palettegen=reserve_transparent=on:transparency_color=ffffff [p]; [b][p] paletteuse`])
							.toFormat('webp')
							.save(ran)
					} else if ((isMedia && mek.message.videoMessage.seconds < 11 || isQuotedVideo && mek.message.extendedTextMessage.contextInfo.quotedMessage.videoMessage.seconds < 11) && args.length == 0) {
						const encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						const media = await client.downloadAndSaveMediaMessage(encmedia)
						ran = getRandom('.webp')
						reply(ind.wait())
						await ffmpeg(`./${media}`)
							.inputFormat(media.split('.')[1])
							.on('start', function (cmd) {
								console.log(`Started : ${cmd}`)
							})
							.on('error', function (err) {
								console.log(`Error : ${err}`)
								fs.unlinkSync(media)
								tipe = media.endsWith('.mp4') ? 'video' : 'gif'
								reply(ind.stikga())
							})
							.on('end', function () {
								console.log('Finish')
								buff = fs.readFileSync(ran)
								client.sendMessage(from, buff, sticker, {quoted: mek})
								fs.unlinkSync(media)
								fs.unlinkSync(ran)
							})
							.addOutputOptions([`-vcodec`,`libwebp`,`-vf`,`scale='min(320,iw)':min'(320,ih)':force_original_aspect_ratio=decrease,fps=15, pad=320:320:-1:-1:color=white@0.0, split [a][b]; [a] palettegen=reserve_transparent=on:transparency_color=ffffff [p]; [b][p] paletteuse`])
							.toFormat('webp')
							.save(ran)
							} else {
						reply(`Kirim gambar dengan caption ${prefix}sticker atau reply/tag gambar`)
					}
					break
				case prefix+'tts':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 2) return reply(`${prefix}tts id halo orang aneh\n*itu contoh*`)
					var codelang = args[0]
               	 var tekstts = body.slice(5+codelang.length)
					reply(`mencari mulut nya google ...`)
					client.updatePresence(from, Presence.recording)
					ttsnya = await getBuffer(`http://translate.google.com/translate_tts?tl=${codelang}&q="${tekstts}"&client=tw-ob`)
					client.sendMessage(from, ttsnya, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
					break
				case prefix+'katakan':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`${prefix}katakan halo\n*itu contoh*`)
					reply(`mencari mulut nya mbak google ...`)
					client.updatePresence(from, Presence.recording)
					ttsnya = await getBuffer(`http://translate.google.com/translate_tts?tl=id&q="${q}"&client=tw-ob`)
					client.sendMessage(from, ttsnya, audio, {mimetype: 'audio/mp4', quoted: mek, ptt: true})
					break
				case prefix+'codelang':
				case prefix+'kodebhs':
					
               	 if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(listbahasa)
					break
				case `${prefix}restart`:
				case `${prefix}reboot`:
				if (!isOwner) return reply(ind.ownerb())
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				try {
					botdipateni = {
					text: `@${sender.split("@")[0]} ingin me-restart / mulai ulang bot ini\n\n_membutuhkan waktu lama_`,
					contextInfo: { mentionedJid: [sender] }
					}
					reply(botdipateni)
					const gtts = require('./lib/gtts')(args[0])
					dtt = body.slice(2)
					ranm = getRandom('.mpmpnan')
					rano = getRandom('.ogg')
					dtt.length > 300
					? reply('lah teks nya kepanjangan bambang')
					: gtts.save(ranm, dtt, function() {
						exec(`ffmpeg -i ${ranm} -ar 48000 -vn -c:a libopus ${rano}`, (err) => {
							fs.unlinkSync(ranm)
							buff = fs.readFileSync(rano)
							if (err) return reply(ind.stikga())
							client.sendMessage(from, buff, audio, {quoted: mek, ptt:true})
							fs.unlinkSync(rano)
						})
					})
					exec(`node rifki.js`)
					} catch {
						reply('Gagal Mematikan Bot')
					}
					await limitAdd(sender)
					break
				case prefix+'crash':
				case prefix+'bunuhbot':
				case prefix+'forceclose':
				if (!isOwner) return reply(ind.ownerb())
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				try {
					botdipateni = {
					text: `@${sender.split("@")[0]} ingin me-restart / mulai ulang bot ini\n\n_membutuhkan waktu lama_`,
					contextInfo: { mentionedJid: [sender] }
					}
					reply(botdipateni)
					const gtts = require('./lib/gtts')(args[0])
					dtt = body.slice(2)
					ranm = getRandom('.mpmpnan')
					rano = getRandom('.ogg')
					dtt.length > 300
					? reply('lah teks nya kepanjangan bambang')
					: gtts.save(ranm, dtt, function() {
						exec(`ffmpeg -i ${ranm} -ar 48000 -vn -c:a libopus ${rano}`, (err) => {
							fs.unlinkSync(ranm)
							buff = fs.readFileSync(rano)
							if (err) return reply(ind.stikga())
							client.sendMessage(from, buff, audio, {quoted: mek, ptt:true})
							fs.unlinkSync(rano)
						})
					})
					reply('Gagal Mematikan Bot')
					} catch {
						reply('Gagal Mematikan Bot')
					}
					await limitAdd(sender)
					break
				case `${prefix}ketik`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(`${body.slice(7)}`)
					await limitAdd(sender)
					break
				case `${prefix}setprefix`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return
					prefix = args[0]
					prefixnya = {
					text: `Prefix diubah menjadi\n\n ${prefix}\n\n\nDiubah oleh @${sender.split("@")[0]}`,
					contextInfo: { mentionedJid: [sender] }
					}
					client.sendMessage(from, prefixnya, text, {quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(from ? { remoteJid: from } : {}) }, message: { conversation: `berhasil` }}})
					break 
				case `${prefix}tiktokstalk`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						if (args.length < 1) return client.sendMessage(from, '  ?', text, {quoted: mek})
						let { user, stats } = await tiktod.getUserProfileInfo(args[0])
						reply(ind.wait())
						teks = `*ID* : ${user.id}\n*Username* : ${user.uniqueId}\n*Nickname* : ${user.nickname}\n*Followers* : ${stats.followerCount}\n*Followings* : ${stats.followingCount}\n*Posts* : ${stats.videoCount}\n*Luv* : ${stats.heart}\n`
						buffer = await getBuffer(user.avatarLarger)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: teks})
					await limitAdd(sender)
					break
				case prefix+'pptiktok':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
						if (args.length < 1) return client.sendMessage(from, '  ?', text, {quoted: mek})
						try {
						let { user, stats } = await tiktod.getUserProfileInfo(args[0])
						reply(ind.wait())
						buffer = await getBuffer(user.avatarLarger)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: `${user.nickname}`})
						} catch (e) {
							reply(`*LOG ERROR*\n${monosp}${e}${monosp}`)
						}
					await limitAdd(sender)
					break
                 case `${prefix}linkgrup`:
                 case `${prefix}linkgroup`:
				    
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (!isGroup) return reply(ind.groupo())
				    if (!isBotGroupAdmins) return reply(ind.badmin())
				    linkgc = await client.groupInviteCode (from)
				    yeh = `https://chat.whatsapp.com/${linkgc}\n\nlink Group *${groupName}*`
				    client.sendMessage(from, yeh, text, {quoted: mek})
			        await limitAdd(sender)
					break
		case prefix+'listsurah': // ARUGAZ
          	 	  
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
            if (!isGroup) return reply(`Perintah ini hanya bisa di gunakan dalam group!`)
            try {
                axios.get('https://raw.githubusercontent.com/ArugaZ/scraper-results/main/islam/surah.json')
                .then((response) => {
                    let hehex = '* DAFTAR SURAH *\n\n___________________________\n'
                    let nmr = 1
                    for (let i = 0; i < response.data.data.length; i++) {
                        hehex += nmr + '. ' +  monospace(response.data.data[i].name.transliteration.id.toLowerCase()) + '\n'
                        nmr++
                            }
                        hehex += '___________________________'
                    reply(hehex)
                })
            } catch(err) {
                reply(err)
            }
            break
        case prefix+'infosurah': // ARUGAZ
            
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
            if (args.length == 1) return reply(`Kirim perintah *#infosurah [ Nama Surah ]*\nContoh : *#infosurah al-fatihah*`)
                var responseh = await axios.get('https://raw.githubusercontent.com/ArugaZ/scraper-results/main/islam/surah.json')
                var { data } = responseh.data
                var idx = data.findIndex(function(post, index) {
                if((post.name.transliteration.id.toLowerCase() == args[1].toLowerCase())||(post.name.transliteration.en.toLowerCase() == args[1].toLowerCase()))
                    return true;
                });
                try {
                    var pesan = "* INFORMASI SURAH *\n\n___________________________\n\n"
                    pesan = pesan + " *Nama* : "+ data[idx].name.transliteration.id + "\n" + " *Asma* : " +data[idx].name.short+"\n"+" *Arti* : "+data[idx].name.translation.id+"\n"+" *Jumlah ayat* : "+data[idx].numberOfVerses+"\n"+" *Nomor surah* : "+data[idx].number+"\n"+"Jenis : "+data[idx].revelation.id+"\n"+" *Keterangan* : "+data[idx].tafsir.id
                    pesan += '\n\n___________________________'
                    reply(pesan)
                    limitAdd(serial)
                }catch{
                    reply(`Data tidak ditemukan, atau nama surah salah\n\n*LOG ERROR*\n${monosp}${e}${monosp}`)
                }
            break
        case prefix+'tafsir': // ARUGAZ
            
                    if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
            if (args.length == 1) return reply(`Kirim perintah *#tafsir [ Nama Surah ] [ Ayat ]*\nContoh : *#tafsir al-fatihah 2*`)
                var responsh = await axios.get('https://raw.githubusercontent.com/ArugaZ/scraper-results/main/islam/surah.json')
                var {data} = responsh.data
                var idx = data.findIndex(function(post, index) {
                if((post.name.transliteration.id.toLowerCase() == args[1].toLowerCase())||(post.name.transliteration.en.toLowerCase() == args[1].toLowerCase()))
                    return true;
                });
            try{
                nmr = data[idx].number
                if(!isNaN(nmr)) {
                var responsih = await axios.get('https://api.quran.sutanlab.id/surah/'+nmr+"/"+args[2])
                    var {data} = responsih.data
                    pesan = ""
                    pesan = pesan + "* TAFSIR *\n\nTafsir Q.S. "+data.surah.name.transliteration.id+":"+args[2]+"\n\n"
                    pesan = pesan + data.text.arab + "\n\n"
                    pesan = pesan + "_" + data.translation.id + "_" + "\n\n" +data.tafsir.id.long
                    pesan += '\n\n___________________________'
                    reply(pesan)
                    limitAdd(serial)
                }
            }catch{
                reply('Data tidak ditemukan, mungkin nama surah/ayat salah')
            }
            break
        	case prefix+'listonline':
        		if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (!isGroup) return reply(ind.groupo())
        		let ido = args && /\d+\-\d+@g.us/.test(args[0]) ? args[0] : from
			    let online = [...Object.keys(client.chats.get(ido).presences), client.user.jid]
			    client.sendMessage(from, 'Yang online & centang biru:\n' + online.map(v => '- @' + v.replace(/@.+/, '')).join`\n`, text, { quoted: mek,
  			  contextInfo: { mentionedJid: online }
			    })
				break
				case `${prefix}tagall`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isGroup) return reply(ind.groupo())
					members_id = []
					teks = (args.length > 1) ? body.slice(8).trim() : ''
					teks += '\n\n'
					for (let mem of groupMembers) {
						teks += ` @${mem.jid.split('@')[0]}\n`
						members_id.push(mem.jid)
					}
					mentions(teks, members_id, true)
					break
				case `${prefix}clearall`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (!isOwner) return reply(ind.ownerb())
					anu = await client.chats.all()
					client.setMaxListeners(25)
					for (let _ of anu) {
						client.deleteChat(_.jid)
					}
					reply(ind.clears())
					break
			       case `${prefix}block`:
				 client.updatePresence(from, Presence.composing) 
				 client.chatRead (from)
					if (!isGroup) return reply(ind.groupo())
					client.blockUser (`${body.slice(7)}@c.us`, "add")
					client.sendMessage(from, `perintah Diterima, memblokir ${body.slice(7)}@c.us`, text)
					break
                    case `${prefix}unblock`:
					if (!isGroup) return reply(ind.groupo())
					if (!isOwner) return reply(ind.ownerb())
				    client.blockUser (`${body.slice(9)}@c.us`, "remove")
					client.sendMessage(from, `?? ??,  ${body.slice(9)}@c.us`, text)
					break
				case `${prefix}leave`: 
				case `${prefix}kickbot`:
				if (!isGroup) return reply(ind.groupo())
				await reply('saya akan keluar')
				setTimeout( () => {
				client.leaveGroup(groupId)
				}, 1000) // 1000 = 1detik,
				await limitAdd(sender)
					break
				case `${prefix}clearbc`: 
					if (!isOwner) return reply(ind.ownerb()) 
						for (let dibc of _registered) {
							client.deleteChat(dibc.id)
						reply('menghapus bc untuk bot saja\nbukan untuk semua orang')
					}
					await limitAdd(sender)
					break
				case prefix+'setpp':
				case prefix+'ubah.ikon':
                        if (!isGroup) return reply(ind.groupo())
                       if (!isGroupAdmins) return reply(ind.admin())
                        if (!isBotGroupAdmins) return reply(ind.badmin())
                        if (isMedia && !mek.message.videoMessage || isQuotedImage) {
                       media = await client.downloadAndSaveMediaMessage(mek)
                         await client.updateProfilePicture (from, media)
                        reply('Berhasil mengubah ikon grup')
                        } else {
                         reply(ind.badmin())
                        }
					break
				case `${prefix}add`:
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (args.length < 1) return reply('Yang mau di add jin ya?')
					if (args[0].startsWith('08')) return reply('Gunakan kode negara mas')
					try {
						num = `${args[0].replace(/ /g, '')}@s.whatsapp.net`
						client.groupAdd(from, [num])
					} catch (e) {
						console.log('Error :', e)
						reply('Gagal menambahkan target, mungkin karena di private')
					}
					await limitAdd(sender)
					break
				case `${prefix}culik10`:
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (args.length < 1) return reply(`Codechat nya mana\nSilahkan ketik ${prefix}codechat`)
					culiklist = []
					for (let culik of groupMembers) {
						teks += `** ${culik.jid.split('@')[0]}\n`
						culiklist.push(culik.jid)
					}
						nyulik1 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik1)
						nyulik2 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik2)
						nyulik3 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik3)
						nyulik4 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik4)
						nyulik5 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik5)
						nyulik6 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik6)
						nyulik7 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik7)
						nyulik8 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik8)
						nyulik9 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik9)
						nyulik10 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik10)
					await limitAdd(sender)
					break
				case `${prefix}culik20`:
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (args.length < 1) return reply(`Codechat nya mana\nSilahkan ketik ${prefix}codechat`)
					culiklist = []
					for (let culik of groupMembers) {
						teks += `** ${culik.jid.split('@')[0]}\n`
						culiklist.push(culik.jid)
					}
						nyulik1 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik1)
						nyulik2 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik2)
						nyulik3 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik3)
						nyulik4 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik4)
						nyulik5 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik5)
						nyulik6 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik6)
						nyulik7 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik7)
						nyulik8 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik8)
						nyulik9 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik9)
						nyulik10 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik10)
						nyulik11 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik11)
						nyulik12 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik12)
						nyulik13 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik13)
						nyulik14 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik14)
						nyulik15 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik15)
						nyulik16 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik16)
						nyulik17 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik17)
						nyulik18 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik18)
						nyulik19 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik19)
						nyulik20 = culiklist[Math.floor(Math.random() * culiklist.length)]
						client.groupAdd(args[0], nyulik20)
					await limitAdd(sender)
					break
					case `${prefix}grup`:
					case `${prefix}grub`:
					case `${prefix}group`:
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (args[0] === 'buka') {
					    reply(`*BERHASIL MEMBUKA GROUP*`)
						client.groupSettingChange(from, GroupSettingChange.messageSend, false)
					} else if (args[0] === 'tutup') {
						reply(`*BERHASIL MENUTUP GROUP`)
						client.groupSettingChange(from, GroupSettingChange.messageSend, true)
					}
					await limitAdd(sender)
					break      
            case `${prefix}admin`:
            case `${prefix}owner`:
            case `${prefix}creator`:
            	
                    if (!isRegistered) return reply(ind.noregis())
vcard = 'BEGIN:VCARD\n' 
            + 'VERSION:3.0\n' 
            + `FN:${devName}\n`
            + `ORG: ${ownerName};\n`
            + `TEL;type=CELL;type=VOICE;waid=${nomowner}:${nomowner}\n`
            + 'END:VCARD' 
                  client.sendMessage(from, {displayname: "Jeff", vcard: vcard}, MessageType.contact, { quoted: mek})
                  await limitAdd(sender)
					break    
           case `${prefix}setname`:
           	
                    if (!isRegistered) return reply(ind.noregis())
                if (!isGroup) return reply(ind.groupo()) 
				try {
                client.groupUpdateSubject(from, `${body.slice(9)}`)
                client.sendMessage(from, 'Succes, Ganti Nama Grup', text, {quoted: mek})
                } catch (e) {
							reply(err)
							reply(ind.badmin())
						}
						await limitAdd(sender)
					break
				case prefix+'chatlist':
				case prefix+'chatmu':
				case prefix+'listchat':
                    if (!isRegistered) return reply(ind.noregis())
          	      if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					client.updatePresence(from, Presence.composing)  
                    if (!isRegistered) return reply(ind.noregis())
                
					kae = await client.chats.all()
					teks = `${tanda}\n  *INI CHATKU* \n`
					for (let i of kae) {
						teks += `** ${i.jid}\nName: ${i.name}\nCount: ${i.count}\nt: ${i.t}\nMute: ${i.mute}\nMessage: ${i.message}\nModify tag: ${i.modify_tag}\nSpam: ${i.spam}\n\n`
					}
					reply(teks.trim())
					console.log(kae)
					await limitAdd(sender)
					break
               case `${prefix}setdesc`:
            	if (!isRegistered) return reply(ind.noregis())
                if (!isGroup) return reply(ind.groupo())
                client.groupUpdateDescription(from, `${body.slice(9)}`)
                client.sendMessage(from, 'Succes, Ganti Deskripsi Grup', text, {quoted: mek})
                await limitAdd(sender)
					break
           case `${prefix}demote`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (mentioned.length > 1) {
						teks = 'Berhasil Demote\n'
						for (let _ of mentioned) {
							teks += `@${_.split('@')[0]}\n`
						}
						mentions(teks, mentioned, true)
						client.groupRemove(from, mentioned)
					} else {
						mentions(`Berhasil Demote @${mentioned[0].split('@')[0]} Menjadi Member Group!`, mentioned, true)
						client.groupDemoteAdmin(from, mentioned)
					}
					break
					await limitAdd(sender)
					break
				case prefix+'promote.me':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))				
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					try {
						num = `${sender}`
						client.groupMakeAdmin(from, [num])
					} catch (e) {
						console.log('Error :', e)
						reply('dahlah, saya nggak bisa')
					}
					break
				case `${prefix}promote`:
                    if (!isRegistered) return reply(ind.noregis())
					if (!isGroup) return reply(ind.groupo())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (mentioned.length > 1) {
						teks = 'Berhasil Promote\n'
						for (let _ of mentioned) {
							teks += `@${_.split('@')[0]}\n`
						}
						mentions(from, mentioned, true)
						client.groupRemove(from, mentioned)
					} else {
						mentions(`Berhasil Promote @${mentioned[0].split('@')[0]} Sebagai Admin Group!`, mentioned, true)
						client.groupMakeAdmin(from, mentioned)
					}
					await limitAdd(sender)
					break
				case `${prefix}kickme`:
					if (!isGroup) return reply(ind.groupo())
                    if (!isRegistered) return reply(ind.noregis())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					reply(`Perintah diterima, saya akan mengeluarkan anda dari grup ini`)
					num = `${sender}`
					client.groupRemove(from, [num])
					break
				case prefix+'edotense':
				case prefix+'endotense':
				case prefix+'edotensei':
					if (!isGroup) return reply(ind.groupo())
                    if (!isRegistered) return reply(ind.noregis())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return reply('Tag target yang ingin di tendang!')
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					mentions(`Perintah di terima, di edotensei : @${mentioned[0].split('@')[0]}`, mentioned, true)
					.then(() => {client.groupRemove(from, mentioned)})
					.then(() => {reply(`Menghidupkan kembali`)})
					setTimeout( () => {
					client.groupAdd(from, mentioned)
					}, 5000) // 1000 = 1s,
					break
				case `${prefix}kick`:
					if (!isGroup) return reply(ind.groupo())
				try {
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (budy.includes(`@${me.jid.split('@')[0]}`)) return client.groupRemove(from, mentioned)
                    if (!isRegistered) return reply(ind.noregis())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return reply('Tag target yang ingin di tendang!')
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (isBotGroupAdmins) {
					mentions(`Perintah di terima !!, mengeluarkan : @${mentioned[0].split('@')[0]}`, mentioned, true)
					setTimeout( () => {
						kick = fs.readFileSync(`./fauzan.rifki.m/kick1.webp`)
            			client.sendMessage(from, kick, sticker, {quoted: mek})
					}, 1000) // 1000 = 1s,
					setTimeout( () => {
						kick = fs.readFileSync(`./fauzan.rifki.m/kick2.webp`)
            			client.sendMessage(from, kick, sticker, {quoted: mek})
					}, 5000) // 1000 = 1s,
					setTimeout( () => {
						kick = fs.readFileSync(`./fauzan.rifki.m/kick3.webp`)
            			client.sendMessage(from, kick, sticker, {quoted: mek})
            		}, 9000) // 1000 = 1s,
					setTimeout( () => {
						client.groupRemove(from, mentioned)
						tuduh(`${me.jid}`, `Perintah di terima !!, mengeluarkan : @${mentioned[0].split('@')[0]}`, `Kelamaan woy, keburu di kick orang`)
						tuduh(`${me.jid}`, `Kelamaan woy, keburu di kick orang`, `pakai *#kickfast*\njika pengen cepat`)
					}, 13000) // 1000 = 1s,
					}
					await limitAdd(sender)
				} catch {
					reply(`Yang mau di kick apa ?\nApakah si kebal itu yang di kick ?`)
				}
					break
				case `${prefix}kickfast`:
				case `${prefix}fastkick`:
					if (!isGroup) return reply(ind.groupo())
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (budy.includes(`@${me.jid.split('@')[0]}`)) return client.groupRemove(from, mentioned)
                    if (!isRegistered) return reply(ind.noregis())
					if (!isBotGroupAdmins) return reply(ind.badmin())
					if (mek.message.extendedTextMessage === undefined || mek.message.extendedTextMessage === null) return reply('Tag target yang ingin di tendang!')
					mentioned = mek.message.extendedTextMessage.contextInfo.mentionedJid
					if (mentioned.length > 1) {
						teks = 'Perintah di terima, mengeluarkan :\n'
						for (let _ of mentioned) {
							teks += `@${_.split('@')[0]}\n`
						}
						mentions(teks, mentioned, true)
						client.groupRemove(from, mentioned)
					} else {
						mentions(`Perintah di terima, mengeluarkan : @${mentioned[0].split('@')[0]}`, mentioned, true)
						client.groupRemove(from, mentioned)
					}
					await limitAdd(sender)
					break
				case `${prefix}listadmin`:
				case prefix+'listadmins':
				case prefix+'listadmin':
				case prefix+'daftaradmin':
				case prefix+'admingrup':
                    if (!isRegistered) return reply(ind.noregis())
					if (!isGroup) return reply(ind.groupo())
					teks = `List admin of group *${groupMetadata.subject}*\nTotal : ${groupAdmins.length}\n\n`
					no = 0
					for (let admon of groupAdmins) {
						no += 1
						teks += `[${no.toString()}] @${admon.split('@')[0]}\n`
					}
					mentions(teks, groupAdmins, true)
					await limitAdd(sender)
					break
				case prefix+'infopesan':
				case prefix+'infomessage':
               	 if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					console.log(encmedia)
					reply(JSON.stringify(encmedia, null, 2))
					await limitAdd(sender)
					break
			case prefix+'afk': // by Slavyan
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
                if (isAfkOn) return reply(ind.afkOnAlready())
                const reason = q ? q : '_Nggak punya alasan_'
                addAfkUser(sender, time, reason, _afk)
				reply(ind.afkOn(namaneuser(sender), reason))
                await limitAdd(sender)
           	 break
			case `${prefix}toimg`:
			case `${prefix}stikimg`:
                if (!isRegistered) return reply(ind.noregis())
				if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
				if (!isQuotedSticker) return reply('tidak ada sticker')
					reply(ind.wait())
					encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
					media = await client.downloadAndSaveMediaMessage(encmedia)
					ran = getRandom('.png')
					exec(`ffmpeg -i ${media} ${ran}`, (err) => {
						fs.unlinkSync(media)
						if (err) return reply(ind.stikga())
						buffer = fs.readFileSync(ran)
						client.sendMessage(from, buffer, image, {quoted: mek, caption: '??  '})
						fs.unlinkSync(ran)
					})
					await limitAdd(sender)
					break
				case `*`:
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply('Textnya mana um?')
					simijawab = await fetchJson(`http://lolhuman.herokuapp.com/api/simi?apikey=${LolKey}&text=${q}`, {method: 'get'})
					//if (simijawab.error) return reply('Simi error kak')
					reply(`${tanda}\n${simijawab.result}`)
					await limitAdd(sender)
					break
				case prefix+'tebakgambar':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					tebgmb = await fetchJson(`https://videfikri.com/api/tebakgambar`, {method: 'get'})
					bufftbkgmb = await getBuffer(tebgmb.result.soal_gbr)
					setTimeout( () => {
					client.sendMessage(from, `${tanda}\n*➸ Jawaban :* `+tebgmb.result.jawaban, text, {quoted: mek}) // ur cods
					}, 30000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_10 Detik lagi…_', text) // ur cods
					}, 20000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_20 Detik lagi_…', text) // ur cods
					}, 10000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_30 Detik lagi_…', text) // ur cods
					}, 1000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, bufftbkgmb, image, { caption: `${tanda}\n_Jelaskan Apa Maksud Gambar Ini_`, quoted: mek }) // ur cods
					}, 0) // 1000 = 1s,
					break
				case prefix+'caklontong':
                	if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/funkuis&apikey=${VhtearKey}`, {method: 'get'})
					setTimeout( () => {
					client.sendMessage(from, '*➸ Jawaban :* '+anu.result.jawaban+'\n'+anu.result.desk, text, {quoted: mek}) // ur cods
					}, 30000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_10 Detik lagi…_', text) // ur cods
					}, 20000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_20 Detik lagi_…', text) // ur cods
					}, 10000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_30 Detik lagi_…', text) // ur cods
					}, 1000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, anu.result.soal, text, { quoted: mek }) // ur cods
					}, 0) // 1000 = 1s,
					break
				case prefix+'family100':
					if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					anu = await fetchJson(`https://api.vhtear.com/family100&apikey=${VhtearKey}`, {method: 'get'})
					setTimeout( () => {
					client.sendMessage(from, '*➸ Jawaban :* '+anu.result.jawaban, text, {quoted: mek}) // ur cods
					}, 30000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_10 Detik lagi…_', text) // ur cods
					}, 20000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_20 Detik lagi_…', text) // ur cods
					}, 10000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, '_30 Detik lagi_…', text) // ur cods
					}, 1000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, anu.result.soal, text, { quoted: mek }) // ur cods
					}, 0) // 1000 = 1s,
					break
				case prefix+'testime':
				case prefix+'teswaktu':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					reply(`*Silahkan isi formulir ini dibawah ini*\nJika tidak kamu akan saya keluarkan dari grup yang sama dengan saya`)
					setTimeout( () => {
					client.updatePresence(from, Presence.recording)
					lagutapi = fs.readFileSync('./fauzan.rifki.m/tapiboong.m4a')
					client.sendMessage(from, lagutapi, audio, {mimetype: 'audio/mp4', ptt: true})
					}, 35000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, `anda telah dikeluarkan`, text, {quoted: mek}) // ur cods
					}, 30000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, `${tanda}\n_10 Detik lagi�_`, text, {quoted: mek}) // ur cods
					}, 20000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, `${tanda}\n_20 Detik lagi_�`, text, {quoted: mek}) // ur cods
					}, 10000) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, `${tanda}\n_30 Detik lagi_�`, text, {quoted: mek}) // ur cods
					}, 2500) // 1000 = 1s,
					setTimeout( () => {
					client.sendMessage(from, `${tanda}\n_Tes waktu akan dimulai_�`, text, {quoted: mek}) // ur cods
					}, 0) // 1000 = 1s,
					await limitAdd(sender)
					break
                case `${prefix}leveling`:
				if (!isRegistered) return reply(ind.noregis())
                if (args.length < 1) return reply('Boo :')
                if (args[0] === 'on') {
                    if (isLevelingOn) return reply('*fitur level sudah aktif sebelum nya*')
                    _leveling.push(from)
                    fs.writeFileSync('./database/group/leveling.json', JSON.stringify(_leveling))
                     reply(ind.lvlon())
                } else if (args[0] === 'off') {
                    _leveling.splice(from, 1)
                    fs.writeFileSync('./database/group/leveling.json', JSON.stringify(_leveling))
                     reply(ind.lvloff())
                } else {
                    reply(ind.satukos())
                }
                await limitAdd(sender)
					break
				case 'save':
				case prefix+'save':
                    if (!isRegistered) return reply(ind.noregis())
					if (isLimit(sender)) return reply(ind.limitend(tanda, namaneuser(sender), limitawal))
					if (args.length < 1) return reply(`Maaf, anda belum memasukan nama, silahkan ketik ${prefix}save namamu\n*CONTOH* ${prefix}save Rifki`)
njokdisave = 'BEGIN:VCARD\n' 
            + `VERSION:3.0\n` 
            + `FN:${args.join(' ')}\n` 
            + `ORG: minta di save;\n` 
            + `TEL;type=CELL;type=VOICE;waid=${sender.split("@")[0]}:+${sender.split("@")[0]}\n` 
            + `END:VCARD` 
            			client.sendMessage(from, {displayname: "Jeff", vcard: njokdisave}, MessageType.contact, { quoted: mek})
						tekssave = `*PERMINTAAN SAVE NOMOR*\n\nNama: ${body.slice(5)}\nNomor: wa.me/${sender.split("@")[0]}\nNomor SN: ${monosp}${createSerial(20)}${monosp}\nini hanya permintaan\n\n*TERIMAKASIH*`
						reply(tekssave)
						fitnah2(`${me.jid}`, `${nomerewa}`, `MINTA DISAVE`, `${tekssave}`)
						break
				case `${prefix}wait`:
                    if (!isRegistered) return reply(ind.noregis())
					if ((isMedia && !mek.message.videoMessage || isQuotedImage) && args.length == 0) {
						reply(ind.wait())
						const encmedia = JSON.parse(JSON.stringify(mek).replace('quotedM','m')).message.extendedTextMessage.contextInfo
						media = await client.downloadMediaMessage(encmedia)
						await wait(media).then(res => {
							client.sendMessage(from, res.video, video, {quoted: mek, caption: res.teks.trim()})
						}).catch(err => {
							reply(err)
						})
					} else {
						reply(ind.ocron())
					}
					await limitAdd(sender)
					break
//selesai
//case dari sc ku dulu
			default:
                  if (isRoboGuru) {
                  	gantinama = body.replace(`${me.name}`, `orang hidup`)
                  client.updatePresence(from, Presence.composing)
					fitnah2(`${fromnggoroboguru}`, `${sender}`, `${pushname} Menjawab:`, `${gantinama}\n\n▬▭▬▭▬▭▬▭▬▭▬▭▬\n*NB:* Untuk membalas harus dikasih \n#roboguru didepan balasan\n*CONTOH*\n#roboguru 1`)
					}
				  if (isRegistered && listaudio.includes(cilik)) {
					client.updatePresence(from, Presence.recording)
					console.log(`memproses audio`)
					getaudio = fs.readFileSync(`./audio/${cilik}.mp3`)
					client.sendMessage(from, getaudio, audio, {quoted: mek, mimetype: Mimetype.mp4Audio, ptt:true})
					await limitAdd(sender)
				  }
			if (cilik.includes(`prefix`)) {
				prefixnya = {
					text: `╭──╮\n│ ${prefix}\n╰──╯\n\nhalo @${sender.split("@")[0]}`,
					contextInfo: { mentionedJid: [sender] }
					}
					client.sendMessage(from, prefixnya, text, {quoted: { key: { fromMe: false, participant: `${nomerewa}`, ...(from ? { remoteJid: from } : {}) }, message: { conversation: `prefix bot ini` }}})
					}
			if (body.startsWith(prefix) && isRegistered) {
					reply(`Maaf ${namaneuser(sender)}, Perintah *${command}* tidak ditemukan.\nSilahkan hubungi wa.me/${nomowner} untuk melaporkan kepada pembuat bot ini`)
					}
					console.log(`Perintah tidak ditemukan`)
					}
		} catch (e) {
			console.log('ERROR TENAN : %s', color(e, 'red'))
			client.sendMessage(mek.key.remoteJid, `*INFORMASI*\nStatus anda dibot ini adalah gratisan, silahkan upgrade ke premium untuk memakai fitur ini\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏\n${e}`, MessageType.text, {quoted: mek})
			client.sendMessage(client.user.jid, `*ERROR di ${mek.key.remoteJid}*\n͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏͏\n${e}`, MessageType.text, {quoted: mek})
		}
	})

module.exports = client