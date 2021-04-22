var express = require('express');
var router = express.Router();

router.get('/', (req, res) => {
    res.redirect('wa.me/62895803265350')
})

module.exports = router