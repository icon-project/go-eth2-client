// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package capella_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/assert"

	"github.com/attestantio/go-eth2-client/spec/capella"
)

func TestLightClientFinalityUpdateJSON(t *testing.T) {
	const (
		format        = `{"attested_header":%s,"finalized_header":%s,"finality_branch":%s,"sync_aggregate":%s,"signature_slot":%s}`
		header        = `{"beacon":{"slot":"4943744","proposer_index":"222870","parent_root":"0x0c594acb2c7ec3564590fd2feb6724cfcf786faf51fe2a284154516c2903c153","state_root":"0x237962d02698b2f5f37f3a7c43dfae0e2fe28e103225237bc7f09938c8527eaa","body_root":"0xff42d5726526628ce27c4ca89172ccf5c562adbfec64c22d494b6f8bd03dc034"},"execution":{"parent_hash":"0x17f4eeae822cc81533016678413443b95e34517e67f12b4a3a92ff6b66f972ef","fee_recipient":"0x58e809c71e4885cb7b3f1d5c793ab04ed239d779","state_root":"0x3d6e230e6eceb8f3db582777b1500b8b31b9d268339e7b32bba8d6f1311b211d","receipts_root":"0xea760203509bdde017a506b12c825976d12b04db7bce9eca9e1ed007056a3f36","logs_bloom":"0x0c803a8d3c6642adee3185bd914c599317d96487831dabda82461f65700b2528781bdadf785664f9d8b11c4ee1139dfeb056125d2abd67e379cabc6d58f1c3ea304b97cf17fcd8a4c53f4dedeaa041acce062fc8fbc88ffc111577db4a936378749f2fd82b4bfcb880821dd5cbefee984bc1ad116096a64a44a2aac8a1791a7ad3a53d91c584ac69a8973daed6daee4432a198c9935fa0e5c2a4a6ca78b821a5b046e571a5c0961f469d40e429066755fec611afe25b560db07f989933556ce0cea4070ca47677b007b4b9857fc092625f82c84526737dc98e173e34fe6e4d0f1a400fd994298b7c2fa8187331c333c415f0499836ff0eed5c762bf570e67b44","prev_randao":"0x76ff751467270668df463600d26dba58297a986e649bac84ea856712d4779c00","block_number":"2983837628677007840","gas_limit":"6738255228996962210","gas_used":"5573520557770513197","timestamp":"1744720080366521389","extra_data":"0xc648","base_fee_per_gas":"88770397543877639215846057887940126737648744594802753726778414602657613619599","block_hash":"0x42c294e902bfc9884c1ce5fef156d4661bb8f0ff488bface37f18c3e7be64b0f","transactions_root":"0x8457d0eb7611a621e7a094059f087415ffcfc91714fc184a1f3c48db06b4d08b","withdrawals_root":"0x000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"},"execution_branch":["0x65af40980fe7dfc2f8587fd1d75044f8adcf8e0e8b142363f5bf3bce21e66bb5","0x26648104944ae0085548cea356ebdd0c5c4b73aa440bcaf0c7b2821325b28f66","0xf97cbc51dd5b8ffffb73783e6938e3eee934448eaa08c9f50e136cb00635cf9f","0xf2adfbbfc2a4e45f01f90752b069b5fcd136b89dfa473dacbea52f6fefc3936c"]}`
		branch        = `["0x65af40980fe7dfc2f8587fd1d75044f8adcf8e0e8b142363f5bf3bce21e66bb5","0x26648104944ae0085548cea356ebdd0c5c4b73aa440bcaf0c7b2821325b28f66","0xf97cbc51dd5b8ffffb73783e6938e3eee934448eaa08c9f50e136cb00635cf9f","0xf2adfbbfc2a4e45f01f90752b069b5fcd136b89dfa473dacbea52f6fefc3936c","0x8b32360624a233863dc23f40fa08618cf49651e1b556e21b2f992d12a6cd84c2","0x42ad9040275048b9209b72096dd0d3551e08962a04dfc0602dd1459c4a165367"]`
		syncAggregate = `{"sync_committee_bits":"0xfffffffffffbffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff","sync_committee_signature":"0x97e006ecbe9df2f082eb450e1c07ace045da0d4e367f453170bfb32911e72fc9f08237d348e99b3500531c8cba770fc119d844c22950c094d860cfa784ba237debe681e55875994a75c72689d9e289f72c8bb7559ae91b3788e5e769aee0705a"}`
		slot          = `"1234"`
	)

	tests := []struct {
		name  string
		input []byte
		err   string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "invalid JSON: json: cannot unmarshal array into Go value of type capella.lightClientFinalityUpdateJSON",
		},
		{
			name:  "AttestedHeaderMissing",
			input: []byte(fmt.Sprintf(`{"finalized_header":%s,"finality_branch":%s,"sync_aggregate":%s,"signature_slot":%s}`, header, branch, syncAggregate, slot)),
			err:   "attested_header missing",
		},
		{
			name:  "FinalizedHeaderMissing",
			input: []byte(fmt.Sprintf(`{"attested_header":%s,"finality_branch":%s,"sync_aggregate":%s,"signature_slot":%s}`, header, branch, syncAggregate, slot)),
			err:   "finalized_header missing",
		},
		{
			name:  "FinalityBranchMissing",
			input: []byte(fmt.Sprintf(`{"attested_header":%s,"finalized_header":%s,"sync_aggregate":%s,"signature_slot":%s}`, header, header, syncAggregate, slot)),
			err:   "finality_branch missing",
		},
		{
			name:  "FinalityBranchEmpty",
			input: []byte(fmt.Sprintf(format, header, header, "[]", syncAggregate, slot)),
			err:   "finality_branch length cannot be 0",
		},
		{
			name:  "FinalityBranchWrongType",
			input: []byte(fmt.Sprintf(format, header, header, "true", syncAggregate, slot)),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field lightClientFinalityUpdateJSON.finality_branch of type []string",
		},
		{
			name:  "FinalityBranchWrongValueType",
			input: []byte(fmt.Sprintf(format, header, header, "[123]", syncAggregate, slot)),
			err:   "invalid JSON: json: cannot unmarshal number into Go struct field lightClientFinalityUpdateJSON.finality_branch of type string",
		},
		{
			name:  "FinalityBranchWrongValueFormat",
			input: []byte(fmt.Sprintf(format, header, header, `["invalid"]`, syncAggregate, slot)),
			err:   "invalid value for finality_branch[0]: encoding/hex: invalid byte: U+0069 'i'",
		},
		{
			name:  "FinalityBranchWrongValueFormat",
			input: []byte(fmt.Sprintf(format, header, header, `["0x12acde"]`, syncAggregate, slot)),
			err:   "invalid length of finality_branch[0]",
		},
		{
			name:  "SyncAggregateMissing",
			input: []byte(fmt.Sprintf(`{"attested_header":%s,"finalized_header":%s,"finality_branch":%s,"signature_slot":%s}`, header, header, branch, slot)),
			err:   "sync_aggregate missing",
		},
		{
			name:  "SignatureSlotMissing",
			input: []byte(fmt.Sprintf(`{"attested_header":%s,"finalized_header":%s,"finality_branch":%s,"sync_aggregate":%s}`, header, header, branch, syncAggregate)),
			err:   "signature_slot missing",
		},
		{
			name:  "SignatureSlotWrongType",
			input: []byte(fmt.Sprintf(format, header, header, branch, syncAggregate, "true")),
			err:   "invalid JSON: json: cannot unmarshal bool into Go struct field lightClientFinalityUpdateJSON.signature_slot of type string",
		},
		{
			name:  "SignatureSlotInvalid",
			input: []byte(fmt.Sprintf(format, header, header, branch, syncAggregate, `"-1"`)),
			err:   "invalid value for slot: strconv.ParseUint: parsing \"-1\": invalid syntax",
		},
		{
			name:  "Good",
			input: []byte(fmt.Sprintf(format, header, header, branch, syncAggregate, slot)),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res capella.LightClientFinalityUpdate
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				assert.Equal(t, string(test.input), string(rt))
				assert.Equal(t, string(rt), res.String())
			}
		})
	}
}