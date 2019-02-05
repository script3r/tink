// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

package hybrid

// ECIESHKDFRecipientKem represents a HKDF-based KEM (key encapsulation mechanism)
// for ECIES recipient.
type ECIESHKDFRecipientKem struct {
	recipientPrivateKey *ECPrivateKey
}

// decapsulate uses the KEM to generate a new HKDF-based key.
func (s *ECIESHKDFRecipientKem) decapsulate(kem []byte, hashAlg string, salt []byte, info []byte, keySize uint32, pointFormat string) ([]byte, error) {
	ephemeralPvt, err := GenerateECDHKeyPair(s.recipientPrivateKey.PublicKey.Curve)
	if err != nil {
		return nil, err
	}

	secret, err := ComputeSharedSecret(&ephemeralPvt.PublicKey.Point, s.recipientPrivateKey)
	if err != nil {
		return nil, err
	}
	i := append(kem, secret...)

	return computeHKDF(hashAlg, i, salt, info, keySize)
}