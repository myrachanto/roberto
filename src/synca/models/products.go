package models

import (
	"time"

	httperrors "github.com/myrachanto/erroring"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name"`
	Url           string             `json:"url"`
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	Meta          string             `json:"meta"`
	Altertag      string             `json:"altertag"`
	Footer        string             `json:"footer"`
	Code          string             `json:"code"`
	Source        string             `json:"source"`
	Majorcategory string             `json:"majorcat"`
	Category      string             `json:"category"`
	Newarrivals   string             `json:"newarrivals"`
	Subcategory   string             `json:"subcategory"`
	Oldprice      float64            `json:"oldprice"`
	Newprice      float64            `json:"newprice"`
	Likes         int64              `json:"likes"`
	Buyprice      float64            `json:"buyprice"`
	Picture       string             `json:"picture"`
	Quantity      float64            `json:"quantity"`
	Services      []Service          `json:"services"`
	Supercategory string             `json:"supercategory"`
	Desc          string             `json:"desc"`
	Shopalias     string             `json:"shopalias"`
	Rating        *Rating            `json:"rating"`
	Rates         []Rating           `json:"rates"`
	Images        []Picture          `json:"images"`
	Tag           []Tag              `json:"tag"`
	Featured      bool               `json:"featured"`
	Promotion     bool               `json:"promotion"`
	Hotdeals      bool               `json:"hotdeals"`
	Base          Base               `json:"base,omitempty"`
}
type Rating struct {
	Author      string `json:"author"`
	Bestrate    int64  `json:"bestrate"`
	Rate        int64  `json:"rate"`
	Description string `json:"description"`
	TotalCount  int    `json:"totalcount"`
}
type Newarrivals struct {
	Product []*Product `json:"product"`
}
type Service struct {
	Productcode string `json:"productcode"`
	Name        string `json:"name"`
	Price       string `json:"price"`
}
type Majorcat struct {
	Name string `json:"name"`
}
type Supercat struct {
	Name string `json:"name"`
}
type Categs struct {
	Name string `json:"name"`
}
type Tag struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
type Picture struct {
	Productcode string `json:"productcode"`
	Name        string `json:"name"`
}
type Color struct {
	Productcode string `json:"productcode"`
	Name        string `json:"name"`
}

type Results struct {
	Data  []*Product `json:"data"`
	Total int64      `json:"total"`
}

func (product Product) Validate() httperrors.HttpErr {
	if product.Name == "" {
		return httperrors.NewNotFoundError("Invalid Name")
	}
	if product.Title == "" {
		return httperrors.NewNotFoundError("Invalid title")
	}
	// if product.Description == "" {
	// 	return httperrors.NewNotFoundError("Invalid Description")
	// }
	return nil
}

type Base struct {
	Created_At time.Time  `bson:"created_at"`
	Updated_At time.Time  `bson:"updated_at"`
	Delete_At  *time.Time `bson:"deleted_at"`
}
