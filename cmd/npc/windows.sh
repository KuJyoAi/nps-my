#/bin/sh
go env -w CGO_ENABLED=0
go env -w GOARCH=amd64
go env -w GOOS=windows
mkdir npc
go build -ldflags="-s -w" -a -v -o win_amd64_npc.exe npc.go
##upx win_amd64_npc.exe
mv win_amd64_npc.exe npc
go env -w GOOS=linux
go build -ldflags="-s -w" -a -v -o linux_amd64_npc npc.go
##upx linux_amd64_npc
mv linux_amd64_npc npc
go env -w GOOS=darwin
go build -ldflags="-s -w" -a -v -o darwin_amd64_npc npc.go
#upx darwin_amd64_npc
mv darwin_amd64_npc npc

go env -w GOARCH=386
go env -w GOOS=windows
go build -ldflags="-s -w" -a -v -o win_amd86_npc.exe npc.go
#upx win_amd86_npc.exe
mv win_amd86_npc.exe npc
go env -w GOOS=linux
go build -ldflags="-s -w" -a -v -o linux_amd86_npc npc.go
#upx linux_amd86_npc
mv linux_amd86_npc npc
go env -w GOOS=darwin
go build -ldflags="-s -w" -a -v -o darwin_amd86_npc npc.go
#upx darwin_amd86_npc
mv darwin_amd86_npc npc

go env -w GOARCH=arm64
go env -w GOOS=windows
go build -ldflags="-s -w" -a -v -o win_arm64_npc.exe npc.go
#upx win_arm64_npc.exe
mv win_arm64_npc.exe npc
go env -w GOOS=linux
go build -ldflags="-s -w" -a -v -o linux_arm64_npc npc.go
#upx linux_arm64_npc
mv linux_arm64_npc npc
go env -w GOOS=darwin
go build -ldflags="-s -w" -a -v -o darwin_arm64_npc npc.go
#upx darwin_arm64_npc
mv darwin_arm64_npc npc
zip -r Nps_Client.zip npc