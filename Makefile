clean:
	rm -f bin/*

generate:
	go generate ./...

test: generate
	go test ./... -cover

function-population-generate:
	gcloud functions deploy \
		--runtime=go116 \
    	--trigger-http \
    	--source=. \
     	--entry-point=GeneratePopulation \
    	--region=europe-west1 \
    	--allow-unauthenticated \
		populationgenerate

function-roster-import:
	gcloud functions deploy \
		--runtime=go116 \
    	--trigger-resource rosters \
		--trigger-event google.storage.object.finalize \
    	--source=. \
     	--entry-point=ImportRoster \
    	--region=europe-west1 \
		rosterimport