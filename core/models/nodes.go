package models

type NodeInfo struct {
	ID                string
	Datacenter        string
	Name              string
	Status            string
	StatusDescription string
}

type NodeStats struct {
	Memory           *HostMemoryStats
	CPU              []*HostCPUStats
	DiskStats        []*HostDiskStats
	Uptime           uint64
	CPUTicksConsumed float64
}

type HostMemoryStats struct {
	Total     uint64
	Available uint64
	Used      uint64
	Free      uint64
}

type HostCPUStats struct {
	CPU    string
	User   float64
	System float64
	Idle   float64
}

type HostDiskStats struct {
	Size              uint64
	Used              uint64
	Available         uint64
	UsedPercent       float64
	InodesUsedPercent float64
}
