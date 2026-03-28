package services

import (
	"errors"
	"wallet/services/accounts/internal/models"
)

type AccountRepository interface {
	Save(account models.Account) error
	FindByID(id string) (*models.Account, error)
	Update(account models.Account) error
}

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(r AccountRepository) *AccountService {
	return &AccountService{repo: r}
}

func (s *AccountService) CreateAccount(userID string) (*models.Account, error) {
	account := models.Account{
		ID:      userID, // simplificado
		UserID:  userID,
		Balance: 0,
	}

	err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AccountService) GetAccount(id string) (*models.Account, error) {
	return s.repo.FindByID(id)
}

func (s *AccountService) Deposit(id string, amount float64) error {
	acc, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	acc.Balance += amount
	return s.repo.Update(*acc)
}

func (s *AccountService) Withdraw(id string, amount float64) error {
	acc, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if acc.Balance < amount {
		return errors.New("insufficient funds")
	}

	acc.Balance -= amount
	return s.repo.Update(*acc)
}
