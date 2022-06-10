# appdaemon
android app process daemon with golang.
# build
```
set GOARCH=arm
set GOOS=linux
go build -o appdaemon ./main/main.go
```
# use
```
adb push appdaemon /data/local/tmp/
adb shell
su
cd /data/local/tmp/
chmod +x ./appdaemon
./appdaemon your.app.package.name  //single app
./appdaemon your.app.package.name/your.app.package.name  //mutil app
./appdaemon your.app.package.name/your.app.package.name& //run in the background
// socket listen port daemon
./appdaemon -s your.app.package.name:port  //single app
./appdaemon -s your.app.package.name:port/your.app.package.name:port  //mutil app
./appdaemon -s your.app.package.name:port/your.app.package.name:port& //run in the background
```
