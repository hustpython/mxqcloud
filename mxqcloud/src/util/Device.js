let Device = "Windows PC"
let userAgent = navigator.userAgent
// 获取微信版本
let m1 = ""

// 苹果手机
if (userAgent.includes('iPhone') || userAgent.includes('iPad')) {
    // 获取设备名
    if (userAgent.includes('iPad')) {
        Device = 'iPad'
    } else {
        Device = 'iPhone'
    }
}

// 安卓手机
if (userAgent.includes('Android')) {
    // 获取设备名
    m1 = userAgent.match(/Android.*; ?(.*(?= AppleWebKit))/)
    if (m1 && m1.length > 1) {
        Device = m1[1].replace(')', "")
    }
}

export function GetDevice() {

    return Device;
}