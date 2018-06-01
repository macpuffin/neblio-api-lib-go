/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GetTransactionInfoResponseVout struct {

	// Value of the output in NEBL satoshi
	Value float32 `json:"value,omitempty"`

	// Output index
	N float32 `json:"n,omitempty"`

	ScriptPubKey *GetTransactionInfoResponsePreviousOutput `json:"scriptPubKey,omitempty"`

	Tokens []GetTransactionInfoResponseTokens `json:"tokens,omitempty"`

	// Whether this output has now been used
	Used bool `json:"used,omitempty"`

	// Blockheight of this transaction
	Blockheight float32 `json:"blockheight,omitempty"`

	// Blockheight this output was used in
	UsedBlockheight float32 `json:"usedBlockheight,omitempty"`

	// TXID this output was used in
	UsedTxid string `json:"usedTxid,omitempty"`
}
