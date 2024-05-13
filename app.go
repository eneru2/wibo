package main

import (
	"context"
	_ "embed"
	"os/exec"
	"syscall"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	// "github.com/adrg/xdg"
)

//go:embed ffmpeg/ffmpeg.exe
var ffmpeg_embed string

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// configFilePath, err := xdg.ConfigFile("wibo/config.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// configFilePath, err = xdg.SearchConfigFile("wibo/config.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func (a *App) MinimiseApp() {
	runtime.WindowMinimise(a.ctx)
}

func (a *App) CloseApp() {
	runtime.Quit(a.ctx)
}

func (a *App) Open_Input_File_Dialog(name string) string {
	s, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{

	})
	if err != nil {
		return "error"
	}

	return s
}

func (a *App) Open_Output_Dir_Dialog(name string) string {
	s, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{

	})
	if err != nil {
		return "error"
	}

	return s
}

func (a *App) ConvertImage(input_file string, format []string, scale string, output_location string, output_schema string) string {
	// ffmpeg_location := "G:/Andrés/Download/ffmpeg-7.0-essentials_build/bin/ffmpeg"
	// ffmpeg_location := "./ffmpeg.exe"
	ffmpeg_location := "ffmpeg"
	ffmpeg_scale := "scale=" + scale + ":-1"
	for i := 0; i < len(format); i++ {
		ffmpeg_output := ""
		if output_schema == "" {
			ffmpeg_output = output_location + "/input_file." + format[i]
		} else {
			ffmpeg_output = output_location + "/" + output_schema + "." + format[i]
		}

		switch(format[i]){
			case "avif":
				ffmpeg_command := exec.Command(
					ffmpeg_location, "-i", input_file, "-still-picture", "1",
					"-pix_fmt", "yuv420p", "-vf", ffmpeg_scale, "-crf", "25", ffmpeg_output)
				
				ffmpeg_command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				stdout, err := ffmpeg_command.Output()
								
				if err != nil {
					return "error"
				}
				
				if string(stdout) == "dog" {
					
				}
				
			case "webp":
				ffmpeg_command := exec.Command(
					ffmpeg_location, "-i", input_file, "-vf", ffmpeg_scale, "-quality", "82", ffmpeg_output)

				ffmpeg_command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				stdout, err := ffmpeg_command.Output()
								
				if err != nil {
					return "error"
				}
				if string(stdout) == "dog" {
					
				}
				break
			case "jpg":
				ffmpeg_command := exec.Command(
					ffmpeg_location, "-i", input_file, "-vf", ffmpeg_scale, ffmpeg_output)

				ffmpeg_command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				stdout, err := ffmpeg_command.Output()
								
				if err != nil {
					return "error"
				}

				if string(stdout) == "dog" {
					
				}
				break
		}
	}

	// ffmpeg_command := exec.Command("G:/Andrés/Download/ffmpeg-7.0-essentials_build/bin/ffmpeg", "-i", "G:/Andrés/Download/ffmpeg-7.0-essentials_build/bin/ape.webp", "G:/Andrés/Download/ffmpeg-7.0-essentials_build/bin/sas2.jpg")
	return "success"
}