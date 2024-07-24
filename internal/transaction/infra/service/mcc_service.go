package service

import "github.com/guths/caju-transaction-approver/internal/transaction/domain"

type MccService struct {
	mccRepo domain.MccRepository
}
