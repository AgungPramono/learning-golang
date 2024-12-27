package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	golang_databases "golang-database"
	"golang-database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_databases.GetConnection())

	comment := entity.Comment{
		Email:   "repository@mail.com",
		Comment: "test repository",
	}
	result, err := commentRepository.Insert(context.Background(), comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUpdateComment(t *testing.T) {
	commentRepository := NewCommentRepository(golang_databases.GetConnection())
	comment := entity.Comment{
		Email:   "email.baru@mail.com",
		Comment: "update data baru",
		Id:      int32(10),
	}

	_, err := commentRepository.Update(context.Background(), comment)
	if err != nil {
		return
	}
	//cari lagi berdasarkan id
	newResult, err := commentRepository.FindById(context.Background(), comment.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(newResult)
}

func TestDeleteComment(t *testing.T) {
	commentRepository := NewCommentRepository(golang_databases.GetConnection())

	id := int32(13)
	result, err := commentRepository.Delete(context.Background(), id)
	if err != nil {
		return
	}
	if result == true {
		fmt.Println("delete success")
	} else {
		fmt.Println("delete fail")
	}
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_databases.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 22)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_databases.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
