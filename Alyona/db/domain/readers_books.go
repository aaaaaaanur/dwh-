package domain

type ReadersBooks struct {
	ID             uint
	TitleBook      string
	FullNameReader string
	TakenAt        string
	ReturnedAt     *string
}
