package app

import (
		"github.com/spf13/cobra"
	"github.com/termie/go-shutil"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)
var src string

func Run() {

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),

		//Run: func(cmd *cobra.Command, args []string) {
		//	fmt.Println("Print: " + strings.Join(args, " "))
		//},
		Run: func(cmd *cobra.Command, args []string) {
			//copyfile("C:\\code\\src\\user","C:\test")
			src := args[0]
			log.Info(src)
		},
	}

	var rootCmd = &cobra.Command{Use: "app111"}
	rootCmd.AddCommand(cmdPrint)
	rootCmd.Execute()
	log.Info(src)
}

func copyfile( src_dir,tar_dir string  ) {
	var listfile []string
	//os.Getenv("dir")
	filepath.Walk(src_dir, func(path string, info os.FileInfo, err error) error {
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
		err := shutil.CopyFile(string(a), tar_dir+new_filename, false)
		if err != nil {
			log.Error(err)
		}
	}

	println(listfile[0])
}
