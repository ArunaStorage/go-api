package examples

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	v1Storage "github.com/ScienceObjectsDB/go-api/sciobjsdb/api/storage/services/v1"
	"google.golang.org/grpc"
)

func upload() {
	projectID := "<id>"

	var conn grpc.ClientConnInterface

	//This examples assumes that an api_token is used. Therefor project creation is handled via Website
	//projectClient := v1Storage.NewProjectServiceClient(conn)
	datasetClient := v1Storage.NewDatasetServiceClient(conn)
	objectGroupClient := v1Storage.NewDatasetObjectsServiceClient(conn)
	loadClient := v1Storage.NewObjectLoadServiceClient(conn)

	dataset, _ := datasetClient.CreateDataset(context.Background(), &v1Storage.CreateDatasetRequest{
		Name:        "testdataset",
		Description: "A test dataset",
		ProjectId:   projectID,
	})

	objectGroup, _ := objectGroupClient.CreateObjectGroup(context.Background(), &v1Storage.CreateObjectGroupRequest{
		Name:              "testobjectgroup",
		Description:       "A test objectgroup",
		DatasetId:         dataset.GetId(),
		IncludeObjectLink: true, // Returns presigned upload links for all objects. For large objects please use the multipart upload
		Objects: []*v1Storage.CreateObjectRequest{
			&v1Storage.CreateObjectRequest{
				Filename:   "foo.txt",
				Filetype:   "txt",
				ContentLen: 9,
			},
			&v1Storage.CreateObjectRequest{
				Filename:   "foo.bin",
				Filetype:   "bin",
				ContentLen: 9,
			},
		},
	})

	// The link order is the same as the request order in the createobjectgrouprequest
	for _, object := range objectGroup.ObjectLinks {
		link := object.Link
		putRequest, _ := http.NewRequest("PUT", link, bytes.NewBufferString("abcdefghi"))
		http.DefaultClient.Do(putRequest)

		//This is currently not required but will be used in the future to perform checks on metadata, execute additional tasks and so on...
		objectGroupClient.FinishObjectUpload(context.Background(), &v1Storage.FinishObjectUploadRequest{
			Id: object.ObjectId,
		})
	}

	objectGroup2, _ := objectGroupClient.GetObjectGroup(context.Background(), &v1Storage.GetObjectGroupRequest{
		Id: objectGroup.GetObjectGroupId(),
	})

	for _, object := range objectGroup2.GetObjectGroup().Objects {
		downloadLink, _ := loadClient.CreateDownloadLink(context.Background(), &v1Storage.CreateDownloadLinkRequest{
			Id: object.Id,
		})

		resp, _ := http.Get(downloadLink.GetDownloadLink())
		data, _ := ioutil.ReadAll(resp.Body)
		println(data)
	}
}