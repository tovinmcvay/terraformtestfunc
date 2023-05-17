package reusabletest

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// CheckRequiredParameters checks that the required parameters for testing Terraform code are present
func CheckRequiredParameters(t *testing.T, resourceName string, outputValue string) {
	if resourceName == "" {
		t.Errorf("Resource name is required")
	}

	if outputValue == "" {
		t.Errorf("Output name is required")
	}
}

// CheckRequiredOutputParams for terraform Options
func CheckRequiredOutputParams(t *testing.T, terraformOptions *terraform.Options, outputName string) {
	if terraformOptions == nil {
		t.Errorf("terraformOptions is required inside terrarform.Output() parameter")
	}

	if t == nil {
		t.Errorf("t *testing.t is required inside terraform.Output() parameter")
	}

	if outputName == "" {
		t.Errorf("Output name is required inside terraform.Output() parameter")
	}
}

// CheckRequiredFunctions checks that the required functions for testing Terraform code are present
func CheckRequiredFunctions(t *testing.T, terraformOptions *terraform.Options, path string) {
	var (
		requiredFunctions = []string{"InitAndApply", "Destroy"}
		missingFunctions  []string
		TerraformDir      = ""
	)

	// Check that each required function is present in the Terraform options
	for _, functionName := range requiredFunctions {
		_, found := reflect.TypeOf(terraformOptions).MethodByName(functionName)
		if !found {
			missingFunctions = append(missingFunctions, functionName)
		}
	}
	if path != TerraformDir {
		t.Errorf("The specified path is not corrext or exist")
	}
}

// Check if the testing object has all required assertions
func CheckAssertionType(t *testing.T, outputName string) {
	var (
		assertionNames = []string{
			"assert.Equal",
			"assert.Contains",
		}
	)
	for _, assertionName := range assertionNames {
		assertionFunc := reflect.ValueOf(t).Elem().FieldByName(assertionName)
		if !assertionFunc.IsValid() {
			t.Errorf("assertion function %s not found in testing object", assertionName)
		}
		if assertionFunc.Type().NumIn() == 0 {
			t.Errorf("assertion function %s in testing object must have at least one parameter", assertionName)
		}
		if assertionFunc.Type().In(0).Name() != "T" {
			t.Errorf("first parameter of assertion function %s in testing object must have type *testing.T", assertionName)
		}
	}
}

// Utility function to get the name of the calling function
func getCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
