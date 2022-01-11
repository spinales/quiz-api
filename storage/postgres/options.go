package postgres

import (
	"github.com/lib/pq"
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
		Answers:  pq.StringArray{"12", "7", "10", "0"},
	})

	db.Create(&models.Question{
		Content: "Si un bebé nace en Colombia, pero a los dos años se va a Ecuador, ¿dónde le crecen los dientes?",
		Answers: pq.StringArray{"En la boca", "Colombia", "Ecuador", "Venezuela"},
	})

	db.Create(&models.Question{
		Content:  "Estás corriendo en una carrera y adelantas a la persona que está en segundo lugar, ¿en qué posición pasas a estar?",
		ImageUrl: "",
		Answers:  pq.StringArray{"Segundo", "Primero", "Tercero", "Quinto"},
	})

	db.Create(&models.Question{
		Content:  "La palabra París comienza con “P” y termina con “T”, ¿cierto o falso?",
		ImageUrl: "",
		Answers:  pq.StringArray{"Cierto", "Falso"},
	})

	db.Create(&models.Question{
		Content:  "¿Si un tren eléctrico se mueve hacia el norte a 100 km/h y sopla el viento hacia el oeste a 10 km/h, hacia dónde irá el humo?",
		ImageUrl: "",
		Answers:  pq.StringArray{"No tendra humo", "Norte", "Sur", "Oeste"},
	})
}

func EraseData(DB *gorm.DB) {
	DB.Delete(&models.User{})
	DB.Delete(&models.Question{})
}
