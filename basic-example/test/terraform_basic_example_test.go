package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
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

// The tests in this folder are not example usage of Terratest. Instead, this is a regression test to ensure the
// formatting rules work with an actual Terraform call when using more complex structures.

func TestTerraformFormatNestedOneLevelList(t *testing.T) {
	t.Parallel()

	testList := [][]string{
		[]string{random.UniqueId()},
	}

	options := GetTerraformOptionsForFormatTests(t)
	options.Vars["example_any"] = testList

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)
	outputMap := terraform.OutputForKeys(t, options, []string{"example_any"})
	actualExampleList := outputMap["example_any"]
	AssertEqualJson(t, actualExampleList, testList)
}

func TestTerraformFormatNestedTwoLevelList(t *testing.T) {
	t.Parallel()

	testList := [][][]string{
		[][]string{[]string{random.UniqueId()}},
	}

	options := GetTerraformOptionsForFormatTests(t)
	options.Vars["example_any"] = testList

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)
	outputMap := terraform.OutputForKeys(t, options, []string{"example_any"})
	actualExampleList := outputMap["example_any"]
	AssertEqualJson(t, actualExampleList, testList)
}

func TestTerraformFormatNestedMultipleItems(t *testing.T) {
	t.Parallel()

	testList := [][]string{
		[]string{random.UniqueId(), random.UniqueId()},
		[]string{random.UniqueId(), random.UniqueId(), random.UniqueId()},
	}

	options := GetTerraformOptionsForFormatTests(t)
	options.Vars["example_any"] = testList

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)
	outputMap := terraform.OutputForKeys(t, options, []string{"example_any"})
	actualExampleList := outputMap["example_any"]
	AssertEqualJson(t, actualExampleList, testList)
}

func TestTerraformFormatNestedOneLevelMap(t *testing.T) {
	t.Parallel()

	testMap := map[string]map[string]string{
		"test": map[string]string{
			"foo": random.UniqueId(),
		},
	}

	options := GetTerraformOptionsForFormatTests(t)
	options.Vars["example_any"] = testMap

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)
	outputMap := terraform.OutputForKeys(t, options, []string{"example_any"})
	actualExampleMap := outputMap["example_any"]
	AssertEqualJson(t, actualExampleMap, testMap)
}

func TestTerraformFormatNestedTwoLevelMap(t *testing.T) {
	t.Parallel()

	testMap := map[string]map[string]map[string]string{
		"test": map[string]map[string]string{
			"foo": map[string]string{
				"bar": random.UniqueId(),
			},
		},
	}

	options := GetTerraformOptionsForFormatTests(t)
	options.Vars["example_any"] = testMap

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)
	outputMap := terraform.OutputForKeys(t, options, []string{"example_any"})
	actualExampleMap := outputMap["example_any"]
	AssertEqualJson(t, actualExampleMap, testMap)
}

func TestTerraformFormatNestedMultipleItemsMap(t *testing.T) {
	t.Parallel()

	testMap := map[string]map[string]string{
		"test": map[string]string{
			"foo": random.UniqueId(),
			"bar": random.UniqueId(),
		},
		"other": map[string]string{
			"baz": random.UniqueId(),
			"boo": random.UniqueId(),
		},
	}

	options := GetTerraformOptionsForFormatTests(t)
	options.Vars["example_any"] = testMap

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)
	outputMap := terraform.OutputForKeys(t, options, []string{"example_any"})
	actualExampleMap := outputMap["example_any"]
	AssertEqualJson(t, actualExampleMap, testMap)
}

func TestTerraformFormatNestedListMap(t *testing.T) {
	t.Parallel()

	testMap := map[string][]string{
		"test": []string{random.UniqueId(), random.UniqueId()},
	}

	options := GetTerraformOptionsForFormatTests(t)
	options.Vars["example_any"] = testMap

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)
	outputMap := terraform.OutputForKeys(t, options, []string{"example_any"})
	actualExampleMap := outputMap["example_any"]
	AssertEqualJson(t, actualExampleMap, testMap)
}

func GetTerraformOptionsForFormatTests(t *testing.T) *terraform.Options {
	exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "infra")

	// Set up terratest to retry on known failures
	maxTerraformRetries := 3
	sleepBetweenTerraformRetries := 5 * time.Second
	retryableTerraformErrors := map[string]string{
		// `terraform init` frequently fails in CI due to network issues accessing plugins. The reason is unknown, but
		// eventually these succeed after a few retries.
		".*unable to verify signature.*":             "Failed to retrieve plugin due to transient network error.",
		".*unable to verify checksum.*":              "Failed to retrieve plugin due to transient network error.",
		".*no provider exists with the given name.*": "Failed to retrieve plugin due to transient network error.",
		".*registry service is unreachable.*":        "Failed to retrieve plugin due to transient network error.",
		".*connection reset by peer.*":               "Failed to retrieve plugin due to transient network error.",
	}

	terraformOptions := &terraform.Options{
		TerraformDir:             exampleFolder,
		Vars:                     map[string]interface{}{},
		NoColor:                  true,
		RetryableTerraformErrors: retryableTerraformErrors,
		MaxRetries:               maxTerraformRetries,
		TimeBetweenRetries:       sleepBetweenTerraformRetries,
	}
	return terraformOptions
}

// The value of the output nested in the outputMap returned by OutputForKeys uses the interface{} type for nested
// structures. This can't be compared to actual types like [][]string{}, so we instead compare the json versions.
func AssertEqualJson(t *testing.T, actual interface{}, expected interface{}) {
	actualJson, err := json.Marshal(actual)
	require.NoError(t, err)
	expectedJson, err := json.Marshal(expected)
	require.NoError(t, err)
	assert.Equal(t, actualJson, expectedJson)
}
