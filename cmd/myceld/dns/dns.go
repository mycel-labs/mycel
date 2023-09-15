package dns

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
	resolver "github.com/mycel-domain/mycel/x/resolver/types"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

type grpcService struct {
	grpcConn *grpc.ClientConn
}

func (s *grpcService) QueryDnsToMycelResolver(domain string, recordType string) (ip net.IP) {

	domain = strings.Trim(domain, ".")
	division := strings.Index(domain, ".")

	argName := domain[:division]
	argParent := domain[division+1:]

	queryClient := resolver.NewQueryClient(s.grpcConn)

	params := &resolver.QueryDnsRecordRequest{
		DomainName:    argName,
		DomainParent:  argParent,
		DnsRecordType: recordType,
	}

	res, err := queryClient.DnsRecord(context.Background(), params)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil
	}

	if found := res.Value != nil; found {

		value := res.Value.Value

		switch recordType {
		case "A", "AAAA":
			ip = net.ParseIP(value)
		// case "CNAME", "MX", "NS":
		// return value
		// case "TXT":
		// return []string{value}
		default:
			return nil
		}
	}
	log.Printf("Mycel: %s %s %s", domain, recordType, ip)
	return ip
}

func (s *grpcService) QueryDnsToDefaultResolver(domain string, recordType string) (ip net.IP) {
	// LookupIP
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil
	}

	if recordType == "A" {
		for _, i := range ips {
			if ipv4 := i.To4(); ipv4 != nil {
				ip = ipv4
			}
		}
	} else if recordType == "AAAA" {
		for _, i := range ips {
			if ipv4 := i.To4(); ipv4 == nil {
				ip = i
			}
		}
	}
	log.Printf("Default: %s %s %s", domain, recordType, ip)
	return ip
}

func (s *grpcService) QueryDns(domain string, recordType string) (ip net.IP) {
	ip = s.QueryDnsToMycelResolver(domain, recordType)
	if ip == nil {
		ip = s.QueryDnsToDefaultResolver(domain, recordType)
	}
	return ip
}

func (s *grpcService) HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		ip := s.QueryDns(domain, "A")
		if ip == nil {
			break
		}
		msg.Answer = append(msg.Answer, &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 600},
			A:   ip,
		})
	}

	w.WriteMsg(&msg)
}

func RunDnsServer(nodeAddress string, listenPort int) error {
	grpcConn, err := grpc.Dial(nodeAddress,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to node: %v", err)
	}
	defer grpcConn.Close()

	grpc := grpcService{
		grpcConn: grpcConn,
	}

	dns.HandleFunc(".", grpc.HandleDNSRequest)
	server := &dns.Server{Addr: fmt.Sprintf(":%d", listenPort), Net: "udp"}
	fmt.Printf("Starting DNS server at %s\n", server.Addr)
	err = server.ListenAndServe()
	return err
}

// DnsCommand returns command to start DNS server
func DnsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns",
		Short: "Run DNS server",
		RunE: func(cmd *cobra.Command, _ []string) (err error) {
			listenPort, _ := cmd.Flags().GetInt("port")
			nodeAddress, _ := cmd.Flags().GetString("nodeAddress")
			err = RunDnsServer(nodeAddress, listenPort)
			return err
		},
	}
	cmd.PersistentFlags().Int("port", 53, "Port to listen on")
	cmd.PersistentFlags().String("nodeAddress", "localhost:9090", "Mycel node address")
	return cmd
}
