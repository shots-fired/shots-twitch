package hooker

type streamerEncoding struct {
	name string
	id string
}

type (
	Hooker interface {
		AddStreamer(name string) error
		AddStreamers(names []string) []error
		RemoveStreamer(name)
	}

	hooker struct {
		clientID             string
		streamerEncodings []streamerEncoding
	}
)

func NewHooker(clientID string) Hooker {
	return hooker{
		clientID:          clientID,
		streamerEncodings: []streamerEncoding,
	}
}

func (h hooker)AddStreamers(names []string) []error{
	var errors []error
	for _, name := range names {
		errors = append(errors, h.AddStreamer(name))
	}
	return errors
}

func (h hooker)AddStreamer(name string) error{

}

func (h hooker) RemoveStreamer(interface{}) {
	panic("implement me")
}