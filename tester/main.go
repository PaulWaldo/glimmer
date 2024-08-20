package glimmer

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gopkg.in/masci/flickr.v3"
	"gopkg.in/masci/flickr.v3/groups"
)

type clientData struct {
	ApiKey           string
	ApiSecret        string
	OAuthToken       string
	OAuthTokenSecret string
}

func writeToFile(filename string, c flickr.FlickrClient) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	d := clientData{
		ApiKey:           c.ApiKey,
		ApiSecret:        c.ApiSecret,
		OAuthToken:       c.OAuthToken,
		OAuthTokenSecret: c.OAuthTokenSecret,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: for pretty printing
	err = encoder.Encode(d)
	if err != nil {
		return err
	}
	return nil
}

func readFromFile(filename string) (*flickr.FlickrClient, error) {
	var d clientData

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&d)
	if err != nil {
		return nil, err
	}

	c := flickr.NewFlickrClient(d.ApiKey, d.ApiSecret)
	c.OAuthToken = d.OAuthToken
	c.OAuthTokenSecret = d.OAuthTokenSecret

	return c, nil
}

// func authenticate() (*flickr.FlickrClient, error) {
// 	client := flickr.NewFlickrClient("817ed1840777a3bba356c0e95c8b1fcf", "896e9b0ddeff8441")
// 	// Works!
// 	// client := flickr.NewFlickrClient("aba7e80c0aae6a896446b2046f1fbddb", "ec104b69e6afcfe2")

// 	// first, get a request token
// 	requestTok, err := flickr.GetRequestToken(client)
// 	if err != nil {
// 		return nil, fmt.Errorf("GetRequestToken: %w", err)
// 	}

// 	// build the authorizatin URL
// 	url, err := flickr.GetAuthorizeUrl(client, requestTok)
// 	if err != nil {
// 		return nil, fmt.Errorf("GetAuthorizeUrl: %s", err)
// 	}

// 	// ask user to hit the authorization url with
// 	// their browser, authorize this application and coming
// 	// back with the confirmation token
// 	fmt.Println("Authorize at ", url)
// 	fmt.Println("Enter confirmation code")
// 	var code string
// 	_, err = fmt.Scanln(&code)
// 	if err != nil {
// 		return nil, fmt.Errorf("reading confirmation code: %s", err)
// 	}

// 	// finally, get the access token, setup the client and start making requests
// 	accessTok, err := flickr.GetAccessToken(client, requestTok, code)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting access token: %s", err)
// 	}
// 	client.OAuthToken = accessTok.OAuthToken
// 	client.OAuthTokenSecret = accessTok.OAuthTokenSecret
// 	return client, nil
// }

type Update struct {
	Type      string
	ID        string
	Name      string
	Timestamp time.Time
	Content   string
}

// func main() {
// 	_, err := readFromFile("client.json")
// 	var auth *glimmer.Authorize
// 	if err != nil {
// 		fmt.Println("Got error ", err, "  Hoping this is just the first run")
// 		secrets := glimmer.Secrets{ApiKey: "817ed1840777a3bba356c0e95c8b1fcf", ApiSecret: "896e9b0ddeff8441"}
// 		auth = glimmer.NewAuth(secrets)
// 		url, err := auth.GetUrl()
// 		if err != nil {
// 			fmt.Println("Error getting auth URL: ", err)
// 			return
// 		}

// 		var code string
// 		_, err = fmt.Scanln(&code)
// 		if err != nil {
// 			fmt.Println("reading confirmation code: %s", err)
// 			return
// 		}

// 		// finally, get the access token, setup the client and start making requests
// 		accessTok, err := flickr.GetAccessToken(auth.Client, requestTok, code)
// 		if err != nil {
// 			return nil, fmt.Errorf("getting access token: %s", err)
// 		}
// 		client.OAuthToken = accessTok.OAuthToken
// 		client.OAuthTokenSecret = accessTok.OAuthTokenSecret
// 		return client, nil

// 		fmt.Println("Auth URL: ", url)
// 		err = writeToFile("client.json", *auth.Client)
// 		if err != nil {
// 			fmt.Println("Error writing JSON: ", err)
// 		}
// 	}

// 	// cl, err := glimmer.GetContactList(client)
// 	// if err != nil {
// 	// 	fmt.Println("Getting Contact List: ", err)
// 	// }
// 	// fmt.Printf("Contacts: %v\n", cl.Contacts)

// 	updates := []Update{}

// 	// Get groups
// 	groupResp, err := groups.GetGroups(auth.Client, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error getting groups:", err)
// 		return
// 	}

// 	for _, group := range groupResp.Groups {
// 		groupUpdates := getGroupUpdates(auth.Client, group.ID)
// 		updates = append(updates, groupUpdates...)
// 	}

// 	// // 	// Get people you follow (this is a placeholder, as there's no direct method for this)
// 	// cl, err := glimmer.GetContactList(client)
// 	// if err != nil {
// 	// 	fmt.Println("Getting Contact List: ", err)
// 	// 	return
// 	// }
// 	// // 	followedPeople := getFollowedPeople(client)
// 	// // followedPeople := cl.Contacts.Contact
// 	// // for _, person := range followedPeople {
// 	// // 	personUpdates := getPersonUpdates(client, person.Id)
// 	// // 	updates = append(updates, personUpdates...)
// 	// // }

// 	// // Sort updates by timestamp
// 	// sort.Slice(updates, func(i, j int) bool {
// 	// 	return updates[i].Timestamp.After(updates[j].Timestamp)
// 	// })

// 	// // Print updates from most to least recent
// 	// for _, update := range updates {
// 	// 	fmt.Printf("%s - %s (%s): %s\n", update.Timestamp, update.Name, update.Type, update.Content)
// 	// }
// }

func getGroupUpdates(client *flickr.FlickrClient, groupID string) []Update {
	updates := []Update{}

	// Get group info
	_ /*groupInfo*/, err := groups.GetInfo(client, groupID)
	if err != nil {
		fmt.Println("Error getting group info:", err)
		return updates
	}

	// Get recent photos from the group
	// photoResp, err := groups.GetPhotos(client, groupID, nil)
	// if err != nil {
	// 	fmt.Println("Error getting group photos:", err)
	// 	return updates
	// }

	// for _, photo := range photoResp.Photos {
	// 	updates = append(updates, Update{
	// 		Type:      "group",
	// 		ID:        groupID,
	// 		Name:      groupInfo.Name,
	// 		Timestamp: time.Unix(int64(photo.LastUpdate), 0),
	// 		Content:   fmt.Sprintf("New photo: %s", photo.Title),
	// 	})
	// }

	return updates
}

// func getFollowedPeople(client *flickr.FlickrClient) []people.Person {
// 	// This is a placeholder function, as there's no direct method to get followed people
// 	// You might need to use a combination of other API calls or store this information separately
// 	return []people.Person{
// 		{Id: "user1", Username: "User 1"},
// 		{Id: "user2", Username: "User 2"},
// 	}
// }

// func getPersonUpdates(client *flickr.FlickrClient, personID string) []Update {
// 	updates := []Update{}

// 	// Get person info
// 	personInfo, err := people.GetInfo(client, personID)
// 	if err != nil {
// 		fmt.Println("Error getting person info:", err)
// 		return updates
// 	}

// 	// Get recent photos from the person
// 	photoResp, err := people.GetPhotos(client, personID, nil)
// 	if err != nil {
// 		fmt.Println("Error getting person photos:", err)
// 		return updates
// 	}

// 	for _, photo := range photoResp.Photos {
// 		updates = append(updates, Update{
// 			Type:      "person",
// 			ID:        personID,
// 			Name:      personInfo.Username,
// 			Timestamp: time.Unix(int64(photo.LastUpdate), 0),
// 			Content:   fmt.Sprintf("New photo: %s", photo.Title),
// 		})
// 	}

// 	return updates
// }

// type Contact struct {
// 	// <contact nsid="12037949629@N01" username="Eric" iconserver="1" realname="Eric Costello" friend="1" family="0" ignored="1" />
// 	Nsid       string `xml:"nsid,attr"`
// 	UserName   string `xml:"username,attr"`
// 	IconServer string `xml:"iconserver,attr"`
// 	RealName   string `xml:"realname,attr"`
// 	Friend     string `xml:"friend,attr"`
// 	Family     string `xml:"family,attr"`
// 	Ignored    string `xml:"ignored,attr"`
// }

// type Contacts struct {
// 	Page    int       `xml:"page,attr"`
// 	Pages   int       `xml:"pages,attr"`
// 	PerPage int       `xml:"perpage,attr"`
// 	Total   int       `xml:"total,attr"`
// 	Contact []Contact `xml:"contact"`
// }

// type GetContactListResponse struct {
// 	flickr.BasicResponse
// 	Contacts Contacts `xml:"contacts"`
// }

// func GetContactList(client *flickr.FlickrClient) (*GetContactListResponse, error) {
// 	client.Init()
// 	client.EndpointUrl = flickr.API_ENDPOINT

// 	client.Args.Set("method", "flickr.contacts.getList")
// 	// client.Args.Set("brand", "nikon")

// 	client.OAuthSign()
// 	response := &GetContactListResponse{}
// 	err := flickr.DoGet(client, response)

// 	if err != nil {
// 		fmt.Printf("Error: %s", err)
// 		// } else {
// 		// 	fmt.Println("Api response:", response.Extra)
// 	}

// 	return response, nil
// }
