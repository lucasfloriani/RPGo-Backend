package router

import (
	"rpgo-backend/handler"
	"rpgo-backend/middleware"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes add routes to the system
func InitializeRoutes(r *gin.Engine) {
	breedGroup := r.Group("/breed")
	{
		breedGroup.GET("/", middleware.Cors, handler.GetBreeds)
		breedGroup.GET("/:breedID", middleware.Cors, handler.GetBreed)
		breedGroup.POST("/", middleware.Cors, handler.CreateBreed)
		breedGroup.PUT("/:breedID", middleware.Cors, handler.UpdateBreed)
		breedGroup.DELETE("/:breedID", middleware.Cors, handler.DeleteBreed)
	}

	characterGroup := r.Group("/character")
	{
		characterGroup.GET("/", middleware.Cors, handler.GetCharacters)
		characterGroup.GET("/:characterID", middleware.Cors, handler.GetCharacter)
		characterGroup.POST("/", middleware.Cors, handler.CreateCharacter)
		characterGroup.PUT("/:characterID", middleware.Cors, handler.UpdateCharacter)
		characterGroup.DELETE("/:characterID", middleware.Cors, handler.DeleteCharacter)
	}

	classGroup := r.Group("/class")
	{
		classGroup.GET("/", middleware.Cors, handler.GetClasses)
		classGroup.GET("/:classID", middleware.Cors, handler.GetClass)
		classGroup.POST("/", middleware.Cors, handler.CreateClass)
		classGroup.PUT("/:classID", middleware.Cors, handler.UpdateClass)
		classGroup.DELETE("/:classID", middleware.Cors, handler.DeleteClass)
	}

	diceGroup := r.Group("/dice")
	{
		diceGroup.GET("/", middleware.Cors, handler.GetDices)
		diceGroup.GET("/:diceID", middleware.Cors, handler.GetDice)
		diceGroup.POST("/", middleware.Cors, handler.CreateDice)
		diceGroup.PUT("/:diceID", middleware.Cors, handler.UpdateDice)
		diceGroup.DELETE("/:diceID", middleware.Cors, handler.DeleteDice)
	}

	equipmentGroup := r.Group("/equipment")
	{
		equipmentGroup.GET("/", middleware.Cors, handler.GetEquipments)
		equipmentGroup.GET("/:equipmentID", middleware.Cors, handler.GetEquipment)
		equipmentGroup.POST("/", middleware.Cors, handler.CreateEquipment)
		equipmentGroup.PUT("/:equipmentID", middleware.Cors, handler.UpdateEquipment)
		equipmentGroup.DELETE("/:equipmentID", middleware.Cors, handler.DeleteEquipment)
	}

	equipmentTypeGroup := r.Group("/equipment-type")
	{
		equipmentTypeGroup.GET("/", middleware.Cors, handler.GetEquipmentTypes)
		equipmentTypeGroup.GET("/:equipmentTypeID", middleware.Cors, handler.GetEquipmentType)
		equipmentTypeGroup.POST("/", middleware.Cors, handler.CreateEquipmentType)
		equipmentTypeGroup.PUT("/:equipmentTypeID", middleware.Cors, handler.UpdateEquipmentType)
		equipmentTypeGroup.DELETE("/:equipmentTypeID", middleware.Cors, handler.DeleteEquipmentType)
	}

	itemGroup := r.Group("/item")
	{
		itemGroup.GET("/", middleware.Cors, handler.GetItens)
		itemGroup.GET("/:itemID", middleware.Cors, handler.GetItem)
		itemGroup.POST("/", middleware.Cors, handler.CreateItem)
		itemGroup.PUT("/:itemID", middleware.Cors, handler.UpdateItem)
		itemGroup.DELETE("/:itemID", middleware.Cors, handler.DeleteItem)
	}

	itemTypeGroup := r.Group("/item-type")
	{
		itemTypeGroup.GET("/", middleware.Cors, handler.GetItemTypes)
		itemTypeGroup.GET("/:itemTypeID", middleware.Cors, handler.GetItemType)
		itemTypeGroup.POST("/", middleware.Cors, handler.CreateItemType)
		itemTypeGroup.PUT("/:itemTypeID", middleware.Cors, handler.UpdateItemType)
		itemTypeGroup.DELETE("/:itemTypeID", middleware.Cors, handler.DeleteItemType)
	}

	monsterGroup := r.Group("/monster")
	{
		monsterGroup.GET("/", middleware.Cors, handler.GetMonsters)
		monsterGroup.GET("/:monsterID", middleware.Cors, handler.GetMonster)
		monsterGroup.POST("/", middleware.Cors, handler.CreateMonster)
		monsterGroup.PUT("/:monsterID", middleware.Cors, handler.UpdateMonster)
		monsterGroup.DELETE("/:monsterID", middleware.Cors, handler.DeleteMonster)
	}

	monsterTypeGroup := r.Group("/monster-type")
	{
		monsterTypeGroup.GET("/", middleware.Cors, handler.GetMonsterTypes)
		monsterTypeGroup.GET("/:monsterTypeID", middleware.Cors, handler.GetMonsterType)
		monsterTypeGroup.POST("/", middleware.Cors, handler.CreateMonsterType)
		monsterTypeGroup.PUT("/:monsterTypeID", middleware.Cors, handler.UpdateMonsterType)
		monsterTypeGroup.DELETE("/:monsterTypeID", middleware.Cors, handler.DeleteMonsterType)
	}

	npcGroup := r.Group("/npc")
	{
		npcGroup.GET("/", middleware.Cors, handler.GetNpcs)
		npcGroup.GET("/:npcID", middleware.Cors, handler.GetNpc)
		npcGroup.POST("/", middleware.Cors, handler.CreateNpc)
		npcGroup.PUT("/:npcID", middleware.Cors, handler.UpdateNpc)
		npcGroup.DELETE("/:npcID", middleware.Cors, handler.DeleteNpc)
	}

	skillGroup := r.Group("/skill")
	{
		skillGroup.GET("/", middleware.Cors, handler.GetSkills)
		skillGroup.GET("/:skillID", middleware.Cors, handler.GetSkill)
		skillGroup.POST("/", middleware.Cors, handler.CreateSkill)
		skillGroup.PUT("/:skillID", middleware.Cors, handler.UpdateSkill)
		skillGroup.DELETE("/:skillID", middleware.Cors, handler.DeleteSkill)
	}

	weaponGroup := r.Group("/weapon")
	{
		weaponGroup.GET("/", middleware.Cors, handler.GetWeapons)
		weaponGroup.GET("/:weaponID", middleware.Cors, handler.GetWeapon)
		weaponGroup.POST("/", middleware.Cors, handler.CreateWeapon)
		weaponGroup.PUT("/:weaponID", middleware.Cors, handler.UpdateWeapon)
		weaponGroup.DELETE("/:weaponID", middleware.Cors, handler.DeleteWeapon)
	}

	weaponTypeGroup := r.Group("/weapon-type")
	{
		weaponTypeGroup.GET("/", middleware.Cors, handler.GetWeaponTypes)
		weaponTypeGroup.GET("/:weaponTypeID", middleware.Cors, handler.GetWeaponType)
		weaponTypeGroup.POST("/", middleware.Cors, handler.CreateWeaponType)
		weaponTypeGroup.PUT("/:weaponTypeID", middleware.Cors, handler.UpdateWeaponType)
		weaponTypeGroup.DELETE("/:weaponTypeID", middleware.Cors, handler.DeleteWeaponType)
	}
}
