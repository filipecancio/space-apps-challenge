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
		{Loc{
			Lat: -14.53322388,
			Lng: -41.483234066,
		}, 1},
		{Loc{
			Lat: -14.53322388,
			Lng: -40.483234066,
		}, 1},
		{Loc{
			Lat: -13.53322388,
			Lng: -40.483234066,
		}, 3},
		{Loc{
			Lat: -14.53322388,
			Lng: -41.483234066,
		}, 2},
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
