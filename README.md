# Piccolo Internship Planner
[![CodeQL](https://github.com/nlamot/piccolo/actions/workflows/codeql-analysis.yml/badge.svg?branch=master)](https://github.com/nlamot/piccolo/actions/workflows/codeql-analysis.yml)

The purpose of the piccolo interenship planner is 
* to provide a way to fit student teachers into the existing schedule of a high school, taking into account
  * the kind of interneship required by the students
  * the requirements and preferences of the teachers in the school
  * ...
* to learn about building genetic algorithms based on cloud functions :)


## Repo structure
Based on https://medium.com/yakka/cloud-functions-in-go-94c1014a6fe4

## Lessons Learned

### Learning Cloud functions
If you want to set IAM rights on your cloud functions, give the Cloud Build Service Account the "cloudfunctions.admin" role. By default in the Cloud Build IAM overview, the system will give the role "cloudfunctions.developer" which doesn't have the necessary rights to manage IAM policies for cloud functions. Go to the general IAM section of GCP and add the role there.

### Working with the Google Cloud Emulators for Google service dependencies
We're working in a gcp functions environment with functions that depend on other google cloud managed services like firestore. In order to setup a test suite for this code, we can make use of the gcloud emulators functionality.

How can we do this?
* Install the gcloud CLI & the beta components
* Install a Java JRE
* Run `gcloud beta emulators firestore start --host-port=localhost` to verify everything is working.

### Generate mocks for all interfaces
We automatically generate all mocks that are used in unit tests, by using mockery. Use the following command to generate all mocks:

```bash
mockery --all --recursive
```
