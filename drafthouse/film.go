package drafthouse

type Film struct {
	FilmID        string   `json:"FilmId"`
	FilmName      string   `json:"FilmName"`
	FilmYear      string   `json:"FilmYear"`
	FilmRating    string   `json:"FilmRating"`
	FilmRuntime   string   `json:"FilmRuntime"`
	FilmAgePolicy string   `json:"FilmAgePolicy"`
	FilmSlug      string   `json:"FilmSlug"`
	Series        []Series `json:"Series"`
}

type Series struct {
	SeriesID   string   `json:"SeriesId"`
	SeriesName string   `json:"SeriesName"`
	Formats    []Format `json:"Formats"`
}

type Format struct {
	FormatID   string    `json:"FormatId"`
	FormatName string    `json:"FormatName"`
	Sessions   []Session `json:"Sessions"`
}

type Session struct {
	SessionID       string `json:"SessionId"`
	SessionTime     string `json:"SessionTime"`
	SessionStatus   string `json:"SessionStatus"`
	SessionSalesURL string `json:"SessionSalesURL"`
	SessionDateTime string `json:"SessionDateTime"`
	SessionType     string `json:"SessionType"`
	SeatsLeft       string `json:"SeatsLeft"`
	SeatingLow      string `json:"SeatingLow"`
}

func (f *Film) GetFilmTimes() []string {
	var filmTimes []string

	for _, series := range f.Series {
		for _, format := range series.Formats {
			for _, session := range format.Sessions {
				filmTimes = append(filmTimes, session.SessionTime)
			}
		}
	}
	return filmTimes
}
