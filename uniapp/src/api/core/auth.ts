/**
 * 登录态管理：会员 token（客人与员工共用同一套会员体系）
 *
 * 说明：小程序端用 uni.login 换 code，后端换 openid 并签发 token；
 *      H5 / 开发环境没有 uni.login，退化为设备码登录（后端以 dev_ 前缀区分）。
 */
const TOKEN_KEY = 'tavern_token'
const DEVICE_KEY = 'tavern_device'

export function getToken(): string {
  return uni.getStorageSync(TOKEN_KEY) || ''
}

export function setToken(token: string) {
  uni.setStorageSync(TOKEN_KEY, token)
}

export function clearToken() {
  uni.removeStorageSync(TOKEN_KEY)
}

function getBaseUrl() {
  return (import.meta.env.VITE_API_BASE_URL as string) || 'http://127.0.0.1:8000/api'
}

/** 取登录凭证：小程序走 uni.login，H5 走设备码 */
async function resolveLoginCode(): Promise<string> {
  // #ifdef MP-WEIXIN
  const code = await new Promise<string>((resolve) => {
    uni.login({
      success: (res: any) => resolve(res?.code || ''),
      fail: () => resolve(''),
    })
  })
  if (code)
    return code
  // #endif

  let device = uni.getStorageSync(DEVICE_KEY)
  if (!device) {
    device = `h5-${Date.now()}`
    uni.setStorageSync(DEVICE_KEY, device)
  }
  return device
}

/** 登录（幂等：已有 token 直接返回） */
export async function ensureLogin(force = false): Promise<string> {
  const exist = getToken()
  if (exist && !force)
    return exist

  const code = await resolveLoginCode()
  const url = `${getBaseUrl()}/tavern/member/login`
  const res: any = await new Promise((resolve, reject) => {
    uni.request({
      url,
      method: 'POST',
      header: { 'Content-Type': 'application/json' },
      data: { code },
      success: resolve,
      fail: reject,
    })
  })

  const body = res?.data
  if (!body || body.code !== 0 || !body.data?.token)
    throw new Error(body?.message || '登录失败')

  setToken(body.data.token)
  return body.data.token
}
