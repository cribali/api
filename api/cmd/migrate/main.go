package main

import (
	"fmt"
	"suffgo/cmd/config"
	"suffgo/cmd/database"
	o "suffgo/internal/options/infrastructure/models"
	p "suffgo/internal/proposals/infrastructure/models"
	r "suffgo/internal/rooms/infrastructure/models"
	s "suffgo/internal/settingsRoom/infrastructure/models"
	m "suffgo/internal/users/infrastructure/models"
	ur "suffgo/internal/userRooms/infrastructure/models"
	e "suffgo/internal/votes/infrastructure/models"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)

	MigrateUser(db)
	MigrateRoom(db)
	MigrateProposal(db)
	MigrateOption(db)
	MigrateRoomSetting(db)
	MigrateVote(db)

	err := MakeConstraints(db)
	if err != nil {
		fmt.Printf("Error al agregar la clave foránea: %v\n", err)
		panic(err.Error())
	}
}

func MigrateUser(db database.Database) error {
	err := db.GetDb().Sync2(new(m.Users))

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Se ha migrado User con exito\n")
	}

	return err
}

func MigrateRoom(db database.Database) error {
	err := db.GetDb().Sync2(new(r.Room), new(ur.UserRoom))

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Se ha migrado Room y UserRoom con exito\n")
	}

	return err
}

func MigrateProposal(db database.Database) error {
	err := db.GetDb().Sync2(new(p.Proposal))

	if err != nil {
		return err
	} else {
		fmt.Printf("Se ha migrado Proposal con exito\n")
	}

	return nil
}

func MigrateOption(db database.Database) error {
	err := db.GetDb().Sync2(new(o.Option))

	if err != nil {
		return err
	} else {
		fmt.Printf("Se ha migrado Option con exito\n")
	}
	return nil
}

func MigrateRoomSetting(db database.Database) error {
	err := db.GetDb().Sync2(new(s.SettingRoom))

	if err != nil {
		return err
	} else {
		fmt.Printf("Se ha migrado RoomSetting con exito\n")
	}

	return nil
}

func MigrateVote(db database.Database) error {
	err := db.GetDb().Sync2(new(e.Vote))

	if err != nil {
		return err
	} else {
		fmt.Printf("Se ha migrado Election con exito\n")
	}

	return nil
}

func MakeConstraints(db database.Database) error {

	_, err := db.GetDb().Exec("ALTER TABLE user_room ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE user_room ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) success\n")
	}

	_, err = db.GetDb().Exec("ALTER TABLE user_room ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES room(id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE user_room ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES room(id) success\n")
	}

	_, err = db.GetDb().Exec("ALTER TABLE room ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE room ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)")
	}

	_, err = db.GetDb().Exec("ALTER TABLE proposal ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES room(id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE proposal ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES room(id) success\n")
	}

	_, err = db.GetDb().Exec("ALTER TABLE setting_room ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES room(id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE setting_room ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES room(id) success\n")
	}

	_, err = db.GetDb().Exec("ALTER TABLE option ADD CONSTRAINT fk_proposal FOREIGN KEY (proposal_id) REFERENCES proposal(id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE option ADD CONSTRAINT fk_proposal FOREIGN KEY (proposal_id) REFERENCES proposal(id) success\n")
	}

	_, err = db.GetDb().Exec("ALTER TABLE vote ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE vote ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) success\n")
	}

	_, err = db.GetDb().Exec("ALTER TABLE vote ADD CONSTRAINT fk_option FOREIGN KEY(option_id) REFERENCES option(id)")
	if err != nil {
		return err
	} else {
		fmt.Printf("ALTER TABLE vote ADD CONSTRAINT fk_option FOREIGN KEY(option_id) REFERENCES option(id) success\n")
	}

	return nil
}
