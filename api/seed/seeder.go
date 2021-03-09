package seed

import (
	"log"

	"github.com/haryanapnx/go-blog-crud/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Username: "hary ganteng",
		Email:    "hary@gmail.com",
		Password: "123123",
	},
	models.User{
		Username: "haryana tampan",
		Email:    "hary.vai@gmail.com",
		Password: "123123",
	},
}

var articles = []models.Article{
	models.Article{
		Title:   "berita hari ini",
		Content: "ini content nya",
	},
	models.Article{
		Title:   "berita hari itu",
		Content: "ini descriptnya nya",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Article{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Article{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		articles[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Article{}).Create(&articles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed articles table: %v", err)
		}
	}

}
