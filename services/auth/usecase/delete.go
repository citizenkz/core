package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/auth/entity"
	"github.com/citizenkz/core/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Delete(ctx context.Context, req *entity.DeleteRequest) (*entity.DeleteResponse, error) {
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

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		u.log.Error("failed to bcrypt.CompareHashAndPassword", slog.String("error", err.Error()))
		return nil, fmt.Errorf("incorrect password")
	}

	// Delete user
	if err := u.storage.DeleteUser(ctx, userID); err != nil {
		u.log.Error("failed to storage.DeleteUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}

	// Send confirmation email
	if err := u.emailService.SendAccountDeleted(user.Email); err != nil {
		u.log.Error("failed to send account deleted email", slog.String("error", err.Error()))
		// Continue even if email fails
	}

	return &entity.DeleteResponse{
		Success: true,
	}, nil
}
