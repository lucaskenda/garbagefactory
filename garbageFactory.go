package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Keep running forever.
func forever() {
	for {
		fmt.Printf("keep alive...\n")
		time.Sleep(time.Second)
	}
}

// Create file.
func writeFile(fName string, fSize int64) error {

	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	const defaultBufSize = 4096
	buf := make([]byte, defaultBufSize)
	buf[len(buf)-1] = '\n'
	w := bufio.NewWriterSize(f, len(buf))

	start := time.Now()
	written := int64(0)
	for i := int64(0); i < fSize; i += int64(len(buf)) {
		nn, err := w.Write(buf)
		written += int64(nn)
		if err != nil {
			return err
		}
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	since := time.Since(start)

	err = f.Close()
	if err != nil {
		return err
	}
	fmt.Printf("- file: %s, written: %dB %dns %.2fGB %.2fs %.2fMB/s\n",
		fName, written, since,
		float64(written)/1000000000, float64(since)/float64(time.Second),
		(float64(written)/1000000)/(float64(since)/float64(time.Second)),
	)
	return nil
}

var size = flag.Int("size", 100, "file size in MiB")
var files = flag.Int("files", 1, "files to create")
var folder = flag.String("folder", "./garbage", "folder to store data")
var keepalive = flag.Bool("keepalive", false, "keep program running")

func main() {
	flag.Parse()

	// File size in mb.
	fSize := int64(*size) * (1024 * 1024)
	fSizeEnv, ok := os.LookupEnv("GARBAGE_FACTORY_FILE_SIZE")
	if ok {
		fSizeEnvInt, err := strconv.Atoi(fSizeEnv)
		if err != nil {
			panic(err)
		}
		fSize = int64(fSizeEnvInt) * (1024 * 1024)
	}

	// The number of files to create.
	fFiles := int(*files)
	fFilesEnv, ok := os.LookupEnv("GARBAGE_FACTORY_FILES_TO_CREATE")
	if ok {
		fFilesEnvInt, err := strconv.Atoi(fFilesEnv)
		if err != nil {
			panic(err)
		}
		fFiles = int(fFilesEnvInt)
	}

	// Base folder.
	fFolder := string(*folder)
	fFolderEnv, ok := os.LookupEnv("GARBAGE_FACTORY_FOLDER")
	if ok {
		fFolder = fFolderEnv
	}

	// Keep alive flag.
	fKeepAlive := bool(*keepalive)
	fKeepAliveEnv, ok := os.LookupEnv("GARBAGE_FACTORY_KEEP_ALIVE")
	if ok {
		fKeepAliveEnvBool, err := strconv.ParseBool(fKeepAliveEnv)
		if err != nil {
			panic(err)
		}
		fKeepAlive = fKeepAliveEnvBool
	}

	fmt.Printf("File size: %dmb\n", fSize/1024/1024)
	fmt.Printf("Files to create: %d\n", fFiles)
	fmt.Printf("Folder path: %s\n", fFolder)
	fmt.Printf("Keep alive: %s\n", strconv.FormatBool(fKeepAlive))
	fmt.Println("Starting files creation:")

	// Create a folder if don't exist.
	os.MkdirAll(fFolder, os.ModePerm)

	// Start creating all the files.
	var i = 0
	for i < fFiles {
		err := writeFile(fFolder+"/load-"+strconv.Itoa(i), fSize)
		if err != nil {
			fmt.Fprintln(os.Stderr, fSize, err)
		}
		i = i + 1
	}

	// Keep alive
	if fKeepAlive {
		go forever()
		select {}
	}

}
