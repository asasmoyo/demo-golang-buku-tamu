package httpsrv

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/asasmoyo/demo-golang-buku-tamu/model"
)

func TestListTamu(t *testing.T) {
	srv := Server{
		DBConnStr: "postgres://postgres:password@127.0.0.1:6432/tamudb",
	}
	srv.Init()

	t.Run("empty table", func(t *testing.T) {
		srv.DB.Exec("delete from tamus;")

		req, _ := http.NewRequest(http.MethodGet, "/api/tamu", nil)
		rec := httptest.NewRecorder()

		srv.listTamu(rec, req)

		var results []model.Tamu
		err := json.NewDecoder(rec.Body).Decode(&results)
		if err != nil {
			t.Error(err)
		}

		if len(results) != 0 {
			t.Error("not empty")
		}
	})

	t.Run("one row", func(t *testing.T) {
		srv.DB.Exec("delete from tamus;")
		srv.DB.Create(&model.Tamu{Name: "buditest", Keperluan: "maintest"})

		req, _ := http.NewRequest(http.MethodGet, "/api/tamu", nil)
		rec := httptest.NewRecorder()

		srv.listTamu(rec, req)

		var results []model.Tamu
		err := json.NewDecoder(rec.Body).Decode(&results)
		if err != nil {
			t.Error(err)
		}

		if len(results) != 1 {
			t.Error("not empty")
		}
	})
}

func TestCreateTamu(t *testing.T) {
	srv := Server{
		DBConnStr: "postgres://postgres:password@127.0.0.1:6432/tamudb",
	}
	srv.Init()

	t.Run("valid input", func(t *testing.T) {
		srv.DB.Exec("delete from tamus;")

		param := url.Values{}
		param.Set("name", "buditest")
		param.Set("keperluan", "maintest")

		req, _ := http.NewRequest(http.MethodPost, "/api/tamu", bytes.NewBufferString(param.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()

		srv.createTamu(rec, req)

		var result model.Tamu
		err := json.NewDecoder(rec.Body).Decode(&result)
		if err != nil {
			t.Error(err)
		}

		if result.Name != "buditest" {
			t.Error("unexpected name")
		}
		if result.Keperluan != "maintest" {
			t.Error("unexpected keperluan")
		}
	})
}
