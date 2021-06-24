package planner

import (
	"context"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/functions/metadata"
	"go.uber.org/dig"
	"piccolo.com/planner/pkg/common/gcp/gcs"
	"piccolo.com/planner/pkg/planning/roster"
)

var rosterContainer *dig.Container
var rosterContainerOnce sync.Once

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

	rosterContainerOnce.Do(func() {
		rosterContainer = roster.NewContainer()
	})

	er := rosterContainer.Invoke(func(service roster.RosterService) {
		service.Create("orguuid", roster.Roster{})
	})

	if er != nil {
		panic(er)
	}
	return nil
}