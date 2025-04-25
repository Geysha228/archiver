package cmd

import (
	"archiver/cmd/lib"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use: "pack",
	Short: "pack your file",
	Run: pack, 
}

const packedExtension  = "vlc"

var ErrEmptyPath = errors.New("not found a path to file")


func pack(_ *cobra.Command, args []string){

	if len(args[0]) == 0 || args[0] == ""{
		handleErr(ErrEmptyPath)
	} 

	filePath := args[0]

	r, err := os.Open(filePath)

	if err != nil {
		handleErr(err)
	}

	defer r.Close()
	
	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := lib.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName (path string) string {
	fileName := filepath.Base(path)
	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)
	return baseName + "." + packedExtension
}

func init(){
	rootCmd.AddCommand(packCmd)
}