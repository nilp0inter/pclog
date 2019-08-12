package main

import "errors"
import "fmt"
import "github.com/manifoldco/promptui"
import "gopkg.in/yaml.v2"
import "log"
import "strconv"

var data = `
name: Convict Conditioning
exercises:
  - name: Squats
    progressions:
      - name: Shoulderstand Squats
        goals:
          - name: Beginner standard
            sets: 1
            reps: 10
          - name: Intermediate standard
            sets: 2
            reps: 25
`

func menu(label string, entry Selectable) (Selectable, error) {
	prompt := promptui.Select{
		Label: label,
		Items: entry.Items(),
	}

	idx, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return entry.GetOption(idx), nil
}

func main() {
	fmt.Println("vim-go")
	plan := Plan{}

	err := yaml.Unmarshal([]byte(data), &plan)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	exercise, err := menu("Select Exercise", plan)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	progression, err := menu("Select Progression", exercise)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	goal, err := menu("Select Goal", progression)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(exercise, progression, goal)

	validate := func(input string) error {
		_, err := strconv.ParseUint(input, 10, 16)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Repetitions",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

	/*
		d, err := yaml.Marshal(&plan)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("--- t dump:\n%s\n\n", string(d))

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)
	*/
}
