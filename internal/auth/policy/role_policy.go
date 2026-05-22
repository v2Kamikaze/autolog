package policy

import (
	"slices"

	"github.com/v2code/autolog/internal/auth"
	"github.com/v2code/autolog/internal/auth/manager"
)

type RequireRole struct {
	roles []string
}

func RequireRolePolicy(roles ...string) auth.Policy[*manager.SessionPrincipal] {
	return &RequireRole{roles: roles}
}

func (p *RequireRole) Check(principal auth.Principal[*manager.SessionPrincipal]) error {
	for _, role := range p.roles {
		if !slices.Contains(principal.Principal().User.Roles, role) {
			return auth.ErrForbidden
		}
	}

	return nil
}
