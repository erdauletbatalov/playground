package main

import (
	"encoding/json"
	"net/http"

	"github.com/chilts/sid"
	"github.com/tarantool/go-tarantool"
)

// main is the entry point of the application. It connects to a Tarantool database,
// creates a space named "geo" with a primary index on the "id" field and a
// geospatial index on the "coordinates" field. It then inserts a sample GeoObject
// into the "geo" space.
// main is the entry point of the application. It connects to a Tarantool database,
// creates a space named "geo" with a primary index on the "id" field and a
// geospatial index on the "coordinates" field. It then inserts a sample GeoObject
// into the "geo" space.
func main() {
	// https://habr.com/en/companies/vk/articles/574542/
	opts := tarantool.Opts{User: "storage", Pass: "password"}
	conn, err := tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create space "geo"

	// Схема данных
	//
	// На одном узле Tarantool находится только одна база данных. Данные
	// складываются в спейсы == таблицы в мире SQL. К данным обязательно строится
	// первичный индекс, а количество вторичных произвольно.

	// Для хранения маркеров сделаем таблицу:

	// id	     coordinates				comment
	// string	 [double, double]	string
	_, err = conn.Call("box.schema.space.create", []interface{}{
		"geo",
		map[string]bool{"if_not_exists": true},
	})

	_, err = conn.Call("box.space.geo:format", [][]map[string]string{
		{
			{"name": "id", "type": "string"},
			{"name": "coordinates", "type": "array"},
			{"name": "comment", "type": "string"},
		},
	})

	// Индексация
	//
	// Для работы со спейсом необходимо создать первичный ключ. Иначе любое
	// действие с данными в спейсе будет создавать ошибку.
	_, err = conn.Call("box.space.geo:create_index", []interface{}{
		"primary",
		map[string]interface{}{
			"parts":         []string{"id"},
			"if_not_exists": true,
		},
	})

	// 	Геоиндекс

	// Для поиска объектов понадобится геоиндекс, который сможет быстро возвращать
	// данные, которые расположены в некотором регионе.

	// Параметры:

	// имя;
	// поле для индекса;
	// тип индекса RTREE;
	// индекс может содержать неуникальные координаты;
	// флаг для игнорирования ошибки при существующем индексе.
	_, err = conn.Call("box.space.geo:create_index", []interface{}{
		"geoidx",
		map[string]interface{}{
			"parts":         []string{"coordinates"},
			"type":          "RTREE",
			"unique":        false,
			"if_not_exists": "true",
		},
	})

	// Reload, чтобы новая схема загрузилась в коннектор
	conn, err = tarantool.Connect("127.0.0.1:3301", opts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Запись данных

	// Для вставки данных я воспользуюсь функцией Golang-коннектора InsertTyped.

	// InsertTyped позволяет вставлять только новые данные и возвращает ошибку,
	// если данные уже существовали.

	// Параметры:

	// имя спейса; данные; переменная для вставленных таплов.

	// Например, здесь я вставляю тапл {"Indisko", {300.073, 148.857}, "Indian
	// Food" } в спейс geo.
	var tuples []GeoObject
	err = conn.InsertTyped("geo", []interface{}{
		"Indisko",
		[]float64{299.073, 148.857},
		"Indian Food"},
		&tuples)

	// var tuples []GeoObject
	err = conn.DeleteTyped("geo", "primary", []interface{}{"Indisko"}, &tuples)

	// Сигнатура select

	// Для запроса данных используем функцию SelectTyped.

	// Параметры:

	// cпейс;
	// индекс;
	// смещение, лучше указывать 0;
	// максимум сколько можно отдавать объектов;
	// направление поиска по индексу;
	// значение индекса для поиска. Для индексов, состоящих из нескольких полей, можно указывать часть значения, начиная с самой старшей позиции;
	// параметр для возврата сериализованных данных.

	// В этим примере я выполняю поиск данных в спейсе geo по индексу geoidx. И ищу
	// только те данные, которые входят (tarantool.IterLe) в заданный регион поиска
	// {0, 0, 300, 400}. Tarantool вернет мне данные, координаты которых лежат в
	// квадрате от вершины 0,0 до вершины 300,400.

	// var tuples []GeoObject
	err = conn.SelectTyped("geo", "geoidx", 0, 10, tarantool.IterLe,
		[]interface{}{0, 0, 300, 400},
		&tuples,
	)

	// В корневом эндпоинте отдаём пользователю фронтенд
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Отдаём маркеры для указанного в url региона
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		rect, ok := r.URL.Query()["rect"]
		if !ok || len(rect) < 1 {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var arr []float64
		err := json.Unmarshal([]byte(rect[0]), &arr)
		if err != nil {
			panic(err)
		}

		// Запрашивает 1000 маркеров, которые находятся в регионе rect
		var tuples []GeoObject
		err = conn.SelectTyped("geo", "geoidx", 0, 1000, tarantool.IterLe,
			arr,
			&tuples)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(tuples)
	})

	// Эндпоинт для сохранения маркера
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		obj := &GeoObject{}
		err := dec.Decode(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Генерируем уникальный идентификатор маркера
		obj.Id = sid.IdHex()
		var tuples []GeoObject
		// Вставляем новый маркер
		err = conn.InsertTyped("geo", []interface{}{obj.Id, obj.Coordinates, obj.Comment}, &tuples)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(tuples)
	})

	// Эндпоинт для удаления маркера
	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		obj := &GeoObject{}
		err := dec.Decode(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Удаляем переданный маркер по его первичному ключу
		var tuples []GeoObject
		err = conn.DeleteTyped("geo", "primary", []interface{}{obj.Id}, &tuples)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		enc.Encode(tuples)
	})

	// Запускаем http сервер на локальном адресе
	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}

}

type GeoObject struct {
	Id          string     `json:"id"`
	Coordinates [2]float64 `json:"coordinates"`
	Comment     string     `json:"comment"`
}
