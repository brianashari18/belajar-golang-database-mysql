package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type mahasiswaRepositoryImpl struct {
	DB *sql.DB
}

func NewMahasiswaRepository(db *sql.DB) MahasiswaRepository {
	return &mahasiswaRepositoryImpl{DB: db}
}

func (r *mahasiswaRepositoryImpl) Insert(ctx context.Context, mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	script := "INSERT INTO mahasiswa(nama, alamat) VALUES(?, ?)"
	result, err := r.DB.ExecContext(ctx, script, mahasiswa.Nama, mahasiswa.Alamat)
	if err != nil {
		return mahasiswa, err
	}

	nim, err := result.LastInsertId()
	if err != nil {
		return mahasiswa, err
	}
	mahasiswa.NIM = int32(nim)
	return mahasiswa, nil
}

func (r *mahasiswaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Mahasiswa, error) {
	script := "SELECT nim, nama, alamat FROM mahasiswa WHERE nim = ? LIMIT 1"
	rows, err := r.DB.QueryContext(ctx, script, id)
	mahasiswa := entity.Mahasiswa{}
	if err != nil {
		return mahasiswa, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&mahasiswa.NIM, &mahasiswa.Nama, &mahasiswa.Alamat)
		return mahasiswa, nil
	} else {
		return mahasiswa, errors.New("NIM : " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (r *mahasiswaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Mahasiswa, error) {
	script := "SELECT nim, nama, alamat FROM mahasiswa"
	rows, err := r.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tabMahasiswa []entity.Mahasiswa
	for rows.Next() {
		mahasiswa := entity.Mahasiswa{}
		rows.Scan(&mahasiswa.NIM, &mahasiswa.Nama, &mahasiswa.Alamat)
		tabMahasiswa = append(tabMahasiswa, mahasiswa)
	}
	return tabMahasiswa, nil
}
