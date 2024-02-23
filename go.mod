module compendium

go 1.22.0

require (
	github.com/bwmarrin/discordgo v0.27.1
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/mentalisit/logger v0.0.0-20240221024243-6f28067f593e
	go.mongodb.org/mongo-driver v1.14.0
	go.uber.org/zap v1.26.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/go-telegram-bot-api/telegram-bot-api => ./pkg/telegram-bot-api/
