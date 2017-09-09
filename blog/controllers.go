package blog

import (
	"github.com/slawek87/GOblog/storage"
	"time"
	"strconv"
)

func CreateBlogPost(post *Blog) (*Blog, error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	db, _ := storage.InitDB()
	db.NewRecord(&post)
	query := db.Create(&post)


	return post, query.Error
}

func ListBlogPosts() ([]Blog, error) {
	allPosts := []Blog{}

	db, _ := storage.InitDB()
	query := db.Find(&allPosts)

	return allPosts, query.Error
}

func GetBlogPost(id string) (*Blog, error) {
	post := Blog{}

	ID, err := strconv.Atoi(id)

	if err != nil {
		return &post, err
	}

	db, _ := storage.InitDB()
	query := db.First(&post, ID)

	return &post, query.Error
}

func UpdateBlogPost(post *Blog, id string) (*Blog, error) {
	ID, err := strconv.Atoi(id)

	if err != nil {
	    return post, err
	}

	record := Blog{}
	record.ID = uint(ID)

	db, _ := storage.InitDB()
	db.First(&record, uint(ID))

	record.Post = post.Post
	record.UpdatedAt = time.Now()

	query := db.Save(&record)

	return &record, query.Error
}

func DeleteBlogPost(id string) (*Blog, error) {
	post := Blog{}

	ID, err := strconv.Atoi(id)

	if err != nil {
		return &post, err
	}

	db, _ := storage.InitDB()
	db.First(&post, ID)
	query := db.Delete(&post)

	return &post, query.Error
}