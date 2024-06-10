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
package main

import "github.com/hostedbrains/buildutil/cmd"

var Version string = "v1.0.2"
var BuildTime string = "2024-05-30T19:43:41Z"
var GitHash string = "5a96c6a"

func main() {
	cmd.Execute(Version, BuildTime, GitHash)
}
