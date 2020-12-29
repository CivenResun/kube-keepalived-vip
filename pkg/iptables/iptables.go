package iptables

import (
	"github.com/aledbf/kube-keepalived-vip/pkg/ipset"
	"github.com/coreos/go-iptables/iptables"
	"k8s.io/klog"
	"os"
)

const mark = "0x6666"

func InitIPtables() {
	ipt, err := iptables.New()
	if err != nil {
		klog.Errorf("New failed: %v", err)
		os.Exit(1)
	}
	err = ipt.Append("nat", "PREROUTING", "-m", "set", "--match-set", ipset.IPSetName, "dst", "-j", "MARK", "--set-mark", mark)
	if err != nil {
		klog.Infof("New failed: %v", err)
		os.Exit(1)
	}
	err = ipt.Append("nat", "POSTROUTING", "-m", "mark", "--mark", mark, "-j", "MASQUERADE")
	if err != nil {
		klog.Infof("New failed: %v", err)
		os.Exit(1)
	}

}
