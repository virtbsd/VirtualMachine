package VirtualMachine

type VirtualMachine interface {
    Start() error
    Stop() error
    Status() string
    CreateSnapshot(snapname string) error
    RestoreSnapshot(snapname string) error
    DeleteSnapshot(snapname string) error
    PrepareHostNetworking() error
    PrepareGuestNetworking() error
    NetworkingStatus() string
    GetPath() string
    IsOnline() bool
    Validate() error
    Persist() error
    Delete() error
    Archive(archivename string) error
}
