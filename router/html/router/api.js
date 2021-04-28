__path = process.cwd()
var express = require('express');
var router = express.Router();
const moment = require("moment-timezone") 
const fs = require("fs") 
const axios = require('axios')
const { spawn, exec } = require("child_process")
var { color, bgcolor } = require(__path + '/lib/color.js');
var { fetchJson } = require(__path + '/lib/fetcher.js');
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

router.get('/heleh', (req, res) => {
	res.send('HELEH')
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