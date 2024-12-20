package cmd

import "testing"

func TestExecute(t *testing.T) {
	type args struct {
		version   string
		buildTime string
		gitHash   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute(tt.args.version, tt.args.buildTime, tt.args.gitHash)
		})
	}
}

func Test_buildModule(t *testing.T) {
	type args struct {
		withLDFlags bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildModule(tt.args.withLDFlags)
		})
	}
}

func Test_check(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.e)
		})
	}
}

func Test_createVersionFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createVersionFile()
		})
	}
}

func Test_excuteBuild(t *testing.T) {
	type args struct {
		flags string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executeBuild(tt.args.flags)
		})
	}
}

func Test_incrementVersionFunc(t *testing.T) {
	type args struct {
		semver string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			incrementVersionFunc(tt.args.semver)
		})
	}
}

func Test_trimFirstRune(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimFirstRune(tt.args.s); got != tt.want {
				t.Errorf("trimFirstRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_incrementVersionFunc1(t *testing.T) {
	type args struct {
		semver string
	}
	tests := []struct {
		name string
		args args
	}{
		//{"Increment patch version", args{"patch"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			incrementVersionFunc(tt.args.semver)
		})
	}
}

func TestExecute1(t *testing.T) {
	type args struct {
		version   string
		buildTime string
		gitHash   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute(tt.args.version, tt.args.buildTime, tt.args.gitHash)
		})
	}
}

func Test_buildModule1(t *testing.T) {
	type args struct {
		withLDFlags bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildModule(tt.args.withLDFlags)
		})
	}
}

func Test_check1(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.e)
		})
	}
}

func Test_createBuildutilConfigFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createBuildutilConfigFile()
		})
	}
}

func Test_createVersionFile1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createVersionFile()
		})
	}
}

func Test_executeBuild(t *testing.T) {
	type args struct {
		flags string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executeBuild(tt.args.flags)
		})
	}
}

func Test_incrementVersionFunc2(t *testing.T) {
	type args struct {
		semver string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			incrementVersionFunc(tt.args.semver)
		})
	}
}

func Test_initConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initConfig()
		})
	}
}

func Test_trimFirstRune1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimFirstRune(tt.args.s); got != tt.want {
				t.Errorf("trimFirstRune() = %v, want %v", got, tt.want)
			}
		})
	}
}
