package utils

import (
	"github.com/Mandala/go-log"
	"os"
)

var (
	Logger = log.New(os.Stdout).WithColor().WithDebug()
)