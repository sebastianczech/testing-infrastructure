package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in infra using Terratest.
func TestTerraformBasicExample(t *testing.T) {
	t.Parallel()

	// -------------------------- GIVEN --------------------------

	expectedText := "test"
	expectedList := []string{expectedText}
	expectedMap := map[string]string{"expected": expectedText}
	expectedVariableWithValueFromLocals := "forum"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		// The path to where our Terraform code is located
		TerraformDir: "../infra",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"example": expectedText,

			// We also can see how lists and maps translate between terratest and terraform.
			"example_list": expectedList,
			"example_map":  expectedMap,
		},

		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	})

	// Clean up resources with "terraform destroy". Using "defer" runs the command at the end of the test, whether the test succeeds or fails.
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// -------------------------- WHEN --------------------------

	// Run "terraform init" and "terraform apply".
	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualTextExample2 := terraform.Output(t, terraformOptions, "example2")
	actualExampleList := terraform.OutputList(t, terraformOptions, "example_list")
	actualExampleMap := terraform.OutputMap(t, terraformOptions, "example_map")
	actualVariableWithValueFromLocals := terraform.Output(t, terraformOptions, "service_name")

	// -------------------------- THEN --------------------------

	// Check the output against expected values.
	// Verify we're getting back the outputs we expect
	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, expectedText, actualTextExample2)
	assert.Equal(t, expectedList, actualExampleList)
	assert.Equal(t, expectedMap, actualExampleMap)
	assert.Equal(t, expectedVariableWithValueFromLocals, actualVariableWithValueFromLocals)
}
