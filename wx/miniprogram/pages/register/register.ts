import { routing } from "../../utils/routing"

// pages/register/register.ts
Page({
  redirectURL: '',
  data: {
    licNo:'123456',
    name: '张三',
    genderIndex: 0,
    genders: ['未知', '男', '女', '保密'],
    licImgURL: '',
    checkImgURL: '/resources/check.png',
    birthDate: '1990-01-01',
    state: 'UNSUBMITTED' as 'UNSUBMITTED' | 'PENDING' | 'VERIFIED',
  },
  onLoad(opt: Record<'redirect', string>) {
    const o: routing.RegisterOpts = opt
    if(o.redirect) {
      this.redirectURL = decodeURIComponent(o.redirect)
    }
  },
  onUploadLic() {
    console.log('onUploadLic')
    wx.chooseImage({
      success: res => {
        // console.log(res)
        if (res.tempFilePaths.length > 0) {
          this.setData({
            licImgURL: res.tempFilePaths[0]
          })
      }}
    })
  },
  onGenderChange(e: any) {
    this.setData({
      genderIndex: e.detail.value,
    })
  },
  onBirthDateChange(e: any) {
    this.setData({
      birthDate: e.detail.value,
    })
  },
  onSubmit() {
    this.setData({
      state: 'PENDING',
    })
    setTimeout(() =>{
      this.onLicVerified()
    }, 3000);
  },
  onResubmit() {
    this.setData({
      state: 'UNSUBMITTED',
      licImgURL: '',
    }
  )},
  onLicVerified() {
    this.setData({
      state: 'VERIFIED', 
    })
    if(this.redirectURL) {
      wx.redirectTo({
        url: this.redirectURL,
      })
    }
  },

})