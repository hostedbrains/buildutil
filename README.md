# Buildutil
Build utility for GO projects

## Overview
This is a utility for building GO projects.  
This utility CLI helps you to build your GO project in a quick and easy way.  
This utility CLI helps you to increment your version information.  
For version numbering Buildutil uses *[Semantic versioning](https://semver.org/)*  
>The buildutil should always be executed from the root of your project folder which yo want to control the version and build your binaries.    


Buildutil allows you to tkae the following actions:  
- Increment your version number:
  - Increment the Major part of the version
  - Increment the Minor part of the version
  - Increment the Patch part of the version
- Build the binary for your project
  - Optionally with LD Flags to set version etc in the binary meta data

## Basic usage
```console
Usage:
  buildutil [flags]

Flags:
  -b, --build            Build the module
      --config string    config file (default is $HOME/.testing.yaml)
  -h, --help             help for buildutil
      --incrementMajor   Increment the major version, default false
      --incrementMinor   Increment the minor version, default false
      --incrementPatch   Increment the patch version, default false
      --initVersion      Create new version file with initial version of 0.0.1
  -o, --output string    Build output e.g. '-o bin/moduleName' (required if build is set)
      --setup            Create new buildutil.yaml file with defaults.
  -v, --version          Print version information
  -f, --withLDFlags      Include LDFlags with the build.
  ```
## Usage detail
> The buildutil should always be executed from the root of your project folder which yo want to control the version and build your binaries.  

### Files required for and by Buildutil
There are two files that are used and required explicitly for Buildutil, namely buildutil.yaml and .version
Firstly buildutil.yaml, is used to configure buildutil and has the following content:
```yaml
buildutil:
  version:
    path: "./"
    file: ".version"
```
`buildutil.version.path` specifies the path to your version file, default is "./"  
`buildutil.version.file` specifies the version file name, default is ".version"  
> #### Note  
> If the buildutil.yaml file does not exist in the current directory then the buildutil.yaml file will be created with the defaults and stored in your current directory  

Secondaly ".version", is user to store your current version number for your project and has the following content:  
```text
v0.0.1
```
> #### Note
> If the ".version" file does not exist in the current directory then buildutil will issue an error message and you can create the ".version" file with the following command:  
> ```console
> buildutil --initVersion
> ```
> This will create the ".version" file with a version of "v0,0,1"

