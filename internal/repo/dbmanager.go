package repo

type DBmanager interface {
	GetShortLink(link string, domainId int) (string, string, error)
	CheckShortLink(shortLlink string) (bool, error)
	SetLink(link Link) (int64, string, error)
	GetFullLink(shortLink string) (string, error)
}
