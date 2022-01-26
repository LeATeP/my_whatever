package creating_module

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)


func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// msg := fmt.Sprintf("Hi %s welcome!", name) // %v
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func Make_Hellos(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("Empty list")
	}	

	// var greetings = map[string]string {}
	greetings := make(map[string]string)

	for _, name := range names {
		greetings[name], _ = Hello(name)
	}
	fmt.Println(greetings)
	return greetings, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	greeting := []string {
		"hi %v. Welcome!",
		"great to see you %v!",
		"hail %v well met",
	}
	return greeting[rand.Intn(len(greeting))]
}

func randInt() int {
	intset := []int {
		1,2,3,
	}
	return intset[rand.Intn(len(intset))]
}
