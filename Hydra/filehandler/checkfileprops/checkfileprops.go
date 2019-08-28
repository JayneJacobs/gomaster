package checkfileprops

import (
	"fmt"
	"gotrain/goMaster/Hydra/filehandler/printerrors"
	"io"
	"os"
)

// GenerateFileStatusReport prints the file properties
func GenerateFileStatusReport(fname string) {
	// Stat returns file info. It will return
	// an error if there is no file.
	filestats, err := os.Stat("test3.txt")

	printerrors.PrintFatalError(err)

	fmt.Println("What's the file name?", filestats.Name())
	fmt.Println("Am I a directroy?", filestats.IsDir())
	fmt.Println("What are the permissions?", filestats.Mode())
	fmt.Println("What's the file size?", filestats.Size())
	fmt.Println("When was the last time the file modified? ", filestats.ModTime())

}

//CopyFile fname1 to fname2
func CopyFile(fname1, fname2 string) {
	fOld, err := os.Open(fname1)
	printerrors.PrintFatalError(err)
	defer fOld.Close()

	fNew, err := os.Create(fname2)
	printerrors.PrintFatalError(err)
	defer fNew.Close()

	//copy bytes from source to destination
	_, err = io.Copy(fNew, fOld)
	printerrors.PrintFatalError(err)

	// Sync commits the current contents of the file to stable storage.
	err = fNew.Sync()
	printerrors.PrintFatalError(err)
}

//GenerateFileStatusReport Gives a description of the file
