package modeltest

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/haryanapnx/go-blog-crud/api/controllers"
	"github.com/haryanapnx/go-blog-crud/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var articleInstance = models.Article{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Database()
	os.Exit(m.Run())

}

func Database() {
	var err error
	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("TestDbUser"),
			os.Getenv("TestDbPassword"),
			os.Getenv("TestDbHost"),
			os.Getenv("TestDbPort"),
			os.Getenv("TestDbName"),
		)
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			os.Getenv("TestDbHost"),
			os.Getenv("TestDbPort"),
			os.Getenv("TestDbUser"),
			os.Getenv("TestDbName"),
			os.Getenv("TestDbPassword"),
		)
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}

	log.Printf("successfully refresh table")

	return nil
}

func seedOneUser() (models.User, error) {
	refreshUserTable()
	user := models.User{
		Username: "test username",
		Email:    "test@gmail.com",
		Password: "123123",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	return user, nil
}

func seedUsers() error {
	users := []models.User{
		models.User{
			Username: "user test 1",
			Email:    "test1@gmail.com",
			Password: "123123",
		},
		models.User{
			Username: "user test 2",
			Email:    "test2@gmail.com",
			Password: "123123",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil

}

func refreshUserAndArticleTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, models.Article{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}, models.Article{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOneArticles() (models.Article, error) {

	err := refreshUserAndArticleTable()
	if err != nil {
		return models.Article{}, err
	}
	user := models.User{
		Username: "Test user A",
		Email:    "usrA@gmail.com",
		Password: "123123",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Article{}, err
	}
	post := models.Article{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Article{}).Create(&post).Error
	if err != nil {
		return models.Article{}, err
	}
	return post, nil
}

func seedUsersAndArticles() ([]models.User, []models.Article, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Article{}, err
	}
	var users = []models.User{
		models.User{
			Username: "test B",
			Email:    "testB@gmail.com",
			Password: "123123",
		},
		models.User{
			Username: "Tree s",
			Email:    "tree@gmail.com",
			Password: "123123",
		},
	}
	var posts = []models.Article{
		models.Article{
			Title:   "Title 1",
			Content: "Hello world 1",
		},
		models.Article{
			Title:   "Title 2",
			Content: "Hello world 2",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Article{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, posts, nil
}
