package model

type Fire struct {
	Loc
	Dangerousness int64
}

type FireResponse struct {
	Loc
	Dangerousness string
}

var (
	firedb = []Fire{
		{Loc{}, 1},
		{Loc{}, 1},
		{Loc{}, 3},
		{Loc{}, 2},
	}
	dangerousnessTranlate = map[int64]string{
		1: "Periculosidade Media",
		2: "Periculosidade Alta",
		3: "Evacuação imediata",
	}
)

type FireService interface{}

type firemodel struct{}

func Newfireservice() *firemodel {
	return &firemodel{}
}

func (fm *firemodel) List() (fr []FireResponse) {
	for _, v := range firedb {
		fr = append(fr, FireResponse{v.Loc, dangerousnessTranlate[v.Dangerousness]})
	}
	return
}
