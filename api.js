__path = process.cwd()
var express = require('express');
var router = express.Router();
const moment = require("moment-timezone") 
const fs = require("fs") 
const fetch = require('node-fetch')
const axios = require('axios')
const chalk = require('chalk')
var brainly = require('brainly-scraper');
const { spawn, exec } = require("child_process")
const color = (text, color) => {
    return !color ? chalk.green(text) : chalk.keyword(color)(text)
}
const bgcolor = (text, bgcolor) => {
	return !bgcolor ? chalk.green(text) : chalk.bgKeyword(bgcolor)(text)
}
const getBuffer = async (url, options) => {
	try {
		options ? options : {}
		const res = await axios({
			method: "get",
			url,
			headers: {
				'DNT': 1,
				'Upgrade-Insecure-Request': 1
			},
			...options,
			responseType: 'arraybuffer'
		})
		return res.data
	} catch (e) {
		res.json({'result':`terjadi kesalahan \n${e}`})
	}
}
const fetchJson = (url, options) => new Promise(async (resolve, reject) => {
    fetch(url, options)
        .then(response => response.json())
        .then(json => {
            // console.log(json)
            resolve(json)
        })
        .catch((err) => {
            reject(err)
        })
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
	res.send('HELEH')
	})
router.get('/login', (req, res) => {
    res.download('https://frmdeveloper.github.io/frmdev/login.html')
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
			const alihkan = (url) => {
				res.redirect(url)
			}
			const unduh = (link) => {
				res.download(link)
			} 
			const reply = (teks) => {
				res.json({'result':teks})
				.catch(e => {
res.json({'result':'ERROR'})
})
			}
			const kirim = (link) => {
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
    	default:
    	reply(`*${command}* \n tidak ditemukan`)
	}
	} catch (e) {
		res.json({'result':`terjadi kesalahan \n${e}`})
	}
})

module.exports = router