const resolve = require('path').resolve

module.exports = {
  target: 'node',
  mode: 'production',
  node: {
    __dirname: false,
    __filename: false,
  },
  entry: [ './src/server.js' ],
  output: {
    filename: "server.js",
    path: resolve(__dirname, 'dist')
  },
  externals:[
    function(context, request, callback) {
      if(request[0] == '.') {
        callback();
      } else {
        callback(null, "require('" + request + "')");
      }
    }
  ]
}