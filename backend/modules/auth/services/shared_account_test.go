package services

import "testing"

func TestProtectedAccountUsesConfiguredSuffix(t *testing.T) {
	service := &Service{cfg: Config{ProtectedEmailSuffix: "@demo.invalid"}}
	if !service.protectedAccount("Support@DEMO.INVALID") {
		t.Fatal("expected demo account to be protected")
	}
	if service.protectedAccount("person@example.com") {
		t.Fatal("unexpected protection outside configured suffix")
	}
	if (&Service{}).protectedAccount("support@demo.invalid") {
		t.Fatal("empty configuration must disable protection")
	}
}
