package test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/gruntwork-io/terratest/modules/terraform"
)



func theTerraformModuleIsDeployed(module string) error {
	var t = &testing.T{}

	modulePath := "../modules/" + module

	options := terraform.Options{
		TerraformDir: modulePath,
	}
	out, err := terraform.InitAndApplyE(t, &options)
	if err != nil {
		fmt.Println(out, err)
		err := fmt.Errorf("[-] Error applying module %s", module)
		return err
	}
	fmt.Println("[+] Module", module, "applied sucessfully")
	return nil
}

func iCanAccessTheInstanceViaSSH() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the terraform module "([^"]*)" is deployed$`, theTerraformModuleIsDeployed)
	s.Step(`^I can access the instance via SSH$`, iCanAccessTheInstanceViaSSH)
}
