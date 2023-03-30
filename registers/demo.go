package registers

import "errors"

type Demo struct {
}

func (d *Demo) Register(info RegisterInfo) {
	//TODO implement me
	panic("implement me")
}

func (d *Demo) RegisterClient(info RegisterInfo) {
	//TODO implement me
	panic("implement me")
}

func (d *Demo) GetIp(serverName string) (info RegisterInfo, err error) {
	if "user-service" == serverName {
		return RegisterInfo{Ip: "localhost", Port: 8089}, nil
	}
	return RegisterInfo{}, errors.New("未找到")
}

func (d *Demo) Close() {
	//TODO implement me
	panic("implement me")
}

func (d *Demo) RegisterInfo() RegisterInfo {
	//TODO implement me
	panic("implement me")
}
