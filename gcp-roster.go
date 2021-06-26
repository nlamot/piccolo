package planner

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/functions/metadata"
	"go.uber.org/dig"
	"piccolo.com/planner/pkg/common/gcp/gcs"
	"piccolo.com/planner/pkg/common/smartschool"
	"piccolo.com/planner/pkg/planning/roster"
)

var rosterContainer *dig.Container
var rosterContainerOnce sync.Once

var smartschoolContainer *dig.Container
var smartschoolContainerOnce sync.Once

// ImportRoster imports a roster to the firestore
func ImportRoster(ctx context.Context, e gcs.GCSEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Event ID: %v\n", meta.EventID)
	log.Printf("Event type: %v\n", meta.EventType)
	log.Printf("Bucket: %v\n", e.Bucket)
	log.Printf("File: %v\n", e.Name)
	log.Printf("Metageneration: %v\n", e.Metageneration)
	log.Printf("Created: %v\n", e.TimeCreated)
	log.Printf("Updated: %v\n", e.Updated)

	smartschoolContainerOnce.Do(func() {
		smartschoolContainer = smartschool.NewContainer()
	})

	var newRoster roster.Roster
	smartschoolError := smartschoolContainer.Invoke(func(manager smartschool.RosterManager){
		client, gcsError := gcs.ProvideStorageClient()
		if gcsError != nil {
			panic(gcsError)
		}
		upload, uploadError := client.Bucket(e.Bucket).Object(e.Name).NewReader(context.Background())
		if uploadError != nil {
			panic(uploadError)
		}
		var rosterError error
		csv := csv.NewReader(upload)
		csv.Comma = ';'
		newRoster, rosterError = manager.ImportCSV(csv)
		if rosterError != nil {
			fmt.Println(rosterError)
			panic(rosterError)
		}
	})

	if smartschoolError != nil {
		panic(smartschoolError)
	}

	rosterContainerOnce.Do(func() {
		rosterContainer = roster.NewContainer()
	})
	er := rosterContainer.Invoke(func(service roster.RosterService) {
		rosterId, rEr := service.Create("hfam", newRoster)
		if rEr != nil {
			panic(rEr)
		}
		fmt.Printf("ID of new Roster: %s", rosterId)
	})

	if er != nil {
		panic(er)
	}
	return nil
}