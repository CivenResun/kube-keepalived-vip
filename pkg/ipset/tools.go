package ipset

import (
	"k8s.io/klog"
	"time"
)

const IPSetName = "hcvip"

func InitVIPSet() {
	for {
		_, err := New(IPSetName, "hash:net", Params{})
		if err != nil {
			klog.Infof("err is %v", err)
		}
		if err != nil {
			klog.Errorf("error add  ipset set  hcvip  %err", err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
}

func RefreshIPSet(setList []string) error {
	mcastdst, err := New(IPSetName, "hash:net", Params{})
	if err != nil {
		klog.Infof("err is %v", err)
		return err
	}
	for _, ip := range setList {
		err = mcastdst.Add(ip, 0)
		if err != nil {
			klog.Errorf("mcastdst add %v err %v", ip, err)
		}
	}

	return nil
}

func DelIPSet(setList []string) error {
	mcastdst, err := New(IPSetName, "hash:net", Params{})
	if err != nil {
		klog.Infof("err is %v", err)
		return err
	}
	for _, ip := range setList {
		err = mcastdst.Del(ip)
		if err != nil {
			klog.Errorf("mcastdst del %v err %v", ip, err)
		}
	}
	return nil
}
