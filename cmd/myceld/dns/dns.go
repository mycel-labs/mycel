package dns

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
	registry "github.com/mycel-domain/mycel/x/registry/types"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

type grpcService struct {
	grpcConn *grpc.ClientConn
}

func (s *grpcService) QueryDNStoMycel(domain string, recordType string) net.IP {

	domain = strings.Trim(domain, ".")
	division := strings.Index(domain, ".")

	argName := domain[:division]
	argParent := domain[division+1:]

	fmt.Printf("query: %v\n", domain)
	fmt.Printf("argName: %v\n", argName)
	fmt.Printf("argParent: %v\n", argParent)

	queryClient := registry.NewQueryClient(s.grpcConn)

	params := &registry.QueryGetSecondLevelDomainRequest{
		Name:   argName,
		Parent: argParent,
	}

	res, err := queryClient.SecondLevelDomain(context.Background(), params)
	if err != nil {
		log.Printf("queryClient.Domain failed: %v", err)
		return nil
	}
	log.Printf("result: %v", res.SecondLevelDomain.DnsRecords)

	if val, found := res.SecondLevelDomain.DnsRecords[recordType]; found {
		fmt.Printf("record: %v", val)
		return net.ParseIP(val.Value)
	} else {
		log.Printf("record not found")
		return nil
	}
}

func (s *grpcService) HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		ip := s.QueryDNStoMycel(domain, "A")
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
			err = RunDnsServer(nodeAddress,listenPort)
			return err
		},
	}
	cmd.PersistentFlags().Int("port", 53, "Port to listen on")
	cmd.PersistentFlags().String("nodeAddress", "localhost:9090", "Mycel node address")
	return cmd
}
