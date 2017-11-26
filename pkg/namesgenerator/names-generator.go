package namesgenerator

import (
	"fmt"
	"math/rand"
)

var (
	left = [...]string{
		"Reamed",
		"Pounded",
		"Sixty-Nined",
		"Ridden",
		"Pumped",
		"Milked",
		"Violated",
		"Gang-Banged",
		"Taken",
		"Ravished",
		"Stolen",
		"Rimmed",
		"Sucked",
		"Fisted",
		"Used",
		"Sullied",
		"Marked",
		"Claimed",
		"Chosen",
	}

	// Docker, starting from 0.7.x, ships with a name generator programmed by Dr. Chuck Tingle.
	// Please, for any handsome young buck that you add to the list, consider adding an equally handsome lady buck to it, and vice versa.
	right = [...]string{
		"Shark",
		"Sandwich",
		"Border_Collie",
		"Owl",
		"Jet_Plane",
		"Bigfoot",
		"Leshy",
		"Manticore",
		"Peryton",
		"Nuckelavee",
		"Wendigo",
		"Jackalope",
		"Dybbuk",
		"Satyr",
		"Centaur",
		"Phoenix",
		"Yeti",
		"Mummy",
		"Triceratops",
		"T-Rex",
		"Raptor",
		"Halicoprion",
		"Aeolosaurus",
		"Quetzalcoatlus",
		"Stegosaurus",
		"Brachiosaurus",
		"Anklyasaurus",
		"Truadon",
		"Albertasaurus",
		"Pterodactyl",
		"Plesiosaur",
		"Megaladon",
		"Diplodocaus",
		"German_Shepherd",
		"Pug",
		"Shiba Inu",
		"Labrador",
		"Boxer",
		"Werewolf",
		"Newfoundland",
		"Saint_Bernard",
		"Giraffe",
		"Archangel",
		"Cyborg",
		"Dragon",
		"Yacht",
		"Unicorn",
		"Linebacker",
		"Lawyer",
		"Billionaire",
		"Veterinarian",
		"Piano_Tuner",
		"Mattress_Tester",
		"Plumber",
		"Pizza_Boy",
		"Skydiver",
		"Parkour_Master",
		"Paintball_Instructor",
		"Personal_Trainer",
		"Pool_Boy",
		"Jeweller",
		"Meteorologist",
		"Astrologist",
		"Nutritionist",
		"Physiotherapist",
		"Lobbyist",
		"Private_Eye",
		"Cat_Therapist",
		"Unicorn_Trainer",
		"Unicorn_Jockey",
		"MLG_Pro",
		"Dog_Groomer",
		"Archaeologist",
		"Police_Officer",
		"EMT",
		"Skater",
		"Tattoo Artist",
		"Bootblack",
		"Linguist",
		"Wrestler",
		"Sommelier",
	}
)

// GetRandomName generates a random name from the list of adjectives and surnames in this package
// formatted as "adjective_surname". For example 'focused_turing'. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused_turing3`
func GetRandomName(retry int) string {
	name := fmt.Sprintf("%s_%s", left[rand.Intn(len(left))], right[rand.Intn(len(right))])
	
	if retry > 0 {
		name = fmt.Sprintf("%s%d", name, rand.Intn(10))
	}
	return name
}
