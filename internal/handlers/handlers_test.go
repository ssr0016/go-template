package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ssr0016/booking/internal/models"
)

type postData struct {
	key   string
	value string
}

var theTest = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	// {
	// 	"home",
	// 	"/",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"about",
	// 	"/about",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"generals",
	// 	"/generals-quarters",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"majors",
	// 	"/majors-suite",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"search-availability",
	// 	"/search-availability",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"contact",
	// 	"/contact",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"make reservation",
	// 	"/make-reservation",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"make reservation",
	// 	"/reservation-summary",
	// 	"GET",
	// 	[]postData{},
	// 	http.StatusOK,
	// },
	// {
	// 	"post-search-availability",
	// 	"/search-availability",
	// 	"POST",
	// 	[]postData{
	// 		{
	// 			key:   "start",
	// 			value: "2020-01-01",
	// 		},
	// 		{
	// 			key:   "end",
	// 			value: "2020-01-02",
	// 		},
	// 	},
	// 	http.StatusOK,
	// },
	// {
	// 	"post-search-availability-json",
	// 	"/make-reservation",
	// 	"POST",
	// 	[]postData{
	// 		{
	// 			key:   "first_name",
	// 			value: "John",
	// 		},
	// 		{
	// 			key:   "last_name",
	// 			value: "Doe",
	// 		},
	// 		{
	// 			key:   "email",
	// 			value: "example@example.com",
	// 		},
	// 		{
	// 			key:   "phone",
	// 			value: "99999",
	// 		},
	// 	},
	// 	http.StatusOK,
	// },
	// {
	// 	"post-search-availability-json",
	// 	"/search-availability-json",
	// 	"POST",
	// 	[]postData{
	// 		{
	// 			key:   "start",
	// 			value: "2020-01-01",
	// 		},
	// 		{
	// 			key:   "end",
	// 			value: "2020-01-02",
	// 		},
	// 	},
	// 	http.StatusOK,
	// },
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTest {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "Generals Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %d want %d", rr.Code, http.StatusOK)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
