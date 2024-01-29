package postgres

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/api/models"
	"main.go/storage"
)

type basketRepo struct {
	db *pgxpool.Pool
}

func NewBasketRepo(db *pgxpool.Pool) storage.IBasketStorage {
	return basketRepo{db: db}
}

func (b basketRepo) Create(basket models.CreateBasket) (string, error) {
	id := uuid.New()

	if _, err := b.db.Exec(context.Background(), `insert into baskets(id, customer_id, total_sum)
	values($1, $2, $3)`, id, basket.CustomerID, basket.TotalSum); err != nil {
		fmt.Println("error is while inserting data", err.Error())
		return "", err
	}

	return id.String(), nil
}

func (b basketRepo) GetByID(key models.PrimaryKey) (models.Basket, error) {
	basket := models.Basket{}

	if err := b.db.QueryRow(context.Background(), `select id, customer_id, total_sum from baskets where id = $1`,
		key.ID).Scan(&basket.ID,
		&basket.CustomerID,
		&basket.TotalSum); err != nil {
		fmt.Println("error is while selecting basket", err.Error())
		return models.Basket{}, err
	}

	return basket, nil
}

func (b basketRepo) GetList(req models.GetListRequest) (models.BasketResponse, error) {
	var (
		baskets           = []models.Basket{}
		count             = 0
		query, countQuery string
		page              = req.Page
		offset            = (page - 1) * req.Limit
		search            = req.Search
	)

	countQuery = `select count(1) from baskets `

	if search != "" {
		countQuery += fmt.Sprintf(` and total_sum ilike '%%%s%%'`, search)
	}
	if err := b.db.QueryRow(context.Background(), countQuery).Scan(&count); err != nil {
		fmt.Println("error is while selecting count", err.Error())
		return models.BasketResponse{}, err
	}

	query = `select id, customer_id, total_sum from baskets `

	if search != "" {
		query += fmt.Sprintf(` and total_sum ilike '%%%s%%'`, search)
	}

	query += `LIMIT $1 OFFSET $2`
	rows, err := b.db.Query(context.Background(), query, req.Limit, offset)
	if err != nil {
		fmt.Println("error is while selecting baskets", err.Error())
		return models.BasketResponse{}, err
	}

	for rows.Next() {
		b := models.Basket{}
		if err = rows.Scan(&b.ID, &b.CustomerID, &b.TotalSum); err != nil {
			fmt.Println("error is while scanning data", err.Error())
			return models.BasketResponse{}, err
		}
		baskets = append(baskets, b)

	}

	return models.BasketResponse{
		Baskets: baskets,
		Count:   count,
	}, nil
}

func (b basketRepo) Update(basket models.UpdateBasket) (string, error) {
	bas := models.Basket{}

	if _, err := b.db.Exec(context.Background(), `update baskets set customer_id = $1, total_sum = $2 where id = $3`, &basket.CustomerID, &basket.TotalSum, &basket.ID); err != nil {
		return "", err
	}

	if err := b.db.QueryRow(context.Background(), `select id, customer_id, total_sum from baskets where id = $1`, basket.ID).Scan(&bas.ID, &bas.CustomerID, &bas.TotalSum); err != nil {
		fmt.Println("error is while selecting ", err.Error())
		return "", err
	}
	return bas.ID, nil
}

func (b basketRepo) Delete(key models.PrimaryKey) error {
	if _, err := b.db.Exec(context.Background(), `delete from basket_products where basket_id=$1`, key.ID); err != nil {
		return err
	}
	if _, err := b.db.Exec(context.Background(), `delete from baskets where id = $1`, key.ID); err != nil {
		return err
	}
	return nil
}

func (b basketProductRepo) AddProducts(basketID string, products map[string]int) error {
	query := `
			insert into basket_products 
			    (id, basket_id, product_id, quantity) 
					values ($1, $2, $3, $4)
`

	for productID, quantity := range products {
		if _, err := b.db.Exec(context.Background(), query, uuid.New(), basketID, productID, quantity); err != nil {
			fmt.Println("Error while adding product to basket_products table", err.Error())
			return err
		}

	}
	return nil
}
