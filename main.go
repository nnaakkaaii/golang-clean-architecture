package main

import (
	"clean_architecture/app/adapters/controller"
	"clean_architecture/app/adapters/presenter"
	"clean_architecture/app/adapters/repository"
	"clean_architecture/app/interfaces/database"
	"clean_architecture/app/interfaces/server"
	"clean_architecture/app/interfaces/view"
	"clean_architecture/app/usecases/interactor"
	"gopkg.in/yaml.v3"
	"os"
)

var configPath = "config/config.yaml"

type Conf struct {
	DBUser string `yaml:"db_user"`
	DBPass string `yaml:"db_pass"`
	DBHost string `yaml:"db_host"`
	DBPort string `yaml:"db_port"`
	DBName string `yaml:"db_name"`
	Port   int    `yaml:"port"`
}

func NewConf(path string) (conf *Conf) {
	conf = &Conf{}
	file, err := os.Open(path)

	if err != nil {
		panic("設定ファイルが存在しません")
	}

	if err := yaml.NewDecoder(file).Decode(&conf); err != nil {
		panic("設定ファイルが読み込めません")
	}
	return
}

func main() {
	conf := NewConf(configPath)

	// interfaces/database
	d := repository.UserDatabase(database.NewUserDatabase(conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName))
	// interfaces/view
	v := presenter.UserView(view.NewView())
	// adapters/repository
	r := interactor.UserRepository(repository.NewUserRepository(&d))
	// adapters/presenter
	p := interactor.UserPresenter(presenter.NewUserPresenter(&v))
	// usecases/interactor
	i := interactor.NewUserInteractor(&r, &p)
	// adapters/controller
	c := controller.NewUserController(i)
	// interfaces/server
	s := server.NewUserServer(c)

	s.Start(conf.Port)
}
