package enochecker

type InfoMessage struct {
	ServiceName     string `json:"serviceName"`
	FlagVariants    uint64 `json:"flagVariants"`
	NoiseVariants   uint64 `json:"noiseVariants"`
	HavocVariants   uint64 `json:"havocVariants"`
	ExploitVariants uint64 `json:"exploitVariants"`
}

type Result string

const (
	ResultOk      = Result("OK")
	ResultError   = Result("INTERNAL_ERROR")
	ResultMumble  = Result("MUMBLE")
	ResultOffline = Result("OFFLINE")
)

type ResultMessage struct {
	Result     Result  `json:"result"`
	Message    *string `json:"message"`
	AttackInfo string  `json:"attackInfo,omitempty"`
	Flag       string  `json:"flag,omitempty"`
}

func NewResultMessageOk() *ResultMessage {
	return &ResultMessage{
		Result: ResultOk,
	}
}

func NewResultMessageMumble(msg string) *ResultMessage {
	return &ResultMessage{
		Result:  ResultMumble,
		Message: &msg,
	}
}

func NewResultMessageOffline(msg string) *ResultMessage {
	return &ResultMessage{
		Result:  ResultOffline,
		Message: &msg,
	}
}

func NewResultMessageError(msg string) *ResultMessage {
	return &ResultMessage{
		Result:  ResultError,
		Message: &msg,
	}
}

type TaskMessageMethod string

const (
	TaskMessageMethodPutFlag  = TaskMessageMethod("putflag")
	TaskMessageMethodGetFlag  = TaskMessageMethod("getflag")
	TaskMessageMethodPutNoise = TaskMessageMethod("putnoise")
	TaskMessageMethodGetNoise = TaskMessageMethod("getnoise")
	TaskMessageMethodHavoc    = TaskMessageMethod("havoc")
	TaskMessageMethodExploit  = TaskMessageMethod("exploit")
)

type TaskMessage struct {
	TaskId         uint64            `json:"taskId"`
	Method         TaskMessageMethod `json:"method"`
	Address        string            `json:"address"`
	TeamId         uint64            `json:"teamId"`
	TeamName       string            `json:"teamName"`
	CurrentRoundId uint64            `json:"currentRoundId"`
	RelatedRoundId uint64            `json:"relatedRoundId"`
	Flag           string            `json:"flag"`
	VariantId      uint64            `json:"variantId"`
	Timeout        uint64            `json:"timeout"`
	RoundLength    uint64            `json:"roundLength"`
	TaskChainId    string            `json:"taskChainId"`
	FlagRegex      string            `json:"flagRegex"`
	FlagHash       string            `json:"flagHash"`
	AttackInfo     string            `json:"attackInfo"`
}
