var express = require('express');
const { response } = require('../app');
var router = express.Router();

/* GET home page. */
router.get('/', function (req, res, next) {
  res.json({
    code: 1, message: "helloworld"
  });
});

module.exports = router;