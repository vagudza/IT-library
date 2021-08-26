// имя папки должно совпадать с именем пакета, используемым в файлах !
package mydatabase

import (
	"database/sql"
	"fmt"

	// for init() of sql driver is need to use _ symbol
	_ "github.com/go-sql-driver/mysql"
)

var dbDriverName string = "mysql"
var dbUser string = "mysql"
var dbPassword string = "mysql"
var dbIp string = "127.0.0.1"
var dbPort string = "3306"
var dbName string = "vitalg"
var sqlString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbIp, dbPort, dbName)

type Project struct {
	Id          int
	Title       string
	Thumbnail   string
	Description string
	Readtime    int
}

// Выбор из БД описания проекта с заданным id
func DBGetProjectByID(projectID string) Project {
	db, err := sql.Open(dbDriverName, sqlString)
	if err != nil {
		panic(err)
	}

	// defer - отложенная функция
	// гарантия, что подключение к базе данных будет закрыто в любом случае вне зависимости от точек выхода из функции
	defer db.Close()
	fmt.Println("Connected to DB - DBGetProjectByID")

	// выборка трех последних записей
	rows, err := db.Query(fmt.Sprintf("SELECT `id`, `title`, `thumbnail`, `description`, `readtime` FROM projects WHERE `id`= %s", projectID))
	if err != nil {
		panic(err)
	}

	var p Project
	for rows.Next() {
		// Scan проверяет, существует ли значение. Заполняем
		err := rows.Scan(&p.Id, &p.Title, &p.Thumbnail, &p.Description, &p.Readtime)
		if err != nil {
			panic(err)
		}
	}
	return p
}

// Выбор из БД трех крайних проектов для отображения на главной странице
func DBGetProjects() []Project {
	db, err := sql.Open(dbDriverName, sqlString)
	if err != nil {
		panic(err)
	}

	// defer - отложенная функция
	// гарантия, что подключение к базе данных будет закрыто в любом случае вне зависимости от точек выхода из функции
	defer db.Close()
	fmt.Println("Connected to DB - DBGetProjects")

	// выборка трех последних записей
	rows, err := db.Query("SELECT `id`, `title`, `thumbnail`, `description`, `readtime` FROM projects ORDER BY id DESC LIMIT 3")
	if err != nil {
		panic(err)
	}
	//defer rows.Close()
	var projects []Project

	for rows.Next() {
		var p Project
		// Scan проверяет, существует ли значение. Заполняем
		err := rows.Scan(&p.Id, &p.Title, &p.Thumbnail, &p.Description, &p.Readtime)
		if err != nil {
			panic(err)
		}
		projects = append(projects, p)
	}

	/*
		for _, p := range projects {
			fmt.Println(p.Id, p.Thumbnail, p.Description, p.Readtime)
		}
	*/
	return projects
}
