package cmd

/*
Copyright © 2024 Hendry Taylor hendry.taylor@icloud.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"errors"
	"fmt"
	"github.com/hostedbrains/toolbox"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "buildutil",
	Short: "This is a utility for building GO projects",
	Long: `This is a utility for building GO projects.
The buildutil CLI helps you to build your GO project in a quick and easy way.
The buildutil CLI allows you to increment your version parts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if incrementMajor {
			fmt.Println("Incrementing major version")
			incrementVersionFunc("major")
		}
		if incrementMinor {
			fmt.Println("Incrementing minor version")
			incrementVersionFunc("minor")
		}
		if incrementPatch {
			fmt.Println("Incrementing patch version")
			incrementVersionFunc("patch")
		}
		if initVersionFile {
			fmt.Println("Creating new version file")
			createVersionFile()
		}
		if setup {
			fmt.Println("Creating new buildutil.yaml file")
			createBuildutilConfigFile()
		}
		if build {
			fmt.Println("Building module")
			fmt.Println("Module output: " + output)

			buildModule(withLDFlags)
		}
		if version {
			fmt.Println("Version information")
			fmt.Println("Version: " + Version)
			fmt.Println("Build Date and Time: " + BuildTime)
			fmt.Println("Git Hash: " + GitHash)
		}
	},
}

var incrementMajor bool
var incrementMinor bool
var incrementPatch bool
var initVersionFile bool
var setup bool
var build bool
var version bool
var withLDFlags bool
var output string
var Version = "v1.0.2"
var BuildTime = "2024-05-30T19:43:41Z"
var GitHash = "5a96c6a"
var versionFile = "./.version"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, buildTime string, gitHash string) {
	// Set version information
	Version = version
	BuildTime = buildTime
	GitHash = gitHash

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)
	initConfig()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.testing.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&incrementMajor, "incrementMajor", "", false, "Increment the major version, default false")
	rootCmd.PersistentFlags().BoolVarP(&incrementMinor, "incrementMinor", "", false, "Increment the minor version, default false")
	rootCmd.PersistentFlags().BoolVarP(&incrementPatch, "incrementPatch", "", false, "Increment the patch version, default false")
	rootCmd.PersistentFlags().BoolVarP(&initVersionFile, "initVersion", "", false, "Create new version file with initial version of 0.0.1")
	rootCmd.PersistentFlags().BoolVarP(&setup, "setup", "", false, "Create new buildutil.yaml file with defaults.")
	rootCmd.PersistentFlags().BoolVarP(&build, "build", "b", false, "Build the module")
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Print version information")
	rootCmd.PersistentFlags().BoolVarP(&withLDFlags, "withLDFlags", "f", false, "Include LDFlags with the build.")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Build output e.g. '-o bin/moduleName' (required if build is set)")

	rootCmd.MarkFlagsRequiredTogether("build", "output")
	//fmt.Printf("Version file: %s\n", viper.GetString("buildutil.version.file"))
	//fmt.Printf("Version file path: %s\n", viper.GetString("buildutil.version.path"))
	versionFile = viper.GetString("buildutil.version.path") + viper.GetString("buildutil.version.file")
	//fmt.Printf("Version File: %s\n", versionFile)

	//rootCmd.MarkPersistentFlagRequired("region")
}

func checkVersionFile() {
	var tools toolbox.Tools

	if tools.CheckFileExist(versionFile) {
		// file exists
		//fmt.Printf("Version File exists: %s\n", versionFile)
	} else {
		// file does not exist
		fmt.Printf("Version File does not exist: %s\n", versionFile)
		fmt.Printf("Init version file flag value: %s\n", initVersionFile)
		fmt.Println("Run the 'buildutil --initVersion' first!")
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func incrementVersionFunc(semver string) {
	checkVersionFile()
	// Read the current version file
	dat, err := os.ReadFile(".version")

	check(err)
	fmt.Println(string(dat))
	// Extract version parts from the string as Major, Minor and Patch
	verString := trimFirstRune(string(dat))
	s := strings.Split(verString, ".")
	major, err := strconv.ParseInt(s[0], 0, 64)
	minor, err := strconv.ParseInt(s[1], 0, 64)
	patch, err := strconv.ParseInt(s[2], 0, 64)

	fmt.Printf("Current version = Major: %d, Minor: %d, Patch: %d\n", major, minor, patch)
	// Increment correct part of version
	switch semver {
	case "major":
		major++
	case "minor":
		minor++
	case "patch":
		patch++
	}
	fmt.Printf("New version = Major: %d, Minor: %d, Patch: %d\n", major, minor, patch)
	// Format version string as Major.Minor.Patch
	versionString := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	// write the new version to the file
	err = os.WriteFile(viper.GetString("buildutil.version.path")+viper.GetString("buildutil.version.file"), []byte(versionString), 0666)
	check(err)
}

func createVersionFile() {
	// Create a new version file with initial version of 0.0.1
	err := os.WriteFile(".version", []byte("v0.0.1"), 0666)
	check(err)
}

func createBuildutilConfigFile() {
	// Get current working directory
	currentWD, err := os.Getwd()
	cobra.CheckErr(err)
	// Search config in home directory with name ".testing" (without extension).
	viper.AddConfigPath(currentWD)
	viper.SetConfigType("yaml")
	viper.SetConfigName("buildutil")
	viper.Set("buildutil.version.path", "./")
	viper.Set("buildutil.version.file", ".version")
	// Save the updated configuration to the file
	err = viper.SafeWriteConfigAs("./buildutil.yaml")
	if err != nil {
		fmt.Printf("Error writing config file, %s", err)
	} else {
		fmt.Println("Configuration written successfully")
	}
}

func buildModule(withLDFlags bool) {
	// Check if LDFlags to be included
	checkVersionFile()
	if withLDFlags {
		// Get the version from the version file
		dat, err := os.ReadFile(viper.GetString("buildutil.version.path") + viper.GetString("buildutil.version.file"))
		if err != nil {
			log.Fatal(err)
		}
		verString := string(dat)
		// Format the build date time
		t := time.Now()
		formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		buildDateTime := formatted
		//fmt.Printf("Build Date Time: %q\n", buildDateTime)
		// get git hash git rev-parse --short HEAD
		cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
		var out strings.Builder
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		gitHashString := strings.TrimRight(out.String(), "\n")
		//fmt.Printf("Git Hash: %q\n", gitHashString)
		ldflags := "-ldflags=-s -X 'main.Version=" + verString + "' -X 'main.BuildTime=" + buildDateTime + "' -X 'main.GitHash=" + gitHashString + "'"
		fmt.Println("LDFlags: " + ldflags)
		// Set the LDFlags for the build
		executeBuild(ldflags)
	} else {
		executeBuild("")
	}

}

func executeBuild(flags string) {
	// Build the module
	cmd := exec.Command("go", "build", flags, "-o", output, ".")
	fmt.Println("Build Command to execute: ", cmd)
	// Check if there is an error running the command
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %v:\n\noutput:\n\n%s\n", err, out)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		var tools toolbox.Tools

		if tools.CheckFileExist(cfgFile) {
			// file exists
			//fmt.Printf("Buildutil config File exists: %s\n", cfgFile)
		} else {
			// file does not exist
			fmt.Printf("Buildutil config File does not exist: %s\n", cfgFile)
			fmt.Fprintln(os.Stderr, "\nConfig File not found, executing the setup to build default buildutil.yaml file!")
			createBuildutilConfigFile()
		}
		viper.SetConfigFile(cfgFile)
	} else {
		// Get current working Directory
		currentWD, err := os.Getwd()
		cobra.CheckErr(err)
		//fmt.Println("Current Working Directory: ", currentWD)
		// Search config in home directory with name ".testing" (without extension).
		viper.AddConfigPath(currentWD)
		viper.SetConfigType("yaml")
		viper.SetConfigName("buildutil")
	}

	viper.AutomaticEnv() // read in environment variables that match
	//fmt.Printf("Setup flag value: %t\n", setup)
	// If a config file is found, read it in.
	if setup == false {
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			var configFileNotFoundError viper.ConfigFileNotFoundError

			if errors.As(err, &configFileNotFoundError) {
				// Config file not found; ignore error if desired
				fmt.Fprintln(os.Stderr, "\nConfig File not found, executing the setup to build default buildutil.yaml file!")
				createBuildutilConfigFile()
			} else {
				// Config file was found but another error was produced
				panic(fmt.Errorf("fatal error config file: %w", err))
			}
		}

	}

	//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}
