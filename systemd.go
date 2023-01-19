package main

import (
	"context"
	"log"
	"strings"

	"github.com/coreos/go-systemd/v22/dbus"
)

type ServiceStatus int

type Service struct {
	Name        string `json:"name"`
	ActiveState string `json:"activestate"`
	SubState    string `json:"substate"`
}

func QueryService(servicename string) Service {
	conn, err := dbus.NewSystemConnectionContext(context.TODO())

	if err != nil {
		log.Default().Println("Can't connect to systemd via D-BUS")
		return Service{}
	}

	defer conn.Close()

	activestate, err := conn.GetUnitPropertyContext(context.TODO(), servicename, "ActiveState")
	if err != nil {
		log.Default().Println(err)
		return Service{}
	}

	substate, err := conn.GetUnitPropertyContext(context.TODO(), servicename, "SubState")
	if err != nil {
		log.Default().Println(err)
		return Service{}
	}

	return Service{
		Name:        servicename,
		ActiveState: strings.Trim(activestate.Value.String(), "\""),
		SubState:    strings.Trim(substate.Value.String(), "\""),
	}
}
