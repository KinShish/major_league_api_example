package test

import (
	_ "major_league_api_example/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	// нужно добавить сессии если используем
	_ "github.com/beego/beego/v2/server/web/session/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	// Обязательно нужно добавить абсолютный путь до проекта
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	// Обязательно нужно добавить подключение к базе
	sqlconn, _ := beego.AppConfig.String("sqlconn")
	orm.RegisterDataBase("test", "postgres", sqlconn)
}

// TestGetUser Провекра гет запроса
func TestGetUser(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/user/2", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\n  \"err\": false,\n  \"data\": {\n    \"Id\": 2,\n    \"Username\": \"Евгений\",\n    \"Password\": \"12345678\"\n  }\n}", w.Body.String())
}
