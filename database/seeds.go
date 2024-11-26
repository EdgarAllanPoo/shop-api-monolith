package database

import (
	"log"

	"github.com/EdgarAllanPoo/shop-api-monolith/models"
)

func SeedProducts() {
	products := []models.Product{
		{Name: "Laptop", Category: "Electronics", Description: "A high-performance laptop", Price: 1500},
		{Name: "Phone", Category: "Electronics", Description: "A smartphone with a great camera", Price: 800},
		{Name: "Tablet", Category: "Electronics", Description: "A lightweight and portable tablet", Price: 600},
		{Name: "Headphones", Category: "Electronics", Description: "Noise-cancelling over-ear headphones", Price: 200},
		{Name: "Smartwatch", Category: "Electronics", Description: "A smartwatch with health tracking features", Price: 250},

		{Name: "Shirt", Category: "Clothing", Description: "A comfortable cotton shirt", Price: 25},
		{Name: "Pants", Category: "Clothing", Description: "Stylish denim pants", Price: 50},
		{Name: "Jacket", Category: "Clothing", Description: "A warm winter jacket", Price: 120},
		{Name: "Shoes", Category: "Clothing", Description: "Durable running shoes", Price: 80},
		{Name: "Hat", Category: "Clothing", Description: "A trendy baseball cap", Price: 15},

		{Name: "Notebook", Category: "Stationery", Description: "A spiral-bound notebook", Price: 5},
		{Name: "Pen", Category: "Stationery", Description: "A pack of gel pens", Price: 10},
		{Name: "Marker", Category: "Stationery", Description: "A set of colorful markers", Price: 8},
		{Name: "Binder", Category: "Stationery", Description: "A durable three-ring binder", Price: 12},
		{Name: "Sticky Notes", Category: "Stationery", Description: "A pack of sticky notes", Price: 3},

		{Name: "Basketball", Category: "Sports", Description: "An official size basketball", Price: 30},
		{Name: "Soccer Ball", Category: "Sports", Description: "A premium quality soccer ball", Price: 25},
		{Name: "Tennis Racket", Category: "Sports", Description: "A lightweight tennis racket", Price: 100},
		{Name: "Yoga Mat", Category: "Sports", Description: "A non-slip yoga mat", Price: 20},
		{Name: "Dumbbells", Category: "Sports", Description: "A set of adjustable dumbbells", Price: 150},
	}

	for _, product := range products {
		DB.Create(&product)
	}

	log.Println("Sample products seeded")
}
