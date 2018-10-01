package spec

type Marshalled interface {
	Namespace() string
	Name() string
	Version() string
	Hash() string
	Marshal() ([]byte, Links, error)
	Unmarshal([]byte, Links) error
}