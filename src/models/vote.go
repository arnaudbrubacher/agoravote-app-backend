package models

type Vote struct {
    ID          string `gorm:"primary_key"`
    Question    string
    Choices     []string `gorm:"type:text[]"`
    Status      string
    TotalVotes  int
    WriteInResponses []WriteInResponse `gorm:"foreignKey:VoteID"`
}

type WriteInResponse struct {
    ID     string `gorm:"primary_key"`
    VoteID string
    Text   string
    Count  int
}