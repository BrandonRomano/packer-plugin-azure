// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arm

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/packer-plugin-azure/builder/azure/common/constants"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestStepGetCertificateShouldFailIfGetFails(t *testing.T) {
	var testSubject = &StepGetCertificate{
		get: func(context.Context, string, string, string, string) (string, error) {
			return "", fmt.Errorf("!! Unit Test FAIL !!")
		},
		say:   func(message string) {},
		error: func(e error) {},
		pause: func() {},
	}

	stateBag := createTestStateBagStepGetCertificate()

	var result = testSubject.Run(context.Background(), stateBag)
	if result != multistep.ActionHalt {
		t.Fatalf("Expected the step to return 'ActionHalt', but got '%d'.", result)
	}

	if _, ok := stateBag.GetOk(constants.Error); ok == false {
		t.Fatalf("Expected the step to set stateBag['%s'], but it was not.", constants.Error)
	}
}

func TestStepGetCertificateShouldPassIfGetPasses(t *testing.T) {
	var testSubject = &StepGetCertificate{
		get:   func(context.Context, string, string, string, string) (string, error) { return "", nil },
		say:   func(message string) {},
		error: func(e error) {},
		pause: func() {},
	}

	stateBag := createTestStateBagStepGetCertificate()

	var result = testSubject.Run(context.Background(), stateBag)
	if result != multistep.ActionContinue {
		t.Fatalf("Expected the step to return 'ActionContinue', but got '%d'.", result)
	}

	if _, ok := stateBag.GetOk(constants.Error); ok == true {
		t.Fatalf("Expected the step to not set stateBag['%s'], but it was.", constants.Error)
	}
}

func TestStepGetCertificateShouldTakeStepArgumentsFromStateBag(t *testing.T) {
	var actualKeyVaultName string
	var actualSecretName string

	var testSubject = &StepGetCertificate{
		get: func(ctx context.Context, subscriptionId string, resourceGroupName string, keyVaultName string, secretName string) (string, error) {
			actualKeyVaultName = keyVaultName
			actualSecretName = secretName

			return "http://key.vault/1", nil
		},
		say:   func(message string) {},
		error: func(e error) {},
		pause: func() {},
	}

	stateBag := createTestStateBagStepGetCertificate()
	var result = testSubject.Run(context.Background(), stateBag)

	if result != multistep.ActionContinue {
		t.Fatalf("Expected the step to return 'ActionContinue', but got '%d'.", result)
	}

	var expectedKeyVaultName = stateBag.Get(constants.ArmKeyVaultName).(string)

	if actualKeyVaultName != expectedKeyVaultName {
		t.Fatal("Expected StepGetCertificate to source 'constants.ArmKeyVaultName' from the state bag, but it did not.")
	}
	if actualSecretName != DefaultSecretName {
		t.Fatal("Expected StepGetCertificate to use default value for secret, but it did not.")
	}

	expectedCertificateUrl, ok := stateBag.GetOk(constants.ArmCertificateUrl)
	if !ok {
		t.Fatalf("Expected the state bag to have a value for '%s', but it did not.", constants.ArmCertificateUrl)
	}

	if expectedCertificateUrl != "http://key.vault/1" {
		t.Fatalf("Expected the value of stateBag[%s] to be 'http://key.vault/1', but got '%s'.", constants.ArmCertificateUrl, expectedCertificateUrl)
	}
}

func createTestStateBagStepGetCertificate() multistep.StateBag {
	stateBag := new(multistep.BasicStateBag)
	stateBag.Put(constants.ArmKeyVaultName, "Unit Test: KeyVaultName")
	stateBag.Put(constants.ArmSubscription, "testSubscription")
	stateBag.Put(constants.ArmResourceGroupName, "testResourceGroupName")
	return stateBag
}
