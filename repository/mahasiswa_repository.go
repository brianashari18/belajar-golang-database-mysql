package repository

import (
	"context"
	"golang-database/entity"
)

type MahasiswaRepository interface {
	Insert(ctx context.Context, mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error)
	FindById(ctx context.Context, id int32) (entity.Mahasiswa, error)
	FindAll(ctx context.Context) ([]entity.Mahasiswa, error)
}
