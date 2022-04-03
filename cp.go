package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// import (
// 	"io"
// 	"log"
// 	"os"
// )

// func mainnn() {

// 	sourceFile, err := os.Open("img/a.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sourceFile.Close()

// 	// Create new file
// 	newFile, err := os.Create("copy/a.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer newFile.Close()

// 	_, err = io.Copy(newFile, sourceFile) // check first var for number of bytes copied
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 	for _, sourceFile := range newFile {
// 	// 	filename = file.Name()
// 	// 	filenameNew = strings.Split(filename, "_")[0]
// 	// 	println(filename, filenameNew)
// 	// 	original_path = "img/" + filename
// 	// 	new_path = "copy/" + filenameNew
// 	// 	rename_file(original_path, new_path)
// 	// }

// }

// func main() {
// 	oldName := "sourceFile"
// 	newName := "newFile"
// 	err := os.Rename(oldName, newName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// // import (
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"os"
// // )

// // func main() {

// // 	dir := "img"+ os.DirEntry()
// // 	dst := "copy"

// // 	srcFile, err := os.Open(dir)
// // 	check(err)
// // 	defer srcFile.Close()

// // 	destFile, err := os.Create(dst) // creates if file doesn't exist
// // 	check(err)
// // 	defer destFile.Close()

// // 	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
// // 	check(err)

// // 	Original_Path := destFile.Name()
// // 	New_Path := Original_Path + "-1.png"
// // 	e := os.Rename(Original_Path, New_Path)
// // 	if e != nil {
// // 		log.Fatal(e)
// // 	}

// // 	err = destFile.Sync()
// // 	check(err)
// // }

// // func check(err error) {
// // 	if err != nil {
// // 		fmt.Println("Error : %s", err.Error())
// // 		os.Exit(1)
// // 	}
// // }

// package main

// import (
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"strings"
// )

// func copy() {

// 	srcFile, err := os.Open("img/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer srcFile.Close()

// 	destFile, err := os.Create("copy/") // creates if file doesn't exist
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer destFile.Close()

// 	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
// func copy(original_path, new_path string) (int64, error) {
// 	sourceFileStat, err := os.Stat(original_path)
// 	if err != nil {
// 		return 0, err
// 	}

// 	if !sourceFileStat.Mode().IsRegular() {
// 		return 0, fmt.Errorf("%s is not a regular file", original_path)
// 	}

// 	source, err := os.Open(original_path)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer source.Close()

// 	destination, err := os.Create(new_path)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer destination.Close()
// 	nBytes, err := io.Copy(destination, source)
// 	return nBytes, err
// }

func copy(original_path string, new_path string) {

	srcFile, err := os.Open(original_path)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(new_path) // creates if file doesn't exist
	if err != nil {
		log.Fatal(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		log.Fatal(err)
	}

}

func rename_file(original_path string, new_path string) {
	e := os.Rename(original_path, new_path)
	if e != nil {
		println(e.Error())
	}
	// println(destFile, "to "+srcFile)
}

func main() {
	files, err := ioutil.ReadDir("img")
	if err != nil {
		println(err.Error())
	}

	var filename string
	var filenameNew string
	var original_path string
	// var destFile string
	// var srcFile string
	var new_path string

	for _, file := range files {
		filename = file.Name()
		filenameNew = strings.Split(filename, "_")[0]
		println(filename, "to "+filenameNew)
		original_path = "img/" + filename
		new_path = "copy/" + filenameNew
		copy(original_path, new_path)
	}
}
