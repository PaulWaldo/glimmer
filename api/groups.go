package api

import (
	"github.com/your/flickr"
)

type GetGroupPhotosResponse struct {
	flickr.BasicResponse
	Photos struct {
		Page    int `xml:"page,attr"`
		Pages   int `xml:"pages,attr"`
		PerPage int `xml:"perpage,attr"`
		Total   int `xml:"total,attr"`
		Photo   []struct {
			ID       string `xml:"id,attr"`
			Owner    string `xml:"owner,attr"`
			Secret   string `xml:"secret,attr"`
			Server   string `xml:"server,attr"`
			Farm     string `xml:"farm,attr"`
			Title    string `xml:"title,attr"`
			IsPublic int    `xml:"ispublic,attr"`
			IsFriend int    `xml:"isfriend,attr"`
			IsFamily int    `xml:"isfamily,attr"`
			Description string `xml:"description"`
			Dates struct {
				Posted          int `xml:"posted,attr"`
				Taken           string `xml:"taken,attr"`
				TakenGranularity int `xml:"takengranularity,attr"`
			} `xml:"dates"`
			Views    int `xml:"views"`
			Comments int `xml:"comments"`
		} `xml:"photo"`
	} `xml:"photos"`
}

func GetGroupPhotos(client *flickr.FlickrClient, groupID string, params map[string]string) (*GetGroupPhotosResponse, error) {
    client.Init()
    client.Args.Set("method", "flickr.groups.pools.getPhotos")
    client.Args.Set("group_id", groupID)
    for k, v := range params {
        client.Args.Set(k, v)
    }
    client.ApiSign()

    var response GetGroupPhotosResponse
    err := flickr.DoPost(client, &response)
    if err != nil {
        return nil, err
    }

    return &response, nil
}
