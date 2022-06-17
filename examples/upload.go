package examples

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	v1 "github.com/ScienceObjectsDB/go-api/sciobjsdb/api/storage/models/v1"
	v1Storage "github.com/ScienceObjectsDB/go-api/sciobjsdb/api/storage/services/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func api_key_example() {
	projectID := "<id>"

	host := "host"
	port := 443

	var tlsConf tls.Config
	credentials := credentials.NewTLS(&tlsConf)
	dialOptions := grpc.WithTransportCredentials(credentials)

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", host, port), dialOptions)
	if err != nil {
		log.Fatalln(err.Error())
	}

	datasetClient := v1Storage.NewDatasetServiceClient(conn)

	mdMap := make(map[string]string)
	mdMap["API_TOKEN"] = "token"
	tokenMetadata := metadata.New(mdMap)

	outContext := metadata.NewOutgoingContext(context.Background(), tokenMetadata)

	dataset, _ := datasetClient.CreateDataset(outContext, &v1Storage.CreateDatasetRequest{
		Name:        "testdataset",
		Description: "A test dataset",
		ProjectId:   projectID,
	})

	println(dataset.Id)

}

func upload() {
	projectID := "<id>"

	var conn grpc.ClientConnInterface

	//This examples assumes that an api_token is used. Therefor project creation is handled via Website
	//projectClient := v1Storage.NewProjectServiceClient(conn)

	//Create the clients for the individual storage components
	datasetClient := v1Storage.NewDatasetServiceClient(conn)
	objectGroupClient := v1Storage.NewDatasetObjectsServiceClient(conn)
	loadClient := v1Storage.NewObjectLoadServiceClient(conn)

	// After project creation, a dataset has to be created.
	dataset, _ := datasetClient.CreateDataset(context.Background(), &v1Storage.CreateDatasetRequest{
		Name:        "testdataset",
		Description: "A test dataset",
		ProjectId:   projectID,
	})

	object1, _ := objectGroupClient.CreateObject(context.Background(), &v1Storage.CreateObjectRequest{
		Filename:  "test",
		Filetype:  "bin",
		DatasetId: dataset.GetId(),
		Labels: []*v1.Label{&v1.Label{
			Key:   "testkey",
			Value: "testvalue",
		}},
	})

	object2, _ := objectGroupClient.CreateObject(context.Background(), &v1Storage.CreateObjectRequest{
		Filename:  "test",
		Filetype:  "meta",
		DatasetId: dataset.GetId(),
		Labels: []*v1.Label{&v1.Label{
			Key:   "testkey",
			Value: "testvalue",
		}},
	})

	link := object1.UploadLink
	putRequest, _ := http.NewRequest("PUT", link, bytes.NewBufferString("abcdefghi"))
	http.DefaultClient.Do(putRequest)

	objectGroupClient.FinishObjectUpload(context.Background(), &v1Storage.FinishObjectUploadRequest{
		Id: object1.Id,
	})

	link2 := object2.UploadLink
	putRequest2, _ := http.NewRequest("PUT", link2, bytes.NewBufferString("abcdefghi"))
	http.DefaultClient.Do(putRequest2)

	objectGroupClient.FinishObjectUpload(context.Background(), &v1Storage.FinishObjectUploadRequest{
		Id: object2.Id,
	})

	// Data objects can be stored in object groups that represent a set of closely related objects.
	// E.g. A data file with an associated index file that points to various features in the data file
	objectGroup, _ := objectGroupClient.CreateObjectGroup(context.Background(), &v1Storage.CreateObjectGroupRequest{
		DatasetId: dataset.GetId(),
		CreateRevisionRequest: &v1Storage.CreateObjectGroupRevisionRequest{
			Name:        "testobjectgroup",
			Description: "a test",
			Labels: []*v1.Label{
				&v1.Label{
					Key:   "testkey",
					Value: "testvalue",
				},
			},
			UpdateObjects: &v1Storage.UpdateObjectsRequests{
				AddObjects: []*v1Storage.AddObjectRequest{
					&v1Storage.AddObjectRequest{
						Id: object1.Id,
					},
				},
			},
			UpdateMetaObjects: &v1Storage.UpdateObjectsRequests{
				AddObjects: []*v1Storage.AddObjectRequest{
					&v1Storage.AddObjectRequest{
						Id: object2.Id,
					},
				},
			},
		},
	})

	objectGroup2, _ := objectGroupClient.GetObjectGroup(context.Background(), &v1Storage.GetObjectGroupRequest{
		Id: objectGroup.GetObjectGroupId(),
	})

	for _, object := range objectGroup2.GetObjectGroup().CurrentRevision.Objects {
		downloadLink, _ := loadClient.CreateDownloadLink(context.Background(), &v1Storage.CreateDownloadLinkRequest{
			Id: object.Id,
		})

		resp, _ := http.Get(downloadLink.GetDownloadLink())
		data, _ := ioutil.ReadAll(resp.Body)
		println(data)
	}
}
