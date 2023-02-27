package driver

import (
	"fmt"
	"sync"

	"github.com/docker/go-plugins-helpers/volume"
)

// IPFS represent the IPFS volume driver
type IPFS struct {
	mountPoint string
	volumes    map[string]string
	m          *sync.Mutex
}

// New create an IPFS driver
func New(ipfsMountPoint string) IPFS {
	d := IPFS{
		mountPoint: ipfsMountPoint,
		volumes:    make(map[string]string),
		m:          &sync.Mutex{},
	}
	return d
}

// Create implements /VolumeDriver.Create
func (d IPFS) Create(r volume.CreateRequest) volume.ErrorResponse {
	// fmt.Printf("Create %v\n", r)
	// d.m.Lock()
	// defer d.m.Unlock()

	// volumeName := r.Name

	// if _, ok := d.volumes[volumeName]; ok {
	// 	return volume.ErrorResponse{}
	// }

	// volumePath := filepath.Join(d.mountPoint, volumeName)

	// _, err := os.Lstat(volumePath)
	// if err != nil {
	// 	fmt.Println("Error", volumePath, err.Error())
	// 	return volume.ErrorResponse{Err: fmt.Sprintf("Error while looking for volumePath %s: %s", volumePath, err.Error())}
	// }

	// d.volumes[volumeName] = volumePath
	// return volume.ErrorResponse{}
	panic("Create")
}

// Path implements /VolumeDriver.Path
func (d IPFS) Path(r volume.PathRequest) volume.PathResponse {
	fmt.Printf("Path %v\n", r)
	fmt.Printf("%v", d.volumes)
	volumeName := r.Name

	if volumePath, ok := d.volumes[volumeName]; ok {
		return volume.PathResponse{
			Mountpoint: volumePath,
		}
	}

	return volume.PathResponse{
		Mountpoint: "",
	}
}

// Remove implements /VolumeDriver.Remove
func (d IPFS) Remove(r volume.RemoveRequest) volume.ErrorResponse {
	// fmt.Printf("Remove %v", r)

	// d.m.Lock()
	// defer d.m.Unlock()

	// volumeName := r.Name

	// if _, ok := d.volumes[volumeName]; ok {
	// 	delete(d.volumes, volumeName)
	// }

	// return volume.ErrorResponse{}
	panic("Remove")
}

// Mount implements /VolumeDriver.Mount
func (d IPFS) Mount(r volume.MountRequest) volume.MountResponse {
	fmt.Printf("Mount %v\n", r)
	volumeName := r.Name

	if volumePath, ok := d.volumes[volumeName]; ok {
		return volume.MountResponse{
			Mountpoint: volumePath,
		}
	}

	return volume.MountResponse{}
}

// Unmount implements /VolumeDriver.Mount
func (d IPFS) Unmount(r volume.UnmountRequest) volume.ErrorResponse {
	fmt.Printf("Unmount %v: nothing to do\n", r)
	return volume.ErrorResponse{}
}

// Get implements /VolumeDriver.Get
func (d IPFS) Get(r volume.GetRequest) volume.GetResponse {
	// fmt.Printf("Get %v\n", r)
	// volumeName := r.Name

	// if volumePath, ok := d.volumes[volumeName]; ok {
	// 	return volume.GetResponse{Volume: &volume.Volume{Name: volumeName, Mountpoint: volumePath}}
	// }

	// return volume.GetResponse{}
	panic("Get")
}

// List implements /VolumeDriver.List
func (d IPFS) List(r volume.GetRequest) volume.ListResponse {
	// fmt.Printf("List %v\n", r)

	// volumes := []*volume.Volume{}

	// for name, path := range d.volumes {
	// 	volumes = append(volumes, &volume.Volume{Name: name, Mountpoint: path})
	// }

	panic("List")
}

// Capabilities implements /VolumeDriver.Capabilities
func (d IPFS) Capabilities(r volume.GetRequest) volume.CapabilitiesResponse {
	// FIXME(vdemeester) handle capabilities better
	// return volume.Response{
	// 	Capabilities: volume.Capability{
	// 		Scope: "local",
	// 	},
	// }
	panic("Capabilities")
}
