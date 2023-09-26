package main

type User struct {
	BaseModel

	// Unique id of the user.
	// In the future, for multiple auth providers this could be:
	// oidc://<ISSUER>/<JWT-Profile.sub>,
	// ldap://<domain>/<GUID>
	ID string `db:"id"`
}
