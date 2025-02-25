import { getSetting, getUserInfo } from "./utils/wxapi"
import { IAppOption } from "./appoption"

let resolveUserInfo: (value: WechatMiniprogram.UserInfo | PromiseLike<WechatMiniprogram.UserInfo>) => void;
let rejectUserInfo: (reason?: any) => void;

// app.ts
App<IAppOption>({
  globalData: {
    userInfo: new Promise((resolve, reject) => {
      resolveUserInfo = resolve
      rejectUserInfo = reject
    }),  
},
  async onLaunch() {
    // // 展示本地存储能力
    // const logs = wx.getStorageSync('logs') || []
    // logs.unshift(Date.now())
    // wx.setStorageSync('logs', logs)
    wx.request({
      url: 'http://10.99.192.85:8080/trip/trip123',
      method: 'GET',
      success: console.log,
      fail: console.error,
    })

    // 登录
    wx.login({
      success: res => {
        console.log(res.code)
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })

      //   userInfo: new Promise((resolve, reject) => {
      // // 获取用户信息 promise版本
  try {
    const setting = await getSetting()
    if (setting.authSetting['scope.userInfo']) {
      const userInfoRes = await getUserInfo()
      resolveUserInfo(userInfoRes.userInfo)
    }
  } catch (err) {
    rejectUserInfo(err)
  }

  //   getSetting().then(res => {
  //     if(res.authSetting['scope.userInfo']) {
  //       return getUserInfo()
  //     }
  //     return undefined
  // }).then(res => {
  //   if(!res) {
  //     return
  //   }
  //   // this.globalData.userInfo = res?.userInfo // 加入?可处理undefined情况
  //   // 通知页面获得了用户信息
  //   resolveUserInfo(res.userInfo)
  // }).catch(rejectUserInfo)
 
  },

    // // 获取用户信息 promise版本
    // getSetting().then(res => {
    //     if(res.authSetting['scope.userInfo']) {
    //       return getUserInfo()
    //     }
    //     return Promise.resolve(undefined)
    // }).then(res => {
    //   if(!res) {
    //     return
    //   }
    //   // this.globalData.userInfo = res?.userInfo // 加入?可处理undefined情况
    //   // 通知页面获得了用户信息
    //   if (this.userInfoReadyCallback) {
    //     this.userInfoReadyCallback(res)
    //   }
    // })

    // // 获取用户信息
    // wx.getSetting({
    //   success: res => {
    //     if (res.authSetting['scope.userInfo']) {
    //       // 已经授权，可以直接调用 getUserInfo 获取头像昵称，不会弹框
    //       wx.getUserInfo({  
    //         success: res => {
    //           console.log(res.userInfo)
    //           // 可以将 res 发送给后台解码出 unionId
    //           this.globalData.userInfo = res.userInfo
    //           // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
    //           // 所以此处加入 callback 以防止这种情况
    //           if (this.userInfoReadyCallback) {
    //             this.userInfoReadyCallback(res)
    //           }
    //         },  
    //     })
    //   }
    // },
    // })
  // },
  resolveUserInfo(userInfo:WechatMiniprogram.UserInfo) {
    resolveUserInfo(userInfo)
  } 
    //   getUserProfile() {
    //   // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    //   wx.getUserProfile({
    //     desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
    //     success: (res) => {
    //       console.log(res)
    //       this.setData({
    //         userInfo: res.userInfo,
    //         hasUserInfo: true
    //       })
    //     }
    //   })
    // },
})
