package checkfileprops

import (
	"fmt"
	"gotrain/goMaster/Hydra/filehandler/printerrors"
	"os"
	"time"
)

// WatchFile takes a file name string and creates a watcher
func WatchFile(fname string) {
	filestat1, err := os.Stat(fname)
	printerrors.PrintFatalError(err)
	for {
		time.Sleep(1 * time.Second)
		filestat2, err := os.Stat(fname)
		printerrors.PrintFatalError(err)
		if filestat1.ModTime() != filestat2.ModTime() {
			fmt.Println("File was modified at ", filestat2.ModTime())
			filestat1, err = os.Stat(fname)
			printerrors.PrintFatalError(err)
		}
	}
}
