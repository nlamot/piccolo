clean:
	rm -f bin/*

generate:
	go generate ./...

piccolo: clean generate
	go build \
		-o bin/piccolo \
		cmd/piccolo/*

test: generate
	go test ./... -cover

function-generate-population: generate
	gcloud functions deploy \
		--runtime=go113 \
    	--trigger-http \
    	--source=. \
     	--entry-point=GeneratePopulation \
    	--region=europe-west1 \
    	--allow-unauthenticated \
		generatepopulation