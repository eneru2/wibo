package main

import (
	"context"
	"os/exec"
	"syscall"

	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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
}

type config struct {
	indenting_amount int8
	include_mime_types bool
	use_absolute_paths bool
	use_new_lines_for_properties bool
	output_formats []string
	avif_crf int8
	webp_quality int8

	output_resolutions []string
	use_height_for_output bool
	disable_tooltips bool
}

func (a *App) InitialConfigRead(send string) {
	configPath := configdir.LocalConfig("wibo")
	err := configdir.MakePath(configPath)
	if err != nil {
		panic(err)
	}
	config, _ := assets.ReadFile(send)
	println(string(config))
}



func (a *App) ReadConfig() map[string]any {
	config := viper.New()
	config.SetConfigFile("config.toml")
	config.ReadInConfig()
	return config.AllSettings()
}

func (a *App) WriteOutputFormats(format string, state bool){
	viper.Set(format, state)
	viper.WriteConfig()
}

func (a *App) MinimiseApp() {
	runtime.WindowMinimise(a.ctx)
}

func (a *App) CloseApp() {
	runtime.Quit(a.ctx)
}

func (a *App) Copy_To_Clipboard(code string) {
	runtime.ClipboardSetText(a.ctx, code)
}

func (a *App) Open_Input_File_Dialog() string {
	input_file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	
	if err != nil {
		return "error"
	}

	return input_file
}

func (a *App) Open_Output_Dir_Dialog() string {
	output_directory, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})

	if err != nil {
		return "error"
	}

	return output_directory
}

func (a *App) ConvertImage(input_file string, format []string, scale string, output_location string, output_schema string) string {
	ffmpeg_location := "./ffmpeg.exe"
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

	return "success"
}