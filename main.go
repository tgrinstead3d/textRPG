package main

import "fmt"

type Tile struct {
	Name        string
	Intro       string
	Description string
	Exits       map[string]*Tile
	Visited     bool
}

type Player struct {
	Name        string
	CurrentTile *Tile
	Inventory   []string
	Bitten      bool
	Health      int
	Dead        bool
}

var input string

func printExits(tile *Tile) {
	fmt.Print("Exits: ")
	for direction := range tile.Exits {
		fmt.Print(direction, " ")
	}
	fmt.Println()
}

var exitApp bool = false

func main() {
	room1 := Tile{Name: "Entrance", Description: "The air-filled entrance cavern features bioluminescent coral, casting a soft glow on rocky walls. Vibrant fish swim through the chamber, water comes in through the cave to the South and three tunnels extend East, West, and North.", Exits: make(map[string]*Tile), Visited: true, Intro: "none"}
	room2 := Tile{Name: "Water-filled Cave", Description: "The room is nothing more than a dark, musty pit filled with oily, dirty water. It's dark and small, but the only exit is a ladder made of live, hungry snakes.", Exits: make(map[string]*Tile), Visited: false, Intro: "none"}
	room3 := Tile{Name: "Coral Pit", Description: "Clear, undisturbed waters dimly glow from the bioluminescence of the cavern entrance. Mysterious aquatic life gently swim and float around. Visibility is low and it creates a tense and uneasy setting.", Exits: make(map[string]*Tile), Visited: false, Intro: "none"}
	room4 := Tile{Name: "Open Cave", Description: "A vast chamber that is illuminated by bioluminescent fungi. Ancient rock formations and towering underwater pillars define a surreal underwater landscape. The expanse is too daunting and open to explore further.", Exits: make(map[string]*Tile), Visited: false, Intro: "none"}
	room5 := Tile{Name: "Ceremonial Chamber", Description: "The chamber is filled with shallow water with a gentle current. A dark and intimidating tunnel extends to the North. The walls of the chamber have delicate patterns with vibrant flora and a tranquil atmosphere. A soft bioluminescent glow is provided by the flora making the room feel inviting and safe.", Exits: make(map[string]*Tile), Visited: false, Intro: "none"}
	room6 := Tile{Name: "Final Room", Description: "Flickering torchlight casts shadows on cold stone walls. The floor, strewn with lifeless bodies, creates a somber scene. The interplay of light and shadows in the unsettling ambiance blurs the distinction between reality and illusion, enhancing the eerie atmosphere of the macabre room.", Exits: make(map[string]*Tile), Visited: false, Intro: "none"}

	room1.Exits["west"] = &room5
	room1.Exits["south"] = &room3
	room1.Exits["north"] = &room4
	room1.Exits["east"] = &room2
	room2.Exits["west"] = &room1
	room3.Exits["north"] = &room1
	room4.Exits["south"] = &room1
	room5.Exits["east"] = &room1
	room5.Exits["north"] = &room6
	room6.Exits["south"] = &room5

	currentTile := &room1
	player := Player{Name: "none", CurrentTile: currentTile, Inventory: []string{}, Bitten: false, Health: 100, Dead: false}

	fmt.Println("Who are you?")
	fmt.Scanln(&player.Name)
	fmt.Println("")
	fmt.Println("Welcome to THE CURRENTS OF FATE, a text based adventure game by Tyler Grinstead.")
	fmt.Println("********************")
	fmt.Println("Exit the game at any time by typing 'exit'.")
	fmt.Print("The journey begins on a serene fishing boat, gently bobbing on the surface of the open sea. ")
	fmt.Print(player.Name)
	fmt.Println(", an avid fisherman, had set out to explore a remote and uncharted region, lured by tales of mythical sea creatures and hidden treasures. The day unfolded with a sense of tranquility as they cast their fishing line into the deep blue, the rhythmic sounds of the waves accompanying their solitary venture.")
	fmt.Println("As your main character regains consciousness in the mysterious underwater cavern, the soft glow of bioluminescent coral casts an ethereal light, revealing the beauty of the submerged world around them. The air in the cavern is surprisingly fresh, and the gentle sound of water lapping against the rocky walls echoes through the chamber.")
	fmt.Println("")
	fmt.Println("The walls of the cavern are adorned with vibrant coral formations in various hues, creating a mesmerizing underwater tapestry. Schools of iridescent fish dart between the coral, their scales reflecting the ambient light. Large anemones sway gracefully with the subtle currents, creating a dance of colors and textures.")
	fmt.Println("")
	fmt.Println("In the South corner of the cavern, there's a serene water hole, its surface shimmering with reflections of the surrounding beauty. The water is crystal clear, inviting your character to explore its depths. As your character observes the water hole, they notice a subtle movement beneath the surface—a mysterious aquatic creature gracefully gliding through the depths.")
	fmt.Println("")
	fmt.Println("Three tunnels beckon in different directions, each offering a unique path for exploration. To the east, the tunnel is adorned with bioluminescent fungi, casting an eerie yet enchanting glow. To the west, a gentle current flows, carrying with it a chorus of underwater whispers. The north tunnel reveals a cavernous expanse with towering underwater rock formations, creating an otherworldly skyline.")
	fmt.Println("")
	fmt.Println("As " + player.Name + " stands at the crossroads of these tunnels, a layer of mystery is blanketing their surroundings. The journey ahead is uncertain, and the underwater world holds secrets waiting to be unveiled. The choice of which path to take is yours, and the adventure unfolds in the depths of this captivating underwater realm.")
	fmt.Println("")

	for {
		if currentTile.Visited == false {
			fmt.Println(currentTile.Intro)
			currentTile.Visited = true
		} else {
			fmt.Println(currentTile.Description)
		}
		printExits(currentTile)
		var input string
		fmt.Scanln(&input)

		if nextTile, ok := currentTile.Exits[input]; ok {
			currentTile = nextTile
		} else if input == "exit" {
			fmt.Println("Thanks for playing!")
			break
		} else {
			fmt.Println("Invalid input")
		}
	}
}
