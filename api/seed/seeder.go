package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/martinyonathann/fullstack/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Martin Yonathan",
		Email:    "martin.yonathan305@gmail.com",
		Password: "password",
	}, models.User{
		Nickname: "jarwo",
		Email:    "jarwo@gmail.com",
		Password: "pass1234",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello Word 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello Word 2",
	},
}

func load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table : %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table : %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table:%v", err)
		}
	}
}
