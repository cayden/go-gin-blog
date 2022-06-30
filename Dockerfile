FROM scratch

WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-blog
COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-blog

EXPOSE 8000
CMD ["./go-gin-blog"]