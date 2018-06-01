/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Object describing token rules
type IssueTokenRequestMetadataRules struct {

	Fees *IssueTokenRequestMetadataRulesFees `json:"fees,omitempty"`

	// Array of objects describing what addresses can hold the token
	Holders []IssueTokenRequestMetadataRulesHolders `json:"holders,omitempty"`

	Expiration *IssueTokenRequestMetadataRulesExpiration `json:"expiration,omitempty"`
}
