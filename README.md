# Buildutil
Build utility for GO projects

## Overview
This is a utility for building GO projects.
This utility CLI helps you to build your GO project in a quick and easy way.
This utility CLI helps you to increment your version information.

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
