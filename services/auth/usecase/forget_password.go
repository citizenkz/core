package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/citizenkz/core/services/auth/entity"
	"github.com/google/uuid"
)

func (u *usecase) ForgetPassword(ctx context.Context, req *entity.ForgetPasswordRequest) (*entity.ForgetPasswordResponse, error) {
	// Check if user exists
	_, err := u.storage.GetUserByEmail(ctx, req.Email)
	if err != nil {
		u.log.Error("failed to storage.GetUserByEmail", slog.String("error", err.Error()))
		return nil, fmt.Errorf("user with this email not found")
	}

	// Generate 6-digit OTP
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))

	// Create attempt record
	attemptID, err := u.storage.CreateAttempt(ctx, req.Email, otp)
	if err != nil {
		u.log.Error("failed to storage.CreateAttempt", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create reset attempt: %w", err)
	}

	// Send OTP via email
	if err := u.emailService.SendOTP(req.Email, otp); err != nil {
		u.log.Error("failed to send OTP email", slog.String("error", err.Error()))
		// Continue even if email fails - user can request another OTP
	}

	return &entity.ForgetPasswordResponse{
		AttemptID: int(attemptID.ID()), // Convert UUID to int for response
		RetryTime: 600,                 // 10 minutes in seconds
	}, nil
}

func (u *usecase) ForgetPasswordConfirm(ctx context.Context, req *entity.ForgetPasswordConfirmRequest) (*entity.ForgetPasswordConfirmResponse, error) {
	// Convert int ID back to UUID
	attemptUUID, err := uuid.Parse(fmt.Sprintf("%d", req.AttemptID))
	if err != nil {
		return nil, fmt.Errorf("invalid attempt ID")
	}

	// Get attempt
	attempt, err := u.storage.GetAttempt(ctx, attemptUUID)
	if err != nil {
		u.log.Error("failed to storage.GetAttempt", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid or expired reset attempt")
	}

	// Verify OTP
	if attempt.Otp != req.OtpCode {
		return nil, fmt.Errorf("invalid OTP code")
	}

	// Get user by email
	user, err := u.storage.GetUserByEmail(ctx, attempt.Email)
	if err != nil {
		u.log.Error("failed to storage.GetUserByEmail", slog.String("error", err.Error()))
		return nil, fmt.Errorf("user not found")
	}

	// Delete the attempt
	if err := u.storage.DeleteAttempt(ctx, attemptUUID); err != nil {
		u.log.Error("failed to delete attempt", slog.String("error", err.Error()))
	}

	return &entity.ForgetPasswordConfirmResponse{
		Profile: *user,
	}, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
