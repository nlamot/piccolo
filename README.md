# piccolo
Home to the Piccolo project

## Repo structure
Based on https://medium.com/yakka/cloud-functions-in-go-94c1014a6fe4

## Lessons Learned

### Learning Cloud functions
If you want to set IAM rights on your cloud functions, give the Cloud Build Service Account the "cloudfunctions.admin" role. By default in the Cloud Build IAM overview, the system will give the role "cloudfunctions.developer" which doesn't have
the necessary rights to manage IAM policies for cloud functions. Go to the general IAM section of GCP and add the role there.