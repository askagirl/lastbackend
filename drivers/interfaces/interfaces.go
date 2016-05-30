package interfaces

import "errors"

type ILog interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	SetDebugLevel()
}

type IStorage interface {
	Write(ILog, string, string) error
	Read(ILog, string) (string, error)
	ListAllFiles(ILog) (map[string]string, error)
	Delete(ILog, string) error
}

type IContainers interface {
	GetContainer(string) (Container, error)

	StartContainer(*Container) error
	StopContainer(*Container) error
	RestartContainer(*Container) error
	RemoveContainer(*Container) error

	PullImage(i Image) error
	BuildImage(opts BuildImageOptions) error

	ListImages() (map[string]Image, error)
	ListContainers() (map[string]Container, error)

	System() (*Node, error)
}

var ErrBucketNotFound error = errors.New("BUCKET_NOT_FOUND")