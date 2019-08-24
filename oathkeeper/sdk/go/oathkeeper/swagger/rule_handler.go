/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import "encoding/json"

type RuleHandler struct {

	// Config contains the configuration for the handler. Please read the user guide for a complete list of each handler's available settings.
	Config json.RawMessage `json:"config,omitempty"`

	// Handler identifies the implementation which will be used to handle this specific request. Please read the user guide for a complete list of available handlers.
	Handler string `json:"handler,omitempty"`
}
