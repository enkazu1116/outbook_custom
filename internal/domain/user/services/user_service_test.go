package services

import (
	"context"
	"errors"
	"testing"
)

func TestFindUserValidation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id    string
		name  string
		email string
		err   error
	}{
		"all empty": {
			err: errors.New(ErrorMessage["search_required"]),
		},
		"id only": {
			id: "user-1",
		},
		"name only": {
			name: "Alice",
		},
		"email only": {
			email: "alice@example.com",
		},
		"multiple fields": {
			id:    "user-1",
			name:  "Alice",
			email: "alice@example.com",
		},
	}

	ctx := context.Background()

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := FindUserValidation(ctx, tt.id, tt.name, tt.email)

			switch {
			case tt.err == nil && err != nil:
				t.Fatalf("unexpected error: %v", err)
			case tt.err != nil && err == nil:
				t.Fatalf("expected error %v, got nil", tt.err)
			case tt.err != nil && err != nil && err.Error() != tt.err.Error():
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}
