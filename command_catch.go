package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
    time.Sleep(1 * time.Second)
    catchResult := rand.Intn(pokemon.BaseExperience)

	if catchResult > 40 {
		fmt.Printf("shake...\n")
        time.Sleep(1 * time.Second)
		fmt.Printf("shake...\n")
        time.Sleep(1 * time.Second)
        fmt.Printf("shake...\n")
        time.Sleep(1 * time.Second)

		fmt.Printf("Oh no, the pokemon broke free!\n")
        return nil

	} else {
		fmt.Printf("shake...\n")
        time.Sleep(1 * time.Second)
		fmt.Printf("shake...\n")
        time.Sleep(1 * time.Second)
		fmt.Printf("shake...\n")
        time.Sleep(1 * time.Second)
		fmt.Printf("%s was caught!\n", pokemon.Name)
	}

    cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}


