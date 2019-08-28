package main

import (
	"bufio"
	"fmt"
	"gotrain/goMaster/Hydra/filehandler/checkfileprops"
	"gotrain/goMaster/Hydra/filehandler/printerrors"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	//open a file for read only
	f1, err := os.Open("testfile3.txt")
	printerrors.PrintFatalError(err)
	defer f1.Close()
	go checkfileprops.WatchFile("testfile3.txt")
	/*
		//Create a new file
		f2, err := os.Create("test2.txt")
		printerrors.PrintFatalError(err)
		defer f2.Close()
	*/

	//open file for read write
	f3, err := os.OpenFile("testfile3.txt", os.O_APPEND|os.O_RDWR, 0666)
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening
	// os.O_CREATE|os.O_RDWR|os.O_WRONLY

	//0666 => Owner: (read & write), Group: (read & write), and other (read & write)
	printerrors.PrintFatalError(err)
	defer f3.Close()

	//rename a file current new
	//err = os.Rename("test1.txt", "test1New.txt")
	//printerrors.PrintFatalError(err)
	//copy a file current location new location
	checkfileprops.CopyFile("testfile3.txt", "test1.txt")
	time.Sleep(1 * time.Second)
	checkfileprops.CopyFile("testfile3.txt", "test2.txt")
	time.Sleep(1 * time.Second)

	//move a file current location new location
	err = os.Rename("./test1.txt", "./testfolder/test2.txt")
	printerrors.PrintFatalError(err)

	//copy a file current location new location
	checkfileprops.CopyFile("test2.txt", "./testfolder/test3.txt")

	//delete a file loction/name
	err = os.Remove("test2.txt")
	printerrors.PrintFatalError(err)

	bytes, err := ioutil.ReadFile("./testfolder/test3.txt")
	fmt.Println(string(bytes)) //string converts the bytes to a readable string
	//3 Scan the file into a slice of bytes
	scanner := bufio.NewScanner(f3)
	count := 0
	//Scan each line of the file
	for scanner.Scan() {
		count++
		fmt.Println("Found line:", count, scanner.Text())
	}
	//copy a file current location new location
	checkfileprops.CopyFile("testfile3.txt", "test1.txt")

	//copy a file current location new location
	checkfileprops.CopyFile("test1.txt", "test3.txt")

	//buffered write, efficient store in memory, saves disk I/O
	writebuffer := bufio.NewWriter(f3)
	for i := 1; i <= 5; i++ {
		//Take the bytes from a buffer and add a  ine
		writebuffer.WriteString(fmt.Sprintln("Added line", i))
	}
	writebuffer.Flush() // this line commits contents of the buffer to disk

	checkfileprops.GenerateFileStatusReport("test3.txt")

	filestat1, err := os.Stat("test3.txt")
	printerrors.PrintFatalError(err)
	for {
		time.Sleep(1 * time.Second)
		filestat2, err := os.Stat("test3.txt")
		printerrors.PrintFatalError(err)
		if filestat1.ModTime() != filestat2.ModTime() {
			fmt.Println("File was modified at", filestat2.ModTime())
			filestat1, err = os.Stat("test3.txt")
			printerrors.PrintFatalError(err)
		}
	}
}
