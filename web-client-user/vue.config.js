module.exports = {
  publicPath: './', //配置为相对路径，这样打包之后的包才能正常使用
  devServer: {
    host: '0.0.0.0',
    hot: true, //是否启用热加载，就是每次更新代码，是否需要重新刷新浏览器才能看到新代码效果
    port: '8888', //服务启动端口
    open: false, //是否自动打开浏览器默认为false
    proxy: {
      //配置http代理
      '/api': {
        // 匹配所有以 '/api'开头的请求路径
        target: 'http://localhost:8080', // 代理目标的基础路径
        // secure: false,  // 如果是https接口，需要配置这个参数
        changeOrigin: true, // 支持跨域
        pathRewrite: {
          // 重写路径: 去掉路径中开头的'/api'
          '^/api': '',
        },
      },
    },
  },
}
