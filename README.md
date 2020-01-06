# emr-sample

Server for a toy electronic medical record system. This project is meant to be used by nobody.

Start the server by running `go run .` and send requests to the server.

```bash
# Add a patient
curl -d name="joe schmoe" -d dob=19901010 http://localhost:8080/patients

# Get a patient
curl http://localhost:8080/patients/0

# Edit a patient
curl -X PATCH -d name="joseph schmoe" -d dob=19901010 http://localhost:8080/patients/0

# Get all patients
curl http://localhost:8080/patients 
```
