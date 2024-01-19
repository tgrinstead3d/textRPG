package main

import (
	"fmt"
)

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
var action string

func printExits(tile *Tile) {
	fmt.Print("Exits: ")
	for direction := range tile.Exits {
		fmt.Print(direction, " ")
	}
	fmt.Println()
}

func main() {
	room1 := Tile{Name: "Entrance", Description: "The air-filled entrance cavern features bioluminescent coral, casting a soft glow on rocky walls. Vibrant fish swim through the chamber, water comes in through the cave to the South and three tunnels extend East, West, and North.", Exits: make(map[string]*Tile), Visited: true, Intro: "You find yourself in a cavernous entrance room, part of an enchanting underwater world. The air is surprisingly fresh, filled with the invigorating scent of the sea. The chamber is well-lit by the soft glow of bioluminescent coral formations, casting an ethereal light on the rocky walls.\n\nThe air-filled portion of the cavern provides a comfortable environment, allowing you to catch your breath. The underwater landscape is adorned with vibrant coral reefs in various hues, creating a visually stunning and serene atmosphere. Schools of iridescent fish gracefully swim through the chamber, their scales reflecting the ambient light.\n\nIn one corner of the cavern, there's a serene water hole to the South with crystal-clear water. The surface remains undisturbed, hinting at hidden depths below. Mysterious aquatic life moves gracefully beneath the water's surface, adding an extra layer of intrigue to the surroundings.\n\nThree tunnels beckon for exploration—one leading east, another west, and the third heading north. Each tunnel is a portal to a different underwater realm, promising unique adventures and discoveries. The choice of which path to take is the first step in your journey through this captivating and mysterious underwater world."}
	room2 := Tile{Name: "Water-filled Cave", Description: "The room is nothing more than a dark, musty pit filled with oily, dirty water. It's dark and small, but the only exit is a ladder made of live, hungry snakes.", Exits: make(map[string]*Tile), Visited: false, Intro: "Upon entering the watering hole, you are greeted by a vast expanse of water, teeming with a fascinating array of sea creatures. On one side, adorable and inviting creatures with vibrant colors playfully dart through the water, their whimsical movements creating a sense of wonder. These cute companions seem to beckon you into the inviting depths, promising hidden delights.\n\nHowever, on the opposite side of the expanse, menacing and intimidating sea creatures lurk in the shadows. Their fierce, predatory gazes and ominous silhouettes send shivers down the spine. Dark and mysterious, these creatures add an element of danger to the otherwise enchanting scene.\n\nThe contrasting mix of adorable and menacing sea creatures creates a dynamic underwater tableau. Colorful schools of charming fish coexist with formidable predators, offering you a choice—to cautiously navigate the inviting waters, guided by the allure of the cute creatures, or to confront the challenges posed by the menacing and intimidating inhabitants. The underwater world unfolds as a captivating, visually stunning, and perilous adventure awaits you. Unfortunately, you do not have the lung capacity to explore the depths of the watering hole, and you must return to the surface to continue your journey through the underwater cavern."}
	room3 := Tile{Name: "Coral Pit", Description: "Clear, undisturbed waters dimly glow from the bioluminescence of the cavern entrance. Mysterious aquatic life gently swim and float around. Visibility is low and it creates a tense and uneasy setting.", Exits: make(map[string]*Tile), Visited: false, Intro: "You discover an unsettling atmosphere... The space is dark and musty, with minimal light penetrating the shadows. The air feels heavy, carrying a distinct dampness.\n\nAs you cautiously explore, you stumble upon an unexpected hazard—a sudden hidden drop in the floor that leads to a pit filled with water. The pit is shrouded in darkness, and the sound of echoing drips reverberates in the confined space. The water below reflects a murky abyss, concealing the depth of the pit.\n\nAlong the wall where the entrance drops in, a chilling discovery awaits— a ladder made entirely of live, angry snakes. The serpents writhe and hiss, creating an unsettling and ominous spectacle. Their scales glisten faintly in the low light, and their eyes gleam with an eerie intensity.\n\nThe ladder of live snakes presents a formidable challenge, and you must carefully navigate this treacherous path to continue your journey through the mysterious underwater cavern. The combination of darkness, mustiness, and the unsettling ladder of serpents adds a layer of tension to this part of the underwater adventure."}
	room4 := Tile{Name: "Open Cave", Description: "A vast chamber that is illuminated by bioluminescent fungi. Ancient rock formations and towering underwater pillars define a surreal underwater landscape. The expanse is too daunting and open to explore further.", Exits: make(map[string]*Tile), Visited: false, Intro: "You discover a vast and awe-inspiring chamber within the underwater cavern. The space is illuminated by clusters of bioluminescent fungi, casting an eerie yet enchanting glow. The walls are adorned with intricate underwater rock formations, creating an otherworldly landscape.\n\nThe chamber opens up to reveal towering underwater pillars and arches, reminiscent of ancient ruins beneath the sea. The rocky floor is dotted with submerged caves and hidden alcoves, inviting exploration. Schools of iridescent fish navigate through the grand expanse, adding a dynamic and vibrant element to the underwater scenery.\n\nAs you venture deeper into the northern chamber, the distant echoes of underwater currents become more pronounced, hinting at the mysteries concealed within the depths. The underwater skyline is adorned with silhouettes of towering rock formations, creating a breathtaking and surreal underwater panorama.\n\nThis northern chamber offers a glimpse into the rich history of the underwater realm, with hidden secrets waiting to be uncovered. The combination of bioluminescent fungi, ancient rock formations, and the promise of undiscovered wonders makes this room a captivating and integral part of your underwater journey."}
	room5 := Tile{Name: "Ceremonial Chamber", Description: "The chamber is filled with shallow water with a gentle current. A dark and intimidating tunnel extends to the North. The walls of the chamber have delicate patterns with vibrant flora and a tranquil atmosphere. A soft bioluminescent glow is provided by the flora making the room feel inviting and safe.", Exits: make(map[string]*Tile), Visited: false, Intro: "In the western chamber, you explore an underwater world characterized by gentle currents and intricate formations. The walls are adorned with delicate patterns created by the subtle flow of water over time. The ambient light reveals a mesmerizing array of colorful underwater flora, creating an enchanting and serene atmosphere.\n\nAs you venture further, the underwater currents guide them through a series of interconnected tunnels, each unveiling hidden wonders. The chamber is a symphony of underwater whispers, hinting at the existence of secret sanctuaries and concealed passages.\n\nThe gentle currents carry you past ethereal underwater gardens, where aquatic life flourishes in harmony with the mesmerizing surroundings. Illuminated by the soft glow of bioluminescent flora, the western chamber offers a tranquil and captivating underwater experience.\n\nThis room provides a stark contrast to the darkness of the east, offering a peaceful and visually stunning journey for you to navigate. The underwater world's delicate beauty is on full display in the western chamber, making it a unique and memorable part of the underwater adventure."}
	room6 := Tile{Name: "Final Room", Description: "Flickering torchlight casts shadows on cold stone walls. The floor, strewn with lifeless bodies, creates a somber scene. The interplay of light and shadows in the unsettling ambiance blurs the distinction between reality and illusion, enhancing the eerie atmosphere of the macabre room.", Exits: make(map[string]*Tile), Visited: false, Intro: "The flickering torchlight, like a dying ember, throws grotesque shadows that writhe and dance across the cold, damp stone walls. Beneath your feet, the uneven flagstones are slick with a viscous liquid that catches the errant glints of flame, revealing a horrifying truth – the floor is not stone, but a macabre mosaic of lifeless bodies, their sightless eyes reflecting the distorted firelight.\n\nThe air is thick with the cloying scent of decay, mingling with the acrid tang of burnt flesh and the metallic tang of blood. The silence is broken only by the erratic drip of moisture from unseen cracks in the ceiling, each drop echoing with a hollow, mournful chime.\n\nAs your eyes adjust to the gloom, you begin to discern the grisly details of your surroundings. The bodies, clad in tattered remnants of finery, are sprawled in unnatural contortions, their faces frozen in silent screams. Some clutch at empty air, their fingers clawed in rigor mortis, while others lie supine, their limbs splayed wide as if flung down in a moment of unimaginable terror.\n\nThe interplay of light and shadow plays tricks on your vision. Grotesque gargoyles carved into the walls seem to leer down at you, their stone fangs glinting in the firelight. Cobwebs, draped like ghostly shrouds, hang from the vaulted ceiling, and in their intricate strands, you swear you can see the spectral forms of those who met their demise in this accursed chamber.\n\nThe line between reality and illusion blurs, the macabre scene taking on a nightmarish quality. Is it the flickering torchlight or are the shadows truly dancing? Do the bodies twitch, or is it merely the play of your own overwrought imagination? The very air seems to crackle with unseen energy, charged with the lingering echoes of screams and the palpable weight of despair.\n\nYou take a hesitant step forward, the crunching of bone beneath your foot a stark reminder of the grim tableau that surrounds you. The stench intensifies, threatening to overwhelm you, and you press a hand to your nose, trying to shut out the overpowering cloying sweetness of decay.\n\nBut even as you turn to flee, a new sound catches your ear – a faint, rhythmic scraping from the far corner of the room. You strain to listen, your heart hammering against your ribs..."}

	room1.Exits["west"] = &room5
	room1.Exits["south"] = &room2
	room1.Exits["north"] = &room4
	room1.Exits["east"] = &room3
	room3.Exits["west"] = &room1
	room2.Exits["north"] = &room1
	room4.Exits["south"] = &room1
	room5.Exits["east"] = &room1
	room5.Exits["north"] = &room6
	room6.Exits["south"] = &room5

	currentTile := &room1
	player := Player{Name: "none", CurrentTile: currentTile, Inventory: []string{"knife", "fishing line", "3 gold coins"}, Bitten: false, Health: 100, Dead: false}

	ending := func() {
		fmt.Println("")
		fmt.Println("The relentless sun beats down on the parched earth, painting the endless dunes in a shimmering haze. Before you, nestled like a emerald jewel in the sun-scorched wasteland, lies the oasis of Al'Sharim. Crystal-clear water shimmers amidst a grove of swaying date palms, their fronds whispering secrets in the warm desert breeze.\n\nAs you approach, the scent of jasmine and spices hangs heavy in the air, carried on tendrils of smoke rising from mudbrick houses nestled beneath the palms. Cobblestone paths wind through the settlement, bustling with life. Merchants hawk their wares in colorful stalls, their voices weaving a melodious hum. Children chase each other, their laughter echoing off the sun-baked walls.\n\nA towering minaret pierces the azure sky, its call to prayer resonating across the oasis. A bustling marketplace lies at its base, overflowing with exotic fruits, woven tapestries, and handcrafted trinkets. Camel caravans rest nearby, their riders huddled in the shade, sharing stories and sipping cool mint tea.\n\nA cool, inviting stream meanders through the heart of the oasis, fed by a hidden spring. Gentle waterfalls cascade over smooth stones, forming turquoise pools that glimmer like sapphires. Lush grasses line the banks, dotted with vibrant wildflowers that dance in the breeze.\n\nBeyond the oasis, the harsh desert landscape resumes. Towering sand dunes, sculpted by relentless winds, rise in the distance, their peaks disappearing into the hazy horizon. A sense of mystery and danger hangs in the air, beckoning the adventurous spirit.\n\nBut within the oasis walls, a sense of peace and tranquility prevails. Al'Sharim offers a haven from the unforgiving desert, a place where weary travelers can rest, refresh, and embark on new adventures. Will you stay and unravel the secrets this enchanting oasis holds, or will you venture beyond its embrace into the vast, sun-scorched unknown?")
		fmt.Println("Thank you for playing!")
	}

	help := func() {
		fmt.Println("To play, type the direction you would like to go (north, south, east, or west) and press enter. \n\nTo interact with objects, type the object name and press enter. \n\nTo view your inventory, type \"inventory\" and press enter. \n\nTo view your health, type \"health\" and press enter. \n\nTo view your current location, type \"location\" and press enter. \n\nTo exit the game, type \"exit\" and press enter. \n\nTo view these instructions again, type \"help\"")
		fmt.Scanln(&input)
	}

	listen := func() {
		if input == "help" {
			help()
		} else if input == "inventory" {
			fmt.Println(player.Inventory)
		} else if input == "health" {
			fmt.Println(player.Health)
		} else if input == "location" {
			fmt.Println(currentTile.Name)
			fmt.Println(currentTile.Description)
			printExits(currentTile)
		}
		fmt.Scanln(&input)
	}

	fmt.Println("Please enter your character's name:")
	fmt.Scanln(&player.Name)
	fmt.Println("To play, type the direction you would like to go (north, south, east, or west) and press enter. \n\nTo interact with objects, type the object name and press enter. \n\nTo view your inventory, type \"inventory\" and press enter. \n\nTo view your health, type \"health\" and press enter. \n\nTo view your current location, type \"location\" and press enter. \n\nTo exit the game, type \"exit\" and press enter. \n\nTo view these instructions again, type \"help\"")
	fmt.Printf("\n\n The sun, a molten orb on the horizon, cast its warm glow across the cerulean expanse of the ocean. Gentle waves lapped at the weathered hull of the \"Salty Dog\", a seafaring vessel as comfortable and familiar to %s as his own skin. Today, however, the\"Salty Dog\" wasn't merely a fishing boat; it was a chariot carrying them towards adventure. Tales of monstrous krakens guarding sunken galleons laden with gold, whispered by weathered sailors in smoke-filled taverns, had spurred %s to explore this remote, uncharted corner of the sea. \n\n Excitement danced in his chest, mirroring the playful glint of sunlight on the water. The air was sweet with the tang of salt and brine, punctuated by the rhythmic creak of the ship's timbers and the soothing lull of the waves. As %s cast his line, the lure danced and twirled on the surface, a silver ballerina against the sapphire stage. Each gentle tug on the line sent a thrill through him, a promise of hidden bounty lurking in the depths. The day stretched on, languid and peaceful. Time seemed to dissolve into the endless blue, punctuated only by the rhythmic rise and fall of the boat and the occasional cry of a lone gull wheeling overhead. \n\n But then, in an instant, the serene tableau shattered. A bolt of lightning, impossibly vivid against the cloudless sky, ripped across the horizon. It was a celestial spear, crackling with raw power, aimed directly at their boat. Before %s could even register the impossible sight, the world exploded in a blinding flash of white. The deafening crack of thunder sent a tremor through the boat, nearly tossing him overboard. His ears sang, his vision blurred, and a jolt of electricity coursed through his body, leaving him tingling and numb. \n\n As the world faded to black, a chilling thought echoed in the void of his mind: they weren't alone in this vast ocean.", player.Name, player.Name, player.Name, player.Name)
	fmt.Println("")
	fmt.Println("You regain consciousness in a mysterious underwater cavern, the soft glow of bioluminescent coral casts an ethereal light, revealing the beauty of the submerged world around you. The air in the cavern is surprisingly fresh, and the gentle sound of water lapping against the rocky walls echoes through the chamber. \n\n The walls of the cavern are adorned with vibrant coral formations in various hues, creating a mesmerizing underwater tapestry. Schools of iridescent fish dart between the coral, their scales reflecting the ambient light. Large anemones sway gracefully with the subtle currents, creating a dance of colors and textures. \n\n In the South corner of the cavern, there's a serene water hole, its surface shimmering with reflections of the surrounding beauty. The water is crystal clear, inviting your character to explore its depths. As your character observes the water hole, they notice a subtle movement beneath the surface—a mysterious aquatic creature gracefully gliding through the depths. \n\n Three tunnels beckon in different directions, each offering a unique path for exploration. To the east, the tunnel is adorned with bioluminescent fungi, casting an eerie yet enchanting glow. To the west, a gentle current flows, carrying with it a chorus of underwater whispers. The north tunnel reveals a cavernous expanse with towering underwater rock formations, creating an otherworldly skyline. \n\n As \" + player.Name + \" stands at the crossroads of these tunnels, a layer of mystery is blanketing their surroundings. The journey ahead is uncertain, and the underwater world holds secrets waiting to be unveiled. The choice of which path to take is yours, and the adventure unfolds in the depths of this captivating underwater realm.")

	for {
		if currentTile.Name == "Final Room" {
			fmt.Println(currentTile.Intro)
			currentTile.Visited = true
			fmt.Println("Suddenly, a guttural moan shatters the silence. It erupts from the far corner, where shadows gather in an impenetrable mass. Your heart skips a beat, and a tremor of fear crawls up your spine. \n\nThen, as if born from the darkness itself, a figure lunges. Skeletal and emaciated, with skin stretched taut over its frame, it moves with a feral grace that belies its decrepitude. Glowing embers burn in its vacant eyes, reflecting the faint torchlight. In its bony hand, it clutches a rusty scythe, its blade dripping with a viscous, black liquid.\n\nThe creature lets out a raspy shriek, a chilling echo in the tomb-like chamber. It raises its scythe, its glowing eyes fixed on you with a predatory hunger. You know, in that instant, that you have stumbled into a nightmare come alive. Your mind races, searching for a plan, for an escape from this grotesque tableau.\n\nBut before you can even react, the scythe whistles through the air, aimed for your heart. You have no choice but to fight. And so, the macabre dance begins, bathed in the flickering light of your torch, the stench of death the only music in this chamber of horrors.")
			fmt.Println("Do you FIGHT or RUN?")
			fmt.Scanln(&input)
			if input == "fight" {
				fmt.Println("You fight the creature")
				fmt.Println("You have the attacker pinned and your position is now advantageous. It pleads with you to stop fight and that it needs to kill you... that you want it to kill you. Do you give up and accept your death or do you finish the job and end the nightmare?")
				fmt.Println("KILL or BE KILLED?")
				fmt.Scanln(&action)
				if action == "kill" {
					fmt.Println("You kill the man... Silence suddenly falls over the room. You attempt to leave the room but you are unable to reach the door. You are stuck here... Eventually you learn that the one you killed was once just like you... instead of accepting death you, instead, killed the attacker and locked yourself into hell. Once you finally are met with another who has entered this room you find yourself unable to control your body, you attack the person relentlessly, until they are dead. This happens over and over again until finally, you meet you are granted the sweet release of death. You're immediately greeted by darkness and complete silence. You begin to feel relief... until... Blinding light begins to pierce through your closed eyelids. You slowly open your eyes and realize you are standing in the middle of a vast desert.")
					ending()
					break
				} else if action == "be killed" {
					fmt.Println("You accept your death, bravely. You close your eyes and feel the world around you disappear into the blackest black you didn't even know existed. Suddenly, though, a piercing bright light begins to pry into your senses. You slowly open your eyes and realize you are standing in the middle of a vast desert.")
					ending()
					break
				} else {
					listen()
					fmt.Println("Invalid input. Please type \"kill\" or \"be killed\".")
				}
			} else if input == "run" {
				fmt.Println("You attempt to run away but you aren't fast enough. You briefly feel a sharp puncture into your back and then all existence quickly snaps away. There is nothing but black... then from the darkness your vision begins to fade in. Blinding light begins to penetrate your senses.")
				ending()
				break
			} else {
				listen()
				fmt.Println("Invalid input")
			}
		}
		if currentTile.Visited == false {
			fmt.Println(currentTile.Intro)
			currentTile.Visited = true
		} else {
			if currentTile.Name != "Final Room" {
				fmt.Println(currentTile.Description)
			}
		}
		fmt.Println(currentTile.Intro)
		currentTile.Visited = true
		if currentTile.Name != "Final Room" {
			printExits(currentTile)
			fmt.Scanln(&input)
		}

		if nextTile, ok := currentTile.Exits[input]; ok {
			currentTile = nextTile
		} else if input == "exit" {
			fmt.Println("Thanks for playing!")
			break
		} else {
			listen()
			fmt.Println("Invalid input")
		}
		if input == "exit" {
			break
		}
	}
}
