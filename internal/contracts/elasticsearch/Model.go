package elasticsearch

type Model interface {
    GetId() string
    ToJson() string
}