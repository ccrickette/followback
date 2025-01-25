package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Simply program to parse and extract list of followers that don't follow you back
//How to Use: Simply go to https://accountscenter.instagram.com/info_and_permissions/->Select "Your information and permissions"-> Choose "Download your information"
// Select Download your information and choose "Some of oyur information". We only need to get a copy of the followers, nothing else.
// Under "Connections" header choose "Followers and Following" - select for entire time to get a full extract.
// This might take a few minutes for Meta to process. This will then prepare a zip file with 2 json files followers_1.json & following.json
// Download the zip file and place the 2 files in the root of this directory and you will get a "difference.txt"

type StringListData struct {
	Href      string `json:"href"`
	Value     string `json:"value"`
	Timestamp int    `json:"timestamp"`
}

// Relationship represents each object in the JSON array
type Relationship struct {
	Title          string           `json:"title"`
	MediaListData  []interface{}    `json:"media_list_data"`
	StringListData []StringListData `json:"string_list_data"`
}

// FollowingData represents the root structure of following.json
type FollowingData struct {
	RelationshipsFollowing []Relationship `json:"relationships_following"`
}

func main() {
	// Load followers_1.json
	followersFile, err := os.ReadFile("followers_1.json")
	if err != nil {
		fmt.Printf("Error reading followers file: %v\n", err)
		return
	}

	// Load following.json
	followingFile, err := os.ReadFile("following.json")
	if err != nil {
		fmt.Printf("Error reading following file: %v\n", err)
		return
	}

	// Parse followers_1.json as an array of Relationship
	var followersData []Relationship
	err = json.Unmarshal(followersFile, &followersData)
	if err != nil {
		fmt.Printf("Error parsing followers JSON: %v\n", err)
		return
	}

	// Parse following.json as FollowingData
	var followingData FollowingData
	err = json.Unmarshal(followingFile, &followingData)
	if err != nil {
		fmt.Printf("Error parsing following JSON: %v\n", err)
		return
	}

	// Extract values from followers data into a set for fast lookup
	followersSet := make(map[string]bool)
	for _, relationship := range followersData {
		for _, stringData := range relationship.StringListData {
			followersSet[stringData.Value] = true
		}
	}

	// Find values in following.json that are not in followers_1.json
	var differences []string
	for _, relationship := range followingData.RelationshipsFollowing {
		for _, stringData := range relationship.StringListData {
			if !followersSet[stringData.Value] {
				differences = append(differences, stringData.Value)
			}
		}
	}

	// Write the differences to a plain text file
	diffFile, err := os.Create("differences.txt")
	if err != nil {
		fmt.Printf("Error creating differences file: %v\n", err)
		return
	}
	defer diffFile.Close()

	for _, diff := range differences {
		_, err := diffFile.WriteString(diff + "\n")
		if err != nil {
			fmt.Printf("Error writing to differences file: %v\n", err)
			return
		}
	}

	fmt.Println("Differences have been written to differences.txt")
}
