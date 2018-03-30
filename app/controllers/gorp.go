package controllers

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/modules/db/app"
	r "github.com/revel/revel"
	"github.com/s4ntos/remote/app/models"
	_ "github.com/ziutek/mymysql/godrv"
	_ "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	"golang.org/x/crypto/bcrypt"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()

	dbDriver := r.Config.StringDefault("db.driver", "sqlite3")

	if dbDriver == "mysql" || dbDriver == "mymysql" {
		Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	} else if dbDriver == "postgres" {
		Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}
	} else {
		Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}
	}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.Role{}).SetKeys(true, "RoleId")
	setColumnSizes(t, map[string]int{
		"Role":        64,
		"Description": 400,
		"Privileges":  200, // To be changed for know only Admin or something else
	})
	t = Dbm.AddTable(models.Roles{})

	t = Dbm.AddTable(models.User{}).SetKeys(true, "UserId")
	t.ColMap("Password").Transient = true
	t.ColMap("Created").Transient = true
	t.ColMap("Profile").Transient = true
	setColumnSizes(t, map[string]int{
		"Email": 200,
	})

	t = Dbm.AddTable(models.Token{}).SetKeys(true, "TokenId")
	setColumnSizes(t, map[string]int{
		"Email": 200,
		"Type":  20,
		"Hash":  16,
	})

	t = Dbm.AddTable(models.Profile{}).SetKeys(true, "ProfileId")
	t.ColMap("User").Transient = true
	setColumnSizes(t, map[string]int{
		"UserName":    64,
		"Name":        100,
		"Summary":     140,
		"Description": 400,
		"PhotoUrl":    200,
	})

	t = Dbm.AddTable(models.Post{}).SetKeys(true, "PostId")
	t.ColMap("DateObj").Transient = true
	t.ColMap("ContentStr").Transient = true
	setColumnSizes(t, map[string]int{
		"Title":   400,
		"Content": 16777212, // mediumblob storage capacity
	})

	// Social components
	Dbm.AddTable(models.Like{}).SetKeys(true, "LikeId")
	Dbm.AddTable(models.Follower{}).SetKeys(true, "FollowerId")

	Dbm.TraceOn("[gorp]", r.INFO)

	// Create tables in datastore if they don't already exist
	Dbm.CreateTablesIfNotExists()

	count, err := Dbm.SelectInt("select count(1) from User")
	if err != nil {
		panic(err)
	}

	// Set up database if we don't have any users inside
	if count == 0 {
		adminUsername := r.Config.StringDefault("admin.username", "admin")
		adminPassword := r.Config.StringDefault("admin.password", "adminuser")

		adminRole := &models.Role{
			Role:        "admin",
			Description: "Admin user",
			Privileges:  "admin",
		}
		if err := Dbm.Insert(adminRole); err != nil {
			panic(err)
		}

		bcryptAdminPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(adminPassword), bcrypt.DefaultCost)
		adminUser := &models.User{
			Email:          "admin@demo.com",
			HashedPassword: bcryptAdminPassword,
			Confirmed:      false,
		}
		if err := Dbm.Insert(adminUser); err != nil {
			panic(err)
		}

		adminProfile := &models.Profile{
			UserId:             adminUser.UserId,
			UserName:           adminUsername,
			Name:               "Admin User",
			Summary:            "Just the admin of the joint",
			Description:        "Yes I'm the Admin",
			PhotoUrl:           "/public/images/admin.png",
			User:               adminUser,
			AggregateFollowing: 0,
		}

		if err := Dbm.Insert(adminProfile); err != nil {
			panic(err)
		}
		adminRoles := &models.Roles{
			RoleId:    adminRole.RoleId,
			ProfileId: adminProfile.ProfileId,
		}
		if err := Dbm.Insert(adminRoles); err != nil {
			panic(err)
		}
	}

}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
