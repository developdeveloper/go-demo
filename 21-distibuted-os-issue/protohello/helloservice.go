package protohello

type HelloService struct{}

func (hs *HelloService) Say(name String, reply *String) error {
	reply.Value = "hi, " + name.GetValue()
	return nil
}
