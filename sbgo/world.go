// Copyright 2020 Vladislav Smirnov

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

// World is used as a general structure of a world
type World struct {
	Size   int                 `json:"size"`
	Points map[int]*WorldPoint `json:"points"`
}

type couple struct {
	Point1 int
	Point2 int
}

/* --- WORLD GENERATION --- */

// optimize or rewrite world generation => DONE

// GenerateWorld create a world of s points
func GenerateWorld(s int) *World {
	wp := make(map[int]*WorldPoint)

	fmt.Println("Generating world... 0%")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < s; i++ {
		wp[i] = &WorldPoint{
			LocType:  rand.Intn(3),
			Position: generatePosition(wp, i),
			Adjacent: make([]int, 0),
		}
	}

	fmt.Println("Generating world... 100%")

	// make sure there are no disjoint graphs
	fmt.Print("Building MST...")
	buildMST(wp)
	fmt.Println(" done.")

	// add more random connections
	edgeDistance := viper.GetFloat64("EdgeDistance")
	for i := 0; i < s-1; i++ {
		fmt.Printf("Generating edges... %d%%\n", 100*(i+1)/s)
		for j := i + 1; j < s; j++ {
			dist := wp[i].Position.Distance(wp[j].Position)
			if dist < edgeDistance {
				wp[i].Adjacent = append(wp[i].Adjacent, j)
				wp[j].Adjacent = append(wp[j].Adjacent, i)
			}
		}
	}

	fmt.Println("Generating edges... 100%")

	w := World{
		Size:   s,
		Points: wp,
	}

	return &w
}

func buildMST(wp map[int]*WorldPoint) {
	v := make([]int, 1)
	v[0] = 0 // set initial visited point 0
	for {
		c := findNearestCouple(v, wp)
		if c.Point1 == -1 {
			break
		}
		wp[c.Point1].Adjacent = append(wp[c.Point1].Adjacent, c.Point2)
		wp[c.Point2].Adjacent = append(wp[c.Point2].Adjacent, c.Point1)
		v = append(v, c.Point2)
	}
}

// for MST
func findNearestCouple(cp []int, wp map[int]*WorldPoint) couple {
	nearestCouple := couple{-1, -1}
	nearestDist := 0.0
	if len(cp) == 0 {
		return nearestCouple
	}
	for i, p := range wp {
		if isInArray(i, cp) {
			continue
		}
		currentID := -1
		dist := 0.0
		for _, c := range cp {
			cdist := p.Position.Distance(wp[c].Position)
			if cdist < dist || currentID == -1 {
				dist = cdist
				currentID = c
			}
		}
		if currentID != -1 && (nearestCouple.Point1 == -1 || dist < nearestDist) {
			nearestCouple.Point1 = currentID
			nearestCouple.Point2 = i
			nearestDist = dist
		}
	}
	return nearestCouple
}

func isInArray(i int, a []int) bool {
	for _, k := range a {
		if i == k {
			return true
		}
	}
	return false
}

// TODO: this may generate the same coordinates for different points
func generatePosition(wp map[int]*WorldPoint, s int) Vector2 {
	v := Vector2{
		X: rand.Intn(1000),
		Y: rand.Intn(1000),
	}

	for !checkDistance(v, wp, s) {
		v = Vector2{
			X: rand.Intn(1000),
			Y: rand.Intn(1000),
		}
	}

	return v
}

func checkDistance(v Vector2, wp map[int]*WorldPoint, s int) bool {
	if s == 0 {
		return true
	}

	for i := 0; i < s; i++ {
		p, ok := wp[i]
		if !ok {
			fmt.Println("Invalid map access. Perhaps checkDistance size argument is wrong.")
		}

		minimalDistance := viper.GetFloat64("MinimalDistance")
		if p.Position.Distance(v) < minimalDistance {
			return false
		}
	}

	return true
}
