package cmd

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/termie/go-shutil"
)

var src, tar, suffix string

var RootCmd = &cobra.Command{
	Use: "app",
}

var cmdCopy = &cobra.Command{
	Use:   "copy [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		copyfile(src, tar, suffix)
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&src, "src", "C:\\code\\src\\user\\cmd", "Author name for copyright attribution")
	RootCmd.PersistentFlags().StringVar(&tar, "tar", "C:\\test", "Author name for copyright attribution")
	RootCmd.PersistentFlags().StringVar(&suffix, "suffix", ".g", "Author name for copyright attribution")

	// 两个顶层的命令，和一个cmdEcho命令下的子命令cmdTimes
	RootCmd.AddCommand(cmdCopy)

}

func Run() {
	RootCmd.Execute()
}

func copyfile(src_dir, tar_dir, suffix string) {
	var listfile []string
	//os.Getenv("dir")
	ostype := os.Getenv("GOOS") // 获取系统类型
	var strRet string
	if ostype == "windows" {
		strRet = "\\"
	} else if ostype == "linux" || ostype == "darwin" {
		strRet = "/"
	}
	log.Info(strRet)
	filepath.Walk(src_dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		//fmt.Println("file:", info.Name(), "in directory:", path)
		ok := strings.HasSuffix(path, suffix)

		if ok {
			listfile = append(listfile, path)
			println(suffix+"  file", path)
		}
		return nil
	})
	// And again without clearing the files
	for _, a := range listfile {
		log.Info(a)
		new_filename := filepath.Base(a)
		err := shutil.CopyFile(string(a), tar_dir+strRet+new_filename, false)
		log.Info(tar_dir + strRet + new_filename)
		if err != nil {
			log.Error(err)
		}
	}
	log.Info(listfile[0])
}
