package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "net/http"
    "io"
    "encoding/json"
)


type cliCommand struct {
    name        string
    description string
    callback    func(*config) error
}
type config struct {
    Next     *string
    Previous *string
}


var commandRegistry map[string]cliCommand

func commandHelp(cfg *config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")

    for _, cmd := range commandRegistry {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

func commandExit(cfg *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}


type LocationAreaResponse struct {
    Count    int `json:"count"`
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}
func commandMap(cfg *config) error {
    url := "https://pokeapi.co/api/v2/location-area?limit=20"
    if cfg.Next != nil {
        url = *cfg.Next
    }

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var data LocationAreaResponse
    if err := json.Unmarshal(body, &data); err != nil {
        return err
    }

    for _, area := range data.Results {
        fmt.Println(area.Name)
    }

    cfg.Next = data.Next
    cfg.Previous = data.Previous

    return nil
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}
func commandMapb(cfg *config) error {
    if cfg.Previous == nil {
        fmt.Println("you're on the first page")
        return nil
    }

    url := *cfg.Previous

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var data LocationAreaResponse
    if err := json.Unmarshal(body, &data); err != nil {
        return err
    }

    for _, area := range data.Results {
        fmt.Println(area.Name)
    }

    cfg.Next = data.Next
    cfg.Previous = data.Previous

    return nil
}


func main() {
    cfg := &config{}

    commandRegistry = map[string]cliCommand{
        "exit": {
			name: 		 "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
			name: 		 "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
			name: 		 "map",
            description: "Displays area locations",
            callback:    commandMap,
        },
        "mapb": {
			name: 		 "mapb",
            description: "Displays previous area locations",
            callback:    commandMapb,
        },
    }

    reader := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        reader.Scan()

        words := cleanInput(reader.Text())
        if len(words) == 0 {
            continue
        }

        commandName := words[0]

        cmd, exists := commandRegistry[commandName]
        if !exists {
            fmt.Println("Unknown command")
            continue
        }

        if err := cmd.callback(cfg); err != nil {
            fmt.Println(err)
        }
    }
}

