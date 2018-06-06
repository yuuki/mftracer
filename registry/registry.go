package registry

import (
	"log"
	"net"

	mackerel "github.com/mackerelio/mackerel-client-go"
)

// FindIPAddrsByDestServiceAndRoles find IP Addresses of hosts filtered by service and roles.
// TODO: handling status
func FindIPAddrsByDestServiceAndRoles(client *mackerel.Client, service string, roles []string) (
	map[string][]net.IP, error) {
	hosts, err := client.FindHosts(&mackerel.FindHostsParam{
		Service: service,
		Roles:   roles,
	})
	if err != nil {
		return nil, err
	}
	ipaddrsByRole := make(map[string][]net.IP, len(roles))
	for _, role := range roles {
		ipaddrsByRole[service+":"+role] = make([]net.IP, 0)
	}
	for _, host := range hosts {
		for _, roleFullname := range host.GetRoleFullnames() {
			ipaddrs := ipaddrsByRole[roleFullname]
			for _, iface := range host.Interfaces {
				for _, ipv4 := range iface.IPv4Addresses {
					ipaddr := net.ParseIP(ipv4)
					if ipaddr == nil {
						log.Printf("couldn't parse %s as an IPv4 address", ipv4)
					}
					ipaddrs = append(ipaddrs, ipaddr)
				}
			}
			ipaddrsByRole[roleFullname] = ipaddrs
		}
	}
	return ipaddrsByRole, nil
}
