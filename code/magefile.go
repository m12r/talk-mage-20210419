//+build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

var Default = Build

// Build builds the executable and places it in dist
func Build() error {
	mg.Deps(requireGoVersion, Web.Build)

	return sh.RunWith(defaultEnv(), "go", "build", "-o", "dist/server")
}

// Clean cleans the project of the built artifacts
func Clean() error {
	return sh.Rm("dist")
}

// DistClean cleans the project of all built artifacts
func DistClean() {
	mg.Deps(Clean, Web.DistClean)
}

type Web mg.Namespace

// Install installs the javascript dependencies
func (Web) Install() error {
	origDir, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.Chdir("web"); err != nil {
		return err
	}
	defer os.Chdir(origDir)

	isDirty, err := target.Dir("node_modules", "package.json", "package-lock.json")
	if err != nil {
		return err
	}
	if !isDirty {
		return nil
	}

	return sh.RunWith(defaultEnv(), "npm", "ci")
}

// Build builds the javascript project
func (Web) Build() error {
	mg.Deps(Web.Install)

	origDir, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := os.Chdir("web"); err != nil {
		return err
	}
	defer os.Chdir(origDir)

	isDirty, err := target.Dir("dist", "public", "src")
	if err != nil {
		return err
	}
	if !isDirty {
		return nil
	}

	return sh.RunWith(defaultEnv(), "npm", "run", "build")
}

// Clean removes the previously built javascript project
func (Web) Clean() error {
	return sh.Rm("web/dist")
}

// DistClean like clean, but also removes node_modules
func (Web) DistClean() error {
	mg.Deps(Web.Clean)

	return sh.Rm("web/node_modules")
}

func defaultEnv() map[string]string {
	return map[string]string{
		"GO111MODULE": "on",
	}
}
