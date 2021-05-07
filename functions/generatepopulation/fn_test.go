package generatepopulation

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
)

const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"

var request *http.Request
var response *httptest.ResponseRecorder

func TestGeneratePopulation(t *testing.T) {
	givenAValidFirestoreClient()
	givenAValidRequestToGeneratePopulations()
	
	whenTheFunctionIsTriggered()

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Hello, World!", response.Body.String())
}

func TestGeneratePopulationReturnsInternalServerErrorOnMissingFirestoreConnection(t *testing.T) {
	givenThereIsNoFirestoreConnection()
	givenAValidRequestToGeneratePopulations()

	whenTheFunctionIsTriggered()
	
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func givenThereIsNoFirestoreConnection() {
	client = nil
}

func givenAValidRequestToGeneratePopulations() {
	request = httptest.NewRequest("GET", "/", strings.NewReader(`{"name": ""}`))
	request.Header.Add("Content-Type", "application/json")
}

func givenAValidFirestoreClient() {
	var err error
	client, err = firestore.NewClient(context.Background(), "piccolo")
	if err != nil {
		log.Fatalf("firebase.NewClient err: %v", err)
	}
	defer client.Close()
}

func whenTheFunctionIsTriggered() {
	response = httptest.NewRecorder()
	GeneratePopulation(response, request)
}


// TestMain sets up the gcloud firestore emulator to test our integration with firestore.
// It will run before all tests.
func TestMain(m *testing.M) {
	// command to start firestore emulator
	cmd := exec.Command("gcloud", "beta", "emulators", "firestore", "start", "--host-port=localhost")

	// this makes it killable
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// we need to capture its output to know when it's started
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stderr.Close()

	// start her up!
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// ensure the process is killed when we're finished, even if an error occurs
	var result int
	defer func() {
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		os.Exit(result)
	}()

	// we're going to wait until it's running to start
	var wg sync.WaitGroup
	wg.Add(1)

	// by starting a separate go routine
	go func() {
		// reading it's output
		buf := make([]byte, 256, 256)
		for {
			n, err := stderr.Read(buf[:])
			if err != nil {
				// until it ends
				if err == io.EOF {
					break
				}
				log.Fatalf("reading stderr %v", err)
			}

			if n > 0 {
				d := string(buf[:n])

				// only required if we want to see the emulator output
				log.Printf("%s", d)

				// checking for the message that it's started
				if strings.Contains(d, "Dev App Server is now running") {
					wg.Done()
				}

				// and capturing the FIRESTORE_EMULATOR_HOST value to set
				pos := strings.Index(d, FirestoreEmulatorHost+"=")
				if pos > 0 {
					host := d[pos+len(FirestoreEmulatorHost)+1 : len(d)-1]
					os.Setenv(FirestoreEmulatorHost, host)
				}
			}
		}
	}()

	// wait until the running message has been received
	wg.Wait()

	// now it's running, we can run our unit tests
	result = m.Run()
}
