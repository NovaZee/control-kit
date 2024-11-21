//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var Default = Build

var (
	appName  = "control-plane"
	srcDir   = "./cmd/"
	buildDir = "./bin"
	logDir   = "./log"
	pkg      = "./..."
)

// Build builds the binary for the current platform.
// 编译当前平台的二进制文件
func Build() error {
	fmt.Println("Building", appName, "...")
	os.MkdirAll(buildDir, os.ModePerm)
	cmd := exec.Command("go", "build", "-ldflags=-w -s", "-o", filepath.Join(buildDir, appName), srcDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// BuildWindows builds the binary for Windows.
// 编译 Windows 平台的二进制文件
func BuildWindows() error {
	fmt.Println("Building", appName, "for Windows...")
	os.MkdirAll(buildDir, os.ModePerm)
	cmd := exec.Command("go", "build", "-ldflags=-w -s", "-o", filepath.Join(buildDir, appName+".exe"), srcDir)
	cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Run() error {
	fmt.Println("Running", appName, "...")
	cmd := exec.Command(filepath.Join(buildDir, appName))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Test() error {
	fmt.Println("Running tests...")
	cmd := exec.Command("go", "test", pkg, "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Fmt() error {
	fmt.Println("Formatting code...")
	cmd := exec.Command("go", "fmt", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Removes build and log files.
// 删除构建和日志文件
func Clean() error {
	fmt.Println("Cleaning build file...")
	err := os.RemoveAll(buildDir)
	if err != nil {
		return err
	}
	fmt.Println("Cleaning log file...")
	err = os.RemoveAll(logDir)
	if err != nil {
		return err
	}
	return nil
}

func Deps() error {
	fmt.Println("Installing dependencies...")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func UpdateDeps() error {
	fmt.Println("Updating dependencies...")
	cmd := exec.Command("go", "get", "-u", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return Deps()
}

// Release creates a tarball of the binary.
// 编译并打包二进制文件
func Release() error {
	if err := Clean(); err != nil {
		return err
	}
	if err := Build(); err != nil {
		return err
	}
	fmt.Println("Creating release...")
	cmd := exec.Command("tar", "-czvf", filepath.Join(buildDir, appName+".tar.gz"), "-C", buildDir, appName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
