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

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("=== Testing gRPC Auth Service ===\n")

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
		fmt.Printf("   ✅ Register successful!\n")
		fmt.Printf("   User ID: %s\n", registerResp.User.Id)
		fmt.Printf("   Email: %s\n", registerResp.User.Email)
		fmt.Printf("   Name: %s\n", registerResp.User.Name)
		fmt.Printf("   Access Token: %s...\n", registerResp.AccessToken[:20])
		fmt.Printf("   Refresh Token: %s...\n\n", registerResp.RefreshToken[:20])
	}

	// Test 2: Login with the registered user
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
		fmt.Printf("   ✅ Login successful!\n")
		fmt.Printf("   User ID: %s\n", loginResp.User.Id)
		fmt.Printf("   Email: %s\n", loginResp.User.Email)
		fmt.Printf("   Access Token: %s...\n", loginResp.AccessToken[:20])
		fmt.Printf("   Refresh Token: %s...\n", loginResp.RefreshToken[:20])
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
			fmt.Printf("   ✅ RefreshToken successful!\n")
			fmt.Printf("   New Access Token: %s...\n", refreshResp.AccessToken[:20])
			fmt.Printf("   New Refresh Token: %s...\n\n", refreshResp.RefreshToken[:20])

			// Test 4: Logout
			fmt.Println("4. Testing Logout...")
			logoutReq := &pb.LogoutRequest{
				RefreshToken: refreshResp.RefreshToken,
			}

			logoutResp, err := client.Logout(ctx, logoutReq)
			if err != nil {
				fmt.Printf("   ❌ Logout failed: %v\n\n", err)
			} else {
				fmt.Printf("   ✅ Logout successful!\n")
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
		fmt.Printf("   ℹ️  ValidateToken response:\n")
		fmt.Printf("   Valid: %v\n", validateResp.Valid)
		fmt.Printf("   Error: %s\n\n", validateResp.ErrorMessage)
	}

	fmt.Println("=== All tests completed ===")
}
