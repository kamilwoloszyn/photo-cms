package auth

var (
	// Post method
	ResourceCreateIdentity = "/identity"
	//Delete Method: /identities/{id}
	ResourceDeleteIdentity = "/identities/"
	//Get Method: /identities/{id}
	ResourceGetIdentity = "/identities/"
	//Put Method: /identities/{id}
	ResourceUpdateIdentity = "/identities/"
	//Post Method
	ResourceRecoveryLink = "/recovery/link"
	//Get method
	ResourceIdentityList = "/identities"
	//Get method
	ResourceSoftwareVersion = "/version"
	// Get Method.
	// This endpoint returns a HTTP 200 status code when Ory Kratos is accepting incoming HTTP requests. This status does currently not include checks whether the database connection is working.
	ResourceServerHealth = "/health/alive"
	//Get Method
	ResourceServerDbStatus = "/health/ready"
)
