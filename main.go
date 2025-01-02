package main

import (
	_ "embed"

	"github.com/lyfeblocnetwork/lyfebloc-edge/command/root"
	"github.com/lyfeblocnetwork/lyfebloc-edge/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
