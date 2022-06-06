package utils

import (
	"bytes"
	"context"
	"douyin-lite/pkg/configs"
	"errors"
	"os/exec"
	"time"
)

func MakeSnapshot(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cmd := exec.CommandContext(ctx, configs.FFmpegPath,
		"-loglevel", "error",
		"-y",
		"-ss", "1",
		"-t", "1",
		"-i", "file/play/"+id,
		"-vframes", "1",
		"-f", "mjpeg",
		"file/cover/"+id)
	defer cancel()
	var Errbuf bytes.Buffer
	cmd.Stderr = &Errbuf
	err := cmd.Run()
	if err != nil {
		return err
	}
	if Errbuf.Len() != 0 {
		return errors.New(Errbuf.String())
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return nil
}
