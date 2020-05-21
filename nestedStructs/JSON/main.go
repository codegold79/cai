package main

import (
	"encoding/json"
	"fmt"
)

// At first, I tried to adhere to the strict typing system of Go.
// However, it is very unforgiving and inflexible. Call this ProcessV1.
type respTS struct {
	Base struct {
		Inventory struct {
			Clusters []struct {
				Namespaces []struct {
					Nodes []struct {
						Deployments []struct {
							Pods []interface{}
						}
					}
				}
			}
		}
	}
}

func main() {
	// Raw data is a mess. It's a nest in a nest in a nest. Who knows where the
	// desired data really is?
	rawData := data()

	// My first way going about this was unfeasible.
	var response respTS

	err := json.Unmarshal(rawData, &response)

	if err != nil {
		fmt.Println(err)
		return
	}

	// The goal is to get the Phase, Response Code, and Pod Data.
	// processV1 was able to grab pod data, but I gave up when I saw how bad it was.
	processV1(response)

	// Let's try something new. No generics? No problem. Using the empty Go interface.
	fmt.Print("\n\n\n=================================\n")

	// As noted above, let's start with the empty interface.
	var simpleResp interface{}
	err = json.Unmarshal(rawData, &simpleResp)

	if err != nil {
		fmt.Println(err)
		return
	}

	// This process works a whole lot better. And the recursive processing makes
	// this look quite elegant. I think I should write a blog on this.
	processV2(simpleResp)

	// Credits. These three resources helped me to understand how to write
	// processV2 function:
	// https://tour.golang.org/methods/15
	// https://medium.com/code-zen/dynamically-creating-instances-from-key-value-pair-map-and-json-in-go-feef83ab9db2
	// https://eager.io/blog/go-and-json/
}

func processV2(input interface{}, keys ...string) {
	switch v := input.(type) {
	case string:
		fmt.Printf("%s: %q\n", keys[0], v)
	case float64:
		fmt.Printf("%s: %f\n", keys[0], v)
	case map[string]interface{}:
		for key, val := range v {
			processV2(val, key)
		}
	case []interface{}:
		for _, val := range v {
			processV2(val)
		}
	case bool:
		fmt.Println("bool case")
	case nil:
		fmt.Println("null case")
	default:
		fmt.Printf("since all JSON cases were covered, this was probably not from JSON, %T", v)
	}
}

// Ok, this works. A loop for every level of the struct. What a mess.
// Completely inelegant. We need another way.
func processV1(response respTS) {
	for k, v := range response.Base.Inventory.Clusters {
		fmt.Printf("type v: %T", v)
		for k2, v2 := range response.Base.Inventory.Clusters[k].Namespaces {
			fmt.Printf("type v2: %T", v2)
			for k3 := range response.Base.Inventory.Clusters[k].Namespaces[k2].Nodes {
				for k4 := range response.Base.Inventory.Clusters[k].Namespaces[k2].Nodes[k3].Deployments {
					for k5, v := range response.Base.Inventory.Clusters[k].Namespaces[k2].Nodes[k3].Deployments[k4].Pods {
						fmt.Println(k5, v)
					}
				}
			}
		}
	}
}

func data() []byte {
	return []byte(`{
		"base": {
		  "inventory": {
			"clusters": [
			  {
				"namespaces": [
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": [
							  {
								"phase": "Running",
								"queryServiceTS": {
								  "response": 200,
								  "data": "[{\"time\":\"Thu May 21 2020 16:03:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":null},{\"time\":\"Thu May 21 2020 16:03:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":21},{\"time\":\"Thu May 21 2020 16:04:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":24},{\"time\":\"Thu May 21 2020 16:04:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":25},{\"time\":\"Thu May 21 2020 16:05:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":15},{\"time\":\"Thu May 21 2020 16:05:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":24},{\"time\":\"Thu May 21 2020 16:06:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":22},{\"time\":\"Thu May 21 2020 16:06:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":25},{\"time\":\"Thu May 21 2020 16:07:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":27},{\"time\":\"Thu May 21 2020 16:07:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":23},{\"time\":\"Thu May 21 2020 16:08:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":24},{\"time\":\"Thu May 21 2020 16:08:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":18},{\"time\":\"Thu May 21 2020 16:09:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":24},{\"time\":\"Thu May 21 2020 16:09:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":20},{\"time\":\"Thu May 21 2020 16:10:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":26},{\"time\":\"Thu May 21 2020 16:10:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":26},{\"time\":\"Thu May 21 2020 16:11:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":24},{\"time\":\"Thu May 21 2020 16:11:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":25},{\"time\":\"Thu May 21 2020 16:12:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":27},{\"time\":\"Thu May 21 2020 16:12:30 GMT+0000 (Coordinated Universal Time)\",\"Requests\":22},{\"time\":\"Thu May 21 2020 16:13:00 GMT+0000 (Coordinated Universal Time)\",\"Requests\":7}]"
								}
							  }
							]
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  },
						  {
							"pods": []
						  },
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  },
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  },
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  },
					  {
						"deployments": []
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": []
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": []
					  },
					  {
						"deployments": []
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  },
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": []
				  }
				]
			  },
			  {
				"namespaces": [
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  },
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  },
				  {
					"nodes": []
				  },
				  {
					"nodes": [
					  {
						"deployments": []
					  },
					  {
						"deployments": [
						  {
							"pods": []
						  }
						]
					  }
					]
				  }
				]
			  }
			]
		  }
		}
	  }
	`)
}
