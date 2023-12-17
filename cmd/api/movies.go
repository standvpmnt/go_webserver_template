package main
// {
// 	var input struct {
// 		Title   string       `json:"title"`
// 		Year    int32        `json:"year"`
// 		Runtime data.Runtime `json:"runtime"`
// 		Genres  []string     `json:"genres"`
// 	}

// 	// err := json.NewDecoder(r.Body).Decode(&input)
// 	err := app.readJSON(w, r, &input)
// 	if err != nil {
// 		// app.errorResponse(w, r, http.StatusBadRequest, err.Error())
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	movie := &data.Movie{
// 		Title:   input.Title,
// 		Year:    input.Year,
// 		Runtime: input.Runtime,
// 		Genres:  input.Genres,
// 	}

// 	v := validator.New()
// 	if data.ValidateMovie(v, movie); !v.Valid() {
// 		app.failedValidationResponse(w, r, v.Errors)
// 		return
// 	}

// 	err = app.models.Movies.Insert(movie)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}

// 	headers := make(http.Header)
// 	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

// 	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 	}
// }
