package genre

type GenreFormatter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func FormatGenre(genre Genre) GenreFormatter {
	formatter := GenreFormatter{
		ID:   genre.ID.String(),
		Name: genre.Name,
	}

	return formatter
}
