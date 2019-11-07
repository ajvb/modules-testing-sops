package submodule

import "github.com/ajvb/modules-testing-sops/v2"

func SubmoduleFunc(msg string) {
	modules_testing_sops.AnotherExampleFunc(msg, msg)
}
