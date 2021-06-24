clean:
	rm -f bin/*

generate:
	go generate ./...

test: generate
	go test ./... -cover

function-generate-population:
	gcloud functions deploy \
		--runtime=go113 \
    	--trigger-http \
    	--source=. \
     	--entry-point=GeneratePopulation \
    	--region=europe-west1 \
    	--allow-unauthenticated \
		generatepopulation