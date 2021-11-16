package psql

func (c *Creator) Select(query string) error {
	_, err := c.db.Query(query)
	if err != nil {
		c.logError.Print(err)
		return err
	}

	return nil
}

func (c *Creator) Insert(query string) error {
	_, err := c.db.Exec(query)
	if err != nil {
		c.logError.Print(err)
		return err
	}

	return nil
}

func (c *Creator) Update(query string) error {
	_, err := c.db.Exec(query)
	if err != nil {
		c.logError.Print(err)
		return err
	}

	return nil
}

func (c *Creator) Delete(query string) error {
	_, err := c.db.Exec(query)
	if err != nil {
		c.logError.Print(err)
		return err
	}

	return nil
}
