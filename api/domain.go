package api

import(
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/samuel/go-zookeeper/zk"
	conf "bamboo/configuration"
	service "bamboo/services/domain"
)

type Domain struct {
	Config conf.Configuration
	Zookeeper zk.Conn
}

func (d Domain) All(w http.ResponseWriter, r *http.Request) {
	domains, err := service.All(&d.Zookeeper, d.Config.DomainMapping.Zookeeper)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusBadRequest)
	}
	bites, _ := json.Marshal(domains)
	w.Write(bites)
}

func (d Domain) Create(w http.ResponseWriter, r *http.Request) {

}

func (d Domain) Delete(w http.ResponseWriter, r *http.Request) {

}

func (d Domain) Put(w http.ResponseWriter, r *http.Request) {

}


