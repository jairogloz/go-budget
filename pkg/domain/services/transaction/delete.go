package transaction

// Delete removes a transaction from the database.
func (s Service) Delete(userId, transactionID string) error {
	return s.txRepo.Delete(userId, transactionID)
}
