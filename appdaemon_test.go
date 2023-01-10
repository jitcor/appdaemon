package appdaemon

import (
	"context"
	"github.com/shirou/gopsutil/process"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestDaemon(t *testing.T) {
	log.Println("self:",os.Args[0])
	log.Println("self:",filepath.Base(os.Args[0]))
	args:=os.Getenv("APP_DAEMON_ARGS")
	log.Println(strings.Split(args,","))
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			go func() {
				ctx,cancel:=context.WithCancel(context.Background())
				defer cancel()
				if processes,err:=process.ProcessesWithContext(ctx);err!=nil{
					log.Println("error: ",err)
				}else {
					count:=0
					for _, p := range processes {
						if name,err:=p.Name();err!=nil{
							log.Println("error: ",err)
						}else if name==filepath.Base(os.Args[0]){
							count++
						}
					}
					if count<2{
						args:=os.Getenv("APP_DAEMON_ARGS")
						if err=exec.Command(os.Args[0],strings.Split(args,",")...).Start();err!=nil{
							log.Println("error: ",err)
						}
					}
				}
			}()
		}
	}
}