var express = require('express');
var router = express.Router();

router.get('/docs', (req, res) => {
	res.send('HELEH')
}

router.get('/c', async (req, res, next) => {
	try {
	q = req.query.q
	command = req.query.c
switch(command) {
	case 'wa':
    res.redirect('wa.me/62895803265350')
    break
    case ''
    res.redirect('wa.me/62895803265350')
}
} catch (e) {
	res.json({'info':'Terjadi kesalahan'
			'result':e})
}

module.exports = router