package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Topic struct {
	Topic                 string `json:"topic"`
	MyLearnings           string `json:"learnings"`
	Date                  string `json:"date"`
	ProjectsBuild         string `json:"projectsbuild"`
	LeetcodeProblemSolved string `json:"leetcodeproblemsolved"`
}

func main() {
	Topics := []Topic{
		{
			Topic:                 "Dijkstra's Algorithm",
			MyLearnings:           "Understood the shortest path algorithm,how the backend mechanics and math works",
			Date:                  "16.06.2024",
			ProjectsBuild:         "",
			LeetcodeProblemSolved: "",
		},
		{
			Topic:                 "Docker",
			MyLearnings:           "Understood about the images , containers, Stats, Docker Hub, Bridge, Hosts and its commands",
			Date:                  "22.06.2024",
			ProjectsBuild:         "",
			LeetcodeProblemSolved: "",
		},
		{
			Topic:                 "Ollama",
			MyLearnings:           "Set ollama locally on docker , imported llama3 on the docker container, ran it",
			Date:                  "22.06.2024",
			ProjectsBuild:         "",
			LeetcodeProblemSolved: "",
		},
	}

	jsonData, err := json.Marshal(Topics)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %s", err)
	}

	fmt.Println(string(jsonData))
}
