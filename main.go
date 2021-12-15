package main

import (
	"clean_architecture/app/controller"
	"clean_architecture/app/database"
	"clean_architecture/app/interactor"
	"clean_architecture/app/presenter"
	"clean_architecture/app/repository"
	"clean_architecture/app/server"
	"clean_architecture/app/view"
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

	// database
	d := repository.UserDatabaseInterface(database.NewUserDatabase(conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName))
	// view
	v := presenter.UserViewInterface(view.NewUserView())
	// repository
	r := interactor.UserRepositoryInterface(repository.NewUserRepository(&d))
	// presenter
	p := controller.UserPresenterInterface(presenter.NewUserPresenter(&v))
	// interactor
	i := interactor.UserInteractorInterface(interactor.NewUserInteractor(&r))
	// controller
	c := controller.UserControllerInterface(controller.NewUserController(&i, &p))
	// interfaces/server
	s := server.NewUserServer(&c)

	s.Start(conf.Port)
}
