# appdaemon
android app process daemon with golang.
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
```
