// 导入axios实例
import httpRequest from './request'

export function GetDeviceByName() {
    return httpRequest({
        url: '/device',
        method: 'get',
    })
}