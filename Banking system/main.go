package main

type SavingsAccount struct {
	balance int
}

type CheckingAccount struct {
	balance int
}

type InvestmentAccount struct {
	balance int
}

func (s *SavingsAccount) MonthlyInterest() int{
	return (s.balance * 5) / 100 / 12
}

func (s *SavingsAccount) Deposit(amount int) string{
	if amount > 0{
		s.balance += amount
		return  "Success"
	}else {
		return "Amount cannot be negative"
	}
}

func (s *SavingsAccount) Withdraw(amount int)string {
	if s.balance <= 0{
		return "Account balance is not enough"
	}
	if amount <= 0{
		return "Amount cannot be negative"
	}
	if s.balance < amount{
		return "Account balance is not enough" 
	}else{
		s.balance -= amount
		return "Success"
	}
}

func (s *SavingsAccount) CheckBalance() int{
	return s.balance
}

func (s *SavingsAccount) Transfer(receiver Account, amount int) string{
	  if s.balance <= 0 {
		return "Account balance is not enough"
	  }
	  if amount <= 0{
		return "Amount cannot be negative"
	  }
	  if s.balance < amount {
		return "Account balance is not enough"
	  }else {
		switch receiver.(type){
		case *SavingsAccount, *CheckingAccount, *InvestmentAccount:
			s.balance -= amount
			receiver.Deposit(amount)
			return "Success"
		default:
			 return "Invalid receiver account"
		}
	  }
    }
func (s *CheckingAccount) MonthlyInterest() int{
	return (s.balance * 1) / 100 / 12
}

func (s *CheckingAccount) Deposit(amount int) string{
	if amount > 0{
		s.balance += amount
		return  "Success"
	}else {
		return "Amount cannot be negative"
	}
}

func (s *CheckingAccount) Withdraw(amount int)string {
	if s.balance <= 0{
		return "Account balance is not enough"
	}
	if amount <= 0{
		return "Amount cannot be negative"
	}
	if s.balance < amount{
		return "Account balance is not enough" 
	}else{
		s.balance -= amount
		return "Success"
	}
}

func (s *CheckingAccount) CheckBalance() int{
	return s.balance
}

func (s *CheckingAccount) Transfer(receiver Account, amount int) string{
	  if s.balance <= 0 {
		return "Account balance is not enough"
	  }
	  if amount <= 0{
		return "Amount cannot be negative"
	  }
	  if s.balance < amount {
		return "Account balance is not enough"
	  }else {
		switch receiver.(type){
		case *SavingsAccount, *CheckingAccount, *InvestmentAccount:
			s.balance -= amount
			receiver.Deposit(amount)
			return "Success"
		default:
			 return "Invalid receiver account"
		}
	  }
}

func (s *InvestmentAccount) MonthlyInterest() int{
	return (s.balance * 2) / 100 / 12
}

func (s *InvestmentAccount) Deposit(amount int) string{
	if amount > 0{
		s.balance += amount
		return  "Success"
	}else {
		return "Amount cannot be negative"
	}
}

func (s *InvestmentAccount) Withdraw(amount int)string {
	if s.balance <= 0{
		return "Account balance is not enough"
	}
	if amount <= 0{
		return "Amount cannot be negative"
	}
	if s.balance < amount{
		return "Account balance is not enough" 
	}else{
		s.balance -= amount
		return "Success"
	}
}

func (s *InvestmentAccount) CheckBalance() int{
	return s.balance
}

func (s *InvestmentAccount) Transfer(receiver Account, amount int) string{
	  if s.balance <= 0 {
		return "Account balance is not enough"
	  }
	  if amount <= 0{
		return "Amount cannot be negative"
	  }
	  if s.balance < amount {
		return "Account balance is not enough"
	  }else {
		switch receiver.(type){
		case *SavingsAccount, *CheckingAccount, *InvestmentAccount:
			s.balance -= amount
			receiver.Deposit(amount)
			return "Success"
		default:
			 return "Invalid receiver account"
		}
	  }
}

func NewSavingsAccount() *SavingsAccount {
	return  &SavingsAccount{balance: 0}
}

func NewCheckingAccount() *CheckingAccount {
	return &CheckingAccount{balance: 0}
}

func NewInvestmentAccount() *InvestmentAccount {
	return &InvestmentAccount{balance: 0}
}
