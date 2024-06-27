# Buildutil
Build utility for GO projects

## Overview
This is a utility for building GO projects.  
This utility CLI helps you to build your GO project in a quick and easy way.  
This utility CLI helps you to increment your version information.  
For version numbering Buildutil uses *[Semantic versioning](https://semver.org/)*  
>The buildutil should always be executed from the root of your project folder which yo want to control the version and build your binaries.    


Buildutil allows you to take the following actions:  
- Increment your version number:
  - Increment the Major part of the version
  - Increment the Minor part of the version
  - Increment the Patch part of the version
- Build the binary for your project
  - Optionally with LD Flags to set version etc. in the binary metadata

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

Secondly ".version", is user to store your current version number for your project and has the following content:  
```text
v0.0.1
```
> #### Note
> If the ".version" file does not exist in the current directory then buildutil will issue an error message, and you can create the ".version" file with the following command:  
> ```console
> buildutil --initVersion
> ```
> This will create the ".version" file with a version of "v0,0,1"

### --help | -h
Use the --help or -h flag to display the buildutil help.  
>#### Examples
> ```console
> buildutil --help
> buildutil -h
>```
 
### --build | -b
Use the --build or -b flag to build your application or module.  
When you use the --build or -b flag the --output or -o flag is also required to provide the output.  
>#### Examples
> ```console
> buildutil --build --output bin/moduleName
> buildutil -b -o bin/moduleName
>```

### --output | -o
Use the --output or -o flag to specify the out destination of your binary.  
When you use the --build or -b flag the --output or -o flag is also required to provide the output.  
>#### Examples
> ```console
> buildutil --build --output bin/moduleName
> buildutil -b -o bin/moduleName
>```

### --withLDFlags | -f
Use the --withLDFlags or -f flag to build your application or module with LDFlags set.  
When you use the --withLDFlags or -f flag build is done with the following LDFlags:  
```console
-ldflags=-s -X 'main.Version=v0.2.0' -X 'main.BuildTime=2024-06-27T14:23:41' -X 'main.GitHash=701a18f'
main.Version sets the version number within your application or module code to the version in your .version file.
main.BuildTime sets the BuildTime variable in your application or module to the current date and time when executing a build.
main.GitHash sets the GitHash variable in your application or module to the short version of your commit has as obtained by executing:
  git rev-parse --short HEAD
```
>#### Examples
> ```console
> buildutil --build --output bin/moduleName --withLDFlags
> buildutil -b -o bin/moduleName -f
>```

### --config
The --config flag allows you to specify the configuration file for the buildutil, if not specified it will default to ./buildutil.yaml
>#### Example
> ```console
> buildutil --config ./MyModuleConfig.yaml
>```

### --incrementMajor
The --incrementMajor flag causes buildutil to increment the major part of the version.  
So if your current version in the .version file is v0.1.0, and you do a buildutil --incrementMajor  the resulting version will be v1.1.0  
For version numbering Buildutil uses *[Semantic versioning](https://semver.org/)*
>#### Example
> ```console
> buildutil --incrementMajor
>```

### --incrementMinor
The --incrementMinor flag causes buildutil to increment the minor part of the version.  
So if your current version in the .version file is v0.1.0, and you do a buildutil --incrementMinor  the resulting version will be v0.2.0  
For version numbering Buildutil uses *[Semantic versioning](https://semver.org/)*
>#### Example
> ```console
> buildutil --incrementMinor
>```

### --incrementPatch
The --incrementPatch flag causes buildutil to increment the patch part of the version.  
So if your current version in the .version file is v0.1.0, and you do a buildutil --incrementPatch  the resulting version will be v0.1.1  
For version numbering Buildutil uses *[Semantic versioning](https://semver.org/)*z
>#### Example
> ```console
> buildutil --incrementPatch
>```

### --initVersion
The --initVersion flag causes buildutil to create a .version file if it does not exist and sets the inital version to v0.0.1.  
For version numbering Buildutil uses *[Semantic versioning](https://semver.org/)*z
>#### Example
> ```console
> buildutil --initVersion
>```

### --setup
The --setup flag causes buildutil to create new buildutil.yaml file with defaults.
>#### Example
> ```console
> buildutil --setup
>```
> #### buildutil.yaml defaults
> ```yaml
> buildutil:
>    version:
>        file: .version
>        path: ./
>```