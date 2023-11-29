package model

type CmdArgs struct {
	Email     *string
	Password  *string
	SlackHook *string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Nick           string `json:"nick"`
	Created        string `json:"created"`
	IsProUser      string `json:"isProUser"`
	Type           string `json:"type"`
	ConsumedTrial  string `json:"consumedTrial"`
	IsVerified     string `json:"isVerified"`
	Url            string `json:"url"`
	Id             string `json:"id"`
	CountryCode    string `json:"countryCode"`
	IsBanned       bool   `json:"isBanned"`
	ChatBan        bool   `json:"chatBan"`
	IsBotUser      bool   `json:"isBotUser"`
	SuspendedUntil string `json:"suspendedUntil"`
	IsCreator      bool   `json:"isCreator"`
}

type CreateChallengeRequest struct {
	ForbidMoving   bool   `json:"forbidMoving"`
	ForbidRotating bool   `json:"forbidRotating"`
	ForbidZooming  bool   `json:"forbidZooming"`
	TimeLimit      int32  `json:"timeLimit"`
	Map            string `json:"map"`
	StreakType     string `json:"streakType"`
}

type CreateChallengeResponse struct {
	Token string `json:"token"`
}

type MapConfig struct {
	ForbidMoving   bool
	ForbidRotating bool
	ForbidZooming  bool
	TimeLimit      int32
	MapId          string
	Type           string
	StreakType     string
}

type SlackMessageRequest struct {
	Text string `json:"text"`
}
