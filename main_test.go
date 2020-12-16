package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestPreloadWithSelect(t *testing.T) {
	user1 := User{Name: "emin",
		Languages: []Language{{
			Name: "Bosnian",
			Code: "ba",
		},
			{
				Name: "English",
				Code: "en",
			},
		},
	}
	user2 := User{Name: "jinzhu",
		Languages: []Language{{
			Name: "Chinese",
			Code: "zh",
		},
			{
				Name: "English",
				Code: "en",
			},
		},
	}

	DB.Create(&user1)
	DB.Create(&user2)

	var users []User
	//this works
	if err := DB.Preload("Languages").Find(&users); err.Error != nil {
		t.Errorf("failed, got error: %v", err.Error)
	}
	log.Println(users)

	var users2 []User
	//this works also
	if err := DB.Preload("Languages", func(db *gorm.DB) *gorm.DB {
		return db.Select("code", "name")
	}).Find(&users2); err.Error != nil {
		t.Errorf("failed, got error: %v", err.Error)
	}
	log.Println(users2)
}
