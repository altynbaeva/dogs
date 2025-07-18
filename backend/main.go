package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type (
	Dog struct {
		Id           int        `json:"id"`
		Name         string     `json:"name"`
		Gender       string     `json:"gender"`
		Age          int        `json:"age"`
		IsVaccinated bool       `json:"isVaccinated"`
		CreatedAt    time.Time  `json:"createdAt"`
		UpdatedAt    *time.Time `json:"updatedAt"`
	}

	DogFull struct {
		Dog
		PhotoURL    string         `json:"photoURL"`
		Personality DogPersonality `json:"personality"`
	}

	DogPersonality struct {
		FavoriteFood []string `json:"favoriteFood"`
	}
)

type (
	CreateUpdateDogRequest struct {
		Name         string         `json:"name"`
		Gender       string         `json:"gender"`
		Age          int            `json:"age"`
		IsVaccinated bool           `json:"isVaccinated"`
		PhotoURL     string         `json:"photoURL"`
		Personality  DogPersonality `json:"personality"`
	}
)

var dogs []DogFull
var nextDogId = 3

func main() {
	fmt.Println("preparing data")
	prepareData()

	fmt.Println("configuring endpoints")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/dogs", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		switch request.Method {
		case http.MethodGet:
			handleListDogsRequest(writer, request)
		case http.MethodOptions:
			handleOptionsRequest(writer, request)
		case http.MethodPost:
			handleCreateDogRequest(writer, request)
		default:
			writer.WriteHeader(http.StatusBadRequest)
			_, err := writer.Write([]byte("эта ручка не поддерживает этот метод"))
			if err != nil {
				fmt.Println("не удалось записать ответ: " + err.Error())
			}
		}
	})
	mux.HandleFunc("/api/v1/dogs/{id}", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		switch request.Method {
		case http.MethodGet:
			handleGetDogRequest(writer, request)
		case http.MethodPut:
			handleUpdateDogRequest(writer, request)
		case http.MethodOptions:
			handleOptionsRequest(writer, request)
		case http.MethodDelete:
			handleDeleteDogRequest(writer, request)
		default:
			writer.WriteHeader(http.StatusBadRequest)
			_, err := writer.Write([]byte("эта ручка не поддерживает этот метод"))
			if err != nil {
				fmt.Println("не удалось записать ответ: " + err.Error())
			}
		}
	})
	fmt.Println("starting server on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func prepareData() {
	updAt := time.Now().Add(-4 * time.Hour)
	dogs = append(dogs,
		DogFull{
			Dog: Dog{
				Id:           1,
				Name:         "Бобик",
				Gender:       "M",
				Age:          2,
				IsVaccinated: true,
				CreatedAt:    time.Now(),
				UpdatedAt:    nil,
			},
			PhotoURL: "https://images.unsplash.com/photo-1598133894008-61f7fdb8cc3a?q=80&w=988&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			Personality: DogPersonality{
				FavoriteFood: []string{"Сосиски", "Мороженное", "Шашлык"},
			},
		},
		DogFull{
			Dog: Dog{
				Id:           2,
				Name:         "Зина",
				Gender:       "F",
				Age:          1,
				IsVaccinated: false,
				CreatedAt:    time.Now().Add(-1 * 87 * time.Hour),
				UpdatedAt:    &updAt,
			},
			PhotoURL: "https://images.unsplash.com/photo-1537151608828-ea2b11777ee8?q=80&w=994&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
			Personality: DogPersonality{
				FavoriteFood: []string{"Чапи сырная курица", "Паста дракоша", "Двойная порция пиццы"},
			},
		})
}

func handleListDogsRequest(writer http.ResponseWriter, request *http.Request) {
	result := []Dog{}
	for _, d := range dogs {
		result = append(result, Dog{
			Id:           d.Id,
			Name:         d.Name,
			Gender:       d.Gender,
			Age:          d.Age,
			IsVaccinated: d.IsVaccinated,
			CreatedAt:    d.CreatedAt,
			UpdatedAt:    d.UpdatedAt,
		})
	}

	resp, _ := json.Marshal(result)
	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	_, err := writer.Write(resp)
	if err != nil {
		fmt.Println("не удалось записать ответ: " + err.Error())
	}
}

func handleCreateDogRequest(writer http.ResponseWriter, request *http.Request) {
	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(400)
		_, err = writer.Write([]byte("не удалось прочитать тело запроса: " + err.Error()))
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
		return
	}

	req := CreateUpdateDogRequest{}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		writer.WriteHeader(400)
		_, err = writer.Write([]byte("в теле запроса некорректный json: " + err.Error()))
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
		return
	}

	newDog := DogFull{
		Dog: Dog{
			Id:           nextDogId,
			Name:         req.Name,
			Gender:       req.Gender,
			Age:          req.Age,
			IsVaccinated: req.IsVaccinated,
			CreatedAt:    time.Now(),
		},
		PhotoURL:    req.PhotoURL,
		Personality: req.Personality,
	}
	dogs = append(dogs, newDog)
	nextDogId++

	respBytes, _ := json.Marshal(newDog)
	writer.WriteHeader(200)
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(respBytes)
	if err != nil {
		fmt.Println("не удалось записать ответ: " + err.Error())
	}
	return
}

func handleDeleteDogRequest(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	if id == "" {
		writer.WriteHeader(400)
		_, err := writer.Write([]byte("не указан id"))
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
		return
	}

	for i := range dogs {
		if strconv.Itoa(dogs[i].Id) == id {
			dogs = append(dogs[:i], dogs[i+1:]...)
			writer.WriteHeader(200)
			writer.Header().Set("Content-Type", "application/json")
			return
		}
	}

	writer.WriteHeader(404)
	return
}

func handleGetDogRequest(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	if id == "" {
		writer.WriteHeader(400)
		_, err := writer.Write([]byte("не указан id"))
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
	}

	for _, dog := range dogs {
		if strconv.Itoa(dog.Id) != id {
			continue
		}
		respBytes, _ := json.Marshal(dog)
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		_, err := writer.Write(respBytes)
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
		return
	}

	writer.WriteHeader(404)
	return
}

func handleUpdateDogRequest(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	if id == "" {
		writer.WriteHeader(400)
		_, err := writer.Write([]byte("не указан id"))
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
	}

	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(400)
		_, err = writer.Write([]byte("не удалось прочитать тело запроса: " + err.Error()))
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
		return
	}

	req := CreateUpdateDogRequest{}
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		writer.WriteHeader(400)
		_, err = writer.Write([]byte("в теле запроса некорректный json: " + err.Error()))
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
		return
	}

	for i := range dogs {
		if strconv.Itoa(dogs[i].Id) != id {
			continue
		}
		now := time.Now()
		dogs[i] = DogFull{
			Dog: Dog{
				Id:           dogs[i].Id,
				Name:         req.Name,
				Gender:       req.Gender,
				Age:          req.Age,
				IsVaccinated: req.IsVaccinated,
				CreatedAt:    dogs[i].CreatedAt,
				UpdatedAt:    &now,
			},
			PhotoURL:    req.PhotoURL,
			Personality: req.Personality,
		}
		respBytes, _ := json.Marshal(dogs[i])
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "application/json")
		_, err := writer.Write(respBytes)
		if err != nil {
			fmt.Println("не удалось записать ответ: " + err.Error())
		}
		return
	}

	writer.WriteHeader(404)
	return
}

func handleOptionsRequest(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	return
}
