package main

import (
	"fmt"
)

type alert struct {
	text   string
	routes []string
}

type alerts []alert

type alertSet []alert

func main() {
	var coll alertSet

	al := getFirstAlert()
	err := al.addAlertsToCollection(&coll)

	if err != nil {
		fmt.Printf("%v", err)
	}

	al = getSecAndThirdAlert()
	err = al.addAlertsToCollection(&coll)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println(coll)
}

func (als alerts) addAlertsToCollection(set *alertSet) error {
	for _, v := range als {
		*set = append(*set, v)
	}
	return nil
}

func getFirstAlert() alerts {
	a := []alert{alert{
		text: "beware of dragons",
		routes: []string{
			"on the road to cairns",
			"on the path to raven",
		}},
	}

	return a
}

func getSecAndThirdAlert() alerts {
	a := []alert{
		alert{
			text: "high dragon activity during the night",
			routes: []string{
				"dark lane",
			},
		},
		alert{
			text: "hungry dragons in the morning",
			routes: []string{
				"hopping fields",
				"grassy knoll",
				"meadow trail",
			},
		},
	}

	return a
}
