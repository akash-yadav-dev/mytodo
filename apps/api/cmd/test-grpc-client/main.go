package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"mytodo/apps/api/internal/auth/interfaces/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func truncateToken(token string) string {
	if len(token) <= 20 {
		return token
	}
	return token[:20] + "..."
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the gRPC server
	conn, err := grpc.DialContext(
		ctx,
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	fmt.Println("=== Testing gRPC Auth Service ===")

	// Test 1: Register a new user
	fmt.Println("1. Testing Register...")

	registerReq := &pb.RegisterRequest{
		Email:    fmt.Sprintf("test%d@example.com", time.Now().Unix()),
		Password: "password123",
		Name:     "Test User",
	}

	registerResp, err := client.Register(ctx, registerReq)
	if err != nil {
		fmt.Printf("   ❌ Register failed: %v\n\n", err)
	} else {
		fmt.Println("   ✅ Register successful!")
		fmt.Printf("   User ID: %s\n", registerResp.User.Id)
		fmt.Printf("   Email: %s\n", registerResp.User.Email)
		fmt.Printf("   Name: %s\n", registerResp.User.Name)
		fmt.Printf("   Access Token: %s\n", truncateToken(registerResp.AccessToken))
		fmt.Printf("   Refresh Token: %s\n\n", truncateToken(registerResp.RefreshToken))
	}

	// Test 2: Login
	fmt.Println("2. Testing Login...")

	loginReq := &pb.LoginRequest{
		Email:     registerReq.Email,
		Password:  registerReq.Password,
		UserAgent: "test-client",
		IpAddress: "127.0.0.1",
	}

	loginResp, err := client.Login(ctx, loginReq)
	if err != nil {
		fmt.Printf("   ❌ Login failed: %v\n\n", err)
	} else {
		fmt.Println("   ✅ Login successful!")
		fmt.Printf("   User ID: %s\n", loginResp.User.Id)
		fmt.Printf("   Email: %s\n", loginResp.User.Email)
		fmt.Printf("   Access Token: %s\n", truncateToken(loginResp.AccessToken))
		fmt.Printf("   Refresh Token: %s\n", truncateToken(loginResp.RefreshToken))
		fmt.Printf("   Token Type: %s\n", loginResp.TokenType)
		fmt.Printf("   Expires In: %d seconds\n\n", loginResp.ExpiresIn)

		// Test 3: Refresh token
		fmt.Println("3. Testing RefreshToken...")

		refreshReq := &pb.RefreshTokenRequest{
			RefreshToken: loginResp.RefreshToken,
		}

		refreshResp, err := client.RefreshToken(ctx, refreshReq)
		if err != nil {
			fmt.Printf("   ❌ RefreshToken failed: %v\n\n", err)
		} else {
			fmt.Println("   ✅ RefreshToken successful!")
			fmt.Printf("   New Access Token: %s\n", truncateToken(refreshResp.AccessToken))
			fmt.Printf("   New Refresh Token: %s\n\n", truncateToken(refreshResp.RefreshToken))

			// Test 4: Logout
			fmt.Println("4. Testing Logout...")

			logoutReq := &pb.LogoutRequest{
				RefreshToken: refreshResp.RefreshToken,
			}

			logoutResp, err := client.Logout(ctx, logoutReq)
			if err != nil {
				fmt.Printf("   ❌ Logout failed: %v\n\n", err)
			} else {
				fmt.Println("   ✅ Logout successful!")
				fmt.Printf("   Message: %s\n\n", logoutResp.Message)
			}
		}
	}

	// Test 5: ValidateToken
	fmt.Println("5. Testing ValidateToken...")

	validateReq := &pb.ValidateTokenRequest{
		AccessToken: "test-token",
	}

	validateResp, err := client.ValidateToken(ctx, validateReq)
	if err != nil {
		fmt.Printf("   ❌ ValidateToken failed: %v\n\n", err)
	} else {
		fmt.Println("   ℹ️  ValidateToken response:")
		fmt.Printf("   Valid: %v\n", validateResp.Valid)
		fmt.Printf("   Error: %s\n\n", validateResp.ErrorMessage)
	}

	fmt.Println("=== All tests completed ===")
}
