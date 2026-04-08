// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flex

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// TerraformError represents a structured error for Terraform operations
type TerraformError struct {
	err          error
	message      string
	resourceName string
	operation    string
}

// TerraformErrorf creates a new TerraformError with formatted message
func TerraformErrorf(err error, message string, resourceName string, operation string) *TerraformError {
	return &TerraformError{
		err:          err,
		message:      message,
		resourceName: resourceName,
		operation:    operation,
	}
}

// GetDebugMessage returns a formatted debug message
func (te *TerraformError) GetDebugMessage() string {
	return fmt.Sprintf("Resource: %s, Operation: %s, Error: %s", te.resourceName, te.operation, te.message)
}

// GetDiag returns a diag.Diagnostics with the error information
func (te *TerraformError) GetDiag() diag.Diagnostics {
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("%s %s failed", te.resourceName, te.operation),
			Detail:   te.message,
		},
	}
}

// Made with Bob
