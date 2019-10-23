package test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

var t = &testing.T{}
var m string

func theTerraformModuleIsDeployed(module string) error {
	m = module
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

func cleanup() error {
	modulePath := "../modules/" + m
	options := terraform.Options{
		TerraformDir: modulePath,
	}
	out, err := terraform.DestroyE(t, &options)
	if err != nil {
		fmt.Println(out, err)
		err := fmt.Errorf("[-] Error applying module %s", m)
		return err
	}
	fmt.Println("[+] Module", m, "destroyed sucessfully")
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the terraform module "([^"]*)" is deployed$`, theTerraformModuleIsDeployed)
	s.Step(`^I can access the instance via SSH$`, iCanAccessTheInstanceViaSSH)

	s.AfterScenario(func(interface{}, error) {
		cleanup()
	})
}
