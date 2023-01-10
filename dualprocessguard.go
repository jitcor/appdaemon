package appdaemon

import (
	"fmt"
	"github.com/Humenger/go-devcommon/dcmd"
	"github.com/shirou/gopsutil/process"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type DualProcessGuard struct {
	interval time.Duration
	task     func()
}
func (that*DualProcessGuard)Start(task func())  {
	that.task = task
	if that.interval ==0{
		that.interval =10*time.Second
	}
	if !strings.HasSuffix(os.Args[0],"_s"){
		that.enterMainProcess()
	}else {
		that.enterSubProcess()
	}
}
func (that *DualProcessGuard) SetInterval(interval time.Duration) *DualProcessGuard {
	that.interval =interval
	return that
}
func (that *DualProcessGuard)enterMainProcess() {
	go that.daemonSub()
	that.task()
}

func (that *DualProcessGuard)enterSubProcess() {
	go that.daemonMain()
	that.task()
}

func (that *DualProcessGuard)subProcessStartup() {
	dcmd.Exec("cp",fmt.Sprintf("-r %s %s_s",os.Args[0],os.Args[0]),new(error))
	if len(os.Args)==1{
		exec.Command(os.Args[0]+"_s").Start()
	}else {
		exec.Command(os.Args[0]+"_s",os.Args[1:]...).Start()
	}
	dcmd.Exec("rm",fmt.Sprintf("-rf %s_s",os.Args[0]),new(error))
	log.Println("sub process start: ",os.Args[0]+"_s")
}
func (that *DualProcessGuard)mainProcessStartup() {
	if len(os.Args)==1{
		exec.Command(os.Args[0][:len(os.Args[0])-2]).Start()
	}else {
		exec.Command(os.Args[0][:len(os.Args[0])-2],os.Args[1:]...).Start()
	}

	log.Println("main process start: ",os.Args[0][:len(os.Args[0])-2])
}

func (that *DualProcessGuard)daemonMain() {
	mainProcessRunning :=false
	for {
		select {
		case <-time.After(that.interval):
			if processes,err:=process.Processes();err!=nil{
				log.Println("daemonSub: error",err)
			}else {
				mainProcessRunning =false
				for _, p := range processes {
					if name,err:=p.Name();err!=nil{
						log.Println("daemonMain: error: ",err)
					}else if name+"_s"==filepath.Base(os.Args[0]){
						log.Println("sub process:",os.Args[0]," found main process:",name)
						mainProcessRunning=true
					}
				}
				if !mainProcessRunning {
					that.mainProcessStartup()
				}

			}
		}
	}
}
func (that *DualProcessGuard)daemonSub() {
	subProcessRunning:=false
	for {
		select {
		case <-time.After(that.interval):
			if processes,err:=process.Processes();err!=nil{
				log.Println("daemonSub: error",err)
			}else {
				subProcessRunning=false
				for _, p := range processes {
					if name,err:=p.Name();err!=nil{
						log.Println("daemonSub: error: ",err)
					}else if name==filepath.Base(os.Args[0])+"_s"{
						log.Println("process:",os.Args[0]," found sub process:",name)
						subProcessRunning=true
						break
					}
				}
				if !subProcessRunning{
					that.subProcessStartup()
				}
			}
		}
	}
}
