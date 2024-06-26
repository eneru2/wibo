// +build !windows
package main

import (
	"os/exec"
)

func (a *App) ConvertImage(input_file string, format []string, scale string, output_location string, output_schema string, avif_crf string, webp_quality string, output_height bool) {
	var ffmpeg_output string
	var ffmpeg_scale string

	if output_height {
		ffmpeg_scale = "scale=" + "-1:" + scale
	} else {
		ffmpeg_scale = "scale=" + scale + ":-1"
	}

	// ffmpeg_command := ffmpeg_location + " -i " + input_file + " -vf " + ffmpeg_scale
	
	for i := 0; i < len(format); i++ {
		if output_schema == "" {
			ffmpeg_output = output_location + "/input_file." + format[i]
			} else {
				ffmpeg_output = output_location + "/" + output_schema + "." + format[i]
			}

		args := []string{"-i", input_file, "-vf", ffmpeg_scale}

		switch(format[i]){
			case "avif":
				args = append(args, "-still-picture", "1", "-pix_fmt", "yuv420p", "-crf", avif_crf, ffmpeg_output)
				// ffmpeg_command += " -still-picture " + "1 " + "-pix_fmt " + "yuv420p " + "-crf " + avif_crf + " " + ffmpeg_output
			case "webp":
				args = append(args, "-quality", webp_quality, ffmpeg_output)
				// ffmpeg_command += " -quality" + " 82 " + ffmpeg_output
			case "jpg":
				args = append(args, ffmpeg_output)
				// ffmpeg_command += ffmpeg_output
		}

		ffmpeg_exec := exec.Command("ffmpeg", args...)
		_, err := ffmpeg_exec.Output()
								
		if err != nil {
			panic(err)
		}
	}
}
