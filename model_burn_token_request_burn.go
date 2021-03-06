/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type BurnTokenRequestBurn struct {
	// Amount of tokens to burn
	Amount float32 `json:"amount,omitempty"`
	// Unique token id we are burning
	TokenId string `json:"tokenId,omitempty"`
}
