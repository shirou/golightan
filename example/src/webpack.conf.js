var webpack = require('webpack')
var path = require('path')

module.exports = {
    context: __dirname,
    entry: {
        default: './index.ts',
    },
    output: {
      path: path.resolve(__dirname, "../assets/js"),
      filename: 'app.js',
    },
    resolve: {
        extensions: ['.js'],
        modules: [
            path.resolve(__dirname, "node_modules")
        ]
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                loader: 'ts-loader',
                include: [__dirname],
                options: {
                    presets: ['es2015'],
                    "plugins": [
                    ]
                }
            }
        ]
    },
    plugins: [
      new webpack.optimize.AggressiveMergingPlugin()
    ]
}
