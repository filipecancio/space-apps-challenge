package model

type Report struct {
	User        string  `json:"user"`
	Kind        int64   `json:"kind"`
	Cause       []int64 `json:"cause"`
	Description string  `json:"description"`
	Loc         `json:"loc"`
	Status      int64 `json:"status"`
}

type ReportResponse struct {
	Kind        string   `json:"kind"`
	Cause       []string `json:"cause"`
	Description string   `json:"description"`
	Loc         `json:"loc"`
	Status      string `json:"status"`
}

type reportmodel struct{}

type ReportSercide interface{}

var (
	reportsdb = []Report{
		{
			User:        "c67af4cc-6543-11ee-8c99-0242ac120002",
			Kind:        1,
			Cause:       []int64{1},
			Description: "asasa",
			Loc: Loc{
				Lat: 0,
				Lng: 0,
			},
			Status: 1,
		},
	}
	statusTranslate = map[int64]string{
		1: "Incendio relatado",
		2: "Problema resolvido",
		3: "Relato não confirmado",
	}
	kindTranslate = map[int64]string{
		1: "Incêndio",
		2: "Risco de Incêndio",
		3: "Danos Ambientais",
	}
	causeTranslate = map[int64]map[int64]string{
		1: {
			1: "Relato de Incêndio",
			2: "Incêndio se Alastrando",
			3: "Início de Incêndio",
		},
		2: {
			1: "Área de Risco de Incêndio por Conta da Seca",
			2: "Combustíveis perigosos",
			3: "Outros distúrbios",
		},
		3: {
			1: "Erosão do Solo",
			2: "Restauração da Vegetação",
			3: "Outros",
		},
	}
)

func Newreportservice() *reportmodel {
	return &reportmodel{}
}

func (rm *reportmodel) Create(r Report) int {
	reportsdb = append(reportsdb, r)
	return len(reportsdb) + 1
}

func (rm *reportmodel) List(user string) (l []ReportResponse) {
	l = make([]ReportResponse, 0)
	for _, r := range reportsdb {
		if user == r.User {
			causes := make([]string, 0)
			for _, c := range r.Cause {
				causes = append(causes, causeTranslate[r.Kind][c])
			}
			l = append(l, ReportResponse{
				Status:      statusTranslate[r.Status],
				Loc:         r.Loc,
				Kind:        kindTranslate[r.Kind],
				Cause:       causes,
				Description: r.Description,
			})
		}
	}
	return
}
