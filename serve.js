//start webpack

const path = require('path');
const express = require('express');
const webpack = require('webpack');
const webpackMiddleware = require('webpack-dev-middleware');
const webpackHotMiddleware = require('webpack-hot-middleware');
const config = require('./webpack.config');

module.exports = () => {
  //create express
  const app = express();

  const isProduction = process.env.NODE_ENV === 'production';

  config.plugins = [
    new webpack.DefinePlugin({
      'process.env': { NODE_ENV: JSON.stringify(process.env.NODE_ENV) },
    }),
  ];

  //setup hot reload
  config.plugins.push(new webpack.HotModuleReplacementPlugin());
  //setup no errors-plugin
  config.plugins.push(new webpack.NoEmitOnErrorsPlugin());
  //override entry for hot reload
  config.entry = ['webpack-hot-middleware/client', config.entry];

  //return compiler instance
  const compiler = webpack(config);
  //stats output config
  const statsConfig = {
    colors: true,
    hash: false,
    timings: true,
    chunks: false,
    chunksModules: false,
    modules: false,
  };

  app.use(
    webpackMiddleware(compiler, {
      publicPath: config.output.publicPath,
      contentBase: 'src',
      stats: statsConfig,
    })
  );
  app.use(webpackHotMiddleware(compiler));

  //serve static files
  app.use(express.static(__dirname));
  //serve index
  app.get('*', (req, res) => res.sendFile(path.join(__dirname, 'index.html')));
  //start server
  app.listen(3000, err => {
    if (err) {
      console.log(err);
    }
    console.info('===> Listening to port 3000');
  });
};
