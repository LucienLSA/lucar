// pages/driving/driving.ts

import { routing } from "../../utils/routing"

// 将秒时间格式转化为特定的时间字符串
function formatDuration(sec: number) {
  const padString = (n: number) => 
    n < 10 ? '0' + n.toFixed(0) : n.toFixed(0)
  const h = Math.floor(sec / 3600) // 将浮点数格式化为整数
  sec -= h * 3600
  const m = Math.floor(sec / 60)
  sec -= m * 60
  const s = Math.floor(sec)
  return `${padString(h)}:${padString(m)}:${padString(s)}`
}

function formatFee(cents: number) {
  return (cents / 100).toFixed(2) // 计算分
}

// 上线后的费用是服务器给出价格或者将费用传给用户进行计算
const centPerSec = 0.8 

Page({
  timer: undefined as number | undefined,
  data: {
    location: {
      latitude: 32.92,
      longitude: 117.13
    },
    scale: 14,
    elapsed: '00:00:00',
    fee: '0.00',
  },
  onLoad(opt: Record<'trip_id', string>) {
    const o: routing.DrivingOpts = opt
    console.log('current trip:', o.trip_id)
    this.setupLocationUpdator(),
    this.setupTimer()
  },
  onUnload() {
    wx.stopLocationUpdate()
    if(this.timer) {
      clearInterval(this.timer)
    }
  },
  setupLocationUpdator() {
    wx.startLocationUpdate({
      fail: console.error,
    })
    wx.onLocationChange((loc) => {
      // console.log('location changed:',loc)
      this.setData({
        location: {
          latitude: loc.latitude,
          longitude: loc.longitude
        }
      })
    })
  },
  setupTimer() {
    let elapsedSec = 0 // 数字后加上单位
    let cents = 0
    this.timer = setInterval(() => {
      elapsedSec += 1
      cents += centPerSec
      this.setData({
        elapsed: formatDuration(elapsedSec),
        fee: formatFee(cents)
      })
    },1000)
  },
  onEndTripTap(){
    wx.redirectTo({
      url: routing.mytrips()
    })
  }
})