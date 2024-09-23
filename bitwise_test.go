package bitwisecheck

import (
	"fmt"
	"testing"
)

func TestFlags(t *testing.T) {
	var flags Flags

	// Test initial state
	if flags.String() != "00000000" {
		t.Errorf("Expected flags to be 00000000, got %s", flags.String())
	}

	// Test setting flags
	flags.Set(0)
	if flags.String() != "00000001" {
		t.Errorf("Expected flags to be 00000001 after setting flag 0, got %s", flags.String())
	}
	t.Log(flags.String())
	flags.Set(1)
	if flags.String() != "00000011" {
		t.Errorf("Expected flags to be 00000011 after setting flag 1, got %s", flags.String())
	}

	// Test clearing flags
	flags.Clear(0)
	if flags.String() != "00000010" {
		t.Errorf("Expected flags to be 00000010 after clearing flag 0, got %s", flags.String())
	}
	t.Log(flags.String())
	flags.Clear(1)
	if flags.String() != "00000000" {
		t.Errorf("Expected flags to be 00000000 after clearing flag 1, got %s", flags.String())
	}

}

func TestFlagsIsSet(t *testing.T) {
	var flags Flags
	flags.Set(0)
	if !flags.IsSet(0) {
		t.Errorf("Expected flag 0 to be set, got %s", flags.String())
	}
	if flags.IsSet(1) {
		t.Errorf("Expected flag 1 to not be set, got %s", flags.String())
	}
}

func TestFlagsString(t *testing.T) {
	var flags Flags
	flags.Set(0)
	if flags.String() != "00000001" {
		t.Errorf("Expected flags to be 00000001, got %s", flags.String())
	}
}

func TestExample(t *testing.T) {
	var userPermissions Flags

	// Setting permissions
	userPermissions.Set(0) // Set permission 0 (e.g., Read)
	userPermissions.Set(2) // Set permission 2 (e.g., Write)

	// Checking permissions
	if userPermissions.IsSet(0) {
		fmt.Println("User has Read permission")
	}
	if userPermissions.IsSet(1) {
		fmt.Println("User has Edit permission")
	} else {
		fmt.Println("User does not have Edit permission")
	}
	if userPermissions.IsSet(2) {
		fmt.Println("User has Write permission")
	}

	// Clearing a permission
	userPermissions.Clear(0) // Remove Read permission

	// Displaying current permissions in binary
	fmt.Println("Current Permissions:", userPermissions.String()) // Output: Current Permissions: 00000100
}
