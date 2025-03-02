// #region CH15.1
func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return fmt.Errorf("error updating user status: %w", err)

	}
	err = a.track(em.status)
	if err != nil {
		return fmt.Errorf("error tracking user bounce: %w", err)

	}
	return nil
}
// #endregion

// #CH15.2: checkPermission(Admin) will raise an error (admin ist nicht als Wert für ein const vom Typ perm definiert)
// #CH15.3: Es wird kein Fehler geworfen