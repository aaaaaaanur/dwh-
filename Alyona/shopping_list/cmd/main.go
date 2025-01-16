package main

import (
	"awesomeProject/Alyona/shopping_list/domain"
	"awesomeProject/Alyona/shopping_list/repository"
	"awesomeProject/Alyona/shopping_list/ui"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	sh, err := sql.Open("postgres", "postgres://user:secret@localhost:5432/shopping_list?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer sh.Close()

	GoodRepo := repository.GoodRepository{DB: sh}
	ListRepo := repository.ListRepository{DB: sh}

	for {
		ui.MenuMessage()

		text, err := ui.ReadInt()
		if err != nil {
			log.Fatal(err)
		}

		switch text {
		case 1:
			lists, err := ListRepo.FindAll()
			if err != nil {
				log.Fatal(err)
			}

			for _, list := range lists {
				fmt.Printf("%d, %s\n", list.ID, list.Name)
			}

			ui.BackExitMessage()

			back, err := ui.ReadString()
			if err != nil {
				log.Fatal(err)
			}

			if back == "b" {
				continue
			}
			if back == "e" {
				Exit()
			}

		case 2:
			lists, err := ListRepo.FindAll()
			if err != nil {
				log.Fatal(err)
			}

			ui.ListMessage(lists)

			list, err := ui.ReadUint()
			if err != nil {
				log.Fatal(err)
			}

			if list == 0 {
				continue
			}
			if list == 100 {
				Exit()
			}

			goods, err := GoodRepo.FindList(list)
			if err != nil {
				log.Fatal(err)
			}

			for _, good := range goods {
				fmt.Printf("%d, %s\n", good.ID, good.Name)
			}

			ui.BackExitMessage()

			back, err := ui.ReadString()
			if err != nil {
				log.Fatal(err)
			}

			if back == "b" {
				continue
			}
			if back == "e" {
				Exit()
			}

		case 3:
			list := domain.List{}

			ui.CreateListMessage()

			list.Name, err = ui.ReadString()
			if err != nil {
				log.Fatal(err)
			}

			if list.Name == "b" {
				continue
			}
			if list.Name == "e" {
				Exit()
			}

			err = ListRepo.CreateList(list)
			if err != nil {
				log.Fatal(err)
			}

			ui.ListCreatedMessage()
			continue

		case 4:
			lists, err := ListRepo.FindAll()
			if err != nil {
				log.Fatal(err)
			}

			ui.ListMessage(lists)

			good := domain.Good{}

			good.ListID, err = ui.ReadUint()
			if err != nil {
				log.Fatal(err)
			}

			if good.ListID == 0 {
				continue
			}
			if good.ListID == 100 {
				Exit()
			}

			ui.CreateGoodMessage()

			good.Name, err = ui.ReadString()
			if err != nil {
				log.Fatal(err)
			}

			if good.Name == "b" {
				continue
			}
			if good.Name == "e" {
				Exit()
			}

			err = GoodRepo.CreateGood(good)
			if err != nil {
				log.Fatal(err)
			}

			ui.GoodsAddedMessage()
			continue

		case 5:
			lists, err := ListRepo.FindAll()
			if err != nil {
				log.Fatal(err)
			}

			ui.ListMessage(lists)

			list := domain.List{}

			list.ID, err = ui.ReadUint()
			if err != nil {
				log.Fatal(err)
			}

			if list.ID == 0 {
				continue
			}
			if list.ID == 100 {
				Exit()
			}

			err = ListRepo.DeleteList(list.ID)
			if err != nil {
				log.Fatal(err)
			}

			ui.ListDeletedMessage()
			continue

		case 6:
			Exit()
		}
	}
}
