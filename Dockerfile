FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-goji
COPY ./main /go/note-goji
COPY ./application-docker.yaml /go/note-goji

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18084
ENTRYPOINT ["./main"]
