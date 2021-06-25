package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func NewService(swaggerPath string) *Service {
	return &Service{SwaggerYamlFile: swaggerPath}
}

// Service OpenAPI representation
type Service struct {
	Openapi string `json:"openapi"`

	Info struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Version     string `json:"version"`
	} `json:"info"`

	Servers []struct {
		Url string `json:"url"`
	} `json:"servers"`

	Components struct {
	} `json:"components"`

	Router *mux.Router

	// Paths
	Paths map[string]map[string]*PathMethod `json:"paths"`

	// Swagger definition YAML file path
	SwaggerYamlFile string
}

// Read YAML definition by file name
func (s *Service) Read(filename string) *Service {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("file.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(file, s)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return s
}

// Start service
func (s *Service) Start() *Service {

	// Check and load swagger service YAML definition
	if s.SwaggerYamlFile == "" {
		s.SwaggerYamlFile = "swagger.yml"
	}
	s.Read(s.SwaggerYamlFile)

	// Set default routing
	//s.Router = NewRouter()
	s.Router = mux.NewRouter().StrictSlash(true)

	// Regular expression to remove anything but chars
	var re = regexp.MustCompile("[^a-zA-Z]+")

	// Initialize request handlers
	for pattern, path := range s.Paths {
		for methodName, method := range path {
			var funcName = strings.Title(methodName) + strings.Title(re.ReplaceAllString(pattern, ""))
			var handlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var headers = w.Header()
				for k, v := range map[string]string{
					"Content-Type":  "application/json; charset=UTF-8",
					"Cache-Control": "no-store, max-age=0",
				} {
					headers.Set(k, v)
				}

				if methods[funcName] != nil {

					var marshal []byte

					switch res := methods[funcName](method, r).(type) {
					case error:
						w.WriteHeader(http.StatusBadRequest)
						marshal, _ = json.Marshal(res.Error())
					default:
						w.WriteHeader(http.StatusOK)
						marshal, _ = json.Marshal(res)
					}
					_, err := w.Write(marshal)

					if err != nil {
						panic(err)
					}
				}

			})
			s.Router.
				Methods(strings.ToUpper(methodName)).
				Path(pattern).
				Name(funcName).
				Handler(Logger(handlerFunc, funcName))

		}
	}

	log.Printf("Server started")
	var err = http.ListenAndServe(":8080", s.Router)

	if err != nil {
		log.Fatal(err)
	}

	return s
}

type PathMethod struct {
	Description string          `json:"description"`
	Responses   interface{}     `json:"responses"`
	Parameters  []PathParameter `json:"parameters"`
}

type PathParameter struct {
	In       string `json:"in"`
	Name     string `json:"name"`
	Style    string `json:"style"`
	Explode  bool   `json:"explode"`
	Required bool   `json:"required"`
	Schema   struct {
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"schema"`
}
