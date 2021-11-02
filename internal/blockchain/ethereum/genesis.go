// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
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

package ethereum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type IBFTGenesis struct {
	Config        *IBFTGenesisConfig `json:"config"`
	Nonce         string             `json:"nonce"`
	TimeStamp     string             `json:"timestamp"`
	GasLimit      string             `json:"gasLimit"`
	Difficulty    string             `json:"difficulty"`
	MixHash       string             `json:"mixHash"`
	ExtraData     string             `json:"extraData"`
	Coinbase      string             `json:"coinbase"`
	Alloc         map[string]*Alloc  `json:"alloc"`
	AccIngressSM  *AccIngressSM      `json:"0x0000000000000000000000000000000000008888,omitempty"`
	NodeIngressSM *AccIngressSM      `json:"0x0000000000000000000000000000000000009999,omitempty"`
}

type AccIngressSM struct {
	Balance string   `json:"balance"`
	Code    string   `json:"code"`
	Storage *Storage `json:"storage"`
}

type Storage struct {
	Field1 string `json:"0x0000000000000000000000000000000000000000000000000000000000000000"`
	Field2 string `json:"0x0000000000000000000000000000000000000000000000000000000000000001"`
	Field3 string `json:"0x0000000000000000000000000000000000000000000000000000000000000004"`
}

type IBFTGenesisConfig struct {
	ChainId                int    `json:"chainId"`
	ConstantinopleFixBlock int    `json:"constantinoplefixblock"`
	IBFT2                  *IBFT2 `json:"ibft2"`
}

type IBFT2 struct {
	BlockPeriodSeconds    int `json:"blockperiodseconds"`
	EpochLength           int `json:"epochlength"`
	RequestTimeoutSeconds int `json:"requesttimeoutseconds"`
}

type Genesis struct {
	Config     *GenesisConfig    `json:"config"`
	Nonce      string            `json:"nonce"`
	Timestamp  string            `json:"timestamp"`
	ExtraData  string            `json:"extraData"`
	GasLimit   string            `json:"gasLimit"`
	Difficulty string            `json:"difficulty"`
	MixHash    string            `json:"mixHash"`
	Coinbase   string            `json:"coinbase"`
	Alloc      map[string]*Alloc `json:"alloc"`
	Number     string            `json:"number"`
	GasUsed    string            `json:"gasUsed"`
	ParentHash string            `json:"parentHash"`
}

type GenesisConfig struct {
	ChainId             int           `json:"chainId"`
	HomesteadBlock      int           `json:"homesteadBlock"`
	Eip150Block         int           `json:"eip150Block"`
	Eip150Hash          string        `json:"eip150Hash"`
	Eip155Block         int           `json:"eip155Block"`
	Eip158Block         int           `json:"eip158Block"`
	ByzantiumBlock      int           `json:"byzantiumBlock"`
	ConstantinopleBlock int           `json:"constantinopleBlock"`
	PetersburgBlock     int           `json:"petersburgBlock"`
	IstanbulBlock       int           `json:"istanbulBlock"`
	Clique              *CliqueConfig `json:"clique"`
}

type CliqueConfig struct {
	Period int `json:"period"`
	Epoch  int `json:"epoch"`
}

type Alloc struct {
	Balance string   `json:"balance"`
	Code    string   `json:"code,omitempty"`
	Storage *Storage `json:"storage,omitempty"`
}

func CreateGenesisJson(addresses []string) *Genesis {

	extraData := "0x0000000000000000000000000000000000000000000000000000000000000000"
	alloc := make(map[string]*Alloc)

	for _, address := range addresses {
		alloc[address] = &Alloc{
			Balance: "0x200000000000000000000000000000000000000000000000000000000000000",
		}
		extraData = extraData + address
	}
	extraData = strings.ReplaceAll(fmt.Sprintf("%-236s", extraData), " ", "0")

	return &Genesis{
		Config: &GenesisConfig{
			ChainId:             2021,
			HomesteadBlock:      0,
			Eip150Block:         0,
			Eip150Hash:          "0x0000000000000000000000000000000000000000000000000000000000000000",
			Eip155Block:         0,
			Eip158Block:         0,
			ByzantiumBlock:      0,
			ConstantinopleBlock: 0,
			IstanbulBlock:       0,
			Clique: &CliqueConfig{
				Period: 0,
				Epoch:  30000,
			},
		},
		Nonce:      "0x0",
		Timestamp:  "0x60edb1c7",
		ExtraData:  extraData,
		GasLimit:   "0x47b760",
		Difficulty: "0x1",
		MixHash:    "0x0000000000000000000000000000000000000000000000000000000000000000",
		Coinbase:   "0x0000000000000000000000000000000000000000",
		Alloc:      alloc,
		Number:     "0x0",
		GasUsed:    "0x0",
		ParentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
	}
}

func (g *Genesis) WriteGenesisJson(filename string) error {
	genesisJsonBytes, _ := json.MarshalIndent(g, "", " ")
	if err := ioutil.WriteFile(filepath.Join(filename), genesisJsonBytes, 0755); err != nil {
		return err
	}
	return nil
}

func CreateIBFTGenesis(addresses []string) *IBFTGenesis {
	alloc := make(map[string]*Alloc)

	for _, address := range addresses {
		alloc[address] = &Alloc{
			Balance: "0x200000000000000000000000000000000000000000000000000000000000000",
		}

	}
	alloc["0x0000000000000000000000000000000000008888"] = &Alloc{
		Balance: "0",
		Code:    "0x608060405234801561001057600080fd5b506004361061009e5760003560e01c8063936421d511610066578063936421d5146101ca578063a43e04d8146102fb578063de8fa43114610341578063e001f8411461035f578063fe9fbb80146103c55761009e565b80630d2020dd146100a357806310d9042e1461011157806311601306146101705780631e7c27cb1461018e5780638aa10435146101ac575b600080fd5b6100cf600480360360208110156100b957600080fd5b8101908080359060200190929190505050610421565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101196104d6565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561015c578082015181840152602081019050610141565b505050509050019250505060405180910390f35b61017861052e565b6040518082815260200191505060405180910390f35b610196610534565b6040518082815260200191505060405180910390f35b6101b461053a565b6040518082815260200191505060405180910390f35b6102e1600480360360c08110156101e057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561025b57600080fd5b82018360208201111561026d57600080fd5b8035906020019184600183028401116401000000008311171561028f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610544565b604051808215151515815260200191505060405180910390f35b6103276004803603602081101561031157600080fd5b810190808035906020019092919050505061073f565b604051808215151515815260200191505060405180910390f35b610349610a1e565b6040518082815260200191505060405180910390f35b6103ab6004803603604081101561037557600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a2b565b604051808215151515815260200191505060405180910390f35b610407600480360360208110156103db57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610caf565b604051808215151515815260200191505060405180910390f35b60008060001b821161049b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e81525060200191505060405180910390fd5b6002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6060600380548060200260200160405190810160405280929190818152602001828054801561052457602002820191906000526020600020905b815481526020019060010190808311610510575b5050505050905090565b60005481565b60015481565b6000600554905090565b60008073ffffffffffffffffffffffffffffffffffffffff16610568600054610421565b73ffffffffffffffffffffffffffffffffffffffff16141561058d5760019050610735565b600260008054815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663936421d58888888888886040518763ffffffff1660e01b8152600401808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156106a857808201518184015260208101905061068d565b50505050905090810190601f1680156106d55780820380516001836020036101000a031916815260200191505b5097505050505050505060206040518083038186803b1580156106f757600080fd5b505afa15801561070b573d6000803e3d6000fd5b505050506040513d602081101561072157600080fd5b810190808051906020019092919050505090505b9695505050505050565b60008060001b82116107b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e81525060200191505060405180910390fd5b600060038054905011610817576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526047815260200180610e446047913960600191505060405180910390fd5b61082033610caf565b610875576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180610e19602b913960400191505060405180910390fd5b6000600460008481526020019081526020016000205490506000811180156108a257506003805490508111155b15610a135760038054905081146109105760006003600160038054905003815481106108ca57fe5b9060005260206000200154905080600360018403815481106108e857fe5b9060005260206000200181905550816004600083815260200190815260200160002081905550505b600380548061091b57fe5b600190038181906000526020600020016000905590556000600460008581526020019081526020016000208190555060006002600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af600084604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a16001915050610a19565b60009150505b919050565b6000600380549050905090565b60008060001b8311610aa5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e81525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610e8b6022913960400191505060405180910390fd5b610b3433610caf565b610b89576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180610e19602b913960400191505060405180910390fd5b600060046000858152602001908152602001600020541415610be8576003839080600181540180825580915050906001820390600052602060002001600090919290919091505560046000858152602001908152602001600020819055505b816002600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af8284604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a16001905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff1660026000600154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610d235760019050610e13565b60026000600154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe9fbb80836040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610dd557600080fd5b505afa158015610de9573d6000803e3d6000fd5b505050506040513d6020811015610dff57600080fd5b810190808051906020019092919050505090505b91905056fe4e6f7420617574686f72697a656420746f2075706461746520636f6e74726163742072656769737472792e4d7573742068617665206174206c65617374206f6e65207265676973746572656420636f6e747261637420746f20657865637574652064656c657465206f7065726174696f6e2e436f6e74726163742061646472657373206d757374206e6f74206265207a65726f2ea265627a7a7230582041609b4b53a670d9d29d1c024dd9467b05a85c59786466daf08dcc1f75f8f6be64736f6c63430005090032",
		Storage: &Storage{
			Field1: "0x72756c6573000000000000000000000000000000000000000000000000000000",
			Field2: "0x61646d696e697374726174696f6e000000000000000000000000000000000000",
			Field3: "0x0f4240",
		},
	}
	alloc["0x0000000000000000000000000000000000009999"] = &Alloc{
		Balance: "0",
		Code:    "0x608060405234801561001057600080fd5b50600436106100885760003560e01c8063a43e04d81161005b578063a43e04d814610196578063de8fa431146101dc578063e001f841146101fa578063fe9fbb801461026057610088565b80630d2020dd1461008d57806310d9042e146100fb578063116013061461015a5780631e7c27cb14610178575b600080fd5b6100b9600480360360208110156100a357600080fd5b81019080803590602001909291905050506102bc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610103610371565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561014657808201518184015260208101905061012b565b505050509050019250505060405180910390f35b6101626103c9565b6040518082815260200191505060405180910390f35b6101806103cf565b6040518082815260200191505060405180910390f35b6101c2600480360360208110156101ac57600080fd5b81019080803590602001909291905050506103d5565b604051808215151515815260200191505060405180910390f35b6101e46106b4565b6040518082815260200191505060405180910390f35b6102466004803603604081101561021057600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106c1565b604051808215151515815260200191505060405180910390f35b6102a26004803603602081101561027657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610945565b604051808215151515815260200191505060405180910390f35b60008060001b8211610336576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e81525060200191505060405180910390fd5b6002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b606060038054806020026020016040519081016040528092919081815260200182805480156103bf57602002820191906000526020600020905b8154815260200190600101908083116103ab575b5050505050905090565b60005481565b60015481565b60008060001b821161044f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e81525060200191505060405180910390fd5b6000600380549050116104ad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526047815260200180610ada6047913960600191505060405180910390fd5b6104b633610945565b61050b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180610aaf602b913960400191505060405180910390fd5b60006004600084815260200190815260200160002054905060008111801561053857506003805490508111155b156106a95760038054905081146105a657600060036001600380549050038154811061056057fe5b90600052602060002001549050806003600184038154811061057e57fe5b9060005260206000200181905550816004600083815260200190815260200160002081905550505b60038054806105b157fe5b600190038181906000526020600020016000905590556000600460008581526020019081526020016000208190555060006002600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af600084604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a160019150506106af565b60009150505b919050565b6000600380549050905090565b60008060001b831161073b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f436f6e7472616374206e616d65206d757374206e6f7420626520656d7074792e81525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156107c1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610b216022913960400191505060405180910390fd5b6107ca33610945565b61081f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180610aaf602b913960400191505060405180910390fd5b60006004600085815260200190815260200160002054141561087e576003839080600181540180825580915050906001820390600052602060002001600090919290919091505560046000858152602001908152602001600020819055505b816002600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fe3d908a1f6d2467f8e7c8198f30125843211345eedb763beb4cdfb7fe728a5af8284604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a16001905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff1660026000600154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156109b95760019050610aa9565b60026000600154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe9fbb80836040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610a6b57600080fd5b505afa158015610a7f573d6000803e3d6000fd5b505050506040513d6020811015610a9557600080fd5b810190808051906020019092919050505090505b91905056fe4e6f7420617574686f72697a656420746f2075706461746520636f6e74726163742072656769737472792e4d7573742068617665206174206c65617374206f6e65207265676973746572656420636f6e747261637420746f20657865637574652064656c657465206f7065726174696f6e2e436f6e74726163742061646472657373206d757374206e6f74206265207a65726f2ea265627a7a723058206703bdfb54a7a3eb61936f024bb43f91b3a8ce1448dc4d9593458137e30b983f64736f6c63430005090032",
		Storage: &Storage{
			Field1: "0x72756c6573000000000000000000000000000000000000000000000000000000",
			Field2: "0x61646d696e697374726174696f6e000000000000000000000000000000000000",
			Field3: "0x0f4240",
		},
	}
	return &IBFTGenesis{
		Config: &IBFTGenesisConfig{
			ChainId:                1337,
			ConstantinopleFixBlock: 0,
			IBFT2: &IBFT2{
				BlockPeriodSeconds:    1,
				EpochLength:           30000,
				RequestTimeoutSeconds: 10,
			},
		},
		Nonce:      "0x0",
		TimeStamp:  "0x58ee40ba",
		GasLimit:   "0xffffffff",
		Difficulty: "0x1",
		MixHash:    "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
		ExtraData:  "0xf87ea00000000000000000000000000000000000000000000000000000000000000000f854944592c8e45706cc08b8f44b11e43cba0cfc5892cb9406e23768a0f59cf365e18c2e0c89e151bcdedc7094c5327f96ee02d7bcbc1bf1236b8c15148971e1de94ab5e7f4061c605820d3744227eed91ff8e2c8908808400000000c0",
		Coinbase:   "0x0000000000000000000000000000000000000000",
		Alloc:      alloc,
	}
}

func (g *IBFTGenesis) WriteIBFTGenesisJson(filename string) error {
	genesisJsonBytes, _ := json.MarshalIndent(g, "", " ")
	if err := ioutil.WriteFile(filepath.Join(filename), genesisJsonBytes, 0755); err != nil {
		return err
	}
	return nil
}
