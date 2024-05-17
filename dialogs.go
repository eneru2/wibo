package main

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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