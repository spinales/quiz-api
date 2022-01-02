package postgres

import (
	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
	"gorm.io/gorm"
)

func FakeData(db *gorm.DB) {
	pass, _ := util.HashPassword("pass123")

	db.Create(&models.User{
		Username: "admin",
		Email:    "admin@mail.com",
		Password: pass,
		Role:     models.Admin,
	})

	db.Create(&models.Question{
		Content:  "Si en una pecera hay 12 peces y 5 de ellos se ahogan, ¿cuántos peces quedan?",
		ImageUrl: "",
	})

	db.Create([]models.Answer{
		{
			Content:    "12",
			Score:      5,
			QuestionID: 1,
		},
		{
			Content:    "7",
			Score:      2,
			QuestionID: 1,
		},
		{
			Content:    "10",
			Score:      0,
			QuestionID: 1,
		},
		{
			Content:    "0",
			Score:      0,
			QuestionID: 1,
		},
	})

	db.Create(&models.Question{
		Content:  "Si un bebé nace en Colombia, pero a los dos años se va a Ecuador, ¿dónde le crecen los dientes?",
		ImageUrl: "",
	})

	db.Create([]models.Answer{
		{
			Content:    "Colombia",
			Score:      0,
			QuestionID: 2,
		},
		{
			Content:    "En la boca",
			Score:      5,
			QuestionID: 2,
		},
		{
			Content:    "Ecuador",
			Score:      0,
			QuestionID: 2,
		},
		{
			Content:    "Venezuela",
			Score:      0,
			QuestionID: 2,
		},
	})

	db.Create(&models.Question{
		Content:  "Estás corriendo en una carrera y adelantas a la persona que está en segundo lugar, ¿en qué posición pasas a estar?",
		ImageUrl: "",
	})

	db.Create([]models.Answer{
		{
			Content:    "Primero",
			Score:      0,
			QuestionID: 3,
		},
		{
			Content:    "Tercero",
			Score:      0,
			QuestionID: 3,
		},
		{
			Content:    "Segundo",
			Score:      5,
			QuestionID: 3,
		},
		{
			Content:    "Quinto",
			Score:      0,
			QuestionID: 3,
		},
	})

	db.Create(&models.Question{
		Content:  "La palabra París comienza con “P” y termina con “T”, ¿cierto o falso?",
		ImageUrl: "",
	})

	db.Create([]models.Answer{
		{
			Content:    "Cierto",
			Score:      5,
			QuestionID: 4,
		},
		{
			Content:    "Falso",
			Score:      0,
			QuestionID: 4,
		},
	})

	db.Create(&models.Question{
		Content:  "¿Si un tren eléctrico se mueve hacia el norte a 100 km/h y sopla el viento hacia el oeste a 10 km/h, hacia dónde irá el humo?",
		ImageUrl: "",
	})

	db.Create([]models.Answer{
		{
			Content:    "Norte",
			Score:      0,
			QuestionID: 5,
		},
		{
			Content:    "Sur",
			Score:      0,
			QuestionID: 5,
		},
		{
			Content:    "Oeste",
			Score:      0,
			QuestionID: 5,
		},
		{
			Content:    "No tendra humo",
			Score:      5,
			QuestionID: 5,
		},
	})
}
