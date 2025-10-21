package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/auth/entity"
	"github.com/citizenkz/core/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) UpdatePassword(ctx context.Context, req *entity.UpdatePasswordRequest) (*entity.UpdatePasswordResponse, error) {
	// Validate passwords match
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("passwords do not match")
	}

	// Parse user ID from token
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid token")
	}

	// Get user
	user, err := u.storage.GetUserByID(ctx, userID)
	if err != nil {
		u.log.Error("failed to storage.GetUserByID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("user not found")
	}

	// Verify old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		u.log.Error("failed to bcrypt.CompareHashAndPassword", slog.String("error", err.Error()))
		return nil, fmt.Errorf("incorrect old password")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("failed to bcrypt.GenerateFromPassword", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Update password
	updatedUser, err := u.storage.UpdateUserPassword(ctx, userID, string(hashedPassword))
	if err != nil {
		u.log.Error("failed to storage.UpdateUserPassword", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to update password: %w", err)
	}

	// Send confirmation email
	if err := u.emailService.SendPasswordChanged(updatedUser.Email); err != nil {
		u.log.Error("failed to send password changed email", slog.String("error", err.Error()))
		// Continue even if email fails
	}

	return &entity.UpdatePasswordResponse{
		Profile: *updatedUser,
	}, nil
}
