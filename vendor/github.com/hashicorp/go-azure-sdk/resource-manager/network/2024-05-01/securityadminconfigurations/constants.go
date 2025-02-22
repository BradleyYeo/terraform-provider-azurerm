package securityadminconfigurations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddressSpaceAggregationOption string

const (
	AddressSpaceAggregationOptionManual AddressSpaceAggregationOption = "Manual"
	AddressSpaceAggregationOptionNone   AddressSpaceAggregationOption = "None"
)

func PossibleValuesForAddressSpaceAggregationOption() []string {
	return []string{
		string(AddressSpaceAggregationOptionManual),
		string(AddressSpaceAggregationOptionNone),
	}
}

func (s *AddressSpaceAggregationOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddressSpaceAggregationOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddressSpaceAggregationOption(input string) (*AddressSpaceAggregationOption, error) {
	vals := map[string]AddressSpaceAggregationOption{
		"manual": AddressSpaceAggregationOptionManual,
		"none":   AddressSpaceAggregationOptionNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddressSpaceAggregationOption(input)
	return &out, nil
}

type NetworkIntentPolicyBasedService string

const (
	NetworkIntentPolicyBasedServiceAll            NetworkIntentPolicyBasedService = "All"
	NetworkIntentPolicyBasedServiceAllowRulesOnly NetworkIntentPolicyBasedService = "AllowRulesOnly"
	NetworkIntentPolicyBasedServiceNone           NetworkIntentPolicyBasedService = "None"
)

func PossibleValuesForNetworkIntentPolicyBasedService() []string {
	return []string{
		string(NetworkIntentPolicyBasedServiceAll),
		string(NetworkIntentPolicyBasedServiceAllowRulesOnly),
		string(NetworkIntentPolicyBasedServiceNone),
	}
}

func (s *NetworkIntentPolicyBasedService) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkIntentPolicyBasedService(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkIntentPolicyBasedService(input string) (*NetworkIntentPolicyBasedService, error) {
	vals := map[string]NetworkIntentPolicyBasedService{
		"all":            NetworkIntentPolicyBasedServiceAll,
		"allowrulesonly": NetworkIntentPolicyBasedServiceAllowRulesOnly,
		"none":           NetworkIntentPolicyBasedServiceNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkIntentPolicyBasedService(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
