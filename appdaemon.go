package appdaemon

import (
	"github.com/Humenger/go-devcommon/dcshell"
	"log"
	"os"
	"strings"
	"time"
)

func Start()  {
	dcShell :=dcshell.NewDcShell(true,false)
	if len(os.Args)!=2{
		log.Fatalln("appdaemon has must arg")
		return
	}
	packageNames:=strings.Split(os.Args[1],"/")
	log.Println("package:",packageNames)
	for{
		daemonImpl(packageNames, dcShell)
	}
}
func daemonImpl(packages []string,shell *dcshell.DcShell) {
	defer func() {
		if err:=recover();err!=nil{
			log.Println("daemonImpl.err:",err.(error).Error())
		}
	}()
	for _, pkgName := range packages {
		if !shell.LaunchApp(pkgName){
			log.Println("launch app failed:",pkgName)
		}
	}
	time.Sleep(100*time.Millisecond)
}
