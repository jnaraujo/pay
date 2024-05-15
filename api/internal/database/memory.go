package database

type database struct {
}

var DB *database

func Init() {
	DB = &database{}
}

func (db *database) Create() {

}
