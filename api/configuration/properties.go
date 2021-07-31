package configuration

import (
	"br.com.github/raphasalomao/go-marvel-heroes-api/api/client"
	"br.com.github/raphasalomao/go-marvel-heroes-api/api/database"
	"github.com/magiconair/properties"
)

func InitProperties() {
	p := properties.MustLoadFile("application.properties", properties.UTF8)
	
	//database properties
	database.Dbname, _ = p.Get("gorm-io.db.dbname")
	database.Host, _ = p.Get("gorm-io.db.host")
	database.User, _ = p.Get("gorm-io.db.user")
	database.Password, _ = p.Get("gorm-io.db.password")
	database.Port, _ = p.Get("gorm-io.db.port")
	database.Sslmode, _ = p.Get("gorm-io.db.sslmode")

	client.PublicKey, _ = p.Get("marvel.api.public-key")
	client.PrivateKey, _ = p.Get("marvel.api.private-key")
}
