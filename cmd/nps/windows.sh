#/bin/sh
mkdir server
go env -w CGO_ENABLED=0
go env -w GOARCH=amd64
go env -w GOOS=windows
go build -ldflags="-s -w" -a -v -o win_amd64_nps.exe nps.go
#upx win_amd64_nps.exe
mv win_amd64_nps.exe server
go env -w GOOS=linux
go build -ldflags="-s -w" -a -v -o linux_amd64_nps nps.go
#upx linux_amd64_nps
mv linux_amd64_nps server
go env -w GOOS=darwin
go build -ldflags="-s -w" -a -v -o darwin_amd64_nps nps.go
#upx darwin_amd64_nps
mv darwin_amd64_nps server

go env -w GOARCH=386
go env -w GOOS=windows
go build -ldflags="-s -w" -a -v -o win_amd86_nps.exe nps.go
#upx win_amd86_nps.exe
mv win_amd86_nps.exe server
go env -w GOOS=linux
go build -ldflags="-s -w" -a -v -o linux_amd86_nps nps.go
#upx linux_amd86_nps
mv linux_amd86_nps server
go env -w GOOS=darwin
go build -ldflags="-s -w" -a -v -o darwin_amd86_nps nps.go
#upx darwin_amd86_nps
mv darwin_amd86_nps server

go env -w GOARCH=arm64
go env -w GOOS=windows
go build -ldflags="-s -w" -a -v -o win_arm64_nps.exe nps.go
#upx win_arm64_nps.exe
mv win_arm64_nps.exe server
go env -w GOOS=linux
go build -ldflags="-s -w" -a -v -o linux_arm64_nps nps.go
#upx linux_arm64_nps
mv linux_arm64_nps server
go env -w GOOS=darwin
go build -ldflags="-s -w" -a -v -o darwin_arm64_nps nps.go
#upx darwin_arm64_nps
mv darwin_arm64_nps server
zip -r Nps_Server.zip server
# rm -rf server

go env -w CGO_ENABLED=1
go env -w GOOS=darwin
go env -w GOARCH="arm64"
