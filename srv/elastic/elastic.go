package elastic

import (
	"encoding/json"
	"errors"
	proto "github.com/Rakanixu/flag/srv/proto/flag"
	lib "github.com/mattbaird/elastigo/lib"
)

var (
	ErrNotFound = errors.New("flag not found")
	Index       = "flags"
	Type        = "flag"
	Hosts       []string
	conn        *lib.Conn
)

// Init ES connection
func Init() {
	conn = lib.NewConn()
	conn.SetHosts(Hosts)
}

// Create new flag into ES
func Create(cr *proto.CreateRequest) error {
	_, err := conn.Index(Index, Type, cr.Key, nil, cr)

	return err
}

// Read flag from ES
func Read(id string) (*proto.ReadResponse, error) {
	r, err := conn.Get(Index, Type, id, nil)
	if err != nil {
		return nil, err
	}

	var rr *proto.ReadResponse
	if err := json.Unmarshal(*r.Source, &rr); err != nil {
		return nil, err
	}

	return rr, nil
}

// Flip flag value
func Flip(id string) error {
	flag, err := Read(id)
	if err != nil {
		return err
	}

	flag.Value = !flag.Value

	_, err = conn.Index(Index, Type, id, nil, flag)

	return err
}

// Delete flag from ES
func Delete(id string) error {
	_, err := conn.Delete(Index, Type, id, nil)

	return err
}

// List all flags, no pagination
func List() (*proto.ListResponse, error) {
	out, err := lib.Search(Index).Type(Type).Size("10000").From("0").Search("*").Result(conn)
	if err != nil {
		return nil, err
	}

	var flags []*proto.ReadResponse
	for _, hit := range out.Hits.Hits {
		var cr *proto.ReadResponse
		if err = json.Unmarshal(*hit.Source, &cr); err != nil {
			return nil, err
		}

		flags = append(flags, cr)
	}

	lr := &proto.ListResponse{
		Result: flags,
	}

	return lr, nil
}
