package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/termie/go-shutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var listfile []string
	//os.Getenv("dir")
	filepath.Walk("C:\\code\\src\\user", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		//fmt.Println("file:", info.Name(), "in directory:", path)
		ok := strings.HasSuffix(path, ".go")
		if ok {
			listfile = append(listfile, path)
			println("Golang file", path)
		}
		return nil
	})
	// And again without clearing the files
	for _, a := range listfile {
		log.Info(a)
		new_filename := filepath.Base(a)
		err := shutil.CopyFile(string(a), "C:\\test\\"+new_filename, false)
		if err != nil {
			log.Error(err)
		}
	}

	println(listfile[0])
}

/**
// filelist.go
package main

import (
    //"flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

var (
    ostype = os.Getenv("GOOS") // 获取系统类型
)

var listfile []string //获取文件列表

func Listfunc(path string, f os.FileInfo, err error) error {
    var strRet string
    strRet, _ = os.Getwd()
    //ostype := os.Getenv("GOOS") // windows, linux

    if ostype == "windows" {
        strRet += "\\"
    } else if ostype == "linux" {
        strRet += "/"
    }

    if f == nil {
        return err
    }
    if f.IsDir() {
        return nil
    }

    strRet += path //+ "\r\n"

    //用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
    ok := strings.HasSuffix(strRet, ".go")
    if ok {

        listfile = append(listfile, strRet) //将目录push到listfile []string中
    }
    //fmt.Println(ostype) // print ostype
    fmt.Println(strRet) //list the file

    return nil
}

func getFileList(path string) string {
    //var strRet string
    err := filepath.Walk(path, Listfunc) //

    if err != nil {
        fmt.Printf("filepath.Walk() returned %v\n", err)
    }

    return " "
}

func ListFileFunc(p []string) {
    for index, value := range p {
        fmt.Println("Index = ", index, "Value = ", value)
    }
}

func main() {
    //flag.Parse()
    //root := flag.Arg(0)
    //fmt.Println()
    var listpath string
    fmt.Scanf("%s", &listpath)
    getFileList(listpath)
    ListFileFunc(listfile)
    //getFileList(root)

}
**/
