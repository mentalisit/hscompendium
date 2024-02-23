package models

type Identity struct {
	User  User   `json:"user"`
	Guild Guild  `json:"guild"`
	Token string `json:"token"`
}

// структура для хранения в будущем игрока и твина и так же нескольких корпораций
type IdentityGET struct {
	User  []User  `json:"user"`
	Guild []Guild `json:"guilds"`
	Token string  `json:"token"`
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	AvatarURL     string `json:"avatarUrl"`
}

type Guild struct {
	URL  string `json:"url"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
type TechLevel struct {
	Level int   `json:"level"`
	Ts    int64 `json:"ts"`
}
type SyncData struct {
	Token      string
	Ver        int
	InSync     int
	TechLevels map[int]TechLevel
}
