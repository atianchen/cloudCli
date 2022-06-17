package monitor

import "cloudCli/component"

type DetectResult struct {
}

/**
 *
 * @author jensen.chen
 * @date 2022/6/17
 */
type IDetect interface {
	/**
	执行监控
	*/
	Execute(app component.AppInfo) DetectResult
}
