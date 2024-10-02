package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertMahasiswa(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(golangdatabase.GetConnection())

	ctx := context.Background()
	mahasiswa := entity.Mahasiswa{
		Nama:   "Anashari",
		Alamat: "Pondok Pinang",
	}

	result, err := mahasiswaRepository.Insert(ctx, mahasiswa)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByIdMahasiswa(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(golangdatabase.GetConnection())

	result, err := mahasiswaRepository.FindById(context.Background(), 1301223200)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAllMahasiswa(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(golangdatabase.GetConnection())

	results, err := mahasiswaRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
