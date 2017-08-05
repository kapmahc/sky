import enUSAntd from 'antd/lib/locale-provider/en_US'
import zhTWAntd from 'antd/lib/locale-provider/zh_TW'

import 'moment/locale/zh-cn'
import 'moment/locale/zh-tw'

import Cookie from 'js-cookie'

import dataEn from 'react-intl/locale-data/en'
import dataZh from 'react-intl/locale-data/zh'

import enUS from './locales/en-US'
import zhHans from './locales/zh-Hans'
import zhHant from './locales/zh-Hant'


const KEY = 'locale'

export const setLocale = (lng) => {
  localStorage.setItem(KEY, lng)
  Cookie.set(KEY, lng, {expires: 365})
  window.location.reload()
}


export const detectLocale = () => {
  switch (localStorage.getItem(KEY)) {
    case 'zh-Hans':
      return {
        locale: 'zh-Hans',
        antd: null,
        data: dataZh,
        moment: 'zh-cn',
        messages: zhHans,
      }
    case 'zh-Hant':
      return {
        locale: 'zh-Hant',
        antd: zhTWAntd,
        data: dataZh,
        moment: 'zh-tw',
        messages: zhHant,
      }
    default:
      return {
        locale: 'en-US',
        antd: enUSAntd,
        data: dataEn,
        moment: 'en',
        messages: enUS,
      }
  }
}
