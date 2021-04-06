FROM golang 
#添加数据券
VOLUME [ "/usr/local/go/src/static" ]
#将此文件夹的内容复制到...
ADD . /usr/local/go/src/web-server
#工作目录
WORKDIR /usr/local/go/src/web-server
#设置代理
RUN go env -w GOPROXY=https://goproxy.io,direct
#环境变量
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 
#编译Golang
RUN go build
#端口号
EXPOSE 8080