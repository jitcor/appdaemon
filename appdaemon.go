package appdaemon

import (
	"github.com/Humenger/go-devcommon/dcshell"
	"log"
	"os"
	"strings"
	"time"
)

func Start() {
	dcShell := dcshell.NewDcShell(true, checkDebug())
	if len(os.Args) == 2 {
		packageNames := strings.Split(os.Args[1], "/")
		log.Println("packageNames:", packageNames)
		for {
			daemonAppImpl(packageNames, dcShell)
		}
	} else if len(os.Args) == 3 {
		if os.Args[1]!="-s"{
			log.Fatalln("appdaemon -s package:port/package:port")
			return
		}
		packagePorts := strings.Split(os.Args[2], "/")
		log.Println("packagePorts:", packagePorts)
		for {
			daemonServerImpl(packagePorts, dcShell)
		}
	} else {
		log.Fatalln("appdaemon has must arg")
		return
	}

}

func checkDebug() bool {
	for _, env := range os.Environ() {
		if strings.Contains(env,"="){
			key:=strings.Trim(strings.Split(env,"=")[0]," ")
			value:=strings.Trim(strings.Split(env,"=")[1]," ")
			if strings.Contains(strings.ToUpper(key),"ADDEBUG"){
				if strings.Contains(strings.ToLower(value),"true"){
					return true
				}else {
					return false
				}
			}
		}

	}
	return false
}
func daemonAppImpl(packages []string, shell *dcshell.DcShell) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("daemonAppImpl.err:", err.(error).Error())
		}
	}()
	for _, pkgName := range packages {
		if !shell.LaunchAppWhenStopped(pkgName) {
			log.Println("launch app failed:", pkgName)
		}
	}
	time.Sleep(100 * time.Millisecond)
}
func daemonServerImpl(packagePorts []string, shell *dcshell.DcShell) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("daemonServerImpl.err:", err.(error).Error())
		}
	}()
	for _, packagePort := range packagePorts {
		pkg:=strings.Split(packagePort,":")[0]
		port:=strings.Split(packagePort,":")[1]
		if !shell.CheckPortIsListen(port){
			if !shell.LaunchApp(pkg) {
				log.Println("launch app failed:", pkg)
			}
		}

	}
	time.Sleep(100 * time.Millisecond)
}
