package usecaseimpls_test

import (
	"testing"

	"github.com/eduardo-gualberto/go.git/internal/application/gateways/usecaseimpls"
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/userrepo"
)

func TestCreateUser(t *testing.T) {
	testUserRepo := userrepo.NewMockedUserRepo()
	uc := usecaseimpls.NewCreateUserImpl(testUserRepo)

	err := uc.Execute("teste user", "0011233334444")
	if err != nil {
		t.Errorf("expected error to be nil, found: %s\n", err.Error())
	}
}
