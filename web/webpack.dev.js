import path from "path";
import HtmlWebpackPlugin from "html-webpack-plugin";
import WebpackMerge from "webpack-merge";

import baseConfig from "./webpack.base.js";

const __dirname = path.dirname("");

const devConfig = {
  mode: "development",
  devtool: "inline-source-map",
  entry: {
    index: path.resolve(__dirname, "src/web/index.tsx"),
  },
  output: {
    filename: "[name].[chunkhash].js",
    path: path.resolve(__dirname, "dist/web"),
  },
  module: {},
  devServer: {
    compress: true,
    host: "127.0.0.1",
    port: 4000,
    liveReload: true,
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, "template.html"),
      filename: "index.html",
      chunks: ["index"],
    }),
  ],
};

export default WebpackMerge.merge(baseConfig, devConfig);
