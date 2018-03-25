package jobs

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/modules/jobs/app/jobs"
	"github.com/s4ntos/remote/app/controllers"
	"github.com/s4ntos/remote/app/models"
)

// Periodically count the users in the database.
type UserCounter struct{}

func (c UserCounter) Run() {
	users, err := controllers.Dbm.Select(&models.User{},
		`select * from User`)
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d users.\n", len(users))
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1m", UserCounter{})
	})
}
