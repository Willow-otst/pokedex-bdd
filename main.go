package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type cliCommand struct {
    name        string
    description string
    callback    func() error
}

var commandRegistry map[string]cliCommand

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")

    for _, cmd := range commandRegistry {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }

    return nil
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

func main() {
    commandRegistry = map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
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

        if err := cmd.callback(); err != nil {
            fmt.Println(err)
        }
    }
}
