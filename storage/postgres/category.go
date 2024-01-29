package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/api/models"
	"main.go/storage"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.ICategoryStorage {
	return categoryRepo{db: db}
}
func (c categoryRepo) Create(category models.CreateCategory) (string, error) {
	id := uuid.New()
	query := `insert into categories (id, name) values($1, $2)`

	if _, err := c.db.Exec(context.Background(),query, id, category.Name); err != nil {
		fmt.Println("error is while creating category", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (c categoryRepo) GetByID(key models.PrimaryKey) (models.Category, error) {
	category := models.Category{}

	query := `select id, name from categories where id = $1`
	if err := c.db.QueryRow(context.Background(),query, key.ID).Scan(&category.ID, &category.Name); err != nil {
		fmt.Println("error is while getting by id", err.Error())
		return models.Category{}, err
	}

	return category, nil
}

func (c categoryRepo) GetList(request models.GetListRequest) (models.CategoryResponse, error) {
	var (
		query, countQuery string
		count             = 0
		page              = request.Page
		offset            = (page - 1) * request.Limit
		search            = request.Search
		categories        = []models.Category{}
	)

	countQuery = `select count(1) from categories `

	if search != "" {
		countQuery += fmt.Sprintf(` where name ilike '%%%s%%'`, search)
	}

	if err := c.db.QueryRow(context.Background(),countQuery).Scan(&count); err != nil {
		fmt.Println("error is while scanning count", err.Error())
		return models.CategoryResponse{}, err
	}

	query = `select id, name from categories `

	if search != "" {
		query += fmt.Sprintf(` where name ilike '%%%s%%' `, search)
	}

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(context.Background(),query, request.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting categories", err.Error())
		return models.CategoryResponse{}, err
	}

	for rows.Next() {
		cat := models.Category{}
		if err = rows.Scan(&cat.ID, &cat.Name); err != nil {
			fmt.Println("error is while scanning category", err.Error())
			return models.CategoryResponse{}, err
		}
		categories = append(categories, cat)
	}
	return models.CategoryResponse{
		Category: categories,
		Count:    count,
	}, err
}

func (c categoryRepo) Update(category models.UpdateCategory) (string, error) {
	query := `update categories set name = $1 where id = $2`

	if _, err := c.db.Exec(context.Background(),query, &category.Name, &category.ID); err != nil {
		fmt.Println("error is while updating category", err.Error())
		return "", err
	}
	return category.ID, nil
}

func (c categoryRepo) Delete(key models.PrimaryKey) error {
	if _, err := c.db.Exec(context.Background(),`delete from products where category_id=$1`, key.ID); err != nil {
		return err
	}
	query := `delete from categories where id = $1`

	if _, err := c.db.Exec(context.Background(),query, key.ID); err != nil {
		fmt.Println("error is while deleting category", err.Error())
		return err
	}
	return nil
}
