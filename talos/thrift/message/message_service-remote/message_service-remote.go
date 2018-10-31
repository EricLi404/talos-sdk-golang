// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"../../../../thrift"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"../../message"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  PutMessageResponse putMessage(PutMessageRequest request)")
	fmt.Fprintln(os.Stderr, "  GetMessageResponse getMessage(GetMessageRequest request)")
	fmt.Fprintln(os.Stderr, "  GetTopicOffsetResponse getTopicOffset(GetTopicOffsetRequest request)")
	fmt.Fprintln(os.Stderr, "  GetPartitionOffsetResponse getPartitionOffset(GetPartitionOffsetRequest request)")
	fmt.Fprintln(os.Stderr, "  GetPartitionsOffsetResponse getPartitionsOffset(GetPartitionsOffsetRequest request)")
	fmt.Fprintln(os.Stderr, "  GetScheduleInfoResponse getScheduleInfo(GetScheduleInfoRequest request)")
	fmt.Fprintln(os.Stderr, "  Version getServiceVersion()")
	fmt.Fprintln(os.Stderr, "  void validClientVersion(Version clientVersion)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := message.NewMessageServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "putMessage":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PutMessage requires 1 args")
			flag.Usage()
		}
		arg21 := flag.Arg(1)
		mbTrans22 := thrift.NewTMemoryBufferLen(len(arg21))
		defer mbTrans22.Close()
		_, err23 := mbTrans22.WriteString(arg21)
		if err23 != nil {
			Usage()
			return
		}
		factory24 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt25 := factory24.GetProtocol(mbTrans22)
		argvalue0 := message.NewPutMessageRequest()
		err26 := argvalue0.Read(jsProt25)
		if err26 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PutMessage(value0))
		fmt.Print("\n")
		break
	case "getMessage":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetMessage requires 1 args")
			flag.Usage()
		}
		arg27 := flag.Arg(1)
		mbTrans28 := thrift.NewTMemoryBufferLen(len(arg27))
		defer mbTrans28.Close()
		_, err29 := mbTrans28.WriteString(arg27)
		if err29 != nil {
			Usage()
			return
		}
		factory30 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt31 := factory30.GetProtocol(mbTrans28)
		argvalue0 := message.NewGetMessageRequest()
		err32 := argvalue0.Read(jsProt31)
		if err32 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetMessage(value0))
		fmt.Print("\n")
		break
	case "getTopicOffset":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTopicOffset requires 1 args")
			flag.Usage()
		}
		arg33 := flag.Arg(1)
		mbTrans34 := thrift.NewTMemoryBufferLen(len(arg33))
		defer mbTrans34.Close()
		_, err35 := mbTrans34.WriteString(arg33)
		if err35 != nil {
			Usage()
			return
		}
		factory36 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt37 := factory36.GetProtocol(mbTrans34)
		argvalue0 := message.NewGetTopicOffsetRequest()
		err38 := argvalue0.Read(jsProt37)
		if err38 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTopicOffset(value0))
		fmt.Print("\n")
		break
	case "getPartitionOffset":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPartitionOffset requires 1 args")
			flag.Usage()
		}
		arg39 := flag.Arg(1)
		mbTrans40 := thrift.NewTMemoryBufferLen(len(arg39))
		defer mbTrans40.Close()
		_, err41 := mbTrans40.WriteString(arg39)
		if err41 != nil {
			Usage()
			return
		}
		factory42 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt43 := factory42.GetProtocol(mbTrans40)
		argvalue0 := message.NewGetPartitionOffsetRequest()
		err44 := argvalue0.Read(jsProt43)
		if err44 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetPartitionOffset(value0))
		fmt.Print("\n")
		break
	case "getPartitionsOffset":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPartitionsOffset requires 1 args")
			flag.Usage()
		}
		arg45 := flag.Arg(1)
		mbTrans46 := thrift.NewTMemoryBufferLen(len(arg45))
		defer mbTrans46.Close()
		_, err47 := mbTrans46.WriteString(arg45)
		if err47 != nil {
			Usage()
			return
		}
		factory48 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt49 := factory48.GetProtocol(mbTrans46)
		argvalue0 := message.NewGetPartitionsOffsetRequest()
		err50 := argvalue0.Read(jsProt49)
		if err50 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetPartitionsOffset(value0))
		fmt.Print("\n")
		break
	case "getScheduleInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetScheduleInfo requires 1 args")
			flag.Usage()
		}
		arg51 := flag.Arg(1)
		mbTrans52 := thrift.NewTMemoryBufferLen(len(arg51))
		defer mbTrans52.Close()
		_, err53 := mbTrans52.WriteString(arg51)
		if err53 != nil {
			Usage()
			return
		}
		factory54 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt55 := factory54.GetProtocol(mbTrans52)
		argvalue0 := message.NewGetScheduleInfoRequest()
		err56 := argvalue0.Read(jsProt55)
		if err56 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetScheduleInfo(value0))
		fmt.Print("\n")
		break
	case "getServiceVersion":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetServiceVersion requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetServiceVersion())
		fmt.Print("\n")
		break
	case "validClientVersion":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ValidClientVersion requires 1 args")
			flag.Usage()
		}
		arg57 := flag.Arg(1)
		mbTrans58 := thrift.NewTMemoryBufferLen(len(arg57))
		defer mbTrans58.Close()
		_, err59 := mbTrans58.WriteString(arg57)
		if err59 != nil {
			Usage()
			return
		}
		factory60 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt61 := factory60.GetProtocol(mbTrans58)
		argvalue0 := message.NewVersion()
		err62 := argvalue0.Read(jsProt61)
		if err62 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ValidClientVersion(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}