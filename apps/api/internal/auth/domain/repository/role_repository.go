// Package repository defines data access interfaces for the auth domain.
//
// This file defines the contract for role and RBAC data persistence.

package repository

// RoleRepository defines data access methods for role entities.
// In production applications, role repositories typically provide:
// - FindByID(id) - retrieve role by ID
// - FindByName(name) - lookup role by name
// - Create(role) - create new role
// - Update(role) - update existing role
// - Delete(id) - remove role (with constraint checks)
// - List() - retrieve all roles
// - FindByUserID(userID) - get all roles assigned to a user
// - AssignToUser(userID, roleID) - assign role to user
// - RemoveFromUser(userID, roleID) - revoke role from user
// - AddPermission(roleID, permissionID) - add permission to role
// - RemovePermission(roleID, permissionID) - remove permission from role
// - FindSystemRoles() - get built-in system roles

// type RoleRepository interface {
// 	FindByID(id string) (*Role, error)
// 	FindByName(name string) (*Role, error)
// 	Create(role *Role) error
// 	Update(role *Role) error
// 	Delete(id string) error
// 	List() ([]*Role, error)
// 	FindByUserID(userID string) ([]*Role, error)
// 	AssignToUser(userID, roleID string) error
// 	RemoveFromUser(userID, roleID string) error
// 	AddPermission(roleID, permissionID string) error
// 	RemovePermission(roleID, permissionID string) error
// 	FindSystemRoles() ([]*Role, error)
// }
