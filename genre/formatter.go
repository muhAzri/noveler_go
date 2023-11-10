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

func FormatGenres(genres []Genre) []GenreFormatter {
	genresFormatter := []GenreFormatter{}

	for _, genre := range genres {
		genreFormatter := FormatGenre(genre)
		genresFormatter = append(genresFormatter, genreFormatter)
	}

	return genresFormatter
}
