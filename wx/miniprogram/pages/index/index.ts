import { routing } from "../../utils/routing"
import { IAppOption } from "../../appoption"

Page({
  isPageShowing: false,
  data: {
    avatarURL: '',
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: '',
      layerStyle: -1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    location: {
      // latitude: initialLat,
      // longitude: initialLng,
      latitude: 31,
      longitude: 120,
    },
    scale: 10,
    // markers: [] as Marker[],
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: 23.099994,
        longitude: 113.324520,
        width: 50,
        height: 50,
      },
      {
        iconPath: "/resources/car.png",
        id: 1,
        latitude: 23.099994,
        longitude: 113.324520,
        width: 50,
        height: 50,
      }
    ],
    // showCancel: true,
    // showModal: false,
  },
  async onLoad() {
    const userInfo = await getApp<IAppOption>().globalData.userInfo
    this.setData({
      avatarURL: userInfo.avatarUrl,
    })
  },
  onMyLocationTap() {
    wx.getLocation({
      type: 'gcj02',
      success: res => {
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
        })
      },
      fail: () => {
        wx.showToast({
          icon: 'none',
          title: '请前往设置页授权位置信息',
        })
      }
    })
  },
  onShow() {
    this.isPageShowing = true
  }
  ,
  onHide(){
    this.isPageShowing = false
  },
  onMyTripsTap() {
    wx.navigateTo({
      url: routing.mytrips(),
    })
  },

  // moveCars() {
  //   const map = wx.createMapContext("map")
  //   const dest = {
  //     latitude: 23.099994,
  //     longitude: 113.324520,
  //   }
  //   // 持续移动 移动过程放在函数体中
  //   const moveCar = () => {
  //     dest.latitude += 0.1
  //     dest.longitude += 0.1
  //     console.log(`Moving car to ${dest.latitude}, ${dest.longitude}`)
  //     map.translateMarker({
  //       destination:{
  //         latitude: dest.latitude,
  //         longitude: dest.longitude,
  //       },
  //       markerId: 0,
  //       autoRotate: false,
  //       rotate: 0,
  //       duration: 1000,
  //       // animationEnd: () => {
  //       //   if (this.isPageShowing) {
  //       //     moveCar()
  //       //   }
  //       // },
  //     })
  //   }
  //   moveCar()
  // }
  onScanTap() {
    wx.scanCode({
      // onlyFromCamera: true,
      // success: console.log,
      fail: console.error,
      success: async () => {
        // wx.showModal({
        //   title: '身份认证',
        //   content: '需要身份认证才能租车',
        //   success: () => {

        //   }
        // })
        await this.selectComponent('#licModal').showModal()
          // TODO: get car_id from scan result
          const carID = 'car123'
          // const redirectURL = `/pages/lock/lock?car_id=${carID}`
          const redirectURL = routing.lock({
            car_id: carID,
          })
          wx.navigateTo({
            // url: `/pages/register/register?redirect=${encodeURIComponent(redirectURL)}`,
            url: routing.register({
              redirectURL: redirectURL,
            })
          })
      },
      
    })
  },

})