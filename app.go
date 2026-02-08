package main

import (
	"context"
	"io"
	"os"
	"slices"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetPicFitMode() string {
	f, err := os.OpenFile("./config", os.O_RDONLY, 0755)
	if err != nil {
		return "contain"
	}
	content, _ := io.ReadAll(f)
	if slices.Contains([]string{"contain", "cover", "fill"}, string(content)) {
		return string(content)
	} else {
		return "contain"
	}
}
