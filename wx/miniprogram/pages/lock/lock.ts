import { routing } from "../../utils/routing"
// pages/lock/lock.ts
const shareLocationKey = "share_location"
Page({
  data: {
    shareLocation: false,
    avatarURL: '',
    // userInfo: {},
    // hasUserInfo: false,
    // canIUseGetUserProfile: false,
  },
  async onLoad(opt:Record<'car_id', string> ) {
    const o: routing.LockOpts = opt
    console.log('unlocking car', o.car_id)
    
    const userInfo = await getApp<IAppOption>().globalData.userInfo
    this.setData({
      avatarURL: userInfo.avatarUrl,
      shareLocation: wx.getStorageSync(shareLocationKey) || false
    })
    wx.setStorageSync(shareLocationKey, true)
  },
  // onLoad() {
  //   wx.getStorageSync(shareLocationKey),
  //   // if (wx.getUserProfile) {
  //     this.setData({
  //       canIUseGetUserProfile: true,
  //       shareLocation: wx.getStorageSync(shareLocationKey) || false
  //     })
  //   // }
  // },
  // getUserProfile() {
  //   wx.getUserProfile({
  //     desc: '实时获取用户头像',
  //     success: (res: any) => {
  //       this.setData({
  //         avatarURL: res.userInfo.avatarUrl,
  //         // userInfo: res.userInfo,
  //         shareLocation: true,
  //         // hasUserInfo: true
  //       })
  //     }
  //   })
  // },
  onGetUserInfo(e: any) { 
    console.log(e)
    const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
    if (userInfo) {
      getApp<IAppOption>().resolveUserInfo(userInfo)
      this.setData({
          shareLocation: true,
      })
    }
  },
  // getUserInfo(e: any) {
  //   // 不推荐使用getUserInfo获取用户信息，预计自2021年4月13日起，getUserInfo将不再弹出弹窗，并直接返回匿名的用户个人信息
  //   this.setData({
  //     userInfo: e.detail.userInfo,
  //     hasUserInfo: true
  //   })
  // },
  onShareLocation(e:any) {
    const shareLocation: boolean = e.detail.value
    wx.setStorageSync(shareLocationKey, shareLocation)
  },
  onUnlockTap() {
    wx.getLocation({
      type: 'gcj02',
      success: loc => {
        console.log('starting a trip', {
          location: {
            latitude: loc.latitude,
            longitude: loc.longitude,
          },
          avatarURL: this.data.shareLocation ?
          this.data.avatarURL : '',
        })
        const tripID = 'trip456'
        wx.showLoading({
          title: '开锁中',
          mask: true
        })
        setTimeout(() => {
          wx.redirectTo({
            // url: `/pages/driving/driving?trip_id=${tripID}`,
            url: routing.driving({
              trip_id:tripID
            }),
            complete: () => {
              wx.hideLoading()
            }
          })
        },2000)
      },
      fail: () => {
        wx.showToast({
          icon: 'none',
          title: '请前往设置页授权位置信息',
        })
      }
    })
  },

})