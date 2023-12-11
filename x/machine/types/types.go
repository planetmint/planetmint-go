package types

type Entity struct {
	Domain string `json:"domain"`
}

type Contract struct {
	Entity       Entity `json:"entity"`
	IssuerPubkey string `json:"issuer_pubkey"` //nolint:tagliatelle // the format liquid network needs it
	MachineAddr  string `json:"machine_addr"`  //nolint:tagliatelle // the format liquid network needs it
	Name         string `json:"name"`
	Precision    uint64 `json:"precision"`
	Version      uint64 `json:"version"`
}
