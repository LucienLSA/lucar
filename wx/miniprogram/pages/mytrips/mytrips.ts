import { routing } from "../../utils/routing"
// pages/mytrips/mytrips.ts

interface Trip {
  id: string
  start: string
  end: string
  duration: string 
  fee: string
  distance: string
  status: string
}

interface MainItem {
  id: string,
  navId: string,
  navScrollId: string
  data: Trip
}

interface NavItem {
  id: string,
  mainId: string,
  label: string
}

interface MainItemQueryResult{
  id: string,
  top: number,
  dataset: {
    navId: string,
    navScrollId: string
  }
}

Page({
  /**
   * 页面的初始数据
   */
  scrollStates: {
    mainItems:[] as MainItemQueryResult[]
  },
  data: {
    indicatorDots: true, // 下方点
    autoPlay: false, // 自动播放
    interval: 3000, // 播放时间
    duration:500, // 切换时间
    multiItemCount:1, // 显示数量
    prevMargin: '', // 与上一个间隔
    nextMargin: '', // 与下一个间隔
    vertical: false, // 是否垂直
    current: 0, // 当前索引
    // imgUrls: [
    //   'https://img14.360buyimg.com/babel/s1180x940_jfs/t20280217/256570/34/21884/63835/67b29ecdF9bf9a252/c611493158a71423.jpg.avif',
    //   'https://imgcps.jd.com/img-cubic/creative_server_cia_jdcloud/v2/2020217/100026891770/FocusFullshop/CkRqZnMvdDEvMjMxNjQwLzM4LzI1ODg5LzIzNTAwMC82NmM2YzE3NEZiZDRlOTA1Mi82ZjYwODE4NTE2NzJjMTAzLnBuZxIFNTE2LXcwATj5pntCHQoZ5Ye76Z-z6JOd54mZL-aXoOe6v-iAs-acuhABQhAKDOenkuadgOS7tzE1OBACQhAKDOeri-WNs-aKoui0rRAGQgoKBui2heWAvBAHWPr7xND0Ag/cr/s/q.jpg',
    //   'https://imgcps.jd.com/img-cubic/creative_server_cia_jdcloud/v2/2020218/10058917856760/FocusFullshop/CkJqZnMvdDEvMzYwMzIvMzcvMjI1NzcvNzEwMzAvNjZjNjgyNzNGMWE5NTQ4MjgvMmVlNDdkYTllNGNmMDRhMS5wbmcSBTcwNC10MAI4-qZ7QhMKD-WNjuW4neayueeDn-acuhABQhYKEueyvuW9qeS4jeWuuemUmei_hxACQhAKDOeri-WNs-aKoui0rRAGQgoKBuS8mOmAiRAHWPiL57HgpAI/cr/s/q.jpg',
    //   'https://imgcps.jd.com/img-cubic/creative_server_cia_jdcloud/v2/2020219/100063705048/FocusFullshop/CkNqZnMvdDEvNzM4NTIvMzQvMjcyNzkvMjQwNTIyLzY2Yzc4MWE0Rjk2NzI5N2E2L2M1MTg5NWVkOWUzMjhiNWIucG5nEgU3MDMtdDACOPume0ITCg_ogZTmg7PnrJTorrDmnKwQAUIQCgznlYXkuqvkvJjlk4EQAkIQCgznq4vljbPmiqLotK0QBkIHCgPmiqIQB1jY74vi9AI/cr/s/q.jpg',
    //   'https://img14.360buyimg.com/babel/s1180x940_jfs/t20280214/261207/17/19519/30822/67aeb8d7F20b039a8/e2f966c78a6fc044.jpg.avif',
    // ]
    promotionItems: [
      {
        img: 'https://img14.360buyimg.com/babel/s1180x940_jfs/t20280217/256570/34/21884/63835/67b29ecdF9bf9a252/c611493158a71423.jpg.avif',
        promotionID: 1,
      },
      {
        img: 'https://imgcps.jd.com/img-cubic/creative_server_cia_jdcloud/v2/2020217/100026891770/FocusFullshop/CkRqZnMvdDEvMjMxNjQwLzM4LzI1ODg5LzIzNTAwMC82NmM2YzE3NEZiZDRlOTA1Mi82ZjYwODE4NTE2NzJjMTAzLnBuZxIFNTE2LXcwATj5pntCHQoZ5Ye76Z-z6JOd54mZL-aXoOe6v-iAs-acuhABQhAKDOenkuadgOS7tzE1OBACQhAKDOeri-WNs-aKoui0rRAGQgoKBui2heWAvBAHWPr7xND0Ag/cr/s/q.jpg',
        promotionID: 2,
      },
      {
        img: 'https://imgcps.jd.com/img-cubic/creative_server_cia_jdcloud/v2/2020219/100063705048/FocusFullshop/CkNqZnMvdDEvNzM4NTIvMzQvMjcyNzkvMjQwNTIyLzY2Yzc4MWE0Rjk2NzI5N2E2L2M1MTg5NWVkOWUzMjhiNWIucG5nEgU3MDMtdDACOPume0ITCg_ogZTmg7PnrJTorrDmnKwQAUIQCgznlYXkuqvkvJjlk4EQAkIQCgznq4vljbPmiqLotK0QBkIHCgPmiqIQB1jY74vi9AI/cr/s/q.jpg',
        promotionID: 3,
      },
      {
        img:'https://img14.360buyimg.com/babel/s1180x940_jfs/t20280214/261207/17/19519/30822/67aeb8d7F20b039a8/e2f966c78a6fc044.jpg.avif',
        promotionID: 4,
      }
    ],
    tripsHeight:0, // 滚动高度
    navCount: 0, // 导航数量
    avatarURL: '', // 用户头像
    trips: [] as Trip[],
    // scrollTop: 0,
    // scrollIntoView: '',
    mainScroll: '',
    navItems: [] as NavItem[],
    mainItems: [] as MainItem[],
    navSel: '',
    navScroll: '',
  },
  // onSwiperChange(e:any) {
  //   console.log(e)
  // },
  async onLoad() {
    this.populateTrips()
    const userInfo = await getApp<IAppOption>().globalData.userInfo
    this.setData({
      avatarURL: userInfo.avatarUrl,
    })
  },
  onReady() {
    wx.createSelectorQuery().select('#heading')
      .boundingClientRect(rect => {
        const height = wx.getSystemInfoSync().windowHeight - rect.height
      this.setData({
        tripsHeight: height,
        navCount: Math.round(height/50),
      })  
    }).exec()
  },

  populateTrips() {
    // const trips: Trip[] = []
    const mainItems: MainItem[] = []
    const navItems: NavItem[] = []
    let navSel = ''
    let prevNav = ''
    for(let i=0; i<100; i++) {
      const mainId = 'main-' + i
      const navId = 'nav-' + i
      const tripId = (10000 + i).toString()
      if (!prevNav) {
        prevNav = navId
      }
      mainItems.push({
        id: mainId,
        navId: navId,
        navScrollId: prevNav,
        data:{
            id: tripId,
            start: '东方明珠',
            end: '迪士尼',
            duration: "0时44分",
            fee: '100元',
            distance: '27.0公里',
            status: '已完成',
        },
      })
      navItems.push({
        id: navId,
        mainId:mainId,
        label:tripId,
      })
      if(i === 0) {
        navSel = navId
      }
      prevNav = navId
    }
    this.setData({
      mainItems,
      navItems,
      navSel,
    }, () => {
      this.prepareScrollState() 
    })
  },
  prepareScrollState() {
    wx.createSelectorQuery().selectAll('.main-item')
      .fields({
        id: true,
        dataset: true,
        rect: true,
      }).exec(res=>{
        // console.log(res)
        this.scrollStates.mainItems = res[0]
      })
  },
  onPromotionItemTap(e:any) {
    console.log(e)
    const promotionID = e.currentTarget.dataset.promotionID
    if(promotionID) {
      // claims this promotionID
    }
  },
  onRegisterTap() {
    wx.navigateTo({
        url: routing.register(),
    })
  },
  onMainScroll(e: any) {
    // console.log(e)
    const top:number = e.currentTarget?.offsetTop + 
      e.detail?.scrollTop
    if(top === undefined) {
      return
    }
    const selItem = this.scrollStates.mainItems.find(
      v => v.top >= top)
    if (!selItem) {
      return
    }
    this.setData({
      navSel: selItem.dataset.navId,
      navScroll: selItem.dataset.navScrollId
    })
    
  },
  onNavItemTap(e: any) {
    const mainId:string = e.currentTarget?.dataset?.mainId
    const navId:string = e.currentTarget?.id
    // console.log(mainId)
    if (mainId && navId) {
      this.setData({
        mainScroll: mainId,
        navSel: navId,
      })
    }
  },
  onGetUserInfo(e: any) {
    const userInfo: WechatMiniprogram.UserInfo = e.detail.userInfo
    if (userInfo) {
        getApp<IAppOption>().resolveUserInfo(userInfo)
        this.setData({
            avatarURL: userInfo.avatarUrl,
        })
    }
},
})

