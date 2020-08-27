package auth

import (
	"context"

	"github.com/covid19cz/erouska-backend/internal/firebase"
)

// Auther is an auth abstraction layer interface
type Auther interface {
	CustomToken(ctx context.Context, uid string) (string, error)
}

// Client to interact with auth API
type Client struct{}

// CustomToken creates a signed custom authentication token with the specified user ID. The resulting JWT can be used in a Firebase client SDK to trigger an authentication flow. See https://firebase.google.com/docs/auth/admin/create-custom-tokens#sign_in_using_custom_tokens_on_clients for more details on how to use custom tokens for client authentication.
func (c *Client) CustomToken(ctx context.Context, uid string) (string, error) {
	client := firebase.FirebaseAuth
	return client.CustomToken(ctx, uid)
}

// MockClient mocks auth client functionaly for unit tests
type MockClient struct{}

// CustomToken creates a signed custom authentication token with the specified user ID.
func (c *MockClient) CustomToken(uid string) (string, error) {
	return "abc", nil
}
