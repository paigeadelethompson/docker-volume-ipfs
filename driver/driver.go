package driver

import (
	"fmt"
	"sync"

	"github.com/docker/go-plugins-helpers/volume"
	"github.com/paigeadelethompson/docker-volume-ipfs/fuse"
	"github.com/paigeadelethompson/docker-volume-ipfs/kubo"
)

type IPFSVolumePlugin struct {
	mountPoint string
	volumes    []string
	m          *sync.Mutex
	fuse       *fuse.FUSEController
	ipfs       *kubo.KuboController
}

func New(ipfsRootMountPoint string, f *fuse.FUSEController, k *kubo.KuboController) IPFSVolumePlugin {
	d := IPFSVolumePlugin{
		mountPoint: ipfsRootMountPoint,
		volumes:    make([]string, 0),
		m:          &sync.Mutex{},
		fuse:       f,
		ipfs:       k,
	}
	return d
}

func (p *IPFSVolumePlugin) Create(req *volume.CreateRequest) error {
	p.volumes = append(p.volumes, req.Name)
	return nil
}

func (p *IPFSVolumePlugin) Get(req *volume.GetRequest) (*volume.GetResponse, error) {
	for _, v := range p.volumes {
		if v == req.Name {
			return &volume.GetResponse{Volume: &volume.Volume{Name: v}}, nil
		}
	}
	return &volume.GetResponse{}, fmt.Errorf("no such volume")
}

func (p *IPFSVolumePlugin) List() (*volume.ListResponse, error) {
	var vols []*volume.Volume
	for _, v := range p.volumes {
		vols = append(vols, &volume.Volume{
			Name: v,
		})
	}
	return &volume.ListResponse{
		Volumes: vols,
	}, nil
}

func (p *IPFSVolumePlugin) Remove(req *volume.RemoveRequest) error {
	for i, v := range p.volumes {
		if v == req.Name {
			p.volumes = append(p.volumes[:i], p.volumes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no such volume")
}

func (p *IPFSVolumePlugin) Path(req *volume.PathRequest) (*volume.PathResponse, error) {
	for _, v := range p.volumes {
		if v == req.Name {
			return &volume.PathResponse{}, nil
		}
	}
	return &volume.PathResponse{}, fmt.Errorf("no such volume")
}

func (p *IPFSVolumePlugin) Mount(req *volume.MountRequest) (*volume.MountResponse, error) {
	for _, v := range p.volumes {
		if v == req.Name {
			return &volume.MountResponse{}, nil
		}
	}
	return &volume.MountResponse{}, fmt.Errorf("no such volume")
}

func (p *IPFSVolumePlugin) Unmount(req *volume.UnmountRequest) error {
	for _, v := range p.volumes {
		if v == req.Name {
			return nil
		}
	}
	return fmt.Errorf("no such volume")
}

func (p *IPFSVolumePlugin) Capabilities() *volume.CapabilitiesResponse {
	return &volume.CapabilitiesResponse{
		Capabilities: volume.Capability{
			Scope: "local",
		},
	}
}
