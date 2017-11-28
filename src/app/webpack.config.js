const path = require('path');
const webpack = require('webpack');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry: {
    app: './resources/js/main.js',
  },
  output: {
    path: path.resolve(__dirname, './public'),
    publicPath: '/public/',
    filename: 'js/[name].js',
    chunkFilename: 'js/[name].js',
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ExtractTextPlugin.extract({
          fallback: 'vue-style-loader',
          use: [
            'css-loader',
          ],
        }),
      },
      {
        test: /\.scss$/,
        use: ExtractTextPlugin.extract({
          fallback: 'vue-style-loader',
          use: [
            'css-loader',
            'sass-loader',
          ],
        }),
      },
      {
        test: /\.sass$/,
        use: ExtractTextPlugin.extract({
          fallback: 'vue-style-loader',
          use: [
            'css-loader',
            'sass-loader?indentedSyntax',
          ],
        }),
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options: {
          loaders: {
            // Since sass-loader (weirdly) has SCSS as its default parse mode, we map
            // the "scss" and "sass" values for the lang attribute to the right configs here.
            // other preprocessors should work out of the box, no loader config like this necessary.
            scss: ExtractTextPlugin.extract({
              fallback: 'vue-style-loader',
              use: [
                'css-loader',
                'sass-loader',
              ],
            }),
            sass: ExtractTextPlugin.extract({
              fallback: 'vue-style-loader',
              use: [
                'css-loader',
                'sass-loader?indentedSyntax',
              ],
            }),
          },
          // other vue-loader options go here
        },
      },
      {
        test: /\.js$/,
        loader: 'babel-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.(png|jpg|gif|svg)$/,
        loader: 'file-loader',
        options: {
          name: '[name].[ext]?[hash]',
        },
      },
    ],
  },
  resolve: {
    alias: {
      vue$: 'vue/dist/vue.esm.js',
    },
    extensions: ['*', '.js', '.vue', '.json'],
  },
  devServer: {
    historyApiFallback: true,
    noInfo: true,
    overlay: true,
    inline: true,
    open: false,
  },
  performance: {
    hints: false,
  },
  devtool: '#eval-source-map',
  plugins: [
    new ExtractTextPlugin({
      filename: 'css/[name].css',
      disable: false,
      allChunks: true,
    }),
    new webpack.optimize.CommonsChunkPlugin({ name: 'vendor', filename: 'js/vendor.js' }),
  ],
};

if (process.env.NODE_ENV === 'production') {
  module.exports.devtool = '#source-map';
  // http://vue-loader.vuejs.org/en/workflow/production.html
  module.exports.plugins = (module.exports.plugins || []).concat([
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: '"production"',
      },
    }),
    new webpack.optimize.UglifyJsPlugin({
      sourceMap: true,
      compress: {
        warnings: false,
      },
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
    }),
  ]);
}
