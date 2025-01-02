package server

import (
	"github.com/lyfeblocnetwork/lyfebloc-edge/consensus"
	consensusDev "github.com/lyfeblocnetwork/lyfebloc-edge/consensus/dev"
	consensusDummy "github.com/lyfeblocnetwork/lyfebloc-edge/consensus/dummy"
	consensusIBFT "github.com/lyfeblocnetwork/lyfebloc-edge/consensus/ibft"
	"github.com/lyfeblocnetwork/lyfebloc-edge/secrets"
	"github.com/lyfeblocnetwork/lyfebloc-edge/secrets/awsssm"
	"github.com/lyfeblocnetwork/lyfebloc-edge/secrets/gcpssm"
	"github.com/lyfeblocnetwork/lyfebloc-edge/secrets/hashicorpvault"
	"github.com/lyfeblocnetwork/lyfebloc-edge/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
