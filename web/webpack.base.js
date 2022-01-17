import CleanWebpackPlugin from "clean-webpack-plugin";

const baseConfig = {
  resolve: {
    extensions: [".js", ".jsx", ".ts", ".tsx"],
    alias: {},
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx|ts|tsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "ts-loader",
        },
      },
      {
        test: /\.less$/,
        use: [
          {
            loader: "style-loader",
          },
          {
            loader: "css-loader",
          },
          {
            loader: "less-loader",
          },
        ],
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"],
      },
      {
        test: /\.(jpg|png|jpeg|gif)$/,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "[name]_[hash].[ext]",
              outputPath: "images/",
            },
          },
        ],
      },
    ],
  },
  plugins: [new CleanWebpackPlugin.CleanWebpackPlugin()],
};

export default baseConfig;
