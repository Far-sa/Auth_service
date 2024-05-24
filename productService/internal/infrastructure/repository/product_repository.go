package repository

import (
	"database/sql"
	"product-service/internal/domain"
	"time"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db}
}

func (r *PostgresProductRepository) CreateProduct(product *domain.Product) error {
	query := `INSERT INTO products (name, description, price, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, time.Now(), time.Now()).Scan(&product.ID)
	return err
}

func (r *PostgresProductRepository) GetProduct(id int64) (*domain.Product, error) {
	query := `SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = $1`
	row := r.db.QueryRow(query, id)
	product := &domain.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *PostgresProductRepository) UpdateProduct(product *domain.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, updated_at = $4 WHERE id = $5`
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, time.Now(), product.ID)
	return err
}

func (r *PostgresProductRepository) DeleteProduct(id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresProductRepository) ListProducts() ([]*domain.Product, error) {
	query := `SELECT id, name, description, price, created_at, updated_at FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
