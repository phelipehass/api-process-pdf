package models

type Diary struct {
	ID         int64  `db:"id"`
	UrlArchive string `db:"url_archive"`
	Cod        int64  `db:"cod"`
}
