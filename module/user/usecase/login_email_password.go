package usecase

import (
	"bankingSystem/module/user/domain"
	"context"
)

func (uc *useCase) LoginEmailPassword(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error) {
	// 1. Find user by email
	user, err := uc.repo.FindByEmail(ctx, dto.Email)

	if err != nil {
		return nil, err
	}

	// 2. Hash & compare password
	if ok := uc.hasher.CompareHashPassword(user.Password(), user.Salt(), dto.Password); !ok {
		return nil, domain.ErrInvalidEmailPassword
	}

	// 3. Gen JWT

	if err != nil {
		return nil, err
	}

	return nil, nil
}
