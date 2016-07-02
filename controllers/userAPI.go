package controllers

func (c *UserController) API_Profile() {
	type user struct {
		UserID string
		Hobby  []string
	}

	u := user{
		"jike",
		[]string{"tom", "jerry"},
	}

	c.Data["json"] = u

	c.ServeJSON()
}
