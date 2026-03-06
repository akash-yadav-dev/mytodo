// Package grpc provides gRPC service implementations for the auth module.
//
// This file implements user management gRPC service.

package grpc

// UserServer implements the gRPC UserService interface.
// In production applications, user gRPC servers typically implement:
// - GetUser(GetUserRequest) - retrieve user by ID
// - ListUsers(ListUsersRequest) stream - paginated user listing
// - UpdateUser(UpdateUserRequest) - update user data
// - DeleteUser(DeleteUserRequest) - remove user
// - GetUserPermissions(PermissionsRequest) - get user's permissions
// - ValidateCredentials(CredentialsRequest) - internal auth validation
