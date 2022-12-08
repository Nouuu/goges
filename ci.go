package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// define the build
	buildPath := "build"
	bin := "goges"
	oses := []string{"linux", "darwin"}
	archs := []string{"amd64", "arm64"}

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer ctx.Done()
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory(".")

	// create empty directory for the build
	output := client.Directory()

	// get `golang` image
	golang := client.Container().From("golang:1.18")

	// mount cloned repository into `golang` image
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	for _, goos := range oses {
		for _, arch := range archs {
			// build binary for each OS/Arch combination
			path := fmt.Sprintf("%s/%s-%s-%s", buildPath, bin, goos, arch)

			build := golang.WithEnvVariable("GOOS", goos)
			build = build.WithEnvVariable("GOARCH", arch)

			// define the application build command
			golang = build.WithExec([]string{"go", "build", "-o", path, "./main.go"})

		}
	}

	// special build for Windows
	path := fmt.Sprintf("%s/%s-windows-amd64.exe", buildPath, bin)
	golang = golang.WithExec([]string{"go", "build", "-o", path, "./main.go"})

	// get reference to build output directory in container
	output = output.WithDirectory(buildPath, golang.Directory(buildPath))

	// write contents of container build/ directory to the host
	_, err = output.Export(ctx, ".")
	if err != nil {
		return err
	}

	return nil
}
