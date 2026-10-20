
go env
go build -gcflags="-N -l=4"  -ldflags=" -extldflags='-static'" -o demo.exe
go build -gcflags="-N -l=4"  -ldflags="-s -w -extldflags='-static'" -o demo2.exe
pause