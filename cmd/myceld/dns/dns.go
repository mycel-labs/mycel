package dns

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	resolver "github.com/mycel-domain/mycel/x/resolver/types"
)

type grpcService struct {
	grpcConn *grpc.ClientConn
}

type DnsRecord interface{}

func (s *grpcService) QueryDnsToMycelResolver(domain string, recordType string) (dnsRecord DnsRecord) {
	domain = strings.Trim(domain, ".")
	division := strings.Index(domain, ".")

	if division < 0 {
		log.Printf("QueryDnsToMycelResolver: %s, %s", "invalid domain format", domain)
		return nil
	}

	argName := domain[:division]
	argParent := domain[division+1:]

	queryClient := resolver.NewQueryClient(s.grpcConn)

	params := &resolver.QueryDnsRecordRequest{
		DomainName:    argName,
		DomainParent:  argParent,
		DnsRecordType: recordType,
	}

	ctx := context.Background()

	res, err := queryClient.DnsRecord(ctx, params)
	if err != nil {
		log.Printf("QueryDnsToMycelResolver: %v", err)
		return nil
	}

	if found := res.Value != nil; found {
		value := res.Value.Value
		switch recordType {
		case "A", "AAAA":
			dnsRecord = net.ParseIP(value)
		case "CNAME", "MX", "NS":
			dnsRecord = value
		case "TXT":
			dnsRecord = []string{value}
		default:
			dnsRecord = nil
		}
	}
	return dnsRecord
}

func (s *grpcService) QueryDnsToDefaultResolver(domain string, recordType string) DnsRecord {
	switch recordType {
	case "A":
		ips, err := net.LookupIP(domain)
		if err != nil {
			return nil
		}
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				return ipv4
			}
		}
	case "AAAA":
		ips, err := net.LookupIP(domain)
		if err != nil {
			return nil
		}
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 == nil {
				return ip
			}
		}
	case "CNAME":
		cname, err := net.LookupCNAME(domain)
		if err != nil {
			return nil
		}
		return cname
	case "MX":
		mxs, err := net.LookupMX(domain)
		if err != nil || len(mxs) == 0 {
			return nil
		}
		return mxs[0].Host
	case "NS":
		nss, err := net.LookupNS(domain)
		if err != nil || len(nss) == 0 {
			return nil
		}
		return nss[0].Host
	case "TXT":
		txts, err := net.LookupTXT(domain)
		if err != nil {
			return nil
		}
		return txts
	default:
		return nil
	}
	return nil
}

func (s *grpcService) QueryDns(domain string, recordType string) (dnsRecord DnsRecord) {
	dnsRecord = s.QueryDnsToMycelResolver(domain, recordType)
	log.Printf("MycelResolver: %s %s %v", domain, recordType, dnsRecord)
	if dnsRecord == nil {
		dnsRecord = s.QueryDnsToDefaultResolver(domain, recordType)
		log.Printf("DefaultResolver: %s %s %v", domain, recordType, dnsRecord)
	}
	return dnsRecord
}

func (s *grpcService) HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	domain := msg.Question[0].Name
	found := false

	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		record := s.QueryDns(domain, "A")
		if ip, ok := record.(net.IP); ok && ip != nil {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 600},
				A:   ip,
			})
			found = true
		}
	case dns.TypeAAAA:
		record := s.QueryDns(domain, "AAAA")
		if ip, ok := record.(net.IP); ok && ip != nil {
			msg.Answer = append(msg.Answer, &dns.AAAA{
				Hdr:  dns.RR_Header{Name: domain, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 600},
				AAAA: ip,
			})
			found = true
		}
	case dns.TypeCNAME:
		record := s.QueryDns(domain, "CNAME")
		if cname, ok := record.(string); ok && cname != "" {
			msg.Answer = append(msg.Answer, &dns.CNAME{
				Hdr:    dns.RR_Header{Name: domain, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 600},
				Target: cname,
			})
			found = true
		}
	case dns.TypeMX:
		record := s.QueryDns(domain, "MX")
		if mx, ok := record.(string); ok && mx != "" {
			msg.Answer = append(msg.Answer, &dns.MX{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeMX, Class: dns.ClassINET, Ttl: 600},
				Mx:  mx,
			})
			found = true
		}
	case dns.TypeNS:
		record := s.QueryDns(domain, "NS")
		if ns, ok := record.(string); ok && ns != "" {
			msg.Answer = append(msg.Answer, &dns.NS{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: 600},
				Ns:  ns,
			})
			found = true
		}
	case dns.TypeTXT:
		record := s.QueryDns(domain, "TXT")
		if txt, ok := record.([]string); ok && len(txt) > 0 {
			msg.Answer = append(msg.Answer, &dns.TXT{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 600},
				Txt: txt,
			})
			found = true
		}
	default:
		break
	}

	if !found {
		msg.SetRcode(r, dns.RcodeNameError)
	}

	err := w.WriteMsg(&msg)
	if err != nil {
		log.Printf("Failed to write message: %v", err)
	}
}

func RunDnsServer(nodeAddress string, listenPort int) error {
	grpcConn, err := grpc.Dial(nodeAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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
