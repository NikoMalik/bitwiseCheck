package bitwisecheck

import "fmt"

func m() {
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
