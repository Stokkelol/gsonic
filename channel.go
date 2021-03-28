package gsonic

type Channel string

const (
	Search  Channel = "search"
	Ingest  Channel = "ingest"
	Control Channel = "control"
)

type base struct {
}
