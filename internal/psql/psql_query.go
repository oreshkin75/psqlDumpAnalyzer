package psql

func (c *Creator) Select(query string) ([]string, error) {
	rows, err := c.db.Query(query)
	if err != nil {
		c.logError.Print(err)
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		c.logError.Print(err)
		return nil, err
	}

	c.logInfo.Print("psql select complete: ", query)
	return columns, nil
}

func (c *Creator) Insert(query string) error {
	_, err := c.db.Exec(query)
	if err != nil {
		c.logError.Print(err)
		return err
	}
	c.logInfo.Print("psql insert complete: ", query)
	return nil
}

func (c *Creator) Update(query string) error {
	_, err := c.db.Exec(query)
	if err != nil {
		c.logError.Print(err)
		return err
	}
	c.logInfo.Print("psql update complete: ", query)
	return nil
}

func (c *Creator) Delete(query string) error {
	_, err := c.db.Exec(query)
	if err != nil {
		c.logError.Print(err)
		return err
	}
	c.logInfo.Print("psql delete complete: ", query)
	return nil
}
