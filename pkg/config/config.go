package config

import (
	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	PanoramaUrlField = field.StringField(
		"panorama-url",
		field.WithRequired(true),
		field.WithDescription("URL of Panorama instance"),
	)
	UsernameField = field.StringField(
		"username",
		field.WithRequired(true),
		field.WithDescription("Username"),
	)
	PasswordField = field.StringField(
		"password",
		field.WithRequired(true),
		field.WithIsSecret(true),
		field.WithDescription("Password"),
	)
	IgnoreBadCertificateField = field.BoolField(
		"ignore-bad-certificate",
		field.WithDefaultValue(false),
		field.WithDescription("Ignore bad certificate. This should be used only for testing purposes."),
	)

	// ConfigurationFields defines the external configuration required for the connector to run.
	ConfigurationFields = []field.SchemaField{
		PanoramaUrlField,
		UsernameField,
		PasswordField,
		IgnoreBadCertificateField,
	}

)

//go:generate go run ./gen
var Config = field.NewConfiguration(ConfigurationFields)
