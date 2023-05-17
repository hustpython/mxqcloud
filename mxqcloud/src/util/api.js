// 导入axios实例
import httpRequest from './request'

export function GetDeviceByName(param) {
    return httpRequest({
        url: '/device?Name=' + param,
        method: "GET",
    })
}

export function AddDevice(param) {
    return httpRequest({
        url: '/device',
        method: "POST",
        data: JSON.stringify(param),
    })
}

export function GetUploadData(param) {
    return httpRequest({
        url: '/upload?Name=' + param,
        method: "GET",
    })
}

export function DelUploadData(param) {
    return httpRequest({
        url: '/upload?Id=' + param,
        method: "DELETE",
    })
}

export function GetAllDownload(param) {
    return httpRequest({
        url: '/download',
        method: "GET",
    })
}