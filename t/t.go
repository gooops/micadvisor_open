package main

import (
	"log"
	"strings"

	"github.com/gooops/micadvisor_open/dockerinspect"
)

var jsonstr = `[
{
    "Id": "d7635149cc889265e29d4793d8152103b318bd1da6c9fd7d22940bcc23ac2925",
    "Created": "2016-08-24T09:38:39.366600989Z",
    "Path": "docker-entrypoint.sh",
    "Args": [
        "redis-server"
    ],
    "State": {
        "Status": "exited",
        "Running": false,
        "Paused": false,
        "Restarting": false,
        "OOMKilled": false,
        "Dead": false,
        "Pid": 0,
        "ExitCode": 0,
        "Error": "",
        "StartedAt": "2016-08-24T09:38:41.773875552Z",
        "FinishedAt": "2016-08-24T11:49:17.554332885Z"
    },
    "Image": "dc5a4c47b1c4434cdd5b89dfc5f6b4a73fd31b821fa82778499cbafbb89c9ef2",
    "ResolvConfPath": "/var/lib/docker/containers/8b5e5657813daa7af115945668422f608700b9cca9ed9e41730ff0ed1104ad76/resolv.conf",
    "HostnamePath": "/var/lib/docker/containers/8b5e5657813daa7af115945668422f608700b9cca9ed9e41730ff0ed1104ad76/hostname",
    "HostsPath": "/var/lib/kubelet/pods/44f773d7-69dd-11e6-9dd9-5254005ccf75/etc-hosts",
    "LogPath": "/var/lib/docker/containers/d7635149cc889265e29d4793d8152103b318bd1da6c9fd7d22940bcc23ac2925/d7635149cc889265e29d4793d8152103b318bd1da6c9fd7d22940bcc23ac2925-json.log",
    "Name": "/k8s_redis.b4ad2849_nginx_default_44f773d7-69dd-11e6-9dd9-5254005ccf75_6becb094",
    "RestartCount": 0,
    "Driver": "devicemapper",
    "ExecDriver": "native-0.2",
    "MountLabel": "",
    "ProcessLabel": "",
    "AppArmorProfile": "",
    "ExecIDs": null,
    "HostConfig": {
        "Binds": [
            "/var/lib/kubelet/pods/44f773d7-69dd-11e6-9dd9-5254005ccf75/volumes/kubernetes.io~secret/default-token-3kh1r:/var/run/secrets/kubernetes.io/serviceaccount:ro",
            "/var/lib/kubelet/pods/44f773d7-69dd-11e6-9dd9-5254005ccf75/etc-hosts:/etc/hosts",
            "/var/lib/kubelet/pods/44f773d7-69dd-11e6-9dd9-5254005ccf75/containers/redis/6becb094:/dev/termination-log"
        ],
        "ContainerIDFile": "",
        "LxcConf": null,
        "Memory": 0,
        "MemoryReservation": 0,
        "MemorySwap": -1,
        "KernelMemory": 0,
        "CpuShares": 2,
        "CpuPeriod": 0,
        "CpusetCpus": "",
        "CpusetMems": "",
        "CpuQuota": 0,
        "BlkioWeight": 0,
        "OomKillDisable": false,
        "MemorySwappiness": null,
        "Privileged": false,
        "PortBindings": null,
        "Links": null,
        "PublishAllPorts": false,
        "Dns": [],
        "DnsOptions": [],
        "DnsSearch": [],
        "ExtraHosts": null,
        "VolumesFrom": null,
        "Devices": null,
        "NetworkMode": "container:8b5e5657813daa7af115945668422f608700b9cca9ed9e41730ff0ed1104ad76",
        "IpcMode": "container:8b5e5657813daa7af115945668422f608700b9cca9ed9e41730ff0ed1104ad76",
        "PidMode": "",
        "UTSMode": "",
        "CapAdd": null,
        "CapDrop": null,
        "GroupAdd": null,
        "RestartPolicy": {
            "Name": "",
            "MaximumRetryCount": 0
        },
        "SecurityOpt": null,
        "ReadonlyRootfs": false,
        "Ulimits": null,
        "Sysctls": null,
        "LogConfig": {
            "Type": "json-file",
            "Config": {}
        },
        "CgroupParent": "",
        "ConsoleSize": [
            0,
            0
        ],
        "VolumeDriver": "",
        "ShmSize": 67108864
    },
    "GraphDriver": {
        "Name": "devicemapper",
        "Data": {
            "DeviceId": "1454",
            "DeviceName": "docker-253:2-1048353-d7635149cc889265e29d4793d8152103b318bd1da6c9fd7d22940bcc23ac2925",
            "DeviceSize": "107374182400"
        }
    },
    "Mounts": [
        {
            "Name": "18968fe6f69c512d4a878be725f8ef683b3a94961990dd1d8dba8d2676ed2a45",
            "Source": "/var/lib/docker/volumes/18968fe6f69c512d4a878be725f8ef683b3a94961990dd1d8dba8d2676ed2a45/_data",
            "Destination": "/data",
            "Driver": "local",
            "Mode": "",
            "RW": true
        },
        {
            "Source": "/var/lib/kubelet/pods/44f773d7-69dd-11e6-9dd9-5254005ccf75/containers/redis/6becb094",
            "Destination": "/dev/termination-log",
            "Mode": "",
            "RW": true
        },
        {
            "Source": "/var/lib/kubelet/pods/44f773d7-69dd-11e6-9dd9-5254005ccf75/etc-hosts",
            "Destination": "/etc/hosts",
            "Mode": "",
            "RW": true
        },
        {
            "Source": "/var/lib/kubelet/pods/44f773d7-69dd-11e6-9dd9-5254005ccf75/volumes/kubernetes.io~secret/default-token-3kh1r",
            "Destination": "/var/run/secrets/kubernetes.io/serviceaccount",
            "Mode": "ro",
            "RW": false
        }
    ],
    "Config": {
        "Hostname": "nginx",
        "Domainname": "",
        "User": "",
        "AttachStdin": false,
        "AttachStdout": false,
        "AttachStderr": false,
        "ExposedPorts": {
            "6379/tcp": {}
        },
        "Tty": false,
        "OpenStdin": false,
        "StdinOnce": false,
        "Env": [
            "FRONTEND_PORT=tcp://10.254.14.125:80",
            "FRONTEND_PORT_80_TCP_PROTO=tcp",
            "MY_INT_SERVICE_PORT_INT_PORT=82",
            "MY_QB_PORT_88_TCP_PROTO=tcp",
            "MY_QB_PORT_88_TCP=tcp://10.254.230.140:88",
            "MY_QB_PORT_88_TCP_ADDR=10.254.230.140",
            "FRONTEND_PORT_80_TCP=tcp://10.254.14.125:80",
            "FRONTEND_PORT_80_TCP_ADDR=10.254.14.125",
            "REDIS_PORT=tcp://10.254.63.9:6005",
            "REDIS_LIEPIN_PORT_6006_TCP_ADDR=10.254.125.124",
            "REDIS_LIEPIN_PORT_6006_TCP_PORT=6006",
            "MY_QB_SERVICE_PORT=99",
            "MY_QB_SERVICE_PORT_INT_PORT=88",
            "KUBERNETES_SERVICE_PORT=443",
            "KUBERNETES_PORT_443_TCP_PORT=443",
            "REDIS_SERVICE_HOST=10.254.63.9",
            "MY_QB_SERVICE_HOST=10.254.230.140",
            "KUBERNETES_SERVICE_HOST=10.254.0.1",
            "LANLUHU_PORT=tcp://10.254.164.220:89",
            "LANLUHU_PORT_89_TCP_PORT=89",
            "REDIS_SERVICE_PORT=6005",
            "LANLUHU_SERVICE_PORT_WEB_PORT=89",
            "KUBERNETES_PORT_443_TCP=tcp://10.254.0.1:443",
            "REDIS_LIEPIN_SERVICE_PORT=6006",
            "MY_QB_PORT_99_TCP=tcp://10.254.230.140:99",
            "MY_INT_PORT_82_TCP=tcp://10.254.124.184:82",
            "REDIS_PORT_6005_TCP=tcp://10.254.63.9:6005",
            "KUBERNETES_PORT=tcp://10.254.0.1:443",
            "KUBERNETES_PORT_443_TCP_ADDR=10.254.0.1",
            "MY_INT_SERVICE_PORT=82",
            "LANLUHU_SERVICE_HOST=10.254.164.220",
            "MY_QB_PORT=tcp://10.254.230.140:99",
            "MY_QB_PORT_88_TCP_PORT=88",
            "FRONTEND_PORT_80_TCP_PORT=80",
            "MY_INT_SERVICE_HOST=10.254.124.184",
            "LANLUHU_SERVICE_PORT=89",
            "MY_QB_SERVICE_PORT_WEB_PORT=99",
            "KUBERNETES_SERVICE_PORT_HTTPS=443",
            "MY_QB_PORT_99_TCP_ADDR=10.254.230.140",
            "MY_INT_PORT_82_TCP_ADDR=10.254.124.184",
            "REDIS_PORT_6005_TCP_ADDR=10.254.63.9",
            "FRONTEND_SERVICE_PORT=80",
            "MY_INT_PORT=tcp://10.254.124.184:82",
            "LANLUHU_PORT_89_TCP_PROTO=tcp",
            "REDIS_LIEPIN_PORT_6006_TCP=tcp://10.254.125.124:6006",
            "FRONTEND_SERVICE_HOST=10.254.14.125",
            "LANLUHU_PORT_89_TCP=tcp://10.254.164.220:89",
            "LANLUHU_PORT_89_TCP_ADDR=10.254.164.220",
            "REDIS_LIEPIN_PORT_6006_TCP_PROTO=tcp",
            "REDIS_LIEPIN_SERVICE_HOST=10.254.125.124",
            "MY_QB_PORT_99_TCP_PROTO=tcp",
            "MY_INT_PORT_82_TCP_PROTO=tcp",
            "REDIS_PORT_6005_TCP_PROTO=tcp",
            "REDIS_LIEPIN_PORT=tcp://10.254.125.124:6006",
            "MY_QB_PORT_99_TCP_PORT=99",
            "KUBERNETES_PORT_443_TCP_PROTO=tcp",
            "MY_INT_PORT_82_TCP_PORT=82",
            "REDIS_PORT_6005_TCP_PORT=6005",
            "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
            "GOSU_VERSION=1.7",
            "REDIS_VERSION=3.2.3",
            "REDIS_DOWNLOAD_URL=http://download.redis.io/releases/redis-3.2.3.tar.gz",
            "REDIS_DOWNLOAD_SHA1=92d6d93ef2efc91e595c8bf578bf72baff397507"
        ],
        "Cmd": [
            "redis-server"
        ],
        "Image": "redis",
        "Volumes": {
            "/data": {}
        },
        "WorkingDir": "/data",
        "Entrypoint": [
            "docker-entrypoint.sh"
        ],
        "OnBuild": null,
        "Labels": {
            "io.kubernetes.container.hash": "b4ad2849",
            "io.kubernetes.container.name": "redis",
            "io.kubernetes.container.restartCount": "0",
            "io.kubernetes.container.terminationMessagePath": "/dev/termination-log",
            "io.kubernetes.pod.name": "nginx",
            "io.kubernetes.pod.namespace": "default",
            "io.kubernetes.pod.terminationGracePeriod": "30",
            "io.kubernetes.pod.uid": "44f773d7-69dd-11e6-9dd9-5254005ccf75"
        }
    },
    "NetworkSettings": {
        "Bridge": "",
        "SandboxID": "",
        "HairpinMode": false,
        "LinkLocalIPv6Address": "",
        "LinkLocalIPv6PrefixLen": 0,
        "Ports": null,
        "SandboxKey": "",
        "SecondaryIPAddresses": null,
        "SecondaryIPv6Addresses": null,
        "EndpointID": "",
        "Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "IPAddress": "",
        "IPPrefixLen": 0,
        "IPv6Gateway": "",
        "MacAddress": "",
        "Networks": null
    }
}
]`

func main() {
	// a, err := mxj.NewMapJson([]byte(str))
	// log.Println(a, err)
	inspect := dockerinspect.Inspect{}
	err := inspect.UnmarshalJSON([]byte(strings.Trim(strings.Trim(dockerData, "["), "]")))
	if err != nil {
		log.Println(err)
	}
	if value, ok := inspect.Config.Labels.(map[string]interface{}); ok {
		var tags []string
		for k, v := range value {
			if s, ok := v.(string); ok {
				tags = append(tags[:], k+"="+s)
			}
		}
		return strings.Join(tags, ",")
	}
	// log.Println(a, err)
}
