

## 安裝 compiler

#### 安裝 protoc CLI 工具
``` 
$ brew install protobuf 
$ protoc --version
```

#### 安裝 protoc-gen-go 可以將 proto buffer 編譯成 Golang 可使用的檔案
```
$ go get github.com/golang/protobuf/protoc-gen-go
```

#### 安裝 grpc-go 後，可以在 Golang 中使用 gRPC
```
$ go get -u google.golang.org/grpc
```

## proto 資料夾下執行
#### 編譯 Protocol Buffer 檔案
```
$ protoc *.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative
```

## 啟用 Golang Server
```
go run ./main.go
```