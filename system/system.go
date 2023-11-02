package system

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
)

// ReportSystemMemory 报告系统内存使用
func ReportSystemMemory(minPercent float64) string {
	v, err := mem.VirtualMemory()
	var info string
	if err != nil {
		info = fmt.Sprintf("danger:监控程序无法获取系统内存 %v", err.Error())
	} else {
		if v.UsedPercent >= minPercent {
			info = fmt.Sprintf("注意系统内存使用超过 %v%%, UsedPercent:%.1f%%\n", minPercent, v.UsedPercent)
		}
	}
	return info
}

// ReportSystemCpu 报告系统cpu使用
func ReportSystemCpu(minPercent float64) string {
	var info string
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		info = fmt.Sprintf("无法获取CPU使用情况：%v\n", err.Error())
	} else {
		if percent[0] >= minPercent {
			info = fmt.Sprintf("注意系统CPU使用超过 %v%%,UsedPercent:%.1f%%\n", minPercent, percent[0])
		}
	}
	return info
}
