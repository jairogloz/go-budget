package account

// Delete deletes an account.
func (s Service) Delete(userId, id string) error {
	return s.repo.Delete(userId, id)
}
