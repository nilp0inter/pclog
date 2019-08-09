package main

import "fmt"
import "gopkg.in/yaml.v2"
import "log"

var data = `
name: The Name of the Plan (e.g. Convict Conditioning)
exercises:
  - name: Exercise Name (e.g. Squats)
    progressions:
      - name: Progression Name (e.g. Shoulderstand Squats)
        goals:
          - name: Goal Name (e.g. Beginner standard)
            sets: 1
            reps: 10
          - name: Intermediate standard
            sets: 2
            reps: 25
`

func main() {
	fmt.Println("vim-go")
	plan := Plan{}
	err := yaml.Unmarshal([]byte(data), &plan)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("plan: \n%v\n\n", plan)

	d, err := yaml.Marshal(&plan)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
